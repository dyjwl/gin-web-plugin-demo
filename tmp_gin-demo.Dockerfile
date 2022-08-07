FROM golang:alpine as builder

RUN mkdir -p /app
WORKDIR /app
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go mod vendor \
    && make build

FROM alpine:latest 

RUN mkdir -p /app  \
    && apk add --no-cache tzdata
WORKDIR /app

COPY --from=0 /app/output/darwin/amd64/gin-demo ./
COPY --from=0 /app/cobra.yaml ./
EXPOSE 8080
ENTRYPOINT ./gin-demo --config cobra.yaml
