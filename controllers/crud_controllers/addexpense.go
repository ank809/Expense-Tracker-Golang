package crud_controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddExpense(c *gin.Context) {
	var expense models.Data

	if err := c.BindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.DateTime = primitive.DateTime(time.Now().Unix())
	expense.ID = primitive.NewObjectID()

	collectionName := "expenses"
	collection := database.OpenCollection(database.Client, collectionName)
	_, err := collection.InsertOne(context.Background(), expense)
	if err != nil {
		c.JSON(400, "Error in saving data")
		return
	}
	c.JSON(200, "Document saved successfully")

}
