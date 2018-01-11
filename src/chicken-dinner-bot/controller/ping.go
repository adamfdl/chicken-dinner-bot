package controller

import (
	"chicken-dinner-bot/constants"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type PingController struct{}

func (this *PingController) Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == constants.COMMAND_PING {
		logrus.WithFields(logrus.Fields{"user_says": constants.COMMAND_PING}).Info("Message from discord client")

		reply := constants.CHICKEN_EMOJI + " pong " + constants.CHICKEN_EMOJI

		s.ChannelMessageSend(m.ChannelID, reply)
		logrus.WithFields(logrus.Fields{"bot_says": reply}).Info("Bot response")
	}
}
