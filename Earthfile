VERSION 0.6
FROM golang:1.17-alpine3.13

test-all:
    BUILD ./activity-log+test
    BUILD +ac-test

### Activity Client
ac-service:
   WORKDIR /activity-log 
   COPY activity-log .

ac-deps:
    FROM +ac-service
    WORKDIR /activity-client
    COPY activity-client/go.mod ./
    RUN go mod download

ac-build:
    FROM +ac-deps
    COPY activity-client .
    RUN go build -o build/activityclient cmd/client/main.go
    SAVE ARTIFACT build/activityclient /activityclient

ac-test-deps:
    FROM earthly/dind
    RUN apk add curl ripgrep

ac-test:
    FROM +ac-test-deps
    COPY +ac-build/activityclient ./activityclient
    COPY activity-client/test.sh .
    WITH DOCKER --load agbell/cloudservices/activityserver=./activity-log+docker
        RUN  docker run -d -p 8080:8080 agbell/cloudservices/activityserver && \
                ./test.sh
    END
