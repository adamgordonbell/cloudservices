#!/usr/bin/env sh
set -e
# set -x

echo "=== Test Reflection API ==="
grpcurl -insecure localhost:8080 describe

echo "=== Test Cert Authority ==="
grpcurl -cacert=./certs/ca.pem localhost:8080 describe

echo "=== Insert Test Data ==="

grpcurl -insecure -d  '{ "description": "christmas eve bike class" }' localhost:8080 api.v1.Activity_Log/Insert

grpcurl -insecure -d  '{ "description": "cross country skiing is horrible and cold" }' localhost:8080 api.v1.Activity_Log/Insert

grpcurl -insecure -d  '{ "description": "sledding with nephew" }' localhost:8080 api.v1.Activity_Log/Insert

echo "=== Test Retrieve Descriptions ==="

grpcurl -insecure -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve

grpcurl -insecure -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'christmas eve bike class'
grpcurl -insecure -d '{ "id": 2 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'cross country skiing'
grpcurl -insecure -d '{ "id": 3 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'sledding'

echo "=== Test List ==="

grpcurl -insecure localhost:8080 api.v1.Activity_Log/List | jq '.activities | length' |  grep -q '3'
grpcurl -insecure -d '{ "offset": 3 }' localhost:8080 api.v1.Activity_Log/List | jq '.activities | length'|  grep -q '0'

echo "=== Test Rest API ==="

echo "=== Insert Test Data ==="
curl -X POST -s localhost:8080/api.v1.Activity_Log/Insert -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

echo "=== Test Retrieve Descriptions ==="
curl -X POST -s localhost:8080/api.v1.Activity_Log/Retrieve -d \
'{ "id": 1 }' | grep -q 'christmas eve bike class'

echo "=== Test List ==="
curl -X POST -s localhost:8080/api.v1.Activity_Log/List -d \
'{ "offset": 0 }' 

echo "Success"