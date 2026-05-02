package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	tokens "server/database/tokens"
	users "server/database/users"
	"time"

	"github.com/lib/pq"
)

type Data struct {
	id                int
	email             string
	username          string
	theme             string
	info              string
	avatar_link       string
	created_at        time.Time
	is_admin          bool
	is_online         bool
	last_seen         time.Time
	is_email_accepted bool
	is_muted_chats_id pq.Int64Array
	is_pinned_chats   pq.Int64Array
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	person := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := person["email"].(string)
	password := person["password"].(string)
	username := person["username"].(string)

	if len(email) < 5 {
		w.WriteHeader(500)
		response := Response{
			Status:  500,
			Message: "Ошибка сервера",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	}
	if len(password) < 6 {
		w.WriteHeader(500)
		return
	}
	if len(username) < 4 {
		w.WriteHeader(500)
		return
	}

	result := users.InputInBasePerson(email, password, username)

	switch result {
	case 200:
		response := Response{
			Status:  200,
			Message: "Пользователь успешно зарегистрирован",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	case 404:
		w.WriteHeader(404)
		return

	default:
		w.WriteHeader(500)
		return
	}
}

func DistributionMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loginGet(w, r)
	case http.MethodPost:
		loginPost(w, r)
	}
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	person := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	email := person["email"].(string)
	password := person["password"].(string)
	result, token := users.Login(email, password)

	if len(email) < 5 {
		w.WriteHeader(500)
		return
	}
	if len(password) < 6 {
		w.WriteHeader(500)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   86400,
	}

	switch result {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно вошли в свой аккаунт",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "%s", jsonData)
		return

	case 400:
		w.WriteHeader(400)
		return

	case 404:
		w.WriteHeader(404)
		return

	default:
		w.WriteHeader(500)
		return
	}
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	status, data := users.GetAllData(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Вы успешно получили свои данные",
			Body:    data,
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
