package crud_controllers

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DecrementAmount(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		c.JSON(400, err.Error())
		return
	}
	jwt_key := []byte(os.Getenv("JWT_SECRETKEY"))

	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, "No cookie found")
			return
		}
		c.JSON(400, err.Error())
		return
	}
	tokenstring := cookie
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, err
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(400, err.Error())
		return
	}
	if !token.Valid {
		c.JSON(http.StatusBadRequest, "Invalid token")
		return
	}
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
