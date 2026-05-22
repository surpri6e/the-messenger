package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database/messages"
	message "server/database/messages"
	tokens "server/database/tokens"
	"strconv"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
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
	token := fmt.Sprint(cookie)
	token = token[6:]
	user_id := tokens.CheckToken(token)

	where_id := messages["where_id"].(float64)
	text := messages["text"].(string)
	Status := messages["status"].(string)
	Type := messages["type"].(string)
	file_link := messages["file_link"].(string)

	status := message.PostMessage(user_id, where_id, text, Status, Type, file_link)

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно отправили сообщение",
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

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

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	mess_id := r.PathValue("message_id")
	if r.Method != http.MethodDelete {
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
	message_id, err := strconv.Atoi(mess_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	status := messages.DeleteMessage(user_id, message_id)

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно удалили сообщение",
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

func GetNewMessages(w http.ResponseWriter, r *http.Request) {
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	id, err := strconv.Atoi(message_id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	timeout := time.After(35 * time.Second)
	checkInterval := time.NewTicker(1 * time.Second)
	defer checkInterval.Stop()
	for {
		select {
		case <-timeout:
			w.WriteHeader(404)
			return
		case <-checkInterval.C:
			status, message := messages.GetNewMessages(id)
			switch status {
			case 200:
				response := Response{
					Status:  200,
					Message: "Вы успешно получили сообщение",
					Body:    message,
				}
				jsonData, err := json.Marshal(response)
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Fprintf(w, "%s", jsonData)
				return
			}
		}
	}
}

func GetMessageChatId(w http.ResponseWriter, r *http.Request) {
	chat_id := r.PathValue("chat_id")
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	id, err := strconv.Atoi(chat_id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	status, messages := messages.GetMessageChatId(id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно получили сообщения",
			Body:    messages,
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
