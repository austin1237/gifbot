FROM golang:1.8

COPY . /go/src/github.com/user/gifbot/app
WORKDIR /go/src/github.com/user/gifbot/app


RUN go get -d -v ./cmd/gifbot/...

RUN go build ./cmd/gifbot/...

CMD if [ ${APP_ENV} = production ]; \
	then \
	gifbot; \
	else \
	go get github.com/pilu/fresh && \
	cd cmd/gifbot && fresh; \
	fi

EXPOSE 8080

