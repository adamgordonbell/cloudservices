#!/usr/bin/env sh
set -e
# set -x

echo "start up"
./grpc-proxy &

echo "=== Insert Test Data ==="
curl -X POST -s localhost:8081/api.v1.ActivityLogService/Insert -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

echo "=== Test Retrieve Descriptions ==="
curl -X POST -s localhost:8081/api.v1.ActivityLogService/Retrieve -d \
'{ "id": 1 }' | grep -q 'christmas eve bike class'

echo "=== Test List ==="
curl -X POST -s localhost:8081/api.v1.ActivityLogService/List -d \
'{ "offset": 0 }' 

echo "Success"