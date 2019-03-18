FROM golang:1.10.8-alpine3.8 as builder
COPY toggleScanningData.go /go
RUN apk add --no-cache --virtual .build-deps git
RUN go get golang.org/x/crypto/ssh/terminal 
RUN cd /go && \
    go build toggleScanningData.go

FROM alpine:3.8
COPY --from=builder /go/toggleScanningData /usr/bin/toggleScanningData
ENTRYPOINT ["toggleScanningData"]
