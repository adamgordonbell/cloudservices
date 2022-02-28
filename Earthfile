VERSION 0.6
FROM golang:1.17-alpine3.13

test-all:
    BUILD ./activity-log+test
    BUILD ./activity-client+test
    BUILD ./grpc-proxy+test

up:
    LOCALLY
    WITH DOCKER \
        --load=./activity-log+docker \
        --load=./grpc-proxy+docker 
        RUN docker rm -f  activity-log || false && \
            docker rm -f  grpc-proxy || false && \
            docker run -d --name activity-log -p 8080:8080 agbell/cloudservices/activity-log && \
            docker run -d --name grpc-proxy -p 8081:8081 agbell/cloudservices/grpc-proxy 
    END

down:
    LOCALLY
    RUN docker stop activity-log 
    RUN docker stop grpc-proxy

