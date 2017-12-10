package main

import (
	"testApi/organisations"
	"testApi/storage"
	"testApi/users"

	"github.com/gin-gonic/gin"
)

func main() {

	db := storage.New()
	db.Init()
	organisations.SetDatabase(db)
	users.SetDatabase(db)

	router := gin.Default()
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/organisation/", organisations.GetAll)
		v1.GET("/organisation/:id/", organisations.Get)
		v1.POST("/organisation", organisations.Create)
		v1.GET("/organisation/:id/divisions", organisations.GetDivision)
		v1.GET("/users", users.GetAll)
		v1.GET("/users/:id", users.Get)
		v1.POST("/users", users.Create)
	}
	router.Run()

}
