package storage

import (
	"fmt"

	"github.ru/GeorgVartanov/todoProject/database/pgdb"
	"github.ru/GeorgVartanov/todoProject/models"
)

// CreateUser ...
func CreateUser(newUser models.User) (models.UserFromDB, error) {
	var dbUser models.UserFromDB
	sql := newUser.InsetSQLString()
	fmt.Println(sql)

	db := pgdb.GetDB()

	tx := db.MustBegin()
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	// err := tx.QueryRowx(`INSERT INTO app_user (email, password, display_name) VALUES ($1, $2, $3) RETURNING id, email, password, display_name`, newUser.Email, newUser.Password, newUser.DisplayName).StructScan(&dbUser)

	err := tx.QueryRowx(sql).StructScan(&dbUser)

	if err != nil {
		return dbUser, err
	}

	// err = tx.Commit()
	// if err != nil {
	// 	return dbUser, err
	// }

	return dbUser, nil

}
