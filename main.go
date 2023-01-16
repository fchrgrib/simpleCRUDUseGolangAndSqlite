package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	models "github.com/gym/CRUD"
	_ "github.com/mattn/go-sqlite3"
	_ "log"
)

func main() {

	r := gin.Default()
	r.GET("/users", models.GetUsersDetail)
	r.GET("/users/:name", models.GetUsers)

	if err := r.Run("localhost:7788"); err != nil {
		panic(err)
	}
}
