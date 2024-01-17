package controller

import (
	"gin/database"
	"gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAllStudents(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func Create(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.ValidateStudents(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSONP(http.StatusCreated, student)
}

func FindById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student model.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var student model.Student
	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, nil)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var student model.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.ValidateStudents(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func FindByCPF(c *gin.Context) {
	cpf := c.Param("cpf")

	var student model.Student
	database.DB.Where(&model.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ShowPageIndex(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func PageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
