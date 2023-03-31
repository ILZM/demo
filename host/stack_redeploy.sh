#!/bin/bash

SHORT=b:
LONG=build:
OPTS=$(getopt -a --options $SHORT --longoptions $LONG -- "$@")

eval set -- "$OPTS"

while :
do
  case "$1" in
    -b | --build )
      BUILD="$2"
      shift 2
      ;;
    *)
      echo "Unexpected option: $1"
      shift;
      break
      ;;
  esac
done

NETWORK="stack"
DOCKER_COMPOSE_LIST_FILE="docker_compose_list.txt"

trap "exit" INT

echo "Creating docker network $NETWORK"
docker network inspect $NETWORK >/dev/null 2>&1 || \
    docker network create --driver bridge $NETWORK

echo "Creating services from $DOCKER_COMPOSE_LIST_FILE"
while read docker_compose_file; do
    docker-compose -f $docker_compose_file down

    params=()
    if [[ -z $BUILD ]]; then
        params+=(--build)
    fi
    docker-compose -f $docker_compose_file up -d "${params[@]}"
done < <(sed -e 's/[[:space:]]*#.*// ; /^[[:space:]]*$/d' "$DOCKER_COMPOSE_LIST_FILE")
