#!/bin/sh
echo "Build fcm service"
GOOS=linux GOARCH=amd64 go build -o app.exe main.go
echo "Restart fcm service"
systemctl restart fcm-service
systemctl status fcm-service