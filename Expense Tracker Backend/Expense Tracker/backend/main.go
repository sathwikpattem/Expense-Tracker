package main

import (
	"github.com/expense-tracker/backend/database"
	"github.com/expense-tracker/backend/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	// Create Gin router
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // In production, specify your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Serve static files (your frontend)
	r.Static("/css", "../css")
	r.Static("/js", "../js")
	r.LoadHTMLGlob("../templates/*")

	// Serve the main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "expense_tracker.html", gin.H{
			"title": "Expense Tracker",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Expense routes
		api.GET("/expenses", handlers.GetExpenses)
		api.POST("/expenses", handlers.CreateExpense)
		api.DELETE("/expenses/:id", handlers.DeleteExpense)

		// Category routes
		api.GET("/categories", handlers.GetCategories)
		api.POST("/categories", handlers.CreateCategory)

		// Summary route
		api.GET("/summary", handlers.GetExpenseSummary)
		
		// Analytics routes
		api.GET("/analytics", handlers.GetExpenseAnalytics)
		api.GET("/analytics/monthly", handlers.GetMonthlyExpenses)
		api.GET("/analytics/weekly", handlers.GetWeeklyExpenses)
	}

	log.Println("Server starting on :8080")
	log.Println("Frontend available at: http://localhost:8080")
	log.Println("API available at: http://localhost:8080/api")
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
