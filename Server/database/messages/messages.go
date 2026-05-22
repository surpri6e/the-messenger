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
	Id           int       `json:"id"`
	User_id      int       `json:"user_id"`
	Where_id     int       `json:"where_id"`
	Text         string    `json:"text"`
	Status       string    `json:"status"`
	Is_pinned    bool      `json:"is_pinned"`
	Created_at   time.Time `json:"created_at"`
	Is_changed   bool      `json:"is_changed"`
	Is_forwarded bool      `json:"is_forwarded"`
	Type         string    `json:"type"`
	File_link    string    `json:"file_link"`
}

func errorHandler(err error) {
	if err != nil {
		log.Println(err)
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
		log.Println("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func PostMessage(user_id int, where_id float64, text string, status string, Type string, file_link string) int {
	dbMessPost := os.Getenv("DB_MESSAGES_POST_MESSAGE")
	dbThereisChat := os.Getenv("DB_IS_THERE_CHAT")
	if dbMessPost == "" {
		log.Println("dbMessPost is empty")
		return 500
	}
	if dbThereisChat == "" {
		log.Println("dbThereischat is empty")
		return 500
	}
	var temp int
	err := database.QueryRow(dbThereisChat, where_id).Scan(&temp)
	if err == sql.ErrNoRows {
		_, err = database.Exec(dbMessPost, user_id, where_id, text, status, Type, file_link, times())
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		_, err = database.Exec(dbMessPost, user_id, where_id, text, status, Type, file_link, times())
		if err != nil {
			errorHandler(err)
			return 500
		} else {
			return 200
		}
	}
}

func GetMessage(message_id int, user_id int) (int, Message) {
	dbGetMessage := os.Getenv("DB_MESSAGES_GET_MESSAGE")
	var message Message
	if dbGetMessage == "" {
		log.Println("dbGetMessage is empty")
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
	dbUpdateTextMess := os.Getenv("DB_UPDATE_TEXT_MESSAGE")
	dbUpdatePinnedMess := os.Getenv("DB_UPDATE_PINNED_MESSAGE")
	dbUpdateForwardedMess := os.Getenv("DB_UPDATE_FORWARDED_MESSAGE")
	dbValidUser := os.Getenv("DB_VALID_USER")
	var temp_id int

	if dbUpdateTextMess == "" {
		log.Println("dbUpdateTextMess is empty")
		return 500
	}
	if dbUpdatePinnedMess == "" {
		log.Println("dbUpdatePinnedMess is empty")
		return 500
	}
	if dbUpdateForwardedMess == "" {
		log.Println("dbUpdateForwardMess is empty")
	}
	if dbValidUser == "" {
		log.Println("dbValidUser is empty")
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
		if is_pinned == "" && is_forwarded == "" {
			_, err = database.Exec(dbUpdateTextMess, text, true, message_id)
			if err != nil {
				errorHandler(err)
				return 500
			}
		}
		if text == "" && is_forwarded == "" {
			_, err = database.Exec(dbUpdatePinnedMess, is_pinned, message_id)
			if err != nil {
				errorHandler(err)
				return 500
			}
		}
		if is_pinned == "" && text == "" {
			_, err = database.Exec(dbUpdateForwardedMess, is_forwarded, message_id)
			if err != nil {
				errorHandler(err)
				return 500
			}
		}
		return 200
	}
	return 500
}

func DeleteMessage(user_id int, message_id int) int {
	dbValidUser := os.Getenv("DB_VALID_USER")
	dbDeleteMess := os.Getenv("DB_DELETE_MESSAGE")
	var temp_id int
	if dbValidUser == "" {
		log.Println("dbValidUser is empty")
	}
	if dbDeleteMess == "" {
		log.Println("dbDeleteMess is empty")
	}

	err := database.QueryRow(dbValidUser, message_id).Scan(&temp_id)
	if err == sql.ErrNoRows {
		errorHandler(err)
		return 404
	}
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500
	}

	if temp_id == user_id {
		_, err = database.Exec(dbDeleteMess, message_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	}
	return 500
}

func GetNewMessages(message_id int) (int, Message) {
	dbGetWhereId := os.Getenv("DB_GET_CHAT_ID_MESSAGES")
	dbGetMess := os.Getenv("DB_GET_NEW_MESSAGES")
	var message Message
	var where_id int
	if dbGetWhereId == "" {
		log.Println("dbGetWhereId is empty")
		return 500, message
	}
	if dbGetMess == "" {
		log.Println("dbGetMess is empty")
		return 500, message
	}
	err := database.QueryRow(dbGetWhereId, message_id).Scan(&where_id)
	if err != nil {
		errorHandler(err)
		return 500, message
	}
	err = database.QueryRow(dbGetMess, where_id, message_id).Scan(
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
	if err != nil {
		errorHandler(err)
		return 500, message
	}
	return 200, message
}

func GetMessageChatId(chat_id int) (int, []Message) {
	var messages []Message
	dbGetMessages := os.Getenv("DB_GET_ALL_MESSAGES")
	if dbGetMessages == "" {
		log.Println("dbGetMessages is empty")
		return 500, messages
	}
	rows, err := database.Query(dbGetMessages, chat_id)
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500, messages
	}
	if err == sql.ErrNoRows {
		return 404, messages
	}
	defer rows.Close()
	for rows.Next() {
		var tempMessage Message
		if err := rows.Scan(
			&tempMessage.Id,
			&tempMessage.User_id,
			&tempMessage.Where_id,
			&tempMessage.Text,
			&tempMessage.Status,
			&tempMessage.Is_pinned,
			&tempMessage.Created_at,
			&tempMessage.Is_changed,
			&tempMessage.Is_forwarded,
			&tempMessage.Type,
			&tempMessage.File_link,
		); err != nil {
			return 500, messages
		}
		messages = append(messages, tempMessage)
	}
	if rows.Err() == sql.ErrNoRows {
		return 404, messages
	}
	return 200, messages

}
