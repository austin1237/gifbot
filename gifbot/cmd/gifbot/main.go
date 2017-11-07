package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/user/gifbot/app/discordbot"
	"github.com/user/gifbot/app/router"
)

// Variables used for environment variables

var (
	Token string
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
	httpRouter := router.SetUp()
	discordbot.Start(Token)
	http.ListenAndServe(":8080", httpRouter)
}
