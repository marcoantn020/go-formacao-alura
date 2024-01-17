package routes

import (
	"gin/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func HandlerRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/students", controller.ShowAllStudents)
	r.POST("/students", controller.Create)
	r.GET("/students/:id", controller.FindById)
	r.PUT("/students/:id", controller.Update)
	r.DELETE("/students/:id", controller.Delete)
	r.GET("/students/cpf/:cpf", controller.FindByCPF)

	r.GET("/home", controller.ShowPageIndex)
	r.NoRoute(controller.PageNotFound)

	err := r.Run()
	if err != nil {
		log.Panicln("Error in server", err)
	}
}
