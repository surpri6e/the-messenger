package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	message "server/database/messages"
	tokens "server/database/tokens"
	"strconv"
)

type Response struct {
	Status  int         `json:"status`
	Message string      `json:"message`
	Body    interface{} `json:"body`
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(429)
		return
	}
	messages := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user_id := messages["user_id"].(float64)

	token := fmt.Sprint(cookie)
	token = token[6:]
	status := tokens.CheckToken(token)

	if status != int(user_id) {
		w.WriteHeader(430)
		return
	}

	where_id := messages["where_id"].(float64)
	text := messages["text"].(string)

	if len(text) < 1 {
		w.WriteHeader(500)
		return
	}

	Status := messages["status"].(string)
	Type := messages["type"].(string)
	file_link := messages["file_link"].(string)

	status = message.PostMessage(user_id, where_id, text, Status, Type, file_link)

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно отправили сообщение",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	default:
		w.WriteHeader(500)
		return
	}
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	message_id := r.PathValue("message_id")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(429)
		return
	}
	token := fmt.Sprint(cookie)
	token = token[6:]
	user_id := tokens.CheckToken(token)

	id, err := strconv.Atoi(message_id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	status, data := message.GetMessage(id, user_id)

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно получили сообщение",
			Body:    data,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	case 404:
		w.WriteHeader(404)
		return

	case 430:
		w.WriteHeader(430)
		return

	default:
		w.WriteHeader(500)
		return
	}
}

func PutMessage(w http.ResponseWriter, r *http.Request) {
	message_id := r.PathValue("message_id")
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(429)
		return
	}

	token := fmt.Sprint(cookie)
	token = token[6:]
	user_id := tokens.CheckToken(token)

	messages := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text := messages["text"].(string)
	is_pinned := messages["is_pinned"].(string)
	is_forwarded := messages["is_forwarded"].(string)

	mess_id, err := strconv.Atoi(message_id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	status := message.ChangeMessage(text, is_pinned, is_forwarded, mess_id, user_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно изменили сообщение",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	case 404:
		w.WriteHeader(404)
		return

	case 430:
		w.WriteHeader(430)
		return

	default:
		w.WriteHeader(500)
		return
	}
}
