package router

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusteodorosnts/pizzaria-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.POST("/pizzas", handler.CreatePizzaHandler)
	v1.GET("/pizzas", handler.ListPizzasHandler)
	v1.DELETE("/pizzas/:id", handler.DeletePizzaHandler)
}
