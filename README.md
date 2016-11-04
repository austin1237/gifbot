gifbot
====
[![Build Status](https://travis-ci.org/austin1237/gifbot.svg?branch=master)](https://travis-ci.org/austin1237/gifbot)

![Imgur](https://media.giphy.com/media/uLECAddeoL93q/giphy.gif)<br><br>
A bot that posts gifs in discord chat

### Setup
This assumes your in the project's directory

Download the needed dependencies
```sh
$ go get ./...
```
Build the binary

```sh
go install ./cmd/gifBot
```


### Start
The below example shows how to start the bot using the binary and the bot's token

```sh
$GOPATH/bin/gifBot -t="Bot YOUR_BOT_TOKEN"
```

### Usage
The bot will send a gif whenever someone types in "gif me" followed by what they want a gif of.


## Docker Setup##

Pull the image from the [Docker Hub](https://hub.docker.com/r/austin1237/gifbot/):
```sh
$ docker pull austin1237/gifbot
```

Run the image
```sh
$ docker run -d --name gifbot austin1237/gifbot /go/bin/gifBot -t="Bot YOUR_BOT_TOKEN"
```

### Props
All the contributors over at https://github.com/bwmarrin/discordgo for making an awesome package.<br><br>
https://github.com/Giphy/GiphyAPI for having a public api key<br><br>
