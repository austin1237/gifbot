FROM golang:latest


ADD . /go/src/github.com/austin1237/gifBot

Run go get github.com/austin1237/gifBot/...

RUN go install github.com/austin1237/gifBot/cmd/gifBot
