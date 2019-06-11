FROM golang:1.12

LABEL maintainer="Nicholas Burrell <zero_frost@protonmail.com>"

WORKDIR $GOPATH/src/github/zero-frost/auth-service

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 7777

CMD ["auth-service"]
