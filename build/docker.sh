#!/bin/sh
docker build --no-cache --cpuset-cpus=1 -t fcm-service:latest .
docker-compose down && docker-compose up -d