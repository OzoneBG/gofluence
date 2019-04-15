FROM golang

ARG app_env
ENV APP_ENV app_env

COPY . /go/src/github.com/ozonebg/gofluence
WORKDIR /go/src/github.com/ozonebg/gofluence

RUN go get .

RUN go build -o main .

ARG listen_port=8081
ENV PORT=$listen_port

CMD ["/go/src/github.com/ozonebg/gofluence/main"]
