package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/user/gifbot/app/giphy"
)

// Model of the json response from the giphy /search endpoint
type GiphyResp struct {
	Type string `json:"type"`
	Data []struct {
		Images struct {
			Original struct {
				Url string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}

// Variables used for command line parameters
var (
	Token string
	BotID string
)

func init() {

	Token = os.Getenv("BOT_TOKEN")
	if Token == "" {
		fmt.Println("Discord Token was not provided.")
		os.Exit(1)
	}
}

func main() {
	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New("", "", "Bot "+Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
		panic(err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	m.Content = strings.ToLower(m.Content)
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	gif, err := giphy.GetGif(m.Content)
	// fmt.Println("error is" + err.Error())
	if err == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, gif)
	}

}
