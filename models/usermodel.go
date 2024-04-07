package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type Claims struct {
	Username string             `json:"username"`
	UserId   primitive.ObjectID `json:"userid"`
	jwt.StandardClaims
}
