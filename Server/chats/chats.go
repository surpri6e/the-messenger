package chats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	chats "server/database/chats"
	tokens "server/database/tokens"
	"strconv"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var status int
	var data chats.Data
	cookie, err := r.Cookie("token")
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

	chat := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	first_person_id := chat["first_person_id"].(float64)
	second_person_id := chat["second_person_id"].(float64)

	if float64(user_id) == first_person_id || float64(user_id) == second_person_id {
		status, data = chats.CreateChat(first_person_id, second_person_id)
	} else {
		w.WriteHeader(430)
		return
	}

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Чат успешно создан",
			Body:    data,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	default:
		w.WriteHeader(status)
	}
}

func GetChatData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var data chats.Data
	var status int
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
	temp_chat_id := r.PathValue("chat_id")
	chat_id, err := strconv.Atoi(temp_chat_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	status, data = chats.GetAllDataFromChats(chat_id, user_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно получили информацию о чате",
			Body:    data,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	default:
		w.WriteHeader(status)
	}
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
	temp_chat_id := r.PathValue("chat_id")
	chat_id, err := strconv.Atoi(temp_chat_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	status := chats.DeleteChat(chat_id, user_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Чат успешно удалён",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	default:
		w.WriteHeader(status)
		return
	}
}
