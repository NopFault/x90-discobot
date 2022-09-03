package x90

import (
	"github.com/bwmarrin/discordgo"
)

func WrongDateFormat(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "@"+message.Author.Username+" Wrong date format! it should be e.x.: 2010-01-01 22:00 and dont forget i cant send messages in the past")
}

func WrongShortMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "@"+message.Author.Username+" Maybe message is too short!")
}

func RemindCreated(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "@"+message.Author.Username+" I will remind you!")
}

func RemindUser(session *discordgo.Session, remind Remind) {
	message, _ := session.UserChannelCreate(remind.recipient)
	session.ChannelMessageSend(message.ID, "Hello, i want to remind you: "+remind.message)
}

func sukashenka(session *discordgo.Session, message *discordgo.MessageCreate, str string) {
	session.ChannelMessageSend(message.ChannelID, "@"+message.Author.Username+" "+str)
}
