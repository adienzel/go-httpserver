#!/bin/bash
# run_containers.sh
IMAGE_NAME="fasthttp-server"

if [ -z "$2" ]; then
    echo "Usage: $0 <number_of_containers> <sart port>"
    exit 1
fi

count=$1
port=8082

for i in $(seq 1 $count); do
    docker run -d -e APP_PORT=$port --name fasthttp_$i $IMAGE_NAME
    echo "Started fasthttp_$i on port $port"
    port=$((port + 1))  # Skip one port each time
done
