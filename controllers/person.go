package controllers

import (
	"build-api-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id =?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (db *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	db.DB.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	person.First_Name = first_name
	person.Last_Name = last_name

	db.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := db.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	newPerson.First_Name = firstName
	newPerson.Last_Name = lastName

	err = db.DB.Model(&person).Updates(newPerson).Error

	if err != nil {
		result = gin.H{
			"result": "Update Failed",
		}
	} else {
		result = gin.H{
			"result": "Successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")

	err := db.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = db.DB.Delete(&person).Error

	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
