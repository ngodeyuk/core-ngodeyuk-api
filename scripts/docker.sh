#!/bin/bash

DOCKER_COMPOSE_FILE=./docker-compose.yml

start_containers() {
  echo "Starting containers..."
  docker compose -f $DOCKER_COMPOSE_FILE up -d
}

stop_containers() {
  echo "Stopping containers..."
  docker compose -f $DOCKER_COMPOSE_FILE down
}

if [ "$1" == "up" ]; then
  start_containers
elif [ "$1" == "down" ]; then
  stop_containers
else
  echo "Usage: $0 {up|down}"
  exit 1
fi
