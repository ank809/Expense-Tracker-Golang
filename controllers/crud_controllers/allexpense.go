package crud_controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

func GetAllExpenses(c *gin.Context) {
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
	var expenses []models.Data

	collection_name := "expenses"
	filter := bson.M{"username": claims.Username}
	collection := database.OpenCollection(database.Client, collection_name)
	expenseCursor, err := collection.Find(context.Background(), filter)
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
