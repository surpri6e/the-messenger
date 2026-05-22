package list

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	list "server/database/list"
	tokens "server/database/tokens"
	"time"

	"github.com/lib/pq"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type Chat struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	LastMessage string    `json:"last_message"`
	AvatarLink  string    `json:"avatar_link"`
	UpdatedAt   time.Time `json:"updated_at"`
	UnreadCount int       `json:"unread_count"`
}

type Group struct {
	Id          int           `json:"id"`
	Owner_id    string        `json:"owner_id"`
	Name        string        `json:"name"`
	Info        string        `json:"info"`
	Users_id    pq.Int64Array `json:"users_id"`
	Admins_id   pq.Int64Array `json:"admins_id"`
	Enemies_id  pq.Int64Array `json:"enemies_id"`
	Created_at  string        `json:"created_at"`
	LastMessage string        `json:"last_message"`
	AvatarLink  string        `json:"avatar_link"`
}

func GetUserChatsById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	fmt.Println(cookie)
	if err != nil {
		w.WriteHeader(429)
		return
	}
	token := fmt.Sprint(cookie)
	if token == "nil" {
		w.WriteHeader(500)
		return
	}
	token = token[6:]
	user_id := tokens.CheckToken(token)

	status, chats := list.GetUserChatsById(user_id)
	if chats == nil {
		chats = []list.Chat{}
	}
	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены чаты конкретного пользователя",
			Body:    chats,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	}
	if status == 500 {
		w.WriteHeader(500)
		return
	}
}

func GetUserGroupsById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	fmt.Println(cookie)
	if err != nil {
		w.WriteHeader(429)
		return
	}
	token := fmt.Sprint(cookie)
	if token == "nil" {
		w.WriteHeader(500)
		return
	}
	token = token[6:]
	user_id := tokens.CheckToken(token)

	status, groups := list.GetUserGroupsById(user_id)
	if groups == nil {
		groups = []list.Group{}
	}
	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены группы конкретного пользователя",
			Body:    groups,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	}
	if status == 500 {
		w.WriteHeader(500)
		return
	}
}
