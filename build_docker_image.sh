#!/bin/bash
# build_image.sh
IMAGE_NAME="fasthttp-server"

echo "Building Go binary and Docker image..."
docker build -t $IMAGE_NAME .
