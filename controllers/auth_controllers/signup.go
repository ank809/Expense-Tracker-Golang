package controllers

import (
	"context"
	"net/http"

	"github.com/ank809/Expense-Tracker-Golang/database"
	"github.com/ank809/Expense-Tracker-Golang/helpers"
	"github.com/ank809/Expense-Tracker-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.ID = primitive.NewObjectID()
	ok, value := helpers.CheckUsername(user.Username)
	if !ok {
		c.JSON(http.StatusBadRequest, value)
		return
	}
	ok, value = helpers.VerifyEmail(user.Email)
	if !ok {
		c.JSON(http.StatusBadRequest, value)
		return
	}
	ok, value = helpers.CheckPassword(user.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, value)
		return
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.Password = string(hashed_password)
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	collection.InsertOne(context.Background(), user)
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "User added successfully",
	})
}
