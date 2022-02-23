echo "start up"
./grpc-proxy &


echo "Testing"
curl  -X POST -s localhost:8081/api.v1.Activity_Log/Insert -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'