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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	expense.ID = primitive.NewObjectID()
	expense.DateTime = primitive.DateTime(time.Now().Unix())

	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)
	collection.InsertOne(context.Background(), expense)
	c.JSON(200, "Expense Successfully Saved")
}
