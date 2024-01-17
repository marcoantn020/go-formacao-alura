package main

import (
	"bytes"
	"encoding/json"
	"gin/controller"
	"gin/database"
	"gin/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := model.Student{
		Name: "Student test",
		CPF:  "09876543211",
		RG:   "123456789",
	}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student model.Student
	database.DB.Delete(&student, ID)
}

func TestListAllStudents(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students", controller.ShowAllStudents)
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByCpf(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students/cpf/:cpf", controller.FindByCPF)
	request, _ := http.NewRequest("GET", "/students/cpf/09876543211", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByID(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students/:id", controller.FindById)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	var studentMock model.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Student test", studentMock.Name)
	assert.Equal(t, "09876543211", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
}

func TestDeleteStudentById(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	r := SetupRoutesTest()
	r.DELETE("/students/:id", controller.Delete)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.PUT("/students/:id", controller.Update)
	studentToUpdate := model.Student{
		Name: "Student updated",
		CPF:  "01020304050",
		RG:   "111111111",
	}
	convertStudentToJson, _ := json.Marshal(studentToUpdate)
	path := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PUT", path, bytes.NewBuffer(convertStudentToJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var studentMock model.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Student updated", studentMock.Name)
	assert.Equal(t, "01020304050", studentMock.CPF)
	assert.Equal(t, "111111111", studentMock.RG)
}
