package handlers

import (
	"github.com/expense-tracker/backend/database"
	"github.com/expense-tracker/backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetExpenses retrieves all expenses
func GetExpenses(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, amount, category, date, note, created FROM expenses ORDER BY created DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.Amount, &expense.Category, &expense.Date, &expense.Note, &expense.Created)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan expense"})
			return
		}
		expenses = append(expenses, expense)
	}

	c.JSON(http.StatusOK, expenses)
}

// CreateExpense creates a new expense
func CreateExpense(c *gin.Context) {
	var req models.ExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO expenses (amount, category, date, note) VALUES (?, ?, ?, ?)",
		req.Amount, req.Category, req.Date, req.Note,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Expense created successfully"})
}

// DeleteExpense deletes an expense by ID
func DeleteExpense(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	result, err := database.DB.Exec("DELETE FROM expenses WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

// GetCategories retrieves all categories
func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name FROM categories ORDER BY name")
	if err != nil {
		log.Printf("Error querying categories: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Printf("Error scanning category: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan category"})
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
}

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec("INSERT INTO categories (name) VALUES (?)", req.Name)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Category created successfully"})
}

// GetExpenseSummary returns expense summary by category
func GetExpenseSummary(c *gin.Context) {
	// First, let's try a simpler approach - get all categories first
	rows, err := database.DB.Query("SELECT name FROM categories ORDER BY name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var summaries []models.ExpenseSummary
	for rows.Next() {
		var categoryName string
		err := rows.Scan(&categoryName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan category"})
			return
		}

		// Get total for this category
		var total float64
		err = database.DB.QueryRow(
			"SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE category = ?", 
			categoryName,
		).Scan(&total)
		if err != nil {
			// If no expenses found, total should be 0
			total = 0
		}

		summaries = append(summaries, models.ExpenseSummary{
			Category: categoryName,
			Total:    total,
		})
	}

	c.JSON(http.StatusOK, summaries)
}

// GetMonthlyExpenses returns monthly expense data for the last 12 months
func GetMonthlyExpenses(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			strftime('%Y-%m', date) as month,
			strftime('%Y-%m', date) as month_name,
			SUM(amount) as total,
			COUNT(*) as count
		FROM expenses 
		WHERE date >= date('now', '-12 months')
		GROUP BY strftime('%Y-%m', date)
		ORDER BY month DESC
	`)
	if err != nil {
		log.Printf("Error querying monthly expenses: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monthly expenses"})
		return
	}
	defer rows.Close()

	var monthlyExpenses []models.MonthlyExpense
	for rows.Next() {
		var expense models.MonthlyExpense
		err := rows.Scan(&expense.Month, &expense.MonthName, &expense.Total, &expense.Count)
		if err != nil {
			log.Printf("Error scanning monthly expense: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan monthly expense"})
			return
		}
		
		// Format month name (e.g., "2025-08" -> "August 2025")
		if len(expense.Month) >= 7 {
			monthNames := []string{"", "January", "February", "March", "April", "May", "June",
				"July", "August", "September", "October", "November", "December"}
			year := expense.Month[:4]
			monthNum := expense.Month[5:7]
			if monthNum[0] == '0' {
				monthNum = monthNum[1:]
			}
			if monthIndex := parseMonth(monthNum); monthIndex > 0 && monthIndex <= 12 {
				expense.MonthName = monthNames[monthIndex] + " " + year
			}
		}
		
		monthlyExpenses = append(monthlyExpenses, expense)
	}

	c.JSON(http.StatusOK, monthlyExpenses)
}

// GetWeeklyExpenses returns weekly expense data for the last 7 weeks
func GetWeeklyExpenses(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			strftime('%Y-W%W', date) as week,
			date(date, 'weekday 0', '-6 days') as week_start,
			date(date, 'weekday 0') as week_end,
			SUM(amount) as total,
			COUNT(*) as count
		FROM expenses 
		WHERE date >= date('now', '-7 weeks')
		GROUP BY strftime('%Y-W%W', date)
		ORDER BY week DESC
	`)
	if err != nil {
		log.Printf("Error querying weekly expenses: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weekly expenses"})
		return
	}
	defer rows.Close()

	var weeklyExpenses []models.WeeklyExpense
	for rows.Next() {
		var expense models.WeeklyExpense
		var weekStart, weekEnd string
		err := rows.Scan(&expense.Week, &weekStart, &weekEnd, &expense.Total, &expense.Count)
		if err != nil {
			log.Printf("Error scanning weekly expense: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan weekly expense"})
			return
		}
		
		// Format week range (e.g., "Aug 11-17, 2025")
		expense.WeekRange = formatWeekRange(weekStart, weekEnd)
		
		weeklyExpenses = append(weeklyExpenses, expense)
	}

	c.JSON(http.StatusOK, weeklyExpenses)
}

// GetExpenseAnalytics returns comprehensive expense analytics
func GetExpenseAnalytics(c *gin.Context) {
	var analytics models.ExpenseAnalytics

	// Get total expenses
	err := database.DB.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM expenses").Scan(&analytics.TotalExpenses)
	if err != nil {
		log.Printf("Error getting total expenses: %v", err)
		analytics.TotalExpenses = 0
	}

	// Get average daily expenses (last 30 days)
	err = database.DB.QueryRow(`
		SELECT COALESCE(AVG(daily_total), 0) FROM (
			SELECT SUM(amount) as daily_total 
			FROM expenses 
			WHERE date >= date('now', '-30 days')
			GROUP BY date
		)
	`).Scan(&analytics.AverageDaily)
	if err != nil {
		analytics.AverageDaily = 0
	}

	// Get average weekly expenses (last 12 weeks)
	err = database.DB.QueryRow(`
		SELECT COALESCE(AVG(weekly_total), 0) FROM (
			SELECT SUM(amount) as weekly_total 
			FROM expenses 
			WHERE date >= date('now', '-12 weeks')
			GROUP BY strftime('%Y-W%W', date)
		)
	`).Scan(&analytics.AverageWeekly)
	if err != nil {
		analytics.AverageWeekly = 0
	}

	// Get average monthly expenses (last 12 months)
	err = database.DB.QueryRow(`
		SELECT COALESCE(AVG(monthly_total), 0) FROM (
			SELECT SUM(amount) as monthly_total 
			FROM expenses 
			WHERE date >= date('now', '-12 months')
			GROUP BY strftime('%Y-%m', date)
		)
	`).Scan(&analytics.AverageMonthly)
	if err != nil {
		analytics.AverageMonthly = 0
	}

	// Get monthly data (last 6 months)
	monthlyRows, err := database.DB.Query(`
		SELECT 
			strftime('%Y-%m', date) as month,
			SUM(amount) as total,
			COUNT(*) as count
		FROM expenses 
		WHERE date >= date('now', '-6 months')
		GROUP BY strftime('%Y-%m', date)
		ORDER BY month DESC
	`)
	if err == nil {
		defer monthlyRows.Close()
		for monthlyRows.Next() {
			var monthly models.MonthlyExpense
			monthlyRows.Scan(&monthly.Month, &monthly.Total, &monthly.Count)
			
			// Format month name
			if len(monthly.Month) >= 7 {
				monthNames := []string{"", "January", "February", "March", "April", "May", "June",
					"July", "August", "September", "October", "November", "December"}
				year := monthly.Month[:4]
				monthNum := monthly.Month[5:7]
				if monthNum[0] == '0' {
					monthNum = monthNum[1:]
				}
				if monthIndex := parseMonth(monthNum); monthIndex > 0 && monthIndex <= 12 {
					monthly.MonthName = monthNames[monthIndex] + " " + year
				}
			}
			
			analytics.MonthlyData = append(analytics.MonthlyData, monthly)
		}
	}

	// Get weekly data (last 8 weeks)
	weeklyRows, err := database.DB.Query(`
		SELECT 
			strftime('%Y-W%W', date) as week,
			date(date, 'weekday 0', '-6 days') as week_start,
			date(date, 'weekday 0') as week_end,
			SUM(amount) as total,
			COUNT(*) as count
		FROM expenses 
		WHERE date >= date('now', '-8 weeks')
		GROUP BY strftime('%Y-W%W', date)
		ORDER BY week DESC
	`)
	if err == nil {
		defer weeklyRows.Close()
		for weeklyRows.Next() {
			var weekly models.WeeklyExpense
			var weekStart, weekEnd string
			weeklyRows.Scan(&weekly.Week, &weekStart, &weekEnd, &weekly.Total, &weekly.Count)
			weekly.WeekRange = formatWeekRange(weekStart, weekEnd)
			analytics.WeeklyData = append(analytics.WeeklyData, weekly)
		}
	}

	// Get category breakdown
	categoryRows, err := database.DB.Query(`
		SELECT category, SUM(amount) as total 
		FROM expenses 
		GROUP BY category 
		ORDER BY total DESC
	`)
	if err == nil {
		defer categoryRows.Close()
		for categoryRows.Next() {
			var summary models.ExpenseSummary
			categoryRows.Scan(&summary.Category, &summary.Total)
			analytics.CategoryBreakdown = append(analytics.CategoryBreakdown, summary)
		}
	}

	// Get top 5 categories
	topCategoryRows, err := database.DB.Query(`
		SELECT category, SUM(amount) as total 
		FROM expenses 
		GROUP BY category 
		ORDER BY total DESC 
		LIMIT 5
	`)
	if err == nil {
		defer topCategoryRows.Close()
		for topCategoryRows.Next() {
			var summary models.ExpenseSummary
			topCategoryRows.Scan(&summary.Category, &summary.Total)
			analytics.TopCategories = append(analytics.TopCategories, summary)
		}
	}

	c.JSON(http.StatusOK, analytics)
}

// Helper functions
func parseMonth(monthStr string) int {
	switch monthStr {
	case "1", "01":
		return 1
	case "2", "02":
		return 2
	case "3", "03":
		return 3
	case "4", "04":
		return 4
	case "5", "05":
		return 5
	case "6", "06":
		return 6
	case "7", "07":
		return 7
	case "8", "08":
		return 8
	case "9", "09":
		return 9
	case "10":
		return 10
	case "11":
		return 11
	case "12":
		return 12
	default:
		return 0
	}
}

func formatWeekRange(startDate, endDate string) string {
	// Simple formatting - in production you'd want better date parsing
	if len(startDate) >= 10 && len(endDate) >= 10 {
		return startDate[5:10] + " to " + endDate[5:10] + ", " + endDate[:4]
	}
	return startDate + " to " + endDate
}
