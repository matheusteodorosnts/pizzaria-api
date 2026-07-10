package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusteodorosnts/pizzaria-api/models"
)

func DeletePizzaHandler(ctx *gin.Context) {
	pizzaId := ctx.Param("id")
	if pizzaId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "pizzaId is required"})
		return
	}

	pizza := models.Pizza{}

	// Find pizza by ID
	if err := db.First(&pizza, "id = ?", pizzaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "pizza not found"})
		return
	}

	// Delete pizza
	if err := db.Delete(&pizza).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
