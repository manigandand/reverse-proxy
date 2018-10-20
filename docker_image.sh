#! /bin/bash

DOCKER_IMAGE=manigandanjeff/recipe_proxy_server

if [[ "$(sudo docker images -q $DOCKER_IMAGE 2> /dev/null)" == "" ]]; then
    echo "$DOCKER_IMAGE not found. installing it..."
    sudo docker pull manigandanjeff/recipe_proxy_server    
fi

echo "==> Running docker image  ..."
sudo docker run -i -t -p 8080:8080 $DOCKER_IMAGE