# Event Repository Application Deployment Guide

This guide explains how to deploy the Event Repository gRPC application with PostgreSQL using Kubernetes and Helm.

## Prerequisites

- Kubernetes cluster up and running
- Helm installed (v3+)
- kubectl configured to connect to your cluster
- Docker installed for building the application image

## Deployment Steps

### 1. Build the Application Docker Image

```bash
# Run the Docker build script
./scripts/docker_build.sh
```

This script builds the application Docker image using the Dockerfile at `pkg/app/Dockerfile`.

### 2. Deploy PostgreSQL with Helm

The PostgreSQL database is deployed using the Bitnami Helm chart with custom values:

```bash
# Deploy PostgreSQL
helm install postgresql ./pkg/db/postgresql -f ./k8s/custom-values.yaml
```

The custom values configure:
- Database name: `event_repository`
- Username: `event_repo_user`
- Password: `change-me-in-production` (should be changed for production)

### 3. Deploy the Application

Apply the Kubernetes manifests to deploy the application:

```bash
# Apply ConfigMap
kubectl apply -f ./k8s/configmap.yaml

# Apply Secrets
kubectl apply -f ./k8s/secrets.yaml

# Apply Deployment
kubectl apply -f ./k8s/deployment.yaml

# Apply Service
kubectl apply -f ./k8s/service.yaml
```

### 4. Verify Deployment

Check that all components are running correctly:

```bash
# Check PostgreSQL pods
kubectl get pods -l app.kubernetes.io/name=postgresql

# Check application pods
kubectl get pods -l app=event-repository

# Check services
kubectl get svc postgresql event-repository
```

### 5. Access the Application

You can access the gRPC application by port-forwarding the service:

```bash
kubectl port-forward svc/event-repository 50051:50051
```

Then use a gRPC client to connect to localhost:50051.

## Configuration

### Environment Variables

The application uses the following environment variables:
- `APP_ENV`: Application environment (`dev`, `prod`, or `test`)
- `TEST_DB_URL`: PostgreSQL connection URL for test environment
- `DEV_DB_URL`: PostgreSQL connection URL for development environment
- `PROD_DB_URL`: PostgreSQL connection URL for production environment

These are configured in the Kubernetes ConfigMap and Secret.

### Database Connection

The database connection URLs are formatted as:
```
postgresql://username:password@hostname:port/database?sslmode=disable
```

For example:
```
postgresql://event_repo_user:change-me-in-production@postgresql.default.svc.cluster.local:5432/event_repository?sslmode=disable
```

## Production Considerations

For production deployments:

1. **Change Passwords**: Update the passwords in `custom-values.yaml` and `secrets.yaml`
2. **Resource Limits**: Adjust resource limits in `deployment.yaml` and `custom-values.yaml`
3. **Persistence**: Configure appropriate storage classes for PostgreSQL persistence
4. **Backups**: Enable and configure the backup section in `custom-values.yaml`
5. **High Availability**: Consider setting `architecture: replication` in `custom-values.yaml` for HA setup

## Automated Deployment

For convenience, you can use the provided deployment script:

```bash
./scripts/deploy.sh
```

This script automates all the steps above.

## Troubleshooting

If you encounter issues:

1. Check pod logs: `kubectl logs -l app=event-repository`
2. Check PostgreSQL logs: `kubectl logs -l app.kubernetes.io/name=postgresql`
3. Verify database connectivity: `kubectl exec -it <pod-name> -- psql -U event_repo_user -d event_repository -h postgresql`
