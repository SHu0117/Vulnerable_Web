package database

import "fmt"

func (db *appdbimpl) NewUser(u User) (User, error) {
	query := fmt.Sprintf("INSERT INTO usersTable (username, password) VALUES ('%s', '%s')", u.Username, u.Password)
	res, err := db.c.Exec(query)
	if err != nil {
		return u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = uint64(lastInsertID)
	return u, nil
}
