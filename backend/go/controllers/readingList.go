package controllers

import (
	"net/http"
	"reading-list/models"
	"reading-list/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET Reading List
// readingList/:id
func FindReadingList(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	readingList, err := service.FindReadingList(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": readingList})
}

// POST CreateReadingList
func CreateReadingList(c *gin.Context) {

	// validate input
	var input models.CreateReadingListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readingList := service.CreateReadingList(input)

	c.JSON(http.StatusOK, gin.H{"data": readingList})

}

//PATCH readingList/:id
func UpdateReadingList(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var readingList, err = service.FindReadingList(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reading List Item not found"})
		return
	}

	var input models.UpdateReadingListInput

	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readingList, err = service.UpdateReadingList(readingList, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": readingList})
}
