#!/bin/bash
# clean_containers.sh

containers=$(docker ps -a --filter "name=fasthttp_" --format "{{.ID}}")

if [ -z "$containers" ]; then
    echo "No fasthttp containers to remove."
    exit 0
fi

echo "Stopping and removing fasthttp containers..."
docker rm -f $containers
