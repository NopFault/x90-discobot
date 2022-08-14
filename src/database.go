package x90

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var config_params Params = GetConfig()

func DatabaseInit() {
	if _, err := os.Stat(config_params.Database); err != nil {
		CreateDatabase(config_params)
	}

	CreateTable()
}

func CreateInstance() *sql.DB {
	sqliteDatabase, _ := sql.Open("sqlite3", config_params.Database)
	return sqliteDatabase
}

func CreateDatabase(config_params Params) {

	file, err := os.Create(config_params.Database)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
}

func CreateTable() {
	db := CreateInstance()

	createReminderTableSQL := `CREATE TABLE IF NOT EXISTS reminder (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"recipient" TEXT,
		"date" DATETIME,
		"message" TEXT,
		"status" BOOLEAN DEFAULT FALSE,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	statement, err := db.Prepare(createReminderTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	defer db.Close()
}

func AddReminder(date string, reminder_str string, message *discordgo.MessageCreate) {
	db := CreateInstance()

	insertReminder := `INSERT INTO reminder(username, recipient, date, message, status) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertReminder)

	if err != nil {
		log.Fatalln(err.Error())
	}

	validate, _ := time.ParseInLocation("2006-01-02 15:04", date, time.Local)

	reminder_str = strings.Replace(reminder_str, "'", " ", -1)
	reminder_str = strings.Replace(reminder_str, "\"", " ", -1)

	_, err = statement.Exec(message.Author.Username, message.Author.ID, validate.Unix(), reminder_str, false)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()
}

func UnsetReminder(reminder Remind) {
	db := CreateInstance()
	db.Exec("UPDATE reminder SET status=true WHERE id=?", strconv.Itoa(reminder.id))
	defer db.Close()

}

func SelectActiveReminders() Reminders {
	db := CreateInstance()

	var reminders Reminders
	t := strconv.FormatInt(time.Now().Unix(), 10)

	row, err := db.Query("SELECT id, username, recipient, date, message FROM reminder WHERE status=false and date < " + t)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var username string
		var recipient string
		var date string
		var message string

		err = row.Scan(&id, &username, &recipient, &date, &message)
		if err != nil {
			log.Println("KLAIDA: ", err)
		}

		r := Remind{
			id:        id,
			username:  username,
			recipient: recipient,
			date:      date,
			message:   message,
		}

		reminders = append(reminders, r)

	}

	defer db.Close()
	return reminders
}
