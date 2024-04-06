package crud_controllers

import (
	"context"

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
		c.JSON(400, err.Error())
		return
	}
	if err := expenseCursor.All(context.Background(), &expenses); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, expenses)
}
