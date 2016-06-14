FROM ubuntu:14.04

ADD target/go-webserver.go-linux-amd64  /go-webserver.go-linux-amd64
COPY static/ /static/

CMD ["/go-webserver.go-linux-amd64"]

EXPOSE 8080
