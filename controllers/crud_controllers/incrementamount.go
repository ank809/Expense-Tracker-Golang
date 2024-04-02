package crud_controllers

import (
	"context"
	"time"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func InceremteAmount(c *gin.Context) {
	var expense models.Data
	if err := c.BindJSON(&expense); err != nil {
		c.JSON(400, err)
		return
	}
	var foundExpenses models.Data
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, "Invalid id")
		return
	}
	filter := bson.M{"_id": objectID}
	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)
	err = collection.FindOne(context.Background(), filter).Decode(&foundExpenses)
	if err != nil {
		c.JSON(400, err)
		return
	}

	foundExpenses.InceremtedAmount = expense.InceremtedAmount
	foundExpenses.TotalAmount = foundExpenses.TotalAmount + foundExpenses.InceremtedAmount
	foundExpenses.DateTime = primitive.DateTime(time.Now().Unix())

	update := bson.M{"$set": foundExpenses}
	count, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		c.JSON(400, err)
		return
	}
	if count.MatchedCount == 0 {
		c.JSON(400, "Document Not found")
		return
	}
	c.JSON(200, "Document updated successfully")

}
