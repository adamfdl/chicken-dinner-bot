package controller

import (
	"fmt"
	"strings"

	"chicken-dinner-bot/constants"
	"chicken-dinner-bot/database/redis"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type AddPlayerController struct{}

func (*AddPlayerController) AddPlayer(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, constants.COMMAND_ADD_NEW_PLAYER) {
		logrus.WithFields(logrus.Fields{"user_says": m.Content}).Info("Message from discord client")

		message := strings.Split(m.Content, " ")
		if len(message) < 1 {
			s.ChannelMessageSend(m.ChannelID, "I need your pubg nickname bitch")
			return
		}

		// Add new player
		if err := redis.GetPUBGLeaderboardOperator().AddNewPlayer(m.Author.ID, message[1]); err != nil {
			s.ChannelMessageSend(m.ChannelID, "It's not you. It's me. Server problems")
			return
		}

		bot_message := fmt.Sprintf("Thank you <@%s> your request has been received! I will notify my master to approve", m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, bot_message)
	}
}
