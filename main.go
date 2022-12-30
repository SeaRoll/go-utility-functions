package main

import (
	"context"
	"my_utility_functions/dao_example/dao"
	"my_utility_functions/dao_example/models"
	"my_utility_functions/migrator"
)

func main() {
	db := migrator.ConnectToSqlite("test.db")
	migrator.RunMigration(db)

	newUser := &models.UserDB{
		Id:       "4",
		Username: "test",
		Password: "password",
		Enabled:  true,
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	err = dao.CreateUser(tx, newUser)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	users, err := dao.GetAllUsers(tx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		println(user.Id, user.Username, user.Password, user.Enabled)
	}
}
