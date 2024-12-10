# Go Web Tools Management System

A web-based tool management system built with Go, Gin Framework, and MySQL. This application allows users to manage tools and their categories with full CRUD operations.

## Features
- Tool Management (Create, Read, Update, Delete)
- Category Management
- Relational Database Structure
- Responsive Web Interface
- Data Validation
- Clean Architecture

## Technology Stack
- Backend: Go (Gin Framework)
- Database: MySQL
- Frontend: HTML, Bootstrap 4
- JavaScript: jQuery, DataTables

## Project Structure
```
PBKK_GO_WEBTOOLS/
├── main.go               # Main application file
├── .env                 # Environment configuration
├── templates/           # HTML templates
│   ├── Index.html       # Tool listing page
│   ├── Show.html        # Tool details view
│   ├── Edit.html        # Tool edit form
│   ├── New.html         # New tool form
│   ├── Categories.html  # Category listing
│   └── ...
├── scripts/
│   └── scripts.sql      # Database schema
└── vendor/              # Dependencies
```

## Database Schema
### Tables
1. `tools`
   - id (PK)
   - name
   - category_id (FK)
   - url
   - rating
   - notes
   - created_at

2. `categories`
   - id (PK)
   - name
   - description
   - created_at

## Setup Instructions

### Prerequisites
- Go 1.21 or later
- MySQL 9.x
- Git

### Installation Steps
1. Clone the repository
```bash
git clone [repository-url]
cd PBKK_GO_WEBTOOLS
```

2. Set up the environment file (.env)
```env
DATABASE_USERNAME=your_username
DATABASE_PASSWORD=your_password
DATABASE_NAME=gowtdb
DATABASE_SERVER=localhost
DATABASE_PORT=3306
```

3. Create the database and tables
```bash
# Login to MySQL
mysql -u root -p

# Create database and tables
mysql> source scripts/scripts.sql
```

4. Install dependencies
```bash
go mod tidy
```

5. Run the application
```bash
go run main.go
```

6. Access the application at `http://localhost:8080`

## Key Features and Implementation

### Tool Management
- List all tools with pagination and search
- Create new tools with category selection
- Update existing tools
- Delete tools
- View detailed tool information

### Category Management
- Create and manage categories
- Associate tools with categories
- Category-based filtering

### Data Validation
- Input validation for required fields
- Rating range validation (1-5)
- URL format validation

## API Routes
```
GET    /                         # Home page, list all tools
GET    /show                     # Show tool details
GET    /new                      # New tool form
GET    /edit                     # Edit tool form
POST   /insert                   # Create new tool
POST   /update                   # Update tool
GET    /delete                   # Delete tool
GET    /categories              # List categories
GET    /categories/new          # New category form
POST   /categories/insert       # Create new category
GET    /categories/edit         # Edit category form
POST   /categories/update       # Update category
GET    /categories/delete       # Delete category
```

## Development Notes

### Key Algorithms
1. **Database Connection Pool**
   - Efficient connection management
   - Connection reuse
   - Proper error handling

2. **CRUD Operations**
   - Prepared statements for security
   - Transaction handling
   - Error handling and validation

3. **Template Rendering**
   - Nested template support
   - Dynamic data binding
   - Cross-site scripting protection

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License
[Your License]

# GOWT
Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL, Docker.

### Prerequisites

* Golang, preferably the latest version (1.16).
* MySQL Database
* Docker (optional)

Run below command and install dependencies

```
go mod download
```

3. Create database on MySQL

```
CREATE DATABASE gowtdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE gowtdb;

CREATE TABLE tools (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  category varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  url varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  rating int(11) DEFAULT NULL,
  notes text COLLATE utf8_unicode_ci,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```

4. Create a .env file with the variables listed bellow and change values as needed

```
DATABASE_NAME="gowtdb"
DATABASE_USERNAME="user"
DATABASE_PASSWORD="pass"
DATABASE_SERVER="localhost"
DATABASE_PORT="3306"
```

6. Run the application

```
make run
```

## Deployment

1. Create an executable

```
make build
```

2. Run the application

```
./out/bin/gowt
```
## Create Docker image

1. To build and tag your image locally

```
make docker-build
```

2. To push your image to registry

```
make docker-release
```

## Run Docker image locally

```
make docker-run
```
