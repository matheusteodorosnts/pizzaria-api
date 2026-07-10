package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/matheusteodorosnts/pizzaria-api/models"
)

func CreatePizzaHandler(ctx *gin.Context) {
	request := &CreatePizzaRequest{}

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pizza := &models.Pizza{
		ID:          uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	if err := db.Create(&pizza).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, pizza)
}
