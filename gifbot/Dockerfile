FROM golang:1.9

# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/user/gifbot/app
COPY . /go/src/github.com/user/gifbot/app
WORKDIR /go/src/github.com/user/gifbot/app

# Install dependencies
RUN dep ensure 

RUN go install ./cmd/gifbot/...

CMD if [ ${APP_ENV} != dev ]; \
	then \
	gifbot; \
	else \
    dep ensure && \
	go get github.com/pilu/fresh && \
	cd cmd/gifbot && fresh; \
	fi

EXPOSE 8080

