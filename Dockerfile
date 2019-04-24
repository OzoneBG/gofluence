# Use golang
FROM golang:1.12-alpine

# Copy the source files
COPY . /go/src/github.com/ozonebg/gofluence

# Use the current dir as workdir
WORKDIR /go/src/github.com/ozonebg/gofluence

# BUild the project
RUN go build -o main .

# Run the project
# CMD ["/go/src/github.com/ozonebg/gofluence/main"]
