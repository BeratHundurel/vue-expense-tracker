# Vue Expense Tracker with Go API

This is a simple expense tracker application built with Vue.js for the frontend and Golang with Microsoft SQL Server for the backend API. It enables users to manage their transactions, including adding new transactions, deleting existing ones, and viewing transaction history.

## Features

-  Add new transactions (income or expense).
-  Delete existing transactions.
-  View transaction history.
-  Automatic saving of transactions to local storage and database.

## Technologies Used

### Frontend:

-  Vue.js
-  Composition API
-  HTML/CSS
-  TypeScript
-  Axios

### Backend:

-  Golang
-  Microsoft SQL Server
-  Gin Framework

## Getting Started

To run the application locally, you need to set up both the frontend and the backend:

### Frontend:

1. Clone the repository:

   ```bash
   git clone https://github.com/BeratHundurel/vue-expense-tracker.git
   ```

2. Navigate into the frontend directory:

   ```bash
   cd client
   ```

3. Install dependencies:

   ```bash
   npm install
   ```

4. Run the development server:

   ```bash
   npm run dev
   ```

5. Open your browser and go to `http://localhost:5174` to access the frontend application.

### Backend:

2. Navigate into the backend directory:

   ```bash
   cd api
   ```

3. Configure your Microsoft SQL Server database connection in the backend code.

4. Build and run the backend server:

   ```bash
   go run .
   ```

## Usage

-  Adding a transaction: Enter the transaction name and amount, then click the "Add Transaction" button.
-  Deleting a transaction: Click the X icon next to the transaction you wish to remove.
