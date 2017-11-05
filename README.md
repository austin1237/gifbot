gifbot
====
[![CircleCI](https://circleci.com/gh/austin1237/gifbot.svg?style=svg)](https://circleci.com/gh/austin1237/gifbot)<br>
![Imgur](https://media.giphy.com/media/uLECAddeoL93q/giphy.gif)<br><br>
A bot that posts gifs in discord chat

## Dependencies
You must have the following installed/configured for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)


### Start
The below example shows how to start the bot 

```sh
export BOT_TOKEN="YOUR_DISCORD_TOKEN" && docker-compose up
```

### Usage
The bot will send a gif whenever someone types in "gif me" followed by what they want a gif of.


### Props
All the contributors over at https://github.com/bwmarrin/discordgo for making an awesome package.<br><br>
https://github.com/Giphy/GiphyAPI for having a public api key<br><br>
