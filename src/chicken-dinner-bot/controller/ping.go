package controller

import (
	"chicken-dinner-bot/constants"

	"github.com/bwmarrin/discordgo"
)

type PingController struct{}

func (this *PingController) Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == constants.BOT_PING {
		reply := constants.CHICKEN_EMOJI + " pong " + constants.CHICKEN_EMOJI
		s.ChannelMessageSend(m.ChannelID, reply)
	}
}
