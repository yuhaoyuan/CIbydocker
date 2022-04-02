FROM golang:1.14

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2

RUN apt-get update
RUN apt install -y redis-server
RUN apt-get install -y mariadb-server mariadb-client

RUN nohup redis-server &

COPY mysql_init.sh /go/mysql_init.sh
RUN /go/mysql_init.sh

COPY ci_noti.sh /go/ci_noti.sh