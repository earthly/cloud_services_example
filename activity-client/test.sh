#!/usr/bin/env sh
set -e

echo "start up"
while ! nc -z localhost 8080; do   
  sleep 1 # wait for 1 second before check again
  echo "Waiting for service to start..."
done

echo "=== Add Records ==="
./activity-client -add "overhead press: 70lbs"
./activity-client -add "20 minute walk"

echo "=== Retrieve Records ==="
./activity-client -get 1 | grep "overhead press"
./activity-client -get 2 | grep "20 minute walk"

echo "=== List Records ==="
./activity-client -list
./activity-client -list  | grep "overhead press"
./activity-client -list  | grep "20 minute walk"
