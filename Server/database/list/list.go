package lists

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Chat struct {
	Id               int    `json:"id"`
	First_person_id  int    `json:"first_person_id"`
	Second_person_id int    `json:"second_person_id"`
	Created_at       string `json:"created_at"`
}

type Group struct {
	Id         int           `json:"id"`
	Owner_id   string        `json:"owner_id"`
	Name       string        `json:"name"`
	Info       string        `json:"info"`
	Users_id   pq.Int64Array `json:"users_id"`
	Admins_id  pq.Int64Array `json:"admins_id"`
	Enemies_id pq.Int64Array `json:"enemies_id"`
	Created_at string        `json:"created_at"`
	AvatarLink string        `json:"avatar_link"`
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

func GetUserChatsById(userId int) (int, []Chat) {

	dbQuery := os.Getenv("DB_GET_USER_CHATS")
	if dbQuery == "" {
		log.Println("DB_GET_USER_CHATS is empty")
		return 500, nil
	}

	rows, err := database.Query(dbQuery, userId)
	if err != nil {
		log.Printf("Error getting chats: %v", err)
		return 500, nil
	}
	defer rows.Close()

	var chats []Chat

	for rows.Next() {
		var chat Chat
		err := rows.Scan(
			&chat.Id,
			&chat.First_person_id,
			&chat.Second_person_id,
			&chat.Created_at,
		)
		if err != nil {
			log.Printf("Error scanning chat: %v", err)
			continue
		}
		chats = append(chats, chat)
	}

	return 200, chats
}

func GetUserGroupsById(userId int) (int, []Group) {

	dbQuery := os.Getenv("DB_GET_USER_GROUPS")
	if dbQuery == "" {
		log.Println("DB_GET_USER_GROUPS is empty")
		return 500, nil
	}

	rows, err := database.Query(dbQuery, userId)
	if err != nil {
		log.Printf("Error getting groups: %v", err)
		return 500, nil
	}
	defer rows.Close()

	var groups []Group

	for rows.Next() {
		var group Group
		err := rows.Scan(
			&group.Id,
			&group.Owner_id,
			&group.Name,
			&group.Info,
			&group.Users_id,
			&group.Admins_id,
			&group.Enemies_id,
			&group.Created_at,
			&group.AvatarLink,
		)
		if err != nil {
			log.Printf("Error scanning chat: %v", err)
			continue
		}
		groups = append(groups, group)
	}

	return 200, groups
}
