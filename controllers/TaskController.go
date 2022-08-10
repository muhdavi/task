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
	c.Redirect(http.StatusFound, "/notes")
}

func (TaskController) Index(c *gin.Context) {
	db := database.Instance()

	query := c.DefaultQuery("q", "")
	var notes []models.Task
	db.Find(&notes, query)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data successfully!",
	//	"notes":   notes,
	//})

	c.HTML(http.StatusOK, "notes/index", gin.H{
		"title": "Notes Index",
		"notes": notes,
	})
}

func (TaskController) Detail(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Task
	result := db.First(&note, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"note":    note,
		})
	}
}

func (TaskController) Create(c *gin.Context) {
	db := database.Instance()
	content := c.PostForm("content")

	db.Create(&models.Task{Content: content})

	//c.JSON(http.StatusCreated, gin.H{
	//	"message": "Data created!",
	//})

	c.Redirect(http.StatusFound, "/notes")
}

func (TaskController) Delete(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	db.Delete(&models.Task{}, id)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data deleted!",
	//})

	c.Redirect(http.StatusFound, "/notes")
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

	var note models.Task
	db.Find(&note, id)

	note.IsDone = !note.IsDone

	db.Save(&note)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data updated!",
	//})

	c.Redirect(http.StatusFound, "/notes")
}
