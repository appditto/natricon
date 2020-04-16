#!/bin/bash

if [ -z "$1" ]; then
    echo "Specify a tag/version to build image"
    exit 1
fi

# Build and deploy new container
docker build -t go-natricon:$1 .
docker tag go-natricon:$1 docker-registry.appditto.com/go-natricon:$1
docker push docker-registry.appditto.com/go-natricon:$1