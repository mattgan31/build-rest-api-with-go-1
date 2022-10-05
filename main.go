package main

import (
	"praktek1/config"
	"praktek1/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := config.DBInit()
	// InDB := &controller.InDB{DB: db}
	ctr := controller.New(db)

	router := gin.Default()

	router.GET("/person/:id", ctr.GetPerson)
	router.GET("/persons", ctr.GetPersons)
	router.POST("/person", ctr.CreatePerson)
	router.PUT("/person", ctr.UpdatePerson)
	router.DELETE("/person/:id", ctr.DeletePerson)
	router.Run(":3000")
}
