#! /bin/bash

DEPENDENCY_GLIDE=glide
if ! type "$DEPENDENCY_GLIDE" > /dev/null; then
  echo "$DEPENDENCY_GLIDE not found. installing it..."
  curl https://glide.sh/get | sh
fi

SERVER_BIN=recipe_proxy_server

# Initial setup
make deps test build-server
# make build-server

# run server
echo "==> Running server ..."
bash -c "source .env"
./$SERVER_BIN