FROM golang:1.12

WORKDIR /go/src/github.com/drish/tr

ENV GO111MODULE on

COPY . .

RUN go mod download
# RUN go build -v -o /go/src/github.com/drish/tr/tr ./cmd/tr/