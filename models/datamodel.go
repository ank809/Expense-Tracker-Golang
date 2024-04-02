package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID                primitive.ObjectID `bson:"_id"`
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	TotalAmount       int                `json:"totalamount"`
	InceremtedAmount  int                `json:"incrementedamount"`
	DecrementedAmount int                `json:"decrementedamount"`
	DateTime          primitive.DateTime `json:"date_time"`
}
