#!/bin/bash

if ! [ -x "$(command -v docker-compose)" ]; then
  echo "Error: docker-compose is not installed"
  exit 1
fi

docker-compose stop $@
docker-compose rm -f $@
docker-compose build $@
docker-compose create $@
docker-compose start $@
docker rmi `docker images | grep '<none>' | awk '{ print $3 }'`
