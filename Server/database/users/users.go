package users

import (
	"database/sql"
	"log"
	"os"
	token "server/tokens"
	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type SearchData struct {
	Id                int       `json:"id"`
	Username          string    `json:"username"`
	Info              string    `json:"info"`
	Avatar_link       string    `json:"avatar_link"`
	Created_at        time.Time `json:"created_at"`
	Is_online         bool      `json:"is_online"`
	Is_email_accepted bool      `json:"is_email_accepted"`
	Last_seen         time.Time `json:"last_seen"`
}

type Data struct {
	Id                int           `json:"id"`
	Email             string        `json:"email"`
	Username          string        `json:"username"`
	Theme             string        `json:"theme"`
	Info              string        `json:"info"`
	Avatar_link       string        `json:"avatar_link"`
	Created_at        time.Time     `json:"created_at"`
	Is_admin          bool          `json:"is_admin"`
	Is_online         bool          `json:"is_online"`
	Last_seen         time.Time     `json:"last_seen"`
	Is_email_accepted bool          `json:"is_email_accepted"`
	Is_muted_chats_id pq.Int64Array `json:"is_muted_chats_id"`
	Is_pinned_chats   pq.Int64Array `json:"is_pinned_chats"`
}

func errorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

func times() time.Time {
	return time.Now()
}

var database *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	errorHandler(err)
	connToPG := os.Getenv("DB_CONNECT_TO_BASEDATA")
	if connToPG == "" {
		log.Println("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func InputInBasePerson(email string, password string, username string) int {
	var existingLogin string
	dbAccPars := os.Getenv("DB_CHECK_USER_EXIST")
	dbInsertPars := os.Getenv("DB_REGISTER_USER")
	dbGetUserId := os.Getenv("DB_USER_GET_USER_ID")

	if dbAccPars == "" {
		log.Println("dbAccPars is empty")
		return 500
	}
	if dbInsertPars == "" {
		log.Println("dbInsertPers is empty")
		return 500
	}
	if dbGetUserId == "" {
		log.Println("dbGetUserID is empty")
		return 500
	}

	err := database.QueryRow(dbAccPars, email).Scan(&existingLogin)
	empty := ""
	empty_mas := []int{0}
	if err == sql.ErrNoRows {
		_, err := database.Exec(dbInsertPars, email, password, username, empty, empty, empty, times(), times(), empty_mas, empty_mas)

		if err != nil {
			errorHandler(err)
			return 500
		}

		var id int
		err = database.QueryRow(dbGetUserId, email).Scan(&id)

		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	}
	return 404
}

func Login(email string, password string) (int, string) {
	dbPassPars := os.Getenv("DB_USER_PASSWORD_PARSER")
	dbGetUserId := os.Getenv("DB_USER_GET_USER_ID")
	dbGetToken := os.Getenv("DB_GET_TOKEN")
	dbTokenAdd := os.Getenv("DB_TOKEN_ADD")
	dbTokenPars := os.Getenv("DB_GET_TOKEN_TOKEN")
	var id int
	var existingPassword string
	if dbPassPars == "" {
		log.Println("dbPassPars is empty")
		return 500, ""
	}
	if dbGetUserId == "" {
		log.Println("dbGetUserId is empty")
		return 500, ""
	}
	if dbTokenAdd == "" {
		log.Println("dbTokenAdd is empty")
		return 500, ""
	}
	if dbTokenPars == "" {
		log.Println("dbTokenPars is empty")
		return 500, ""
	}
	err := database.QueryRow(dbPassPars, email).Scan(&existingPassword)
	if err == sql.ErrNoRows {
		return 404, ""
	}

	if existingPassword == password {
		err = database.QueryRow(dbGetUserId, email).Scan(&id)
		if err != nil {
			errorHandler(err)
			return 500, ""
		}
		token := token.Crypto(email)
		_, err = database.Exec(dbTokenPars, token)
		if err != sql.ErrNoRows {
			_, err = database.Exec(dbTokenAdd, id, token)
			if err != nil {
				errorHandler(err)
				return 500, ""
			}
		}
		err = database.QueryRow(dbGetToken, id).Scan(&token)
		if err != nil {
			errorHandler(err)
			return 500, ""
		}
		return 200, token

	} else {
		return 400, ""
	}
}

func GetAllData(user_id int) (int, *Data) {
	dbGetData := os.Getenv("DB_GET_ALL_DATA")
	var data Data
	if dbGetData == "" {
		log.Println("dbGetData is empty")
		return 500, &data
	}
	err := database.QueryRow(dbGetData, user_id).Scan(
		&data.Id,
		&data.Email,
		&data.Username,
		&data.Theme,
		&data.Info,
		&data.Avatar_link,
		&data.Created_at,
		&data.Is_admin,
		&data.Is_online,
		&data.Last_seen,
		&data.Is_email_accepted,
		&data.Is_muted_chats_id,
		&data.Is_pinned_chats)
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500, nil
	}
	if err == sql.ErrNoRows {
		return 404, nil
	}
	return 200, &data
}

func InsertCode(email string, code int) int {
	dbInsertCode := os.Getenv("DB_CODE_ADD")
	dbCheckUser := os.Getenv("DB_CHECK_USER_EXIST")
	var tempEmail string
	if dbInsertCode == "" {
		log.Println(dbInsertCode)
		return 500
	}

	err := database.QueryRow(dbCheckUser, email).Scan(&tempEmail)
	if err == sql.ErrNoRows {
		return 400
	}

	if tempEmail == email {
		_, err = database.Exec(dbInsertCode, email, code)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	}
	return 500
}

// Изменение пароля при помощи восстановления пароля
func EmailChangePass(email string, password string, code int) int {
	dbChangePass := os.Getenv("DB_UPDATE_PASSWORD")
	dbValidationCode := os.Getenv("DB_VALID_CODE")
	var tempCode int
	if dbChangePass == "" {
		log.Println("dbChangePass is empty")
		return 500
	}
	if dbValidationCode == "" {
		log.Println("dbValidationCode is empty")
		return 500
	}

	err := database.QueryRow(dbValidationCode, email).Scan(&tempCode)
	if err == sql.ErrNoRows {
		return 400
	}
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500
	}

	if tempCode == code {
		_, err = database.Exec(dbChangePass, password, email)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 404
	}
}

func SearchUser(username string, user_id int) (int, []SearchData) {
	var data []SearchData
	dbCheckValid := os.Getenv("DB_CHECK_VALID_USER")
	dbSearchUser := os.Getenv("DB_SEARCH_USER_DATA")
	searchPattern := "%" + username + "%"
	if dbCheckValid == "" {
		log.Println("dbCheckValid is empty")
		return 500, data
	}
	if dbSearchUser == "" {
		log.Println("dbSearchUser is empty")
		return 500, data
	}

	rows, err := database.Query(dbSearchUser, searchPattern, user_id)

	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500, data
	}
	defer rows.Close()
	for rows.Next() {
		var tempData SearchData
		if err := rows.Scan(
			&tempData.Id,
			&tempData.Username,
			&tempData.Info,
			&tempData.Avatar_link,
			&tempData.Created_at,
			&tempData.Is_online,
			&tempData.Is_email_accepted,
			&tempData.Last_seen,
		); err != nil {
			errorHandler(err)
			return 404, data
		}
		data = append(data, tempData)
	}
	return 200, data
}

func GetUserDataById(userId int) (int, *SearchData) {
	dbQuery := os.Getenv("DB_GET_USER_BY_ID")
	if dbQuery == "" {
		log.Fatal("dbQuery is empty")
		return 500, nil
	}
	var data SearchData
	err := database.QueryRow(dbQuery, userId).Scan(
		&data.Id,
		&data.Username,
		&data.Info,
		&data.Avatar_link,
		&data.Created_at,
		&data.Is_online,
		&data.Last_seen,
		&data.Is_email_accepted,
	)
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500, nil
	}
	if err == sql.ErrNoRows {
		return 404, nil
	}
	return 200, &data
}

func ChangeUserData(user_id int, theme string, info string, username string) int {
	dbUpdateDatas := os.Getenv("DB_UPDATE_DATAS_USER")

	if dbUpdateDatas == "" {
		log.Println("dbUpdateDatas is empty")
		return 500
	}

	_, err := database.Exec(dbUpdateDatas, theme, info, username, user_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func ExitFromAccount(user_id int) int {
	dbTokenDelete := os.Getenv("DB_DELETE_TOKEN")
	if dbTokenDelete == "" {
		log.Println("dbTokenDelete is empty")
		return 500
	}
	_, err := database.Exec(dbTokenDelete, user_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func ChangeAvatarLink(user_id int, avatar_link string) int {
	dbUpdateAvatar := os.Getenv("DB_UPDATE_USER_AVATAR_LINK")
	if dbUpdateAvatar == "" {
		log.Println("dbUpdateAvatar is empty")
		return 500
	}
	_, err := database.Exec(dbUpdateAvatar, avatar_link, user_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func ChangeIsOnline(user_id int) int {
	dbUpdateOnline := os.Getenv("DB_UPDATE_IS_ONLINE")
	if dbUpdateOnline == "" {
		log.Println("dbUpdateOnline is empty")
		return 500
	}
	_, err := database.Exec(dbUpdateOnline, true, time.Now(), user_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func CheckerOnline() {
	CheckOnline := os.Getenv("DB_GET_LAST_SEEN_USER")
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			_, err := database.Exec(CheckOnline)
			if err != nil {
				errorHandler(err)
			}
		}
	}()
}
