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

RUN echo "deb https://mirrors.aliyun.com/ubuntu/ noble main restricted universe multiverse\ndeb-src https://mirrors.aliyun.com/ubuntu/ noble main restricted universe multiverse\ndeb https://mirrors.aliyun.com/ubuntu/ noble-security main restricted universe multiverse\ndeb-src https://mirrors.aliyun.com/ubuntu/ noble-security main restricted universe multiverse\ndeb https://mirrors.aliyun.com/ubuntu/ noble-updates main restricted universe multiverse\ndeb-src https://mirrors.aliyun.com/ubuntu/ noble-updates main restricted universe multiverse\ndeb https://mirrors.aliyun.com/ubuntu/ noble-backports main restricted universe multiverse\ndeb-src https://mirrors.aliyun.com/ubuntu/ noble-backports main restricted universe multiverse" > /etc/apt/sources.list && apt update

# Install python3.12 if PLATFORM is local
RUN apt-get update &&  \
    apt-get install -y python3.12 python3.12-venv python3.12 python3.12-dev ffmpeg pip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.12 1;

RUN pip install pysocks --break-system-packages -i https://mirrors.aliyun.com/pypi/simple/

ENV PLATFORM=$PLATFORM
ENV GIN_MODE=release

# run the server, using sh as the entrypoint to avoid process being the root process
# and using bash to recycle resources
CMD ["/bin/bash", "-c", "/app/entrypoint.sh"]
