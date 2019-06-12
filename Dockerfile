FROM golang:1.12

LABEL maintainer="Nicholas Burrell <zero_frost@protonmail.com>"

WORKDIR $GOPATH/src/github/zero-frost/auth-service

COPY . .

RUN GO111MODULE=on go get -d -v ./...

RUN GO111MODULE=on go install -v ./...

EXPOSE 5300

CMD ["auth-service"]
