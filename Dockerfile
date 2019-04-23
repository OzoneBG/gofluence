# Use golang
FROM golang

COPY . /go/src/github.com/ozonebg/gofluence
WORKDIR /go/src/github.com/ozonebg/gofluence

RUN go get .

RUN go build -o main .

# Setup arguments
# ARG listen_port=8081
# ARG db_user=gofluence
# ARG db_pwd=gofluencer
# ARG db_name=gofluence
# ARG token_pwd=mylonghashedpass


# Setup env variables
# ENV PORT=$listen_port
# ENV DB_USER=$db_user
# ENV DB_PWD=$db_pwd
# ENV DB_NAME=$db_name
# ENV TOKEN_PWD=$token_pwd

# Run the project
CMD ["/go/src/github.com/ozonebg/gofluence/main"]
