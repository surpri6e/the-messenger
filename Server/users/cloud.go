package users

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	tokens "server/database/tokens"
	users "server/database/users"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type YandexStorageMinio struct {
	client     *minio.Client
	bucketName string
}

type Link struct {
	Avatar_link string `json:"avatar_link"`
}

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	public_key := os.Getenv("PUBLIC_KEY")
	private_key := os.Getenv("PRIVATE_KEY")
	if public_key == "" {
		log.Println("piblic_key is empty")
		w.WriteHeader(500)
		return
	}
	if private_key == "" {
		log.Println("private_key is empty")
		w.WriteHeader(500)
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

	err = r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	endpoint := "storage.yandexcloud.net"
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(public_key, private_key, ""),
		Secure: true,
		Region: "ru-central1",
	})
	if err != nil {
		fmt.Fprintf(w, "failed to create client: %w", err)
		return
	}
	file_name := generator(user_id)
	_, err = client.PutObject(r.Context(), "notusavatar", file_name, file, header.Size, minio.PutObjectOptions{
		ContentType:    "image/jpeg",
		CacheControl:   "public, max-age=31536000",
		SendContentMd5: true,
	})
	if err != nil {
		log.Println(err)
		return
	}

	status := users.ChangeAvatarLink(user_id, getImageLink(file_name))
	switch status {
	case 200:
		var link Link
		link.Avatar_link = getImageLink(file_name)
		response := Response{
			Status:  200,
			Message: "Вы успешно изменили информацию",
			Body:    link,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "%s", jsonData)
		return

	default:
		w.WriteHeader(status)
		return
	}
}

func generator(user_id int) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(900000) + 100000
	var file_name string = strconv.Itoa(user_id) + "_" + strconv.Itoa(random) + ".jpg"
	return file_name
}

func getImageLink(file_name string) string {
	return "https://storage.yandexcloud.net/notusavatar/" + file_name
}
