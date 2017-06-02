package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/austin1237/gifbot/giphy"
	"github.com/bwmarrin/discordgo"
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
	Email    string
	Password string
	Token    string
	BotID    string
)

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()

	if Email == "" && Password == "" && Token == "" {
		fmt.Println("Email and Password or Token was not provided.")
		os.Exit(1)
	}

	if Token == "" && Password == "" && Email != "" {
		fmt.Println("Password was not provided.")
		os.Exit(1)
	}

	if Token == "" && Email == "" {
		fmt.Println("Email was not provided.")
		os.Exit(1)
	}

}

func addBotIfNeeded(token string) string {
	if len(token) < 3 {
		return ""
	}
	substring := token[0:2]
	substring = strings.ToLower(token)
	if substring != "bot" {
		token = "Bot " + token
	}
	return token
}

func main() {
	// Create a new Discord session using the provided login information.
	Token = addBotIfNeeded(Token)
	dg, err := discordgo.New(Email, Password, Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
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
