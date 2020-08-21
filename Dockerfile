FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /go/src/github.com/ibronit/challenge-tracker-slack-slash-command/

COPY . .

RUN go mod download

RUN go get github.com/go-delve/delve/cmd/dlv \
    && go get github.com/githubnemo/CompileDaemon