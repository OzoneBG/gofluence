FROM golang

ARG app_env
ENV APP_ENV app_env

COPY . /go/src/github.com/ozonebg/gofluence
WORKDIR /go/src/github.com/ozonebg/gofluence

RUN go get .