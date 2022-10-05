package main

import (
	"build-api-go/config"
	"build-api-go/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	fmt.Println(db)
	inDB := controllers.New(db)

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
}
