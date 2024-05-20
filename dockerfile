FROM golang:1.21.1

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app
COPY . ./
RUN go mod download
RUN go build -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
CMD ["/go/bin/app"]

