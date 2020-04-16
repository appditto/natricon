#!/bin/bash

TAG=$1
if [ -z "$TAG" ]; then
    TAG="latest"
fi

# Build and deploy new container
docker build -t go-natricon:$TAG .
docker tag go-natricon:$TAG docker-registry.appditto.com/go-natricon:$TAG
docker push docker-registry.appditto.com/go-natricon:$TAG