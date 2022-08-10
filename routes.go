package main

import (
	"github.com/gin-gonic/gin"
	"taskgo/controllers"
)

func Routes(r *gin.Engine) {
	taskController := controllers.TaskController{}

	r.GET("/", taskController.Root)

}
