package main

import (
	"github.com/gin-gonic/gin"
	"taskgo/controllers"
)

func Routes(r *gin.Engine) {
	taskController := controllers.TaskController{}

	r.GET("/", taskController.Root)
	r.GET("/tasks", taskController.Index)
	r.POST("/tasks", taskController.Create)
	r.GET("/tasks/:id", taskController.Detail)
	r.POST("/tasks/:id/edit", taskController.Update)
	r.GET("/tasks/:id/edit", taskController.Edit)
	r.GET("/tasks/:id/done", taskController.Done)
	r.GET("/tasks/:id/delete", taskController.Delete)
}
