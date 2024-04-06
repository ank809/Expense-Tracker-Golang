package crud_controllers

import (
	"context"
	"strconv"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DecrementAmount(c *gin.Context) {
	var expense models.Data
	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, "Invalid Id")
		return
	}
	filter := bson.M{"_id": objectId}

	if err := collection.FindOne(context.Background(), filter).Decode(&expense); err != nil {
		c.JSON(400, err.Error())
		return
	}

	value := c.Query("value")
	amount, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	expense.TotalAmount -= amount
	expense.Expenses = append(expense.Expenses, amount)
	update := bson.M{"$set": bson.M{"totalamount": expense.TotalAmount, "expenses": expense.Expenses}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "Document updated")
}
