#!/usr/bin/env sh
set -e
set -v
set -x

echo "start up"
while ! nc -z localhost 8080; do   
  sleep 1 # wait for 1 second before check again
  echo "Waiting for grpc-proxy to start..."
done

./grpc-proxy &

while ! nc -z localhost 8081; do   
  sleep 1 # wait for 1 second before check again
  echo "Waiting for grpc-proxy to start..."
done

echo "=== Insert Test Data ==="
curl -k -X POST -s https://localhost:8081/api.v1.Activity_Log/Insert -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

echo "=== Test Retrieve Descriptions ==="
curl -X POST -s localhost:8081/api.v1.Activity_Log/Retrieve -d \
'{ "id": 1 }' | grep -q 'christmas eve bike class'

echo "=== Test List ==="
curl -X POST -s localhost:8081/api.v1.Activity_Log/List -d \
'{ "offset": 0 }' 

echo "Success"
