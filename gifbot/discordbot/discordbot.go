package discordbot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/user/gifbot/app/giphy"
)

var botID string

func Start(token string) {
	dg, err := discordgo.New("", "", "Bot "+token)
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
	botID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	m.Content = strings.ToLower(m.Content)
	// Ignore all messages created by the bot itself
	if m.Author.ID == botID {
		return
	}

	gif, err := giphy.GetGif(m.Content)
	if err == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, gif)
	}

}
