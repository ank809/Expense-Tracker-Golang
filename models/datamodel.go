package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `json:"username"`
	TotalAmount int                `json:"totalamount"`
	Expenses    []int              `json:"expense"`
	DateTime    primitive.DateTime `json:"date_time"`
}
