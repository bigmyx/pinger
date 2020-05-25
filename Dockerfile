FROM golang:1.7.3
WORKDIR /go/src/github.com/bigmyx/pinger/
ADD . $WORKDIR
RUN go get -d -v ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  

FROM alpine:latest
COPY --from=0 /go/src/github.com/bigmyx/pinger/pinger .
ENTRYPOINT ["/pinger"]