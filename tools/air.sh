#!/bin/bash

PORT=2345

free_port() {
  PID=$(lsof -ti tcp:$PORT)
  if [ -n "$PID" ]; then
    kill -9 $PID
  fi
}

CGO_ENABLED=1 GOOS=linux go build -gcflags "all=-N -l" -o ./tmp/main cmd/main.go
chown root:root ./tmp/main
chmod u+s ./tmp/main
dlv exec ./tmp/main --headless --listen=:2345 --accept-multiclient --api-version=2 --log --accept-multiclient