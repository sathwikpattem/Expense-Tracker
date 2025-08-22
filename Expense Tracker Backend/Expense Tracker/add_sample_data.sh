#!/bin/bash

echo "Adding sample expense data for analytics testing..."

# Sample data for the last few months
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 45.50, "category": "Food", "date": "2025-08-10", "note": "Grocery shopping"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 120.00, "category": "Travel", "date": "2025-08-08", "note": "Gas for car"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 25.99, "category": "Entertainment", "date": "2025-08-05", "note": "Movie tickets"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 85.30, "category": "Food", "date": "2025-07-28", "note": "Restaurant dinner"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 200.00, "category": "Groceries", "date": "2025-07-25", "note": "Weekly groceries"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 75.00, "category": "Entertainment", "date": "2025-07-20", "note": "Concert tickets"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 150.00, "category": "Travel", "date": "2025-07-15", "note": "Train tickets"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 35.80, "category": "Food", "date": "2025-07-12", "note": "Lunch meetings"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 95.00, "category": "Others", "date": "2025-07-08", "note": "Pharmacy"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 60.50, "category": "Food", "date": "2025-06-30", "note": "Coffee and snacks"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 180.00, "category": "Groceries", "date": "2025-06-25", "note": "Monthly shopping"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 45.00, "category": "Entertainment", "date": "2025-06-20", "note": "Streaming services"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 110.00, "category": "Travel", "date": "2025-06-18", "note": "Weekend trip"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 30.75, "category": "Food", "date": "2025-06-15", "note": "Fast food"}'
curl -X POST http://localhost:8080/api/expenses -H "Content-Type: application/json" -d '{"amount": 125.50, "category": "Others", "date": "2025-06-10", "note": "Utilities"}'

echo ""
echo "Sample data added! You can now test the analytics features."
echo "Start the server with: cd backend && ./expense-tracker"
echo "Then open http://localhost:8080 and click the Analytics button."
