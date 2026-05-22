package chats

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Data struct {
	Id               int    `json:"id"`
	First_person_id  int    `json:"first_person_id"`
	Second_person_id int    `json:"second_person_id"`
	Created_at       string `json:"created_at"`
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

func CreateChat(first_id, second_id float64) (int, Data) {
	dbCreateChat := os.Getenv("DB_CHAT_CREATE")
	dbAlreadyExistChat := os.Getenv("DB_ALREADY_EXIST_CHAT")
	dbAllChatData := os.Getenv("DB_SELECT_ALL_CHAT_DATA")
	var id1, id2, temp int
	var data Data
	if dbCreateChat == "" {
		log.Println("dbCreateChat is empty")
		return 500, data
	}
	if dbAlreadyExistChat == "" {
		log.Println("dbAlreadyExistChat is empty")
		return 500, data
	}
	if dbAllChatData == "" {
		log.Println("dbAllChatData is empty")
		return 500, data
	}
	if first_id > second_id {
		id1 = int(second_id)
		id2 = int(first_id)
	} else {
		id1 = int(first_id)
		id2 = int(second_id)
	}
	err := database.QueryRow(dbAlreadyExistChat, id1, id2).Scan(&temp)
	if err == sql.ErrNoRows {
		_, err = database.Exec(dbCreateChat, id1, id2, times())
		if err != nil {
			log.Println(err)
			return 500, data
		}
		err = database.QueryRow(dbAllChatData, id1, id2).Scan(
			&data.Id,
			&data.First_person_id,
			&data.Second_person_id,
			&data.Created_at,
		)
		if err != nil {
			errorHandler(err)
			return 500, data
		}
		return 200, data
	} else {
		return 404, data
	}

}

func GetAllDataFromChats(chat_id, user_id int) (int, Data) {
	dbGetChatData := os.Getenv("DB_SELECT_ALL_DATA_CHAT")
	dbValidToken := os.Getenv("DB_VALID_TOKEN")
	var data Data
	var first_temp, second_temp int

	if dbGetChatData == "" {
		log.Println("dbGetChatData is empty")
		return 500, data
	}
	if dbValidToken == "" {
		log.Println("dbValidToken is empty")
		return 500, data
	}

	err := database.QueryRow(dbValidToken, chat_id).Scan(&first_temp, &second_temp)
	if err == sql.ErrNoRows {
		errorHandler(err)
		return 404, data
	}
	if err != sql.ErrNoRows && err != nil {
		errorHandler(err)
		return 500, data
	}
	if user_id == first_temp || user_id == second_temp {
		err = database.QueryRow(dbGetChatData, chat_id).Scan(
			&data.Id,
			&data.First_person_id,
			&data.Second_person_id,
			&data.Created_at)
		return 200, data
	} else {
		return 430, data
	}

}

func DeleteChat(chat_id, user_id int) int {
	dbDeleteChat := os.Getenv("DB_DELETE_CHAT")
	dbValidToken := os.Getenv("DB_VALID_TOKEN")
	var first_temp, second_temp int
	if dbDeleteChat == "" {
		log.Println("dbDeleteChat is empty")
		return 500
	}
	if dbValidToken == "" {
		log.Println("dbValidToken is empty")
		return 500
	}
	err := database.QueryRow(dbValidToken, chat_id).Scan(&first_temp, &second_temp)
	if err == sql.ErrNoRows {
		errorHandler(err)
		return 404
	}
	if err != sql.ErrNoRows && err != nil {
		errorHandler(err)
		return 500
	}
	if user_id == first_temp || user_id == second_temp {
		_, err = database.Exec(dbDeleteChat, chat_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 430
	}
}
