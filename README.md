<div align="center">
<h1>FCM API Microservices</h1>
Copyright © 2025 AnhLe. All rights reserved.
</div>

# Introduction
Main functionlities of this product are connect client to FCM and push notification to mobile devices.

There are 3 main services, including **fcm-service**, **worker-service** and **logging-monitoring-service**:

1. **fcm-service**: the heart of this product. It take on responsible for authentication and authorization, connect to **Goole FCM** and other functionalities.
2. **worker-service**: take on responsible for getting messages from message queue, handling send and store logs into **MongoDB**
3. **logging-monitoring-service**: get logs from **MongoDB**, handling reports and statistics


# Architecture
- Programing Language: Golang(version 1.22 or above)
- Protocol: HTTP(REST API)
- RDBSM: PostgreSQL
- NoSQL: MongoDB
- Cache: Redis
- Message Queue: RabbitMQ
- Design pattern: repository pattern

![Project Architecture](/assets/images/architecture.png)

# Project structure
- `common`: contains common packages, supporting repeative logices
- `server`: define http server
- `service`: define logic layer
- `repository`: define interfaces that connect and interact with databases

# Commit guidelines
- `{action}-{write briefly content (max 50 characters)}`

# References
- Golang: https://golang.google.cn

- ORM: [Bun](https://bun.uptrace.dev/)
    - Link: https://bun.uptrace.dev
    - Git: https://github.com/uptrace/bun

- Go Generics:
    - Link: https://go.dev/doc/tutorial/generics

- A few libraries:
    - Gin: web server framework
    - Jaeger: tracing service and request
    - Cobra: it helps convert golang to command lines

# Setting FCM
I'm using FCM v1 API, accessing to fcm website and downloading file **<name>.json**
1. Open terminal and run the command

    `cat service-account.json | base64`

2. Copy the outcome and put it to **FCM_CREDENTIAL_BASE64** in file **.env**
