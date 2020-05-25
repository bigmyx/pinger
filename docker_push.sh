#!/bin/bash
IMAGE=bigmyx/pinger
docker build -t $IMAGE
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push $IMAGE