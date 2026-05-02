package main

import (
	"log"
	"net/http"
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
	mux := http.NewServeMux()
	//	mux.HandleFunc("/api/health", checkHealth)
	mux.HandleFunc("/registration", users.Register)
	mux.HandleFunc("/auth", users.DistributionMethod)
	mux.HandleFunc("/messages", messages.PostMessage)
	mux.HandleFunc("GET /messages/{message_id}", messages.GetMessage)
	mux.HandleFunc("PUT /messages/{message_id}", messages.PutMessage)

	// Оборачиваем весь маршрутизатор в CORS
	handler := enableCORSForMux(mux) //BE2681D73EB0903771C9B2DCC76FCFC04768FF7F
	handler = LoggingMiddleware(handler)

	log.Print("Listening on port 8080")
	err := http.ListenAndServeTLS("26.132.220.182:8080", "cert.pem", "key.pem", handler)

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

// func DistributionMethod(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		getMessage(w, r)
// 	case http.MethodPost:
// 		postMessage(w, r)
// 	}
// }

// func postMessage(w http.ResponseWriter, r *http.Request) {
// 	message := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	sender_id := message["Sender_id"].(float64)
// 	user_id := message["User_id"].(float64)
// 	chat_id := message["Chat_id"].(float64)
// 	text := message["Text"].(string)
// 	err = main2.InputMessage(sender_id, user_id, chat_id, text)
// 	if err == nil {
// 		fmt.Fprintf(w, "OK")
// 	}
// 	ResponseWriter()
// }

// func getMessages(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	chat_id := query.Get("chat_id")
// 	// err := json.NewDecoder(r.Body).Decode(&us_ch_id)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }
// 	id, err := strconv.Atoi(chat_id)
// 	result, err := json.Marshal(main2.GetMessages(float64(id)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Fprintf(w, "%s", result)
// }

// func getMessage(w http.ResponseWriter, r *http.Request) {
// 	us_ch_id := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&us_ch_id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	res, err := database.GetMessage(us_ch_id["Sender_id"].(float64), us_ch_id["User_id"].(float64))
// 	if err != nil {
// 		fmt.Fprintf(w, "No message in row")
// 		return
// 	}
// 	result, err := json.Marshal(res)
// 	fmt.Fprintf(w, "%s", result)
// }

// func register(w http.ResponseWriter, r *http.Request) {
// 	person := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&person)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	email := person["Email"].(string)
// 	password := person["Password"].(string)
// 	result := main2.InputInBasePerson(email, password)
// 	if result == true {
// 		fmt.Fprintf(w, "%d", 200)
// 	} else {
// 		fmt.Fprintf(w, "%d", 400)
// 	}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	person := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&person)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	email := person["Email"].(string)
// 	password := person["Password"].(string)
// 	result := main2.CheckLogin(email, password)
// 	if result == true {
// 		fmt.Fprintf(w, "%d", 200)
// 	} else {
// 		fmt.Fprintf(w, "%d", 400)
// 	}
// }

// // Check the server status
// func checkHealth(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "OK")
// }
