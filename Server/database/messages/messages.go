package messages

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Message struct {
	Id           int
	User_id      int
	Where_id     int
	Text         string
	Status       string
	Is_pinned    bool
	Created_at   time.Time
	Is_changed   bool
	Is_forwarded bool
	Type         string
	File_link    string
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func times() time.Time {
	return time.Now()
}

var database *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	errorHandler(err)
	connToPG := os.Getenv("DB_CONNECT_TO_BASEDATA")
	if connToPG == "" {
		log.Fatal("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func PostMessage(user_id float64, where_id float64, text string, status string, Type string, file_link string) int {
	dbMessPost := os.Getenv("DB_MESSAGES_POST_MESSAGE")
	if dbMessPost == "" {
		log.Fatal("dbMessPost is empty")
		return 500
	}
	_, err := database.Exec(dbMessPost, user_id, where_id, text, status, Type, file_link, times())
	if err != nil {
		errorHandler(err)
		return 500
	} else {
		return 200
	}
}

func GetMessage(message_id int, user_id int) (int, Message) {
	dbGetMessage := os.Getenv("DB_MESSAGES_GET_MESSAGE")
	var message Message
	if dbGetMessage == "" {
		log.Fatal("dbGetMessage is empty")
		return 500, message
	}
	err := database.QueryRow(dbGetMessage, message_id).Scan(
		&message.Id,
		&message.User_id,
		&message.Where_id,
		&message.Text,
		&message.Status,
		&message.Is_pinned,
		&message.Created_at,
		&message.Is_changed,
		&message.Is_forwarded,
		&message.Type,
		&message.File_link)

	if message.User_id != 0 && message.User_id != user_id {
		return 430, message
	}
	if err == sql.ErrNoRows {
		return 404, message
	} else {
		return 200, message
	}
}

func ChangeMessage(text, is_pinned, is_forwarded string, message_id int, user_id int) int {
	dbUpdateMess := os.Getenv("DB_UPDATE_MESSAGE")
	dbValidUser := os.Getenv("DB_VALID_USER")
	var temp_id int

	if dbUpdateMess == "" {
		log.Fatal("dbUpdateMess is empty")
		return 500
	}
	if dbValidUser == "" {
		log.Fatal("dbValidUser is empty")
		return 500
	}

	err := database.QueryRow(dbValidUser, message_id).Scan(&temp_id)
	if err == sql.ErrNoRows {
		return 404
	}
	if err != nil {
		return 500
	}

	if user_id == temp_id {
		_, err = database.Exec(dbUpdateMess, text, is_pinned, is_forwarded, message_id)

		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	}
	return 500
}
