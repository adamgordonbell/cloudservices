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
        RUN docker compose up -d
    END

down:
    LOCALLY
    RUN docker compose down

