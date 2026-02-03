Students REST API

📘 Overview
- Students REST API is a backend application developed using Go (Golang) that provides RESTful web services for managing student information. The API supports basic operations such as creating and retrieving student records using standard HTTP methods.
- The application uses SQLite as the database for data storage. TablePlus is used to manage the database visually, and Postman is used to test the API endpoints.
- This project demonstrates REST API development, database connectivity, and request handling using Golang.

✨ Features
- RESTful API design
- Create student records using POST requests
- Retrieve student records using GET requests
- SQLite database integration
- API testing with Postman
- Simple and lightweight backend service

🛠️ Tech Stack
- Language: Go (Golang)
- Database: SQLite
- Database Tool: TablePlus
- API Testing: Postman
- Architecture: REST API

⚙️ Setup & Installation
✅ Prerequisites

Make sure the following tools are installed on your system:
  - Go (version 1.20 or higher)
  - Git
  -Postman
  -TablePlus (optional, for database management)


📦 Install Dependencies
  - go mod tidy
  - go get github.com/go-playground/validator/v10
  - go get github.com/mattn/go-sqlite3
  - go get github.com/ilyakaznacheev/cleanenv

▶️ Run the Application
 - go run cmd/students_api/main.go -config config/local.yaml

After running the command, the server will start on:
 - http://localhost:8082


🧪 Testing the API
- Open Postman
- Create a POST request to /students to add data
- Create a GET request to /students to fetch data
- Send and receive data in JSON format

🗄️Database Details
- Database: SQLite
- Database File: students.db
- Database tables are created and managed using TablePlus


🚀 Future Improvements
- Add PUT and DELETE operations
- Add authentication and authorization
- Implement validation and error handling
- Add Swagger/OpenAPI documentation
- Dockerize the application
