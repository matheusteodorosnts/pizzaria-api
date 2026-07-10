package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusteodorosnts/pizzaria-api/models"
)

func ListPizzasHandler(ctx *gin.Context) {
	pizzas := []models.Pizza{}

	if err := db.Find(&pizzas).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pizzas)
}
