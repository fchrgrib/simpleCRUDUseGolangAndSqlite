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
	r.GET("/users/?age", models.GetAgeUsers)
	//r.GET("/users/:name", models.GetUsers)

	r.POST("/users/:id", models.CreateUser)
	r.POST("/users/location/:coordinate-Location", models.CreateLocation)
	r.POST("/users/school/:coordinate-School", models.CreateSchool)

	r.PATCH("/users/:id", models.UpdateUsers)
	r.PATCH("/users/location/:coordinate-Location", models.UpdateLocation)
	r.PATCH("/users/school/:coordinate-School", models.UpdateSchool)

	r.DELETE("/users/:id", models.DeleteUsers)
	r.DELETE("/users/location/:coordinate-Location", models.DeleteLocation)
	r.DELETE("/users/school/:coordinate-School", models.DeleteSchool)

	if err := r.Run("localhost:7788"); err != nil {
		panic(err)
	}
}
