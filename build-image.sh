#!/bin/bash

TAG=$1
if [ -z "$TAG" ]; then
    TAG="latest"
fi

# Build and deploy new container
docker build -t nuxt-natricon:$TAG .
docker tag nuxt-natricon:$TAG docker-registry.appditto.com/nuxt-natricon:$TAG
docker push docker-registry.appditto.com/nuxt-natricon:$TAG