package models

import (
	"time"
)

// Expense represents an expense record
type Expense struct {
	ID       int       `json:"id" db:"id"`
	Amount   float64   `json:"amount" db:"amount"`
	Category string    `json:"category" db:"category"`
	Date     string    `json:"date" db:"date"`
	Note     string    `json:"note" db:"note"`
	Created  time.Time `json:"created" db:"created"`
}

// Category represents a category
type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// ExpenseRequest represents the request body for creating expenses
type ExpenseRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Date     string  `json:"date" binding:"required"`
	Note     string  `json:"note" binding:"required"`
}

// CategoryRequest represents the request body for creating categories
type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// ExpenseSummary represents expense summary by category
type ExpenseSummary struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

// MonthlyExpense represents monthly expense data
type MonthlyExpense struct {
	Month     string  `json:"month"`      // Format: "2025-08"
	MonthName string  `json:"month_name"` // Format: "August 2025"
	Total     float64 `json:"total"`
	Count     int     `json:"count"`
}

// WeeklyExpense represents weekly expense data
type WeeklyExpense struct {
	Week      string  `json:"week"`       // Format: "2025-W33"
	WeekRange string  `json:"week_range"` // Format: "Aug 11-17, 2025"
	Total     float64 `json:"total"`
	Count     int     `json:"count"`
}

// DailyExpense represents daily expense data
type DailyExpense struct {
	Date  string  `json:"date"`
	Total float64 `json:"total"`
	Count int     `json:"count"`
}

// CategoryMonthlyExpense represents monthly expenses by category
type CategoryMonthlyExpense struct {
	Category string  `json:"category"`
	Month    string  `json:"month"`
	Total    float64 `json:"total"`
	Count    int     `json:"count"`
}

// ExpenseAnalytics represents comprehensive expense analytics
type ExpenseAnalytics struct {
	TotalExpenses     float64                  `json:"total_expenses"`
	AverageDaily      float64                  `json:"average_daily"`
	AverageWeekly     float64                  `json:"average_weekly"`
	AverageMonthly    float64                  `json:"average_monthly"`
	MonthlyData       []MonthlyExpense         `json:"monthly_data"`
	WeeklyData        []WeeklyExpense          `json:"weekly_data"`
	CategoryBreakdown []ExpenseSummary         `json:"category_breakdown"`
	TopCategories     []ExpenseSummary         `json:"top_categories"`
}
