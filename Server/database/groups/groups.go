package users

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func errorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

func times() time.Time {
	return time.Now()
}

type GroupData struct {
	Id         int           `json:"id"`
	Owner_id   int           `json:"owner_id"`
	Name       string        `json:"name"`
	Info       string        `json:"info"`
	Users_id   pq.Int64Array `json:"users_id"`
	Admins_id  pq.Int64Array `json:"admins_id"`
	Enemies_id pq.Int64Array `json:"enemies_id"`
	Created_at string        `json:"created_at"`
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

func CreateGroup(user_id int, name string, users_id []interface{}) int {
	dbGroupCreate := os.Getenv("DB_GROUP_CREATE")
	if dbGroupCreate == "" {
		log.Println("dbGroupCreate is empty")
		return 500
	}
	users_id = append(users_id, user_id)
	var admins_id []int
	admins_id = append(admins_id, user_id)
	_, err := database.Exec(dbGroupCreate, user_id, name, "", users_id, admins_id, times(), "")
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func GetGroupData(group_id int) (int, GroupData) {
	dbGetGroupData := os.Getenv("DB_GET_DATA_GROUP")
	var data GroupData
	if dbGetGroupData == "" {
		log.Println("dbGetGroupData is empty")
		return 500, data
	}
	err := database.QueryRow(dbGetGroupData, group_id).Scan(
		&data.Id,
		&data.Owner_id,
		&data.Name,
		&data.Info,
		&data.Users_id,
		&data.Admins_id,
		&data.Enemies_id,
		&data.Created_at,
	)
	if err == sql.ErrNoRows {
		return 408, data
	}
	if err != nil && err != sql.ErrNoRows {
		errorHandler(err)
		return 500, data
	}
	return 200, data

}

func ChangeUsersId(group_id, user_id int, users_id []interface{}) int {
	dbParsAdmin := os.Getenv("DB_ADMIN_PARSER_GROUP")
	dbChangeUsers := os.Getenv("DB_CHANGE_USERS_ID_GROUP")
	var flag int
	if dbParsAdmin == "" {
		log.Println("dbParsAdmin is empty")
		return 500
	}
	if dbChangeUsers == "" {
		log.Println("dbChangeUsers is empty")
		return 500
	}

	var admins_id pq.Int64Array
	err := database.QueryRow(dbParsAdmin, group_id).Scan(&admins_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	for i := 0; i < len(admins_id); i++ {
		if admins_id[i] == int64(user_id) {
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		_, err := database.Exec(dbChangeUsers, users_id, group_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 408 // Нет прав
	}
}

func ChangeAdminsId(group_id, user_id int, admins_id []interface{}) int {
	dbParsAdmin := os.Getenv("DB_ADMIN_PARSER_GROUP")
	dbChangeAdmins := os.Getenv("DB_CHANGE_ADMINS_ID_GROUP")
	var flag int
	if dbParsAdmin == "" {
		log.Println("dbParsAdmin is empty")
		return 500
	}
	if dbChangeAdmins == "" {
		log.Println("dbChangeAdmins is empty")
		return 500
	}

	var admin_id pq.Int64Array
	err := database.QueryRow(dbParsAdmin, group_id).Scan(&admin_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	for i := 0; i < len(admins_id); i++ {
		if admin_id[i] == int64(user_id) {
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		_, err := database.Exec(dbChangeAdmins, admins_id, group_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 408 // Нет прав
	}
}

func ChangeEnemiesId(group_id, user_id int, enemies_id []interface{}) int {
	dbParsAdmin := os.Getenv("DB_ADMIN_PARSER_GROUP")
	dbChangeEnemies := os.Getenv("DB_CHANGE_ENEMIES_ID_GROUP")
	dbUsersPars := os.Getenv("DB_USERS_PARSER_GROUP")
	dbChangeUsers := os.Getenv("DB_CHANGE_USERS_ID_GROUP")
	var flag int
	var users_id pq.Int64Array
	if dbChangeUsers == "" {
		log.Println("dbChangeUsers is empty")
		return 500
	}
	if dbParsAdmin == "" {
		log.Println("dbParsAdmin is empty")
		return 500
	}
	if dbChangeEnemies == "" {
		log.Println("dbChangeEnemies is empty")
		return 500
	}
	if dbUsersPars == "" {
		log.Println("dbUsersPars is empty")
		return 500
	}
	var admins_id pq.Int64Array
	err := database.QueryRow(dbParsAdmin, group_id).Scan(&admins_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	for i := 0; i < len(admins_id); i++ {
		if admins_id[i] == int64(user_id) {
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		_, err := database.Exec(dbChangeEnemies, enemies_id, group_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		enemiesInt64 := make([]int64, len(enemies_id))
		for i, v := range enemies_id {
			enemiesInt64[i] = int64(v.(float64))
		}
		err = database.QueryRow(dbUsersPars, group_id).Scan(&users_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		enemies := make(map[int64]bool)
		for _, enemyID := range enemiesInt64 {
			enemies[enemyID] = true
		}

		var filteredUsers []int64
		for _, userID := range users_id {
			if !enemies[userID] {
				filteredUsers = append(filteredUsers, userID)
			}
		}
		_, err = database.Exec(dbChangeUsers, filteredUsers, group_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 408 // Нет прав
	}
}

func ChangeNameInfoGroup(group_id, user_id int, name, info string) int {
	dbParsAdmin := os.Getenv("DB_ADMIN_PARSER_GROUP")
	dbChangeNameInfo := os.Getenv("DB_CHANGE_NAME_INFO_GROUP")
	var flag int
	if dbParsAdmin == "" {
		log.Println("dbParsAdmin is empty")
		return 500
	}
	if dbChangeNameInfo == "" {
		log.Println("dbChangeNameInfo is empty")
		return 500
	}

	var admins_id pq.Int64Array
	err := database.QueryRow(dbParsAdmin, group_id).Scan(&admins_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	for i := 0; i < len(admins_id); i++ {
		if admins_id[i] == int64(user_id) {
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		_, err := database.Exec(dbChangeNameInfo, name, info, group_id)
		if err != nil {
			errorHandler(err)
			return 500
		}
		return 200
	} else {
		return 408 // Нет прав
	}
}
