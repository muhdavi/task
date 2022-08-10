package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskgo/database"
	"taskgo/models"
)

type TaskController struct{}

func (TaskController) Root(c *gin.Context) {
	c.Redirect(http.StatusFound, "/tasks")
}

func (TaskController) Index(c *gin.Context) {
	db := database.Instance()

	query := c.DefaultQuery("q", "")
	var tasks []models.Task
	db.Find(&tasks, query)

	c.HTML(http.StatusOK, "tasks/index", gin.H{
		"title": "Tasks Index",
		"tasks": tasks,
	})
}

func (TaskController) Detail(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	result := db.First(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"task":    task,
		})
	}
}

func (TaskController) Create(c *gin.Context) {
	db := database.Instance()
	content := c.PostForm("content")
	person := c.PostForm("person")
	due_date := c.PostForm("duedate")

	db.Create(&models.Task{Content: content, Person: person, DueDate: due_date})

	c.Redirect(http.StatusFound, "/tasks")
}

func (TaskController) Delete(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	db.Delete(&models.Task{}, id)

	c.Redirect(http.StatusFound, "/tasks")
}

func (TaskController) Update(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")

	db.Model(&models.Task{}).
		Where("id  = ?", id).
		Update("content", content)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated!",
	})
}

func (TaskController) Done(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	db.Find(&task, id)

	task.IsDone = !task.IsDone

	db.Save(&task)

	c.Redirect(http.StatusFound, "/tasks")
}
