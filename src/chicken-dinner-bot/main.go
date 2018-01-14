package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"chicken-dinner-bot/controller"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	pingController := &controller.PingController{}
	leaderboardController := &controller.LeaderBoardController{}
	addNewPlayerController := &controller.AddPlayerController{}
	// TODO InitiatePlayer
	// TODO UpdatePlayerScore

	dg.AddHandler(addNewPlayerController.AddPlayer)
	dg.AddHandler(pingController.Ping)
	dg.AddHandler(leaderboardController.LeaderBoard)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
