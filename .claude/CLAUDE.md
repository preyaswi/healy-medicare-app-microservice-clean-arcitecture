# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Healy Medicare App** is a microservices-based healthcare application built with specific clean architecture principles.

**Backend Services (Go):**
- **admin-service**: Administration and management.
- **chat-service**: Real-time chat functionality (Kafka-based).
- **doctor-service**: Doctor management and operations.
- **patient-service**: Patient management and operations.
- **healy-apigateway**: API Gateway routing requests to services.

**Infrastructure:**
- **Kubernetes (K8s)**: Container orchestration.
- **Docker**: Containerization.
- **Kafka/Zookeeper**: Event streaming and coordination.
- **Skaffold**: Continuous development for Kubernetes.

## Tech Stack

### Backend
- **Language**: Go 1.21+
- **Frameworks/Libs**: Gin (likely), GORM (PostgreSQL), Viper (Config), Razorpay, Google OAuth2.
- **Database**: PostgreSQL.
- **Messaging**: Kafka.
- **Authentication**: JWT, Google OAuth.

### Infrastructure
- Docker & Docker Compose
- Kubernetes (Ingress Nginx, Cert Manager)
- Skaffold
- Helm

## Common Commands

### Building & Running
```bash
# Run all services locally via Docker Compose
docker-compose up --build

# Run a specific service (e.g., admin-service)
cd admin-service
go run main.go

# Development with Skaffold (K8s)
skaffold dev
```

### Testing
```bash
# Run tests for all services (from root)
go test ./...

# Run tests for a specific service
cd admin-service
go test ./...
```

### Dependency Management
```bash
# Update dependencies in a service
cd <service-directory>
go mod tidy
```

### Kubernetes & Helm
```bash
# Apply generic K8s manifests
kubectl apply -f k8s/

# Install Ingress Nginx
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx --create-namespace

# Check Pods
kubectl get pods --all-namespaces
```

## Architecture

### Microservices Structure
The project follows a modular microservices architecture defined in `go.work`. Each service (e.g., `admin-service`, `patient-service`) is a self-contained Go module.

### Directory Structure
- `admin-service/`: Admin domain logic.
- `chat-service/`: Chat domain logic.
- `doctor-service/`: Doctor domain logic.
- `patient-service/`: Patient domain logic.
- `healy-apigateway/`: API Gateway.
- `k8s/`: Kubernetes manifests (deployments, services, ingress).
- `docker-compose.yml`: Local development orchestration.

### Key Configuration
- **Port mapping**: Defined in `docker-compose.yml` and K8s services.
- **Environment Variables**: Typically handled via `.env` files or K8s ConfigMaps/Secrets.
