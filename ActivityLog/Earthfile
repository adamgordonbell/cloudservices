VERSION 0.6
FROM golang:1.15-alpine3.13
WORKDIR /activityserver

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY . .
    RUN go build -o build/activityserver cmd/server/main.go
    SAVE ARTIFACT build/activityserver /activityserver

docker:
    COPY +build/activityserver .
    ENTRYPOINT ["/activityserver/activityserver"]
    EXPOSE 8080
    SAVE IMAGE --push agbell/cloudservices/activityserver

test-deps:
    FROM earthly/dind
    RUN apk add curl ripgrep

test:
    FROM +test-deps
    COPY test.sh .
    WITH DOCKER --load agbell/cloudservices/activityserver=+docker
        RUN  docker run -d -p 8080:8080 agbell/cloudservices/activityserver && \
                ./test.sh
    END