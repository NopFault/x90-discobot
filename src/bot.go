package x90

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type BotEngine struct {
	bot     *discordgo.User
	session *discordgo.Session
	db      *sql.DB
}

type Remind struct {
	id        int
	username  string
	recipient string
	date      string
	message   string
}

type Reminders []Remind

var botEngine BotEngine

func BotStart() {
	var err error
	botEngine.session, err = discordgo.New("Bot " + config_params.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botEngine.bot, err = botEngine.session.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botEngine.session.AddHandler(MessageHandler)

	err = botEngine.session.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go HeartBeat(botEngine.session)

	fmt.Println("Bot is running !")
}

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02 15:04", stringDate)
	return err == nil
}

func IsDateValid(stringDate string) bool {
	t, _ := time.ParseInLocation("2006-01-02 15:04", stringDate, time.Local)
	return t.Unix() > time.Now().Unix()
}

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == botEngine.bot.ID {
		return
	}

	if strings.HasPrefix(message.Content, "remind") {
		data := strings.Split(message.Content, " ")
		if len(data) > 3 {
			date := data[1] + " " + data[2]
			reminder := strings.Join(data[3:], " ")

			if len(reminder) > 3 && IsDateValue(date) && IsDateValid(date) {
				// database.go
				AddReminder(date, reminder, message)
				// messages
				RemindCreated(session, message)

			} else if !IsDateValue(date) || !IsDateValid(date) {
				// message.go
				WrongDateFormat(session, message)
			} else {
				// messages.go
				WrongShortMessage(session, message)
			}
		}

	}
}
