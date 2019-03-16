FROM golang:1.10.8-alpine3.8
COPY disableScanningData.go /go
RUN apk add --no-cache --virtual .build-deps git
RUN go get golang.org/x/crypto/ssh/terminal 
RUN cd /go && \
    go build disableScanningData.go
ENTRYPOINT ["/go/disableScanningData"]
