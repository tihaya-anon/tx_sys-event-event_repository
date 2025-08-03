#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR/..
# Deployment script for Event Repository application

# Exit on error
set -e

# Build the application Docker image
echo "Building application Docker image..."
./scripts/docker_build.sh

# Deploy PostgreSQL using Helm
echo "Deploying PostgreSQL using Helm..."
helm install postgresql ./pkg/db/postgresql -f ./k8s/custom-values.yaml

echo "Waiting for PostgreSQL to be ready..."
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=postgresql --timeout=300s

# Apply Kubernetes manifests
echo "Applying Kubernetes manifests..."
kubectl apply -f ./k8s/configmap.yaml
kubectl apply -f ./k8s/secrets.yaml
kubectl apply -f ./k8s/deployment.yaml
kubectl apply -f ./k8s/service.yaml

echo "Waiting for Event Repository application to be ready..."
kubectl wait --for=condition=ready pod -l app=event-repository --timeout=300s

echo "Deployment completed successfully!"
echo "You can access the application using the following command:"
echo "  kubectl port-forward svc/event-repository 8080:80"
echo "Then visit http://localhost:8080 in your browser."
