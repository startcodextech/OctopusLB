#!/bin/bash

CGO_ENABLED=1 GOOS=linux go build -gcflags "all=-N -l" -o ./tmp/main cmd/main.go
chown root:root ./tmp/main
chmod u+s ./tmp/main
dlv exec ./tmp/main --headless --listen=:2345 --accept-multiclient --api-version=2 --log --accept-multiclient