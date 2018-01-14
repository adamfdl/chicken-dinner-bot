package controller

import (
	"chicken-dinner-bot/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type UpdateScoreController struct{}

func (*UpdateScoreController) UpdateScore(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, constants.COMMAND_WIN) {
		// TODO
	}
}
