package main

import (
	controllers "github.com/ank809/Expense-Tracker-Golang/controllers/auth_controllers"
	"github.com/ank809/Expense-Tracker-Golang/controllers/crud_controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", HomeScreen)
	router.GET("/signup", controllers.Signup)
	router.GET("/login", controllers.LoginUser)
	router.GET("/addexpense", crud_controllers.AddExpense)
	router.GET("/delete/:id", crud_controllers.DeleteExpense)
	router.GET("allexpenses", crud_controllers.GetAllExpenses)
	router.GET("increment/:id", crud_controllers.IncrementAmount)
	router.GET("decrement/:id", crud_controllers.DecrementAmount)
	router.Run(":8081")
}

func HomeScreen(c *gin.Context) {
	c.JSON(200, "Welcome to Expense Tracker")
}
