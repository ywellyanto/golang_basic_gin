package main

import (
	// "fmt"

	"golang_basic_gin/config"
	"golang_basic_gin/middlewares"
	"golang_basic_gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("golang basic gin")
	config.InitDB()

	// router
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.GET("/", getHome)

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := v1.Group("/departments").Use(middlewares.Auth())
		{
			department.GET("/", routes.GetDepartments)
			department.GET("/:id", routes.GetDepartmentByID)
			department.POST("/", routes.PostDepartments)
			department.PUT("/:id", routes.PutDepartments)
			department.DELETE("/:id", routes.DeleteDepartments)
		}

		position := v1.Group("/positions")
		{
			position.GET("/", routes.GetPositions)
			position.GET("/:id", routes.GetPositionByID)
			position.POST("/", routes.Postpositions)
			position.PUT("/:id", routes.Putpositions)
			position.DELETE("/:id", routes.Deletepositions)
		}

		employee := v1.Group("/employees")
		{
			employee.GET("/", routes.GetEmployees)
			// employee.GET("/:id", routes.GetEmployeeByID)
			// employee.POST("/", routes.PostEmployees)
			// employee.PUT("/:id", routes.PutEmployees)
			// employee.DELETE("/:id", routes.DeleteEmployees)
		}

		inventory := v1.Group("/inventories")
		{
			inventory.GET("/", routes.GetInventory)
			inventory.GET("/:id", routes.GetInventoryByID)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}

		rental := v1.Group("/rental")
		{
			rental.GET("/", routes.GetRental)
			rental.GET("/employee/:id", routes.GetRentalByEmployeeID)
			rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
			rental.POST("/employee", routes.PostRentalByEmployee)
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
