FROM golang:1.22-alpine as builder

ARG VERSION=unknown

# copy project
COPY . /app

# set working directory
WORKDIR /app

# using goproxy if you have network issues
ENV GOPROXY=https://goproxy.cn,direct

# build
RUN go build \
    -ldflags "\
    -X 'github.com/langgenius/dify-plugin-daemon/internal/manifest.VersionX=${VERSION}' \
    -X 'github.com/langgenius/dify-plugin-daemon/internal/manifest.BuildTimeX=$(date -u +%Y-%m-%dT%H:%M:%S%z)'" \
    -o /app/main cmd/server/main.go

# copy entrypoint.sh
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

FROM ubuntu:24.04

COPY --from=builder /app/main /app/main
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh

WORKDIR /app

# check build args
ARG PLATFORM=local

# Install python3.12 if PLATFORM is local
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y python3.12 python3.12-venv python3.12-dev ffmpeg \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.12 1; \
    && pip install pysocks

ENV PLATFORM=$PLATFORM
ENV GIN_MODE=release

# run the server, using sh as the entrypoint to avoid process being the root process
# and using bash to recycle resources
CMD ["/bin/bash", "-c", "/app/entrypoint.sh"]
