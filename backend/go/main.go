package main

import (
	"reading-list/controllers"
	"reading-list/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/readingList/:id", controllers.FindReadingList)
	r.POST("/readingList", controllers.CreateReadingList)
	r.PATCH("/readingList/:id", controllers.UpdateReadingList)

	r.POST("/readingListItem", controllers.CreateReadingListItem)
	r.GET("/readingListItem/:id", controllers.FindReadingListItem)
	r.PATCH("/readingListItem/:id", controllers.UpdateReadingListItem)

	r.POST("/user", controllers.CreateUser)
	r.POST("/user/restore/:id", controllers.RestoreDeletedUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.GET("/user/:id", controllers.FindUser)
	r.DELETE("/user/:id", controllers.SoftDeleteUser)
	r.DELETE("/user/HardDelete/:id", controllers.HardDeleteUser)

	r.Run()
}
