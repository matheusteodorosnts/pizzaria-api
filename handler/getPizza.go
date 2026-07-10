package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusteodorosnts/pizzaria-api/models"
)

func GetPizzaHandler(ctx *gin.Context) {
	pizzaId := ctx.Param("id")
	if pizzaId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "pizzaId is required"})
		return
	}

	pizza := &models.Pizza{}

	if err := db.First(&pizza, "id = ?", pizzaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "pizza not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pizza": pizza})
}
