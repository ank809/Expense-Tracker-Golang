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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DeleteExpense(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
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
	tokenstr := cookie
	claims := models.Claims{}

	token, err := jwt.ParseWithClaims(tokenstr, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, err
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if !token.Valid {
		c.JSON(http.StatusBadRequest, "Token is invalid")
		return
	}

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
	// if err := collection.FindOneAndDelete(context.Background(), filter); err != nil {
	// 	c.JSON(400, err)
	// 	return
	// }
	// c.JSON(200, "Deleted successfuly")

}
