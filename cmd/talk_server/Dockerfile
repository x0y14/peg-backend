FROM golang:1.18-alpine as builder
WORKDIR /go/src/service

# download dependenses
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# main
COPY cmd/talk_server/main.go ./cmd/talk_server/main.go
# protobuf
COPY gen/ ./gen
# scripts
COPY scripts/ ./scripts
# util
COPY util/ ./util
#interceptor
COPY interceptor/ ./interceptor
# database
COPY db/ ./db

RUN apk add build-base

# build as app
RUN go build -o app ./cmd/talk_server/main.go


FROM alpine:latest
WORKDIR /talk
COPY --from=builder /go/src/service/app ./
CMD ["./app"]
