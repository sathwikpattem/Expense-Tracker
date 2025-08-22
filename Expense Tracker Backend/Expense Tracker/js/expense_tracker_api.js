// API base URL
const API_BASE = '/api';

// Global variables
let expenses = [];
let categories = [];
let expenseSummary = [];
let analyticsData = null;

// Initialize the application
document.addEventListener('DOMContentLoaded', function() {
    loadCategories();
    loadExpenses();
    loadSummary();
});

// API functions
async function apiCall(endpoint, options = {}) {
    try {
        const response = await fetch(`${API_BASE}${endpoint}`, {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        });
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'API call failed');
        }
        
        return await response.json();
    } catch (error) {
        console.error('API Error:', error);
        alert('Error: ' + error.message);
        throw error;
    }
}

// Load categories from backend
async function loadCategories() {
    try {
        categories = await apiCall('/categories');
        updateCategoryDropdown();
    } catch (error) {
        console.error('Failed to load categories:', error);
    }
}

// Load expenses from backend
async function loadExpenses() {
    try {
        expenses = await apiCall('/expenses');
        updateExpenseTable();
    } catch (error) {
        console.error('Failed to load expenses:', error);
    }
}

// Load expense summary from backend
async function loadSummary() {
    try {
        expenseSummary = await apiCall('/summary');
        updateSummaryDisplay();
    } catch (error) {
        console.error('Failed to load summary:', error);
    }
}

// Update category dropdown
function updateCategoryDropdown() {
    const categorySelect = document.getElementById('category');
    categorySelect.innerHTML = '';
    
    categories.forEach(category => {
        const option = document.createElement('option');
        option.value = category.name;
        option.textContent = category.name;
        categorySelect.appendChild(option);
    });
}

// Update expense table
function updateExpenseTable() {
    const table = document.querySelector('.expense-table');
    
    // Remove existing rows except header
    const rows = table.querySelectorAll('tr:not(:first-child)');
    rows.forEach(row => row.remove());
    
    // Add expense rows
    expenses.forEach(expense => {
        const row = document.createElement('tr');
        row.style.borderBottom = '1px solid white';
        row.innerHTML = `
            <td>$${expense.amount.toFixed(2)}</td>
            <td>${expense.category}</td>
            <td>${expense.date}</td>
            <td>${expense.note}</td>
            <td><button onclick="deleteExpense(${expense.id})" class="delete-btn">Delete</button></td>
        `;
        table.appendChild(row);
    });
}

// Update summary display
function updateSummaryDisplay() {
    const summaryContainer = document.querySelector('.available-categories');
    summaryContainer.innerHTML = '';
    
    expenseSummary.forEach(summary => {
        const span = document.createElement('span');
        span.innerHTML = `${summary.category} <span class="sum"> - $${summary.total.toFixed(2)} spent</span>`;
        summaryContainer.appendChild(span);
    });
}

// Add expense function
async function add_expense() {
    const amount = parseFloat(document.getElementsByClassName("form-input")[0].value);
    const categorySelect = document.getElementsByClassName("form-input")[1];
    const date = document.getElementsByClassName("form-input")[2].value;
    const note = document.getElementsByClassName("form-input")[3].value;
    const category = categorySelect.value;

    if (!amount || !date || !note || !category) {
        alert("Please enter all details!");
        return;
    }

    try {
        await apiCall('/expenses', {
            method: 'POST',
            body: JSON.stringify({
                amount: amount,
                category: category,
                date: date,
                note: note
            })
        });

        alert("Successfully added expense!");
        
        // Clear form
        document.getElementsByClassName("form-input")[0].value = '';
        document.getElementsByClassName("form-input")[2].value = '';
        document.getElementsByClassName("form-input")[3].value = '';
        
        // Reload data
        await loadExpenses();
        await loadSummary();
        
        // Hide form
        document.getElementsByClassName("add-new-expense")[0].style.display = "none";
    } catch (error) {
        console.error('Failed to add expense:', error);
    }
}

// Add category function
async function add_category() {
    const newCategoryValue = document.getElementsByName("new-category")[0].value;
    
    if (!newCategoryValue) {
        alert("Please enter a category name!");
        return;
    }

    try {
        await apiCall('/categories', {
            method: 'POST',
            body: JSON.stringify({
                name: newCategoryValue
            })
        });

        alert("Successfully added category!");
        
        // Clear form
        document.getElementsByName("new-category")[0].value = '';
        
        // Reload categories and summary
        await loadCategories();
        await loadSummary();
        
        // Hide form
        document.getElementsByClassName("add-category")[0].style.display = "none";
    } catch (error) {
        console.error('Failed to add category:', error);
    }
}

// Delete expense function
async function deleteExpense(expenseId) {
    if (!confirm('Are you sure you want to delete this expense?')) {
        return;
    }

    try {
        await apiCall(`/expenses/${expenseId}`, {
            method: 'DELETE'
        });

        alert("Expense deleted successfully!");
        
        // Reload data
        await loadExpenses();
        await loadSummary();
    } catch (error) {
        console.error('Failed to delete expense:', error);
    }
}

// Toggle add category form
function add_category_page() {
    const addCategoryForm = document.getElementsByClassName("add-category")[0];
    if (addCategoryForm.style.display === "block") {
        // Hide with animation
        addCategoryForm.classList.remove("show");
        setTimeout(() => {
            addCategoryForm.style.display = "none";
        }, 300);
    } else {
        // Show with animation
        addCategoryForm.style.display = "block";
        setTimeout(() => {
            addCategoryForm.classList.add("show");
        }, 10);
        
        // Smooth scroll to the form and focus on input
        setTimeout(() => {
            addCategoryForm.scrollIntoView({ 
                behavior: 'smooth', 
                block: 'nearest' 
            });
            const categoryInput = addCategoryForm.querySelector('.form-input');
            if (categoryInput) categoryInput.focus();
        }, 100);
    }
}

// Toggle add expense form
function add_expense_page() {
    const addExpenseForm = document.getElementsByClassName("add-new-expense")[0];
    if (addExpenseForm.style.display === "block") {
        // Hide with animation
        addExpenseForm.classList.remove("show");
        setTimeout(() => {
            addExpenseForm.style.display = "none";
        }, 300);
    } else {
        // Show with animation
        addExpenseForm.style.display = "block";
        setTimeout(() => {
            addExpenseForm.classList.add("show");
        }, 10);
        
        // Smooth scroll to the form and focus on first input
        setTimeout(() => {
            addExpenseForm.scrollIntoView({ 
                behavior: 'smooth', 
                block: 'nearest' 
            });
            const firstInput = addExpenseForm.querySelector('.form-input');
            if (firstInput) firstInput.focus();
        }, 100);
    }
}

// Analytics Functions

// Load analytics data from backend
async function loadAnalytics() {
    try {
        analyticsData = await apiCall('/analytics');
        updateAnalyticsDashboard();
    } catch (error) {
        console.error('Failed to load analytics:', error);
    }
}

// Toggle analytics dashboard
async function toggle_analytics() {
    const analyticsDiv = document.querySelector('.analytics-dashboard');
    
    if (analyticsDiv.style.display === 'block') {
        analyticsDiv.style.display = 'none';
    } else {
        analyticsDiv.style.display = 'block';
        // Load analytics data when dashboard is opened
        await loadAnalytics();
    }
}

// Update analytics dashboard with data
function updateAnalyticsDashboard() {
    if (!analyticsData) return;

    // Update summary cards
    document.getElementById('total-expenses').textContent = `$${analyticsData.total_expenses.toFixed(2)}`;
    document.getElementById('daily-average').textContent = `$${analyticsData.average_daily.toFixed(2)}`;
    document.getElementById('weekly-average').textContent = `$${analyticsData.average_weekly.toFixed(2)}`;
    document.getElementById('monthly-average').textContent = `$${analyticsData.average_monthly.toFixed(2)}`;

    // Update monthly chart
    updateMonthlyChart(analyticsData.monthly_data);
    
    // Update weekly chart
    updateWeeklyChart(analyticsData.weekly_data);
    
    // Update category chart
    updateCategoryChart(analyticsData.top_categories);
    
    // Update monthly table
    updateMonthlyTable(analyticsData.monthly_data);
    
    // Update weekly table
    updateWeeklyTable(analyticsData.weekly_data);
}

// Update monthly chart
function updateMonthlyChart(monthlyData) {
    const chartDiv = document.getElementById('monthly-chart');
    
    if (!monthlyData || monthlyData.length === 0) {
        chartDiv.innerHTML = '<div class="chart-placeholder">No monthly data available</div>';
        return;
    }

    const maxAmount = Math.max(...monthlyData.map(item => item.total));
    
    let chartHTML = '';
    monthlyData.forEach(month => {
        const percentage = maxAmount > 0 ? (month.total / maxAmount) * 100 : 0;
        chartHTML += `
            <div class="chart-bar" style="--percentage: ${percentage}%">
                <span class="chart-bar-label">${month.month_name}</span>
                <span class="chart-bar-value">$${month.total.toFixed(2)}</span>
            </div>
        `;
    });
    
    chartDiv.innerHTML = chartHTML;
}

// Update weekly chart
function updateWeeklyChart(weeklyData) {
    const chartDiv = document.getElementById('weekly-chart');
    
    if (!weeklyData || weeklyData.length === 0) {
        chartDiv.innerHTML = '<div class="chart-placeholder">No weekly data available</div>';
        return;
    }

    const maxAmount = Math.max(...weeklyData.map(item => item.total));
    
    let chartHTML = '';
    weeklyData.forEach(week => {
        const percentage = maxAmount > 0 ? (week.total / maxAmount) * 100 : 0;
        chartHTML += `
            <div class="chart-bar" style="--percentage: ${percentage}%">
                <span class="chart-bar-label">${week.week_range}</span>
                <span class="chart-bar-value">$${week.total.toFixed(2)}</span>
            </div>
        `;
    });
    
    chartDiv.innerHTML = chartHTML;
}

// Update category chart
function updateCategoryChart(categoryData) {
    const chartDiv = document.getElementById('category-chart');
    
    if (!categoryData || categoryData.length === 0) {
        chartDiv.innerHTML = '<div class="chart-placeholder">No category data available</div>';
        return;
    }

    const maxAmount = Math.max(...categoryData.map(item => item.total));
    
    let chartHTML = '';
    categoryData.forEach(category => {
        const percentage = maxAmount > 0 ? (category.total / maxAmount) * 100 : 0;
        chartHTML += `
            <div class="chart-bar" style="--percentage: ${percentage}%">
                <span class="chart-bar-label">${category.category}</span>
                <span class="chart-bar-value">$${category.total.toFixed(2)}</span>
            </div>
        `;
    });
    
    chartDiv.innerHTML = chartHTML;
}

// Update monthly table
function updateMonthlyTable(monthlyData) {
    const tbody = document.querySelector('#monthly-table tbody');
    tbody.innerHTML = '';
    
    if (!monthlyData || monthlyData.length === 0) {
        tbody.innerHTML = '<tr><td colspan="4" style="text-align: center;">No monthly data available</td></tr>';
        return;
    }

    monthlyData.forEach(month => {
        const avgPerExpense = month.count > 0 ? (month.total / month.count) : 0;
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${month.month_name}</td>
            <td>$${month.total.toFixed(2)}</td>
            <td>${month.count}</td>
            <td>$${avgPerExpense.toFixed(2)}</td>
        `;
        tbody.appendChild(row);
    });
}

// Update weekly table
function updateWeeklyTable(weeklyData) {
    const tbody = document.querySelector('#weekly-table tbody');
    tbody.innerHTML = '';
    
    if (!weeklyData || weeklyData.length === 0) {
        tbody.innerHTML = '<tr><td colspan="4" style="text-align: center;">No weekly data available</td></tr>';
        return;
    }

    weeklyData.forEach(week => {
        const avgPerExpense = week.count > 0 ? (week.total / week.count) : 0;
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${week.week_range}</td>
            <td>$${week.total.toFixed(2)}</td>
            <td>${week.count}</td>
            <td>$${avgPerExpense.toFixed(2)}</td>
        `;
        tbody.appendChild(row);
    });
}
