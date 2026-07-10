package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusteodorosnts/pizzaria-api/models"
)

func UpdatePizzaHandler(ctx *gin.Context) {
	request := &UpdatePizzaRequest{}
	pizzaId := ctx.Param("id")
	pizza := &models.Pizza{}

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find pizza by ID
	if err := db.First(&pizza, "id = ?", pizzaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "pizza not found"})
		return
	}

	if request.Name != "" {
		pizza.Name = request.Name
	}

	if request.Description != "" {
		pizza.Description = request.Description
	}

	if request.Price >= 0 {
		pizza.Price = request.Price
	}

	// Update pizza
	if err := db.Save(&pizza).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}
