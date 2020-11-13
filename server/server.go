package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todoProject/controllers"
)

// StartServer ...
func StartServer() {
	r := gin.Default()
	MyAPI := r.Group("/api/")
	UserRouter := MyAPI.Group("/user/")
	{
		UserRouter.POST("/create/", controllers.CreateUserController)
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
