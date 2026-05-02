package tokens

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
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

func CheckToken(token string) int {
	dbGetUserId := os.Getenv("DB_TOKEN_GET_USER_ID")
	dbCheckToken := os.Getenv("DB_GET_TOKEN_TOKEN")
	var existingToken string
	var user_id int
	if dbGetUserId == "" {
		log.Fatal("dbGetUserId is empty")
		return 500
	}
	if dbCheckToken == "" {
		log.Fatal("dbCheckToken is empty")
		return 500
	}

	err := database.QueryRow(dbCheckToken, token).Scan(&existingToken)
	if err == sql.ErrNoRows {
		return 430
	}
	if err == nil {
		err := database.QueryRow(dbGetUserId, token).Scan(&user_id)

		if err != nil {
			errorHandler(err)
			return 500
		}

		return user_id
	}
	return 500
}
