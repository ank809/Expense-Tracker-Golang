package crud_controllers

import (
	"context"
	"net/http"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid id")
		return
	}
	filter := bson.M{"_id": ID}
	collection_name := "expenses"
	collection := database.OpenCollection(database.Client, collection_name)

	count, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if count.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, "Document not found")
		return
	}
	c.JSON(200, "Document deleted successfully")
}
