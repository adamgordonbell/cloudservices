#!/usr/bin/env sh
set -e
# set -x

# echo "=== Test Reflection API ==="
grpcurl -plaintext localhost:8080 describe

echo "=== Insert Test Data ==="

grpcurl -plaintext -d  '{ "description": "christmas eve bike class" }' localhost:8080 api.v1.Activity_Log/Insert

grpcurl -plaintext -d  '{ "description": "cross country skiing is horrible and cold" }' localhost:8080 api.v1.Activity_Log/Insert

grpcurl -plaintext -d  '{ "description": "sledding with nephew" }' localhost:8080 api.v1.Activity_Log/Insert

echo "=== Test Retrieve Descriptions ==="

grpcurl -plaintext -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve

grpcurl -plaintext -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'christmas eve bike class'
grpcurl -plaintext -d '{ "id": 2 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'cross country skiing'
grpcurl -plaintext -d '{ "id": 3 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'sledding'

echo "=== Test List ==="

grpcurl -plaintext localhost:8080 api.v1.Activity_Log/List | jq '.activities | length' |  grep -q '3'
grpcurl -plaintext -d '{ "offset": 3 }' localhost:8080 api.v1.Activity_Log/List | jq '.activities | length'|  grep -q '0'

echo "Success"