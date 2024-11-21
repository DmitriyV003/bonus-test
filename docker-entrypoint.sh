#!/bin/bash
set -e

DELAY=5

echo "Delaying startup for $DELAY seconds..."
sleep "$DELAY"

echo "Running database migrations..."
soda migrate

echo "Starting the application..."
exec ./server