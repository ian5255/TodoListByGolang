package router

import (
	"TodoListByGolang/controller"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	// Gin
	r := gin.Default()
	r.GET("/todolist", controller.GetTodoList)
	r.Run() // listen and serve on 0.0.0.0:8080
}
