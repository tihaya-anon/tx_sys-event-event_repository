#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR/..
# Teardown script for Event Repository application

# Exit on error
set -e

# Parse command line arguments
TEARDOWN_DB=false
TEARDOWN_SRV=false
TEARDOWN_ALL=true  # Default to teardown all if no specific flag is provided
NAMESPACE="event-repo"  # Default namespace

usage() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  -d, --db            Teardown only PostgreSQL database"
    echo "  -s, --srv           Teardown only gRPC server"
    echo "  -a, --all           Teardown both database and server (default)"
    echo "  -n, --namespace     Specify Kubernetes namespace (default: event-repo)"
    echo "  -h, --help      Display this help message"
    exit 1
}

while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -d|--db)
        TEARDOWN_DB=true
        TEARDOWN_SRV=false
        TEARDOWN_ALL=false
        shift
        ;;
        -s|--srv)
        TEARDOWN_SRV=true
        TEARDOWN_DB=false
        TEARDOWN_ALL=false
        shift
        ;;
        -a|--all)
        TEARDOWN_ALL=true
        TEARDOWN_DB=false
        TEARDOWN_SRV=false
        shift
        ;;
        -n|--namespace)
        if [[ -z "$2" ]]; then
            echo "Error: --namespace requires a value"
            usage
        fi
        NAMESPACE="$2"
        shift 2
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

# Check if namespace exists and create it if it doesn't
ensure_namespace() {
    local namespace=$1
    if ! kubectl get namespace $namespace > /dev/null 2>&1; then
        echo "Namespace '$namespace' does not exist. Existed namespaces are:"
        kubectl get namespaces
        return 1
    else
        echo "Namespace '$namespace' already exists."
        return 0
    fi
}

# Teardown Kubernetes resources based on selected options
teardown_kubernetes() {
    if ! check_kubernetes; then
        echo "Skipping Kubernetes teardown steps."
        return 1
    fi
    
    if ! ensure_namespace $NAMESPACE; then
        return 1
    fi
    
    # Teardown gRPC server if requested
    if [[ "$TEARDOWN_SRV" == "true" || "$TEARDOWN_ALL" == "true" ]]; then
        echo "Tearing down gRPC server in namespace $NAMESPACE..."
        
        # Delete gRPC server resources
        kubectl delete -f ./k8s/service.yaml --ignore-not-found=true --namespace $NAMESPACE
        kubectl delete -f ./k8s/deployment.yaml --ignore-not-found=true --namespace $NAMESPACE
        kubectl delete -f ./k8s/secrets.yaml --ignore-not-found=true --namespace $NAMESPACE
        kubectl delete -f ./k8s/configmap.yaml --ignore-not-found=true --namespace $NAMESPACE
        kubectl delete -f ./k8s/schema-configmap.yaml --ignore-not-found=true --namespace $NAMESPACE
        
        echo "gRPC server teardown completed successfully!"
    fi
    
    # Teardown PostgreSQL if requested
    if [[ "$TEARDOWN_DB" == "true" || "$TEARDOWN_ALL" == "true" ]]; then
        echo "Tearing down PostgreSQL using Helm in namespace $NAMESPACE..."
        helm uninstall postgresql --wait --namespace $NAMESPACE
        
        echo "PostgreSQL teardown completed successfully!"
    fi
    
    # Delete namespace if it's empty
    if [[ "$TEARDOWN_ALL" == "true" ]]; then
        echo "Deleting namespace $NAMESPACE..."
        kubectl delete namespace $NAMESPACE
        echo "Namespace $NAMESPACE deleted successfully!"
    fi

    echo "Teardown process completed."
    return 0
}

# Execute the teardown
teardown_kubernetes
