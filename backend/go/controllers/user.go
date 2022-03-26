package controllers

import (
	"net/http"
	"reading-list/models"
	"reading-list/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Get /user/:id
func FindUser(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user, err = service.FindUser(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//Post /user/create
func CreateUser(c *gin.Context) {

	// validate input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create User
	user := service.CreateUser(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//PATCH /user/update/:id
func UpdateUser(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user, err = service.FindUser(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var input models.UpdateUserInput

	//validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = service.UpdateUser(user, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//DELETE user/:id
//Delete a book
func SoftDeleteUser(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user, err = service.FindUser(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.SoftDeleteUser(user)

	var payload string

	if err == nil {
		payload = "User deleted successfully"
	} else {
		payload = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{"data": payload})
}

//POST user/:id
// Restore Soft Deleted User
func RestoreDeletedUser(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user, err = service.FindUser(uuid)

	if err.Error() != "User is deleted" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = service.RestoreUser(user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

//DELETE user/:id
//Delete a book
func HardDeleteUser(c *gin.Context) {

	var uuid, uuidErr = uuid.Parse(c.Param("id"))

	if uuidErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user, err = service.FindUser(uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	err = service.HardDeleteUser(user)

	var payload string

	if err == nil {
		payload = "User deleted successfully"
	} else {
		payload = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{"data": payload})
}
