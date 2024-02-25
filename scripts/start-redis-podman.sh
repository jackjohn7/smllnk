#!/bin/bash
# Use this script to start a podman container for a local development database

REDIS_CONTAINER_NAME="smllnk-redis-temp"

if ! [ -x "$(command -v podman)" ]; then
  echo "Podman is not installed. Please install podman and try again.\nInstall podman here: https://podman.io/docs/installation"
  exit 1
fi

if [ "$(podman ps -q -f name=$REDIS_CONTAINER_NAME)" ]; then
  podman start $REDIS_CONTAINER_NAME
  echo "Database container started"
  exit 0
fi

# important env variables from .env 
set -a 
source .env

REDIS_PORT=$(echo $REDIS_URL | awk -F':' '{print $2}')

podman run --rm -d --name $REDIS_CONTAINER_NAME -p $REDIS_PORT:6379 redis --requirepass $REDIS_PASSWORD
