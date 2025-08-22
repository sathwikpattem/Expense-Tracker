#!/bin/bash

echo "Setting up Expense Tracker with Go Backend..."

# Navigate to backend directory
cd backend

# Initialize Go modules and download dependencies
echo "Installing Go dependencies..."
go mod tidy

# Build the application
echo "Building the application..."
go build -o expense-tracker main.go

echo "Setup complete!"
echo ""
echo "To run the application:"
echo "1. cd backend"
echo "2. ./expense-tracker"
echo ""
echo "Then open your browser to http://localhost:8080"
