package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) ExistUsername(username string, password string) error {
	var user User
	query := fmt.Sprintf(`SELECT u.id, u.username, u.password FROM usersTable u WHERE u.username = '%s' AND u.password = '%s'`, username, password)
	err := db.c.QueryRow(query).Scan(&user.ID, &user.Username, &user.Password)
	// err := db.c.QueryRow(`SELECT u.id, u.username, u.password FROM usersTable u WHERE u.username = ? AND u.password = ?`, username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrDataDoesNotExist
		}
	}
	return nil
}
