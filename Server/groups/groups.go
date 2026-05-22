package groups

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	groups "server/database/groups"
	tokens "server/database/tokens"
	"strconv"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	group := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	name := group["name"].(string)
	users_id := group["users_id"].([]interface{})

	status := groups.CreateGroup(user_id, name, users_id)
	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Вы успешно создали группу",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
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

func GetGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data groups.GroupData
	pathValue := r.PathValue("group_id")
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
	group_id, err := strconv.Atoi(pathValue)
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}
	status, data := groups.GetGroupData(group_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Успешно получены данные группы",
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

func ChangeUsersId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("token")
	pathValue := r.PathValue("group_id")
	group_id, err := strconv.Atoi(pathValue)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	group := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	users_id := group["users_id"].([]interface{})
	status := groups.ChangeUsersId(group_id, user_id, users_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Список пользователей успешно изменен",
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

func ChangeAdminsId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("token")
	pathValue := r.PathValue("group_id")
	group_id, err := strconv.Atoi(pathValue)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	group := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	admins_id := group["admins_id"].([]interface{})
	status := groups.ChangeAdminsId(group_id, user_id, admins_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Список администраторов успешно изменен",
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

func ChangeEnemiesId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("token")
	pathValue := r.PathValue("group_id")
	group_id, err := strconv.Atoi(pathValue)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	group := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	enemies_id := group["enemies_id"].([]interface{})
	status := groups.ChangeEnemiesId(group_id, user_id, enemies_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Список заблокированных пользователей успешно изменен",
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

func ChangeNameInfoGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("token")
	pathValue := r.PathValue("group_id")
	group_id, err := strconv.Atoi(pathValue)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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
	if user_id == 430 {
		w.WriteHeader(430)
		return
	}

	group := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	name := group["name"].(string)
	info := group["info"].(string)
	status := groups.ChangeNameInfoGroup(group_id, user_id, name, info)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Данные группы успешно изменены",
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
