package controllers

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
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func LoginUser(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	jwt_key := []byte(os.Getenv("JWT_SECRETKEY"))
	var user models.User

	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"username": user.Username}
	collection.FindOne(context.Background(), filter).Decode(&foundUser)
	expected_password := foundUser.Password
	expiration_time := time.Now().Add(time.Minute * 5)
	err := bcrypt.CompareHashAndPassword([]byte(expected_password), []byte(user.Password))
	if err != nil {
		fmt.Println(expected_password)
		fmt.Print(user.Password)
		c.JSON(http.StatusUnauthorized, "Password is incorrect")
		return

	}

	claims := &models.Claims{
		Username: user.Username,
		UserId:   primitive.NewObjectID(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwt_key)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error in generating token")
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenstring,
		Expires: expiration_time,
	})
	c.JSON(200, gin.H{
		"status":  "succcess",
		"message": "User loggined successfully",
		"token":   tokenstring,
	})
}
