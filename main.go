package main

import (
	"github.ru/GeorgVartanov/todoProject/database/pgdb"
	"github.ru/GeorgVartanov/todoProject/server"
)

func main() {
	server.StartServer()
	db := pgdb.GetDB()
	defer db.Close()
}

// // Person ...
// type Person struct {
// 	ID          uint64 `json:"id"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`
// 	DisplayName string `json:"display_name"`
// }

// func main() {
// 	person := &Person{
// 		ID:          1,
// 		Email:       "admin@gmail.com",
// 		Password:    "password",
// 		DisplayName: "TOD",
// 	}

// 	val := reflect.ValueOf(person).Elem()
// 	fmt.Println(val.NumField())
// 	for i := 0; i < val.NumField(); i++ {
// 		fmt.Println(i)
// 		if val.NumField()-i == 1 {
// 			fmt.Println("last")

// 			valueField := val.Field(i)
// 			typeField := val.Type().Field(i)
// 			tag := typeField.Tag

// 			fmt.Printf("%s, %v, %s \n", typeField.Name, valueField.String(), tag.Get("json"))

// 		}

// 	}
// }
