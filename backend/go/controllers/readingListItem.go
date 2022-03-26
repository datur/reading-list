package controllers

import (
	"net/http"
	"reading-list/models"
	"reading-list/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /readingListItem/:id
func FindReadingListItem(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	readingListItem, err := service.FindReadingListItem(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": readingListItem})
}

// POST CreateReadingListitem
func CreateReadingListItem(c *gin.Context) {
	var input models.CreateReadingListItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readingListItem := service.CreateReadingListItem(input)

	c.JSON(http.StatusOK, gin.H{"data": readingListItem})
}

//PATCH readingListItem/:id
func UpdateReadingListItem(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var readingListItem, err = service.FindReadingListItem(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reading List Item not found"})
		return
	}

	var input models.UpdateReadingListItemInput

	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readingListItem, err = service.UpdateReadingListItem(readingListItem, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": readingListItem})
}
