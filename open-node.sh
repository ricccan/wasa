#!/usr/bin/env sh

sudo docker run -it --rm -v "$(pwd):/src" -u "$(id -u):$(id -g)" --network host --workdir /src/webui node:20 /bin/bash
