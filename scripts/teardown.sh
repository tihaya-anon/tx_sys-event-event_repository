#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR/..
# Teardown script for Event Repository application

# Exit on error
set -e

# Parse command line arguments
TEARDOWN_DB=false
TEARDOWN_APP=false
TEARDOWN_ALL=true  # Default to teardown all if no specific flag is provided

usage() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  --db            Teardown only PostgreSQL database"
    echo "  --app           Teardown only application"
    echo "  --all           Teardown both database and application (default)"
    echo "  -h, --help      Display this help message"
    exit 1
}

while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        --db)
        TEARDOWN_DB=true
        TEARDOWN_APP=false
        TEARDOWN_ALL=false
        shift
        ;;
        --app)
        TEARDOWN_APP=true
        TEARDOWN_DB=false
        TEARDOWN_ALL=false
        shift
        ;;
        --all)
        TEARDOWN_ALL=true
        TEARDOWN_DB=false
        TEARDOWN_APP=false
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

# Teardown Kubernetes resources based on selected options
teardown_kubernetes() {
    if ! check_kubernetes; then
        echo "Skipping Kubernetes teardown steps."
        return 1
    fi
    
    # Teardown application if requested
    if [[ "$TEARDOWN_APP" == "true" || "$TEARDOWN_ALL" == "true" ]]; then
        echo "Tearing down Event Repository application..."
        
        # Delete application resources
        kubectl delete -f ./k8s/service.yaml --ignore-not-found=true
        kubectl delete -f ./k8s/deployment.yaml --ignore-not-found=true
        kubectl delete -f ./k8s/secrets.yaml --ignore-not-found=true
        kubectl delete -f ./k8s/configmap.yaml --ignore-not-found=true
        
        echo "Application teardown completed successfully!"
    fi
    
    # Teardown PostgreSQL if requested
    if [[ "$TEARDOWN_DB" == "true" || "$TEARDOWN_ALL" == "true" ]]; then
        echo "Tearing down PostgreSQL using Helm..."
        helm uninstall postgresql --wait
        
        echo "PostgreSQL teardown completed successfully!"
    fi
    
    echo "Teardown process completed."
    return 0
}

# Execute the teardown
teardown_kubernetes
