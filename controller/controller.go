package controller

import (
	"TodoListByGolang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodoList(c *gin.Context) {
	result, err := service.GetList()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	// jsondata, _ := json.Marshal(result)
	c.JSON(http.StatusOK, result)
}
