#!/bin/bash

NETWORK="stack"
DOCKER_COMPOSE_LIST_FILE="docker_compose_list.txt"

trap "exit" INT

echo "Creating docker network $NETWORK"
docker network inspect $NETWORK >/dev/null 2>&1 || \
    docker network create --driver bridge $NETWORK

echo "Creating services from $DOCKER_COMPOSE_LIST_FILE"
while read docker_compose_file; do
    docker-compose -f $docker_compose_file down
    docker-compose -f $docker_compose_file up -d
done < $DOCKER_COMPOSE_LIST_FILE
