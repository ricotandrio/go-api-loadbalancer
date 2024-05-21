FROM golang:1.21.1-alpine AS builder

RUN apk add --no-cache git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app
COPY . ./
RUN go mod download
RUN go build -o /go/bin/app

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bin/app /go/bin/app

CMD ["/go/bin/app"]
