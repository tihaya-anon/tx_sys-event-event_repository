#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR/..
# Deployment script for Event Repository application

# Exit on error
set -e

# Function to check if Docker image exists
check_and_build_image() {
    local force_rebuild=$1
    local image_name="event_repository_app"
    
    # Check if image exists
    if docker image inspect $image_name:latest >/dev/null 2>&1; then
        if [ "$force_rebuild" = "true" ]; then
            echo "Image $image_name exists, but force rebuild requested..."
            ./scripts/docker_build.sh
        else
            echo "Image $image_name already exists, skipping build..."
        fi
    else
        echo "Image $image_name does not exist, building..."
        ./scripts/docker_build.sh
    fi
}

# Parse command line arguments
FORCE_REBUILD=false
DEPLOY_DB=false
DEPLOY_APP=false
DEPLOY_ALL=true  # Default to deploy all if no specific flag is provided

usage() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  --rebuild       Force rebuild of Docker image"
    echo "  --db            Deploy only PostgreSQL database"
    echo "  --app           Deploy only application"
    echo "  --all           Deploy both database and application (default)"
    echo "  -h, --help      Display this help message"
    exit 1
}

while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        --rebuild)
        FORCE_REBUILD=true
        shift
        ;;
        --db)
        DEPLOY_DB=true
        DEPLOY_APP=false
        DEPLOY_ALL=false
        shift
        ;;
        --app)
        DEPLOY_APP=true
        DEPLOY_DB=false
        DEPLOY_ALL=false
        shift
        ;;
        --all)
        DEPLOY_ALL=true
        DEPLOY_DB=false
        DEPLOY_APP=false
        shift
        ;;
        -h|--help)
        usage
        ;;
        *)
        # Unknown option
        echo "Unknown option: $key"
        usage
        ;;
    esac
done

# Build the application Docker image if needed
check_and_build_image $FORCE_REBUILD

# Check if Kubernetes is accessible
check_kubernetes() {
    if kubectl cluster-info > /dev/null 2>&1; then
        echo "Kubernetes cluster is accessible."
        return 0
    else
        echo "ERROR: Kubernetes cluster is not accessible. Please ensure your cluster is running and properly configured."
        return 1
    fi
}

# Deploy to Kubernetes based on selected options
deploy_to_kubernetes() {
    if ! check_kubernetes; then
        echo "Skipping Kubernetes deployment steps."
        return 1
    fi
    
    # Deploy PostgreSQL if requested
    if [[ "$DEPLOY_DB" == "true" || "$DEPLOY_ALL" == "true" ]]; then
        echo "Deploying PostgreSQL using Helm..."
        helm install postgresql ./pkg/db/postgresql -f ./k8s/custom-values.yaml
        
        echo "Waiting for PostgreSQL to be ready..."
        kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=postgresql --timeout=300s
        
        echo "PostgreSQL deployment completed successfully!"
    fi
    
    # Deploy application if requested
    if [[ "$DEPLOY_APP" == "true" || "$DEPLOY_ALL" == "true" ]]; then
        # Apply Kubernetes manifests
        echo "Applying application Kubernetes manifests..."
        kubectl apply -f ./k8s/configmap.yaml
        kubectl apply -f ./k8s/secrets.yaml
        kubectl apply -f ./k8s/deployment.yaml
        kubectl apply -f ./k8s/service.yaml
        
        echo "Waiting for Event Repository application to be ready..."
        kubectl wait --for=condition=ready pod -l app=event-repository --timeout=300s
        
        echo "Application deployment completed successfully!"
        echo "You can access the gRPC application using the following command:"
        echo "  kubectl port-forward svc/event-repository 50051:50051"
        echo "Then use a gRPC client to connect to localhost:50051"
    fi
    
    echo "Deployment process completed."
    return 0
}

# Execute the deployment
deploy_to_kubernetes
