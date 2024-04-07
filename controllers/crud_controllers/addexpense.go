package crud_controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddExpense(c *gin.Context) {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}
	jwt_key := []byte(os.Getenv("JWT_SECRETKEY"))
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(400, err.Error())
		return
	}
	tokenstr := cookie
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, err
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if !token.Valid {
		c.JSON(http.StatusBadRequest, "Token is invalid")
		return
	}

	var expense models.Data

	if err := c.BindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.DateTime = primitive.DateTime(time.Now().Unix())
	expense.ID = claims.UserId
	expense.Username = claims.Username

	collectionName := "expenses"
	collection := database.OpenCollection(database.Client, collectionName)
	_, err = collection.InsertOne(context.Background(), expense)
	if err != nil {
		c.JSON(400, "Error in saving data")
		return
	}
	c.JSON(200, "Document saved successfully")

}
