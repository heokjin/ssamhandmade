FROM golang:1.14.0-buster

ENV GO111MODULE on

WORKDIR $GOPATH/src/github.com/heokjin/ssamhandmade
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ssamhandmade

ENTRYPOINT ./ssamhandmade

EXPOSE 80
