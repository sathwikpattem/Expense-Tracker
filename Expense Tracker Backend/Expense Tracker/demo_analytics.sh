#!/bin/bash

echo "🚀 Expense Tracker Analytics Demo"
echo "================================="
echo ""

# Check if server is running
if ! curl -s http://localhost:8080/api/categories > /dev/null; then
    echo "❌ Server is not running. Please start it first:"
    echo "   cd backend && ./expense-tracker"
    echo ""
    exit 1
fi

echo "✅ Server is running!"
echo ""

echo "📊 Analytics API Endpoints:"
echo "1. Comprehensive Analytics:"
echo "   http://localhost:8080/api/analytics"
echo ""

echo "2. Monthly Data:"
echo "   http://localhost:8080/api/analytics/monthly"
echo ""

echo "3. Weekly Data:"
echo "   http://localhost:8080/api/analytics/weekly"
echo ""

echo "4. Category Summary:"
echo "   http://localhost:8080/api/summary"
echo ""

echo "🌐 Open the web application:"
echo "   http://localhost:8080"
echo ""
echo "   Click the 'Analytics' button to see the dashboard!"
echo ""

echo "📋 Sample Analytics Data:"
echo "========================"

echo ""
echo "💰 Total Expenses:"
curl -s http://localhost:8080/api/analytics | grep -o '"total_expenses":[^,]*' | cut -d':' -f2

echo ""
echo "📅 Monthly Breakdown:"
curl -s http://localhost:8080/api/analytics/monthly | jq -r '.[] | "\(.month_name): $\(.total) (\(.count) expenses)"' 2>/dev/null || echo "Install jq for better formatting"

echo ""
echo "🏷️  Top Categories:"
curl -s http://localhost:8080/api/analytics | jq -r '.top_categories[] | "\(.category): $\(.total)"' 2>/dev/null || echo "Install jq for better formatting"

echo ""
echo "🎯 Ready to explore! Open http://localhost:8080 in your browser."
