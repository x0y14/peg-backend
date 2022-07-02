package db

import "database/sql"

func GetUserName(db *sql.DB, userId string) (string, error) {
	prep, err := db.Prepare("SELECT user_id FROM peg.user_names where user_id = ?")
	if err != nil {
		return "", err
	}
	defer prep.Close()

	var userName string
	err = prep.QueryRow(userId).Scan(&userName)

	return userName, err
}
