package dao

import (
	"database/sql"
	"my_utility_functions/dao_example/models"
)

var user_queries = map[string]string{
	"create_user":   "INSERT INTO users (id, username, password, enabled) VALUES ($1, $2, $3, $4)",
	"get_all_users": "SELECT * FROM users",
	"get_user":      `SELECT * FROM users WHERE id = $1`,
}

func mapRowsToUser(rows *sql.Rows) ([]models.UserDB, error) {
	var users []models.UserDB
	for rows.Next() {
		var user models.UserDB
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Enabled)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func mapRowsToUserFirst(rows *sql.Rows) (*models.UserDB, error) {
	users, err := mapRowsToUser(rows)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}

func CreateUser(tx *sql.Tx, model *models.UserDB) error {
	stmt, err := tx.Prepare(user_queries["create_user"])
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Id, model.Username, model.Password, model.Enabled)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(tx *sql.Tx) ([]models.UserDB, error) {
	stmt, err := tx.Prepare(user_queries["get_all_users"])
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return mapRowsToUser(rows)
}

func GetUser(tx *sql.Tx, id string) (*models.UserDB, error) {
	stmt, err := tx.Prepare(user_queries["get_user"])
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return mapRowsToUserFirst(rows)
}
