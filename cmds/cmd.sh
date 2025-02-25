# These are commands to build and run docker containers

# Redis
docker run -d --name redis_container -p 6379:6379 redis:alpine

# Build container and publish port for nats image
docker run -d --name nats -p 4222:4222 nats:latest -js