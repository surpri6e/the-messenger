package main

import (
	"log"
	"net/http"
	chats "server/chats"
	checker "server/database/users"
	groups "server/groups"
	list "server/list"
	messages "server/messages"
	users "server/users"
	"time"
)

type Pole struct {
	p1 int
	p2 int
	p3 int
	p4 int
	p5 string
	p6 string
	p7 time.Time
	p8 string
	p9 string
}

func main() {
	checker.CheckerOnline()
	// Users
	mux := http.NewServeMux()
	mux.HandleFunc("/registration", users.Register)
	mux.HandleFunc("/auth", users.DistributionMethod)
	mux.HandleFunc("GET /search/{username}", users.SearchUser)
	mux.HandleFunc("POST /mailcode/{email}", users.SendMail)
	mux.HandleFunc("PUT /mailcode", users.MailChangePass)
	mux.HandleFunc("GET /users/{user_id}", users.GetUserDataById)
	mux.HandleFunc("PUT /users", users.ChangeUserData)
	mux.HandleFunc("POST /users/exit", users.ExitFromAccount)
	mux.HandleFunc("PUT /users/online", users.ChangeIsOnline)
	mux.HandleFunc("POST /users/avatarlink", users.UploadAvatar)
	// Messages
	mux.HandleFunc("/messages", messages.PostMessage)
	mux.HandleFunc("GET /messages/{message_id}", messages.GetMessage)
	mux.HandleFunc("PUT /messages/{message_id}", messages.PutMessage)
	mux.HandleFunc("DELETE /messages/{message_id}", messages.DeleteMessage)
	mux.HandleFunc("GET /checkmessage/{message_id}", messages.GetNewMessages)

	// Chats
	mux.HandleFunc("/chats", chats.CreateChat)
	mux.HandleFunc("GET /chats/{chat_id}", chats.GetChatData)
	mux.HandleFunc("DELETE /chats/{chat_id}", chats.DeleteChat)
	// Groups
	mux.HandleFunc("POST /groups", groups.CreateGroup)
	mux.HandleFunc("PUT /groups/admins/{group_id}", groups.ChangeAdminsId)
	mux.HandleFunc("PUT /groups/users/{group_id}", groups.ChangeUsersId)
	mux.HandleFunc("PUT /groups/enemies/{group_id}", groups.ChangeEnemiesId)
	mux.HandleFunc("PUT /groups/data/{group_id}", groups.ChangeNameInfoGroup)
	mux.HandleFunc("GET /groups/{group_id}", groups.GetGroup)
	// Lists
	mux.HandleFunc("GET /list/chats/user/{user_id}", list.GetUserChatsById)
	mux.HandleFunc("GET /list/{user_id}/groups", list.GetUserGroupsById)
	mux.HandleFunc("GET /list/communication/chat/{chat_id}", messages.GetMessageChatId)

	handler := enableCORSForMux(mux)
	handler = LoggingMiddleware(handler)

	log.Print("Listening on port 8080")
	err := http.ListenAndServeTLS("26.171.27.118:8080", "cert.pem", "key.pem", handler)

	if err != nil {
		log.Fatal(err)
	}
}

func enableCORSForMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📥 %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
