package main

import (
	controllers "github.com/ank809/Expense-Tracker-Golang/controllers/auth_controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", HomeScreen)
	router.GET("/signup", controllers.Signup)
	router.Run(":8081")
}

func HomeScreen(c *gin.Context) {
	c.JSON(200, "Welcome to Expense Tracker")

}
