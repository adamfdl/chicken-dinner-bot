package controller

import (
	"chicken-dinner-bot/constants"
	"chicken-dinner-bot/database/redis"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type LeaderBoardController struct{}

func (*LeaderBoardController) LeaderBoard(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == constants.COMMAND_LEADERBOARD {
		logrus.WithFields(logrus.Fields{"user_says": constants.COMMAND_LEADERBOARD}).Info("Message from discord client")

		var reply string
		if result, err := redis.GetPUBGLeaderboardOperator().RetrieveLeaderBoard(); err != nil {
			reply = ":shit: Server is currently down :shit:"
			s.ChannelMessageSend(m.ChannelID, reply)
		} else {
			if len(result) <= 0 {
				reply = ":shit: There's no player to retrieve in the database :shit:"
				s.ChannelMessageSend(m.ChannelID, reply)
				return
			}

			for i := 0; i < len(result); i++ {
				splittedResult := strings.Split(result[i].Member.(string), ":")
				discordID := fmt.Sprintf("<@%s>", splittedResult[0])

				temp := discordID + "\n"
				for j := 0; j < int(result[i].Score); j++ {
					temp += constants.CHICKEN_EMOJI + " "
				}
				temp += "\n\n"
				reply += temp
			}
			s.ChannelMessageSend(m.ChannelID, reply)
		}
	}
}
