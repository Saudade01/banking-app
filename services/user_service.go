package services

import (
	"banking-app/models"
	"database/sql"
	"errors"
)

func CreateUser(db *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, user.Username, user.Password)
	return err
}

func AuthenticateUser(db *sql.DB, username, password string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	err := db.QueryRow(query, username, password).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	return user, nil
}
