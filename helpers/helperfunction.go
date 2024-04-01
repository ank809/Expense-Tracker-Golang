package helpers

import (
	"context"
	"fmt"
	"unicode"

	"gopkg.in/mgo.v2/bson"

	emailVerifier "github.com/AfterShip/email-verifier"
	"github.com/ank809/Expense-Tracker-Golang/database"
)

func CheckUsername(name string) (bool, string) {
	if name == "" {
		return false, "Name cannot be empty"
	}
	if len(name) < 3 {
		return false, "Length of name should be greater than 3 "
	}
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"username": name}
	countDocs, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, fmt.Sprintf("Error counting documents: %v", err)
	}
	if countDocs > 0 {
		return false, "This username is already taken, choose a new one"
	}
	return true, "OK"
}

var (
	verifer = emailVerifier.NewVerifier()
)

func VerifyEmail(email string) (bool, string) {
	ret, err := verifer.Verify(email)
	if err != nil {
		return false, "verify email address failed"
	}
	if !ret.Syntax.Valid {
		return false, "email address syntax is invalid "
	}
	return true, "email is valid"
}

func CheckPassword(password string) (bool, string) {
	containsUpper := false
	containsLower := false
	containsDigits := false
	containsSpecial := false

	if password == "" {
		return false, "Password cannot be empty"
	}
	if len(password) < 8 {
		return false, "Password length should be greater than 8"
	}
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			containsUpper = true
		} else if unicode.IsLower(ch) {
			containsLower = true
		} else if unicode.IsDigit(ch) {
			containsDigits = true
		} else {
			containsSpecial = true
		}

	}
	if containsDigits && containsLower && containsSpecial && containsUpper {
		return true, "OK"
	} else {
		return false, "Password should have uppercase, lowercase, digit and special character"
	}

}
