#!/usr/bin/env bash

# This script builds the siapool inside a docker,
#  it then creates a docker from scratch containing just the built binary from the previous step
#  and then pushes the resulting image to hub.docker.com
set -e

docker build -t siapoolbuilder .
docker run --rm -v "$PWD":/go/src/github.com/robvanmieghem/siapool --entrypoint go siapoolbuilder build -ldflags '-s' -v -o dist/siapool
docker build -t robvanmieghem/siapool:latest -f DockerfileMinimal .
docker push robvanmieghem/siapool:latest
