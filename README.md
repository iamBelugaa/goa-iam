# IAM Platform - Identity and Access Management Service

**goa-iam** is a Go-based service built using the
[Goa framework](https://goa.design/) that implements authentication and
authorization functionality.

## 📋 Prerequisites

- **Go 1.24.4+**
- **Docker**
- **Kind**
- **kubectl**
- **Make**

## 🚀 Quick Start

### Local Development

```bash
# Install dependencies
make deps

# Install Goa framework
make install-goa

# Run the service locally
make run
```

The service will be available at `http://localhost:8080`

## 🐳 Docker Deployment

### Building Docker Image

```bash
# Build Docker image with default tag (latest)
make dockerize

# Build with custom tag
make dockerize DOCKER_TAG=v1.0.0

# Build with custom image name
make dockerize DOCKER_IMAGE_NAME=my-iam-service DOCKER_TAG=dev
```

### Running Docker Container

```bash
# Build and run container
make docker-run

# Run with custom configuration
make docker-run DOCKER_TAG=v1.0.0
```

The containerized service will be available at `http://localhost:8080`

### Manual Docker Commands

```bash
# Build image manually
docker build -f ./deploy/docker/Dockerfile -t iam-service:latest .

# Run container manually
docker run --rm -p 8080:8080 --name iam-service-container iam-service:latest
```

## ☸️ Kubernetes Deployment

### 1. Setup Kind Cluster

Create a local Kubernetes cluster using Kind:

```bash
# Create Kind cluster with ingress support
kind create cluster --config=./deploy/kind/kind-cluster.yaml

# Verify cluster is running
kubectl cluster-info --context kind-goa-iam-cluster
```

### 2. Install NGINX Ingress Controller

```bash
# Install NGINX Ingress for Kind
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for ingress controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

### 3. Deploy Application to Kubernetes

```bash
# Build the Docker image first (if not already built)
make dockerize

# Load the Docker image into Kind cluster
kind load docker-image iam-service:latest --name goa-iam-cluster

# For custom tags, use:
# make dockerize DOCKER_TAG=v1.0.0
# kind load docker-image iam-service:v1.0.0 --name goa-iam-cluster

# Apply all Kubernetes manifests
kubectl apply -f ./deploy/k8s

# Verify deployment
kubectl get pods,services,ingress

# Port forward for direct access (alternative to ingress)
kubectl port-forward service/goa-iam-service 8080:8080
```

**Important Note**: Kind clusters are isolated from your local Docker images.
You must load any locally built images into Kind using `kind load docker-image`
before deploying.

### 4. Configure Local DNS

To access the service via `iam.goa.com`, add the following entry to your hosts
file:

#### Windows

Edit `C:\Windows\System32\drivers\etc\hosts` as Administrator:

```
127.0.0.1 iam.goa.com
```

#### macOS/Linux

Edit `/etc/hosts` with sudo:

```bash
sudo echo "127.0.0.1 iam.goa.com" >> /etc/hosts
```

### 5. Verify Deployment

```bash
# Test the service
curl http://iam.goa.com/api/v1/users

# Or visit in browser
open http://iam.goa.com/api/v1/users
```

### Cleanup

```bash
# Delete Kubernetes resources
kubectl delete -f ./deploy/k8s

# Delete Kind cluster
kind delete cluster --name goa-iam-cluster
```

## 🔗 API Endpoints

| Method | Endpoint               | Description                       | Authentication |
| ------ | ---------------------- | --------------------------------- | -------------- |
| `POST` | `/api/v1/auth/signup`  | Register a new user               | None           |
| `POST` | `/api/v1/auth/signin`  | Login user and get JWT tokens     | None           |
| `POST` | `/api/v1/auth/signout` | Logout user and invalidate tokens | JWT Required   |

### User Service (`/api/v1/users`)

| Method | Endpoint             | Description       | Authentication |
| ------ | -------------------- | ----------------- | -------------- |
| `GET`  | `/api/v1/users`      | List all users    | None           |
| `GET`  | `/api/v1/users/{id}` | Get user by ID    | None           |
| `POST` | `/api/v1/users`      | Create a new user | None           |

### Example API Calls

#### User Registration

```bash
curl -X POST http://iam.goa.com/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword123",
    "confirmPassword": "securepassword123"
  }'
```

#### User Login

```bash
curl -X POST http://iam.goa.com/api/v1/auth/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword123"
  }'
```
