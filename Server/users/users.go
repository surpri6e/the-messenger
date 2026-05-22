package users

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	tokens "server/database/tokens"
	users "server/database/users"
	"strconv"
	"strings"
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
	var flag int
	var new_email string
	var list [3]string = [3]string{"@gmail.com", "@yandex.ru", "@mail.ru"}
	var new_flag int
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
	for i := 0; i < len(email); i++ {
		if string(email[i]) == "@" {
			flag = 1
		}
		if flag == 1 {
			new_email += string(email[i])
		}
	}
	for _, value := range list {
		if new_email == value {
			new_flag = 1
			break
		}
	}
	if new_flag == 1 {
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
				log.Println(err)
			}
			fmt.Fprintf(w, "%s", jsonData)
			return

		default:
			w.WriteHeader(result)
			return
		}
	} else {
		w.WriteHeader(400)
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
			log.Println(err)
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "%s", jsonData)
		return

	default:
		w.WriteHeader(result)
		return
	}
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	status, data := users.GetAllData(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Вы успешно получили свои данные",
			Body:    data,
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

func SendMail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	email := r.PathValue("email")

	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	FromEmail := os.Getenv("FROM_EMAIL")
	Password := os.Getenv(("PASSWORD"))
	if FromEmail == "" {
		log.Println("FromEmail is empty")
		w.WriteHeader(500)
		return
	}
	if Password == "" {
		log.Println("Password is empty")
		w.WriteHeader(500)
		return
	}
	fromEmail := FromEmail
	password := Password
	smtpHost := "smtp.mail.ru"
	smtpPort := "587"
	toEmail := []string{email}

	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	var message strings.Builder

	message.WriteString(fmt.Sprintf("From: \"Техподдержка\" <%s>\r\n", fromEmail))
	message.WriteString(fmt.Sprintf("To: %s\r\n", email))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", "Код подтверждения"))
	message.WriteString(fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z)))
	message.WriteString(fmt.Sprintf("Message-ID: %s\r\n", code))
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	message.WriteString("Content-Transfer-Encoding: 8bit\r\n")
	message.WriteString("\r\n")
	body := fmt.Sprintf(`Уважаемый пользователь!

Вы запросили код подтверждения для восстановления пароля.
Ваш код: %d

Код действителен в течение 10 минут.
Если вы не запрашивали код, проигнорируйте это сообщение.

--
С уважением,
Команда поддержки мессенджера Notus
%s`, code, time.Now().Format("2006-01-02"))
	message.WriteString(body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, toEmail, []byte(message.String()))
	if err != nil {
		w.WriteHeader(500)
		log.Println("Ошибка отправки письма: ", err)
		return
	}
	status := users.InsertCode(email, code)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Код успешно отправлен",
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
	}
	fmt.Println("Письмо успешно отправлено!")

}

func MailChangePass(w http.ResponseWriter, r *http.Request) {
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
	code := person["code"].(int)

	status := users.EmailChangePass(email, password, code)

	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Пароль успешно изменен",
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
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
	username := r.PathValue("username")
	user_id := tokens.CheckToken(token)
	status, data := users.SearchUser(username, user_id)
	log.Println(data)
	if data == nil {
		status = 404
	}
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Успешно найдены пользователи",
			Body:    data,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	case 404:
		var emptymas []string = []string{}
		response := Response{
			Status:  200,
			Message: "Успешно найдены пользователи",
			Body:    emptymas,
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

func GetUserDataById(w http.ResponseWriter, r *http.Request) {
	temp := r.PathValue("user_id")
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
	user_i := tokens.CheckToken(token)
	if user_i == 430 {
		w.WriteHeader(430)
		return
	}
	user_id, err := strconv.Atoi(temp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(430)
	}
	status, data := users.GetUserDataById(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены данные о конкретном пользователе",
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
	if status != 200 {
		w.WriteHeader(status)
		return
	}
}

func ChangeUserData(w http.ResponseWriter, r *http.Request) {
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

	userData := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	theme := userData["theme"].(string)
	info := userData["info"].(string)
	username := userData["username"].(string)
	status := users.ChangeUserData(user_id, theme, info, username)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы успешно изменили информацию",
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
		w.WriteHeader(status)
	}
}

func ExitFromAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
	status := users.ExitFromAccount(user_id)
	switch status {
	case 200:
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			MaxAge:   -1,
		}

		response := Response{
			Status:  200,
			Message: "Вы успешно вышли из аккаунта",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "%s", jsonData)
		return
	default:
		w.WriteHeader(status)
		return
	}
}

func ChangeIsOnline(w http.ResponseWriter, r *http.Request) {
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
	status := users.ChangeIsOnline(user_id)
	switch status {
	case 200:
		response := Response{
			Status:  200,
			Message: "Вы онлайн",
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
		w.WriteHeader(status)
	}
}
