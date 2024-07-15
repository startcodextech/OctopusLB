#!/bin/bash

# Start systemd in the background
/lib/systemd/systemd --system &

# Wait for systemd to be ready
while ! systemctl is-system-running --quiet; do
    echo "Waiting for systemd to be ready..."
    sleep 1
done

# Navigate to the app directory
cd /app/src

# Start air
air -c .air.toml
