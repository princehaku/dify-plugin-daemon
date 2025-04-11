# Variables
DOCKER_REGISTRY=registry.cn-hangzhou.aliyuncs.com/haku-images
PLUGIN_IMAGE=$(DOCKER_REGISTRY)/dify-plugin-daemon
VERSION=0.0.7-local

# Build Docker images
build-plugin:
	@echo "Building web Docker image: $(PLUGIN_IMAGE):$(VERSION)..."
	docker buildx build --output=type=image --platform=linux/amd64,linux/arm64 -t $(PLUGIN_IMAGE):$(VERSION) -f ./docker/local.dockerfile ./
	@echo "Web Docker image built successfully: $(PLUGIN_IMAGE):$(VERSION)"

# Push Docker images
push-plugin:
	@echo "Pushing web Docker image: $(PLUGIN_IMAGE):$(VERSION)..."
	docker buildx build --push --platform=linux/amd64,linux/arm64 -t $(PLUGIN_IMAGE):$(VERSION) -f ./docker/local.dockerfile ./
	@echo "Web Docker image pushed successfully: $(PLUGIN_IMAGE):$(VERSION)"

# Build all images
build-all: build-plugin

# Push all images
push-all: push-plugin

build-push-api: build-api push-api
build-push-web: build-web push-web

# Build and push all images
build-push-all: build-all push-all
	@echo "All Docker images have been built and pushed."

# Phony targets
.PHONY: build-plugin
