package entity

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// User : user entity
type User struct {
	ID          int        `json:"id"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Fullname    string     `json:"fullname"`
	TSLastLogin int32      `json:"ts_last_login"`
	TSCreate    int32      `json:"ts_create"`
	TSUpdate    int32      `json:"ts_update"`
	Permission  json.RawMessage `json:"permission"`
}

// GetUserByEmail : returns single user
func GetUserByEmail(email string) *User {
	user := new(User)
	row := DB.QueryRow("SELECT * from user WHERE email=?", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Fullname, &user.TSLastLogin, &user.TSCreate, &user.TSUpdate, &user.Permission)
	if err != nil {
		log.Errorln("User SELECT by Email Err: ", err)
		return nil
	}
	return user
}

// GetUserByID : returns single user
func GetUserByID(id int64) *User {
	user := new(User)
	row := DB.QueryRow("SELECT * from user WHERE id=?", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Fullname, &user.TSLastLogin, &user.TSCreate, &user.TSUpdate, &user.Permission)
	if err != nil {
		log.Errorln("User SELECT by ID Err: ", err)
		return nil
	}
	return user
}

// AddUser : insert single user
func AddUser(user User) (int, error) {
	var lastInsert int64
	insForm, err := DB.Prepare("INSERT INTO user (email, password, fullname, ts_last_login, ts_create, ts_update, permission) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	result, err := insForm.Exec(user.Email, user.Password, user.Fullname, user.TSLastLogin, user.TSCreate, user.TSUpdate, user.Permission)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	lastInsert, err = result.LastInsertId()
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return int(lastInsert), err
}
