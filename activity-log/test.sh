#!/usr/bin/env sh
set -e
# set -x

echo "=== GRPC: Test Reflection API ==="
grpcurl -insecure localhost:8080 describe

echo "=== GRPC: Test Cert Authority ==="
grpcurl -cacert=./certs/ca.pem localhost:8080 describe

echo "=== GRPC: Insert Test Data ==="

grpcurl -insecure -d  '{ "activity":{ "description": "christmas eve bike class" } }' localhost:8080 api.v1.ActivityLogService/Insert

grpcurl -insecure -d  '{ "activity":{ "description": "cross country skiing is horrible and cold" } }' localhost:8080 api.v1.ActivityLogService/Insert

grpcurl -insecure -d  '{ "activity":{ "description": "sledding with nephew" } }' localhost:8080 api.v1.ActivityLogService/Insert

echo "=== GRPC: Test Retrieve Descriptions ==="

grpcurl -insecure -d '{ "id": 1 }' localhost:8080 api.v1.ActivityLogService/Retrieve

grpcurl -insecure -d '{ "id": 1 }' localhost:8080 api.v1.ActivityLogService/Retrieve | grep -q 'christmas eve bike class'
grpcurl -insecure -d '{ "id": 2 }' localhost:8080 api.v1.ActivityLogService/Retrieve | grep -q 'cross country skiing'
grpcurl -insecure -d '{ "id": 3 }' localhost:8080 api.v1.ActivityLogService/Retrieve | grep -q 'sledding'

echo "=== GRPC: Test List ==="

grpcurl -insecure localhost:8080 api.v1.ActivityLogService/List | jq '.activities | length' |  grep -q '3'
grpcurl -insecure -d '{ "offset": 3 }' localhost:8080 api.v1.ActivityLogService/List | jq '.activities | length'|  grep -q '0'

echo "=== Test Rest API ==="

echo "=== REST: Insert Test Data ==="
curl -k -X POST -s https://localhost:8080/api.v1.ActivityLogService/Insert -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

echo "=== REST: Test Retrieve Descriptions ==="
curl -k -X POST -s https://localhost:8080/api.v1.ActivityLogService/Retrieve -d \
'{ "id": 1 }' | grep -q 'christmas eve bike class'

echo "=== REST: Test List ==="
curl -k -X POST -s https://localhost:8080/api.v1.ActivityLogService/List -d \
'{ "offset": 0 }' 

echo "Success"