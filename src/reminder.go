package x90

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func HeartBeat(session *discordgo.Session) {

	for range time.Tick(time.Second * 30) {

		var reminders Reminders
		reminders = SelectActiveReminders()

		if len(reminders) > 0 {

			for _, reminder := range reminders {
				RemindUser(session, reminder)
				UnsetReminder(reminder)
			}
		}
	}

}
