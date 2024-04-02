package crud_controllers

import (
	"context"
	"net/http"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllExpenses(c *gin.Context) {
	var expenses []models.Data

	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)

	expenseCursor, err := collection.Find(context.Background(), collection)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := expenseCursor.All(context.Background(), &expenses); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, expenses)
}