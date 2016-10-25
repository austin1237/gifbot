package main

import (
	"strings"
	"flag"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Email    string
	Password string
	Token    string
	BotID    string
)

// Model of the json response from the giphy /search endpoint
type GiphyResp struct{
	Type string `json:"type"`
	Data []struct{
		Images struct{
			Original struct{
				Url string `json:"url"`
				} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided login information.
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

	if strings.Index(m.Content, "gif me") >= 0 {
		gifResp := make(chan string)
		splitArr := strings.Split(m.Content, "gif me")
		keyword := splitArr[1]
		go gifMe(keyword, gifResp)
		gifs := <- gifResp
		_, _ = s.ChannelMessageSend(m.ChannelID, gifs)

	}
}

func gifMe(keyword string, done chan string) {
	var test GiphyResp
	keyword = url.QueryEscape(keyword)
	resp, err := http.Get("http://api.giphy.com/v1/gifs/search?q=" + keyword + "&api_key=dc6zaTOxFJmzC")
	if err != nil {
		done <- "An error occured trying to contact giphy"
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &test)
	if err != nil {
		done <- "An error occured trying to contact giphy"
	}
	done <- test.Data[0].Images.Original.Url
}
