module github.com/adamgordonbell/cloudservices/lambda-grpc

go 1.17

require (
	github.com/adamgordonbell/cloudservices/activity-log v0.0.0-20220510131242-57b6c9025807
	github.com/aws/aws-lambda-go v1.31.1
	github.com/awslabs/aws-lambda-go-api-proxy v0.13.2
	google.golang.org/grpc v1.44.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/adamgordonbell/cloudservices/activity-log => ../activity-log
