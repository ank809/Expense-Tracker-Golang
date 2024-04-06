package crud_controllers

import (
	"context"
	"strconv"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IncrementAmount(c *gin.Context) {
	var expense models.Data
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, "Invalid id")
		return
	}
	filter := bson.M{"_id": objectId}
	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)
	if err := collection.FindOne(context.Background(), filter).Decode(&expense); err != nil {
		c.JSON(400, err.Error())
		return
	}
	value := c.Query("value")
	amount, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(400, "Invalid amount")
		return
	}
	expense.TotalAmount = expense.TotalAmount + amount
	expense.Expenses = append(expense.Expenses, amount)
	update := bson.M{"$set": bson.M{"expenses": expense.Expenses, "totalamount": expense.TotalAmount}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Expense added")
}
