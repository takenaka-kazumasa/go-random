FROM golang:1.15

RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get github.com/google/wire/cmd/wire
RUN go get github.com/golang/mock/gomock
aRUN GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3
