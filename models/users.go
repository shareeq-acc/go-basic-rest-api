package models

import (
	"api/db"
	"api/utils"
	"errors"
)

type User struct {
	ID       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `
	INSERT INTO users(name, email, password)
	VALUES(?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := statement.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM USERS WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var hashedPasword string
	err := row.Scan(&u.ID, &hashedPasword)
	if err != nil {
		return err
	}

	isCorrectPassword := utils.CheckPassword(hashedPasword, u.Password)

	if !isCorrectPassword {
		return errors.New("invalid credentials")
	}

	return nil
}
