# 💰 Expense Tracker with Go Backend

A modern, full-stack expense tracking application with a Go backend, RESTful API, and beautiful responsive frontend featuring comprehensive analytics and data visualization.

## ✨ Features

### 📊 Core Functionality
- ✅ **Add/Edit/Delete expenses** with intuitive forms
- ✅ **Dynamic categorization** with custom category creation
- ✅ **Real-time expense summaries** by category
- ✅ **Smart form positioning** - forms appear at the top for better UX
- ✅ **Auto-focus and smooth animations** for seamless interactions

### 📈 Advanced Analytics Dashboard
- ✅ **Comprehensive Analytics Dashboard** with full-screen overlay
- ✅ **Monthly expense tracking** (last 12 months with trend analysis)
- ✅ **Weekly expense analysis** (last 8 weeks with detailed breakdowns)
- ✅ **Expense averages** (daily, weekly, monthly calculations)
- ✅ **Category insights** with top spending categories and visual charts
- ✅ **Interactive data tables** with sortable and responsive design
- ✅ **Visual charts** with percentage-based bar representations
- ✅ **Summary cards** displaying key financial metrics

### 🎨 Modern UI/UX Design
- ✅ **Glassmorphism design** with backdrop blur effects
- ✅ **Responsive grid layout** optimized for all devices
- ✅ **Smooth animations** and transitions throughout
- ✅ **Mobile-first design** with touch-optimized interactions
- ✅ **Professional color scheme** with gradient backgrounds
- ✅ **Enhanced typography** using Montserrat font family
- ✅ **Auto-scroll and focus** for improved form interactions

### 🔧 Technical Features
- ✅ **RESTful API** with comprehensive endpoints
- ✅ **SQLite database** with optimized schema
- ✅ **CORS enabled** for frontend integration
- ✅ **Error handling** and validation
- ✅ **Sample data scripts** for testing and demonstration

## Technology Stack

- **Backend**: Go (Gin framework)
- **Database**: SQLite
- **Frontend**: HTML, CSS, JavaScript
- **API**: RESTful JSON API

## Project Structure

```
expense-tracker/
├── css/
│   └── expense_tracker.css (Enhanced responsive styling)
├── js/
│   ├── expense_tracker.js (Original client-side version)
│   └── expense_tracker_api.js (API-integrated version)
├── templates/
│   └── expense_tracker.html (Modern responsive UI)
├── backend/
│   ├── main.go (Server with static file serving)
│   ├── expense-tracker (Compiled binary)
│   ├── expenses.db (SQLite database)
│   ├── models/
│   │   └── expense.go (Data models and analytics structs)
│   ├── handlers/
│   │   └── expense.go (API endpoints and analytics)
│   └── database/
│       └── db.go (Database initialization and schema)
├── add_sample_data.sh (Script to add demo data)
├── demo_analytics.sh (Analytics demonstration script)
├── setup.sh (Automated setup script)
└── README.md
```

## Quick Start

### Prerequisites
- Go 1.21 or higher
- Git

### Installation

1. **Clone or setup the project:**
   ```bash
   # If you have the files, just run the setup script
   ./setup.sh
   ```

2. **Start the server:**
   ```bash
   cd backend
   ./expense-tracker
   ```

3. **Add sample data (optional):**
   ```bash
   ./add_sample_data.sh
   ```

4. **Open your browser:**
   Navigate to `http://localhost:8080`

5. **Explore analytics:**
   Click the "Analytics" button to see the dashboard

### Alternative Manual Setup

1. **Install dependencies:**
   ```bash
   cd backend
   go mod tidy
   ```

2. **Build the application:**
   ```bash
   go build -o expense-tracker main.go
   ```

3. **Run the server:**
   ```bash
   ./expense-tracker
   # OR
   go run main.go
   ```

### Demo and Testing

- **Run analytics demo:**
  ```bash
  ./demo_analytics.sh
  ```

- **Add sample expenses:**
  ```bash
  ./add_sample_data.sh
  ```

## API Endpoints

### Expenses
- `GET /api/expenses` - Get all expenses
- `POST /api/expenses` - Create new expense
- `DELETE /api/expenses/:id` - Delete expense

### Categories
- `GET /api/categories` - Get all categories
- `POST /api/categories` - Create new category

### Summary
- `GET /api/summary` - Get expense summary by category

### Analytics
- `GET /api/analytics` - Get comprehensive expense analytics
- `GET /api/analytics/monthly` - Get monthly expense data (last 12 months)
- `GET /api/analytics/weekly` - Get weekly expense data (last 8 weeks)

## API Usage Examples

### Create an expense
```bash
curl -X POST http://localhost:8080/api/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 25.50,
    "category": "Food",
    "date": "2025-08-14",
    "note": "Lunch at restaurant"
  }'
```

### Get comprehensive analytics
```bash
curl -s http://localhost:8080/api/analytics
```

### Get monthly data
```bash
curl -s http://localhost:8080/api/analytics/monthly
```

### Get weekly data
```bash
curl -s http://localhost:8080/api/analytics/weekly
```

## 🎯 User Interface Guide

### Main Dashboard
- **Add Expense**: Click to show form at the top, auto-focuses on amount field
- **Add Category**: Click to show form at the top, auto-focuses on name field
- **Analytics**: Opens full-screen analytics dashboard
- **Expense Table**: View all expenses with delete functionality
- **Category Summary**: Real-time spending totals by category

### Analytics Dashboard
- **Summary Cards**: Total expenses, daily/weekly/monthly averages
- **Monthly Chart**: Visual bar chart of last 6 months spending
- **Weekly Chart**: Visual bar chart of last 8 weeks spending
- **Category Chart**: Top spending categories with visual representation
- **Data Tables**: Detailed monthly and weekly breakdowns
- **Responsive Design**: Optimized for desktop, tablet, and mobile

### Key Features
- ✨ **Smooth Animations**: All interactions have smooth transitions
- 📱 **Mobile Optimized**: Touch-friendly design with responsive layout
- 🎨 **Modern UI**: Glassmorphism design with backdrop blur effects
- ⚡ **Auto-scroll**: Forms automatically scroll into view when opened
- 🔍 **Auto-focus**: Input fields automatically receive focus
- 📊 **Real-time Updates**: Data refreshes automatically after changes

## Frontend Integration

The project includes two JavaScript files:
- `expense_tracker.js` - Original client-side only version
- `expense_tracker_api.js` - API-integrated version for the Go backend

To use the API version, update your HTML to include:
```html
<script src="../js/expense_tracker_api.js"></script>
```

## Database

The application uses SQLite with the following schema:

### Categories Table
- `id` (INTEGER PRIMARY KEY)
- `name` (TEXT UNIQUE)

### Expenses Table
- `id` (INTEGER PRIMARY KEY)
- `amount` (REAL)
- `category` (TEXT)
- `date` (TEXT)
- `note` (TEXT)
- `created` (DATETIME)

## Development

### Adding New Features

1. **Models**: Add new structs in `models/expense.go`
2. **Database**: Update schema in `database/db.go`
3. **Handlers**: Add new endpoints in `handlers/expense.go`
4. **Routes**: Register routes in `main.go`

### Building for Production

```bash
cd backend
go build -o expense-tracker main.go
```

## Configuration

- **Port**: Default is 8080, can be changed in `main.go`
- **Database**: SQLite file `expenses.db` created automatically
- **CORS**: Currently allows all origins (change for production)

## Troubleshooting

1. **Port already in use**: Change the port in `main.go`
2. **Database issues**: Delete `expenses.db` and restart
3. **CORS errors**: Check the CORS configuration in `main.go`

## Next Steps

Consider these enhancements:
- [ ] User authentication
- [ ] Data export (CSV/PDF)
- [ ] Advanced charts with Chart.js or D3.js
- [ ] Budget tracking and alerts
- [ ] Recurring expense management
- [ ] Multi-currency support
- [ ] Docker containerization
- [ ] PostgreSQL for production
- [ ] API rate limiting
- [ ] Input validation improvements
- [ ] Dark/Light theme toggle
- [ ] Email notifications
- [ ] Data backup and restore

## 📸 Screenshots

### Main Dashboard
- Modern glassmorphism design with gradient backgrounds
- Responsive expense table with hover effects
- Smart form positioning at the top of sections

### Analytics Dashboard
- Full-screen analytics overlay with comprehensive data
- Interactive charts with percentage-based visualizations
- Summary cards showing key financial metrics
- Detailed monthly and weekly breakdown tables

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Go Gin framework for the excellent web framework
- SQLite for the lightweight database solution
- Modern CSS techniques for the beautiful UI design
