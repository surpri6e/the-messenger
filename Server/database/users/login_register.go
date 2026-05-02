package users

import (
	"database/sql"
	"log"
	"os"
	token "server/tokens"
	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Data struct {
	Id                int           `json:"id"`
	Email             string        `json:"email"`
	Username          string        `json:"username"`
	Theme             string        `json:"theme"`
	Info              string        `json:"info"`
	Avatar_link       string        `json:"avatar_link"`
	Created_at        time.Time     `json:"created_at"`
	Is_admin          bool          `json:"is_admin"`
	Is_online         bool          `json:"is_online"`
	Last_seen         time.Time     `json:"last_seen"`
	Is_email_accepted bool          `json:"is_email_accepted"`
	Is_muted_chats_id pq.Int64Array `json:"is_muted_chats_id"`
	Is_pinned_chats   pq.Int64Array `json:"is_pinned_chats"`
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
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

func InputInBasePerson(email string, password string, username string) int {
	var existingLogin string
	dbAccPars := os.Getenv("DB_CHECK_USER_EXIST")
	dbInsertPars := os.Getenv("DB_REGISTER_USER")
	dbTokenAdd := os.Getenv("DB_TOKEN_ADD")
	dbGetUserId := os.Getenv("DB_USER_GET_USER_ID")

	if dbAccPars == "" {
		log.Fatal("dbAccPars is empty")
		return 500
	}
	if dbInsertPars == "" {
		log.Fatal("dbInsertPers is empty")
		return 500
	}
	if dbTokenAdd == "" {
		log.Fatal("dbTokenAdd is empty")
		return 500
	}
	if dbGetUserId == "" {
		log.Fatal("dbGetUserID is empty")
		return 500
	}

	err := database.QueryRow(dbAccPars, email).Scan(&existingLogin)
	empty := ""
	empty_mas := []int{0}
	if err == sql.ErrNoRows {
		_, err := database.Exec(dbInsertPars, email, password, username, empty, empty, empty, times(), times(), empty_mas, empty_mas)

		if err != nil {
			errorHandler(err)
			return 500
		}

		var id int
		err = database.QueryRow(dbGetUserId, email).Scan(&id)

		if err != nil {
			errorHandler(err)
			return 500
		}

		token := token.Crypto(email)

		_, err = database.Exec(dbTokenAdd, id, token)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	}
	return 404
}

func Login(email string, password string) (int, string) {
	dbPassPars := os.Getenv("DB_USER_PASSWORD_PARSER")
	dbGetUserId := os.Getenv("DB_USER_GET_USER_ID")
	dbGetToken := os.Getenv("DB_GET_TOKEN")
	var id int
	var token string
	var existingPassword string
	if dbPassPars == "" {
		log.Fatal("dbPassPars is empty")
		return 500, ""
	}
	if dbGetUserId == "" {
		log.Fatal("dbGetUserId is empty")
		return 500, ""
	}

	err := database.QueryRow(dbPassPars, email).Scan(&existingPassword)
	if err == sql.ErrNoRows {
		return 404, ""
	}

	if existingPassword == password {
		err = database.QueryRow(dbGetUserId, email).Scan(&id)
		if err != nil {
			errorHandler(err)
			return 500, ""
		}
		err = database.QueryRow(dbGetToken, id).Scan(&token)
		if err != nil {
			errorHandler(err)
			return 500, ""
		}
		return 200, token

	} else {
		return 400, ""
	}
}

func GetAllData(user_id int) (int, *Data) {
	dbGetData := os.Getenv("DB_GET_ALL_DATA")
	var data Data
	if dbGetData == "" {
		log.Fatal("dbGetData is empty")
		return 500, &data
	}
	err := database.QueryRow(dbGetData, user_id).Scan(
		&data.Id,
		&data.Email,
		&data.Username,
		&data.Theme,
		&data.Info,
		&data.Avatar_link,
		&data.Created_at,
		&data.Is_admin,
		&data.Is_online,
		&data.Last_seen,
		&data.Is_email_accepted,
		&data.Is_muted_chats_id,
		&data.Is_pinned_chats)
	if err != nil {
		errorHandler(err)
		return 500, nil
	}
	return 200, &data
}
