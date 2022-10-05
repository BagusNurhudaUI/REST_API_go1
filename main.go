package main

import (
	"build-api-go/config"
)

func main() {
	config.DBInit()
	// db := config.DBInit()
	// fmt.Println(db)
	// inDB := &controllers.InDB{DB: db}

	// router := gin.Default()

	// router.GET("/person/:id", inDB.GetPerson)
}
