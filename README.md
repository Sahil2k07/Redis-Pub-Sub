# Redis-Pub-Sub

This project explores how Redis Pub/Sub can be used to implement a real-time chat system using WebSockets. It mimics a real-world chat room where messages are published to a Redis channel and received by subscribed clients.

⚠️ Note: This is a proof-of-concept and not a production-ready implementation. However, it effectively demonstrates how Redis can be leveraged as a Pub/Sub mechanism for real-time communication.

## Features

- WebSocket-based real-time messaging

- Redis Pub/Sub for message broadcasting

- Multiple users can join the same chat room

- Echo framework for handling WebSocket connections

## Try it for yourself

### Start with a Redis instance

1. Use a redis instance, better to use a docker container, download `docker` first.

   [Docker Download Page](https://docs.docker.com/engine/install/)

2. Pull the redis image from dockerhub

   ```bash
    docker pull redis
   ```

3. Start the redis container

   ```bash
    docker run --name redis-container -p 6379:6379 -d redis
   ```

   for each consecutive use

   ```bash
    docker start redis-container
   ```

4. If you want to use the Redis-CLI

   ```bash
    docker exec -it redis-container redis-cli
   ```

### Set-Up the project

1. Clone the project

   ```bash
    git clone https://github.com/Sahil2k07/Redis-Pub-Sub.git
   ```

2. Change Directory

   ```bash
    cd Redis-Pub-Sub
   ```

3. Get all the packages

   ```bash
    go mod download
   ```

4. Get the vendor ( optional )

   ```bash
    go mod vendor
   ```

5. Start the project

   ```bash
    go run .
   ```
