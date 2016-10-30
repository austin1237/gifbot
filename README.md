# gifbot
![Imgur](https://media.giphy.com/media/uLECAddeoL93q/giphy.gif)<br><br>
A bot that posts gifs in discord chat

### Setup
Download the needed dependcies

```sh
$ go get
```

### Start
The below example shows how to start the bot using the bot's token

```sh
$ go run gifbot.go -t="Bot YOUR_BOT_TOKEN"
```

### Usage
The bot will send a gif whenever someone types in "gif me" followed by what they want a gif of.


## Docker Setup##

Pull the image from the [Docker Hub](https://hub.docker.com/r/austin1237/gifbot/):
```sh
$ docker pull austin1237/notibot
```

Run the image
```sh
$ docker run -d --name gifbot austin1237/gifbot app -t="Bot YOUR_BOT_TOKEN"
```

### Props
All the contributors over at https://github.com/bwmarrin/discordgo for making an awesome package.<br><br>
https://github.com/Giphy/GiphyAPI for having a public api key<br><br>
https://github.com/dhedegaard/notibot for having clear documentation with a simple docker setup
