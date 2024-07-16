package database

import "fmt"

func (db *appdbimpl) GetUserID(username string) (User, error) {
	var user User
	query := fmt.Sprintf(`SELECT u.id, u.username, u.password FROM usersTable u WHERE u.username = '%s'`, username)
	err := db.c.QueryRow(query).Scan(&user.ID, &user.Username, &user.Password)
	// err := db.c.QueryRow(`SELECT u.id, u.username FROM usersTable u WHERE u.username = ?`, username).Scan(&user.ID, &user.Username)
	if err != nil {
		fmt.Println("SQL Error in ExistUsername:", err) // Log the SQL error
		return user, ErrDataDoesNotExist
	}
	return user, nil
}
