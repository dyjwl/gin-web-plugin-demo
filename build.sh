#!/usr/bin/env bash

docker build --no-cache -f ./Dockerfile -t zhaosir1993/gin-demo:latest .
docker push zhaosir1993/gin-demo:latest