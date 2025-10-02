#!/bin/bash
xhost +local:   # allow containers to access X11 (safe locally)

podman run --rm -it \
  --userns keep-id \
  -e DISPLAY \
  -e XDG_CONFIG_HOME=/tmp \
  -e XDG_CACHE_HOME=/tmp \
  -v /tmp/.X11-unix:/tmp/.X11-unix:ro \
  localhost/go-timer
