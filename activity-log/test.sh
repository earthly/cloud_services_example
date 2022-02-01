#!/usr/bin/env sh
set -e
# set -x

echo "=== Test Reflection API ==="
grpc_cli ls localhost:8080 -l

echo "=== Insert Test Data ==="

curl -X POST -s localhost:8080 -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

curl -X POST -s localhost:8080 -d \
'{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

curl -X POST -s localhost:8080 -d \
'{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'

echo "=== Test Descriptions ==="

grpcurl -plaintext -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve

curl -X GET -s localhost:8080 -d '{"id": 1}' | grep -q 'christmas eve bike class'
curl -X GET -s localhost:8080 -d '{"id": 2}' | grep -q 'cross country skiing'
curl -X GET -s localhost:8080 -d '{"id": 3}' | grep -q 'sledding'

echo "=== Test List ==="

curl -X GET -s localhost:8080/list | jq length |  grep -q '3'
curl -X GET -s localhost:8080/list -d '{"offset": 3}' | jq length |  grep -q '0'

echo "Success"