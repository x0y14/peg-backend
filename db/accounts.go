package db

import (
	typesv1 "backend/gen/types/v1"
	"database/sql"
)

func CreateAccount(db *sql.DB, userId string, email string) (*typesv1.Account, error) {
	prep, err := db.Prepare("INSERT INTO accounts (user_id, email) values (?, ?)")
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	_, err = prep.Exec(userId, email)
	if err != nil {
		return nil, err
	}

	return GetAccount(db, userId)
}

func GetAccount(db *sql.DB, userId string) (*typesv1.Account, error) {
	// usernameは別のテーブルにあるので、結合を行ってます。
	prep, err := db.Prepare("SELECT email, user_name from accounts left join user_names on accounts.user_id = user_names.user_id where accounts.user_id = ?")
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	var email string
	var userNameNullable sql.NullString
	err = prep.QueryRow(userId).Scan(&email, &userNameNullable)

	var userName string
	if userNameNullable.Valid {
		userName = userNameNullable.String
	}

	return &typesv1.Account{
		UserId:   userId,
		Email:    email,
		UserName: userName,
	}, err
}

func IsAccountExists(db *sql.DB, userId string) bool {
	_, err := GetAccount(db, userId)
	if err != nil {
		return false
	}
	return true
}
