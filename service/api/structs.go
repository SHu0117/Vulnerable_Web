package api

import (
	"github.com/SHu0117/WASA-Photo/service/database"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) UserFromDatabase(user database.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Password = user.Password
}

func (u *User) UserToDatabase() database.User {
	return database.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}
