package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Category struct
type Category struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

// Tool struct
type Tool struct {
    Id         int      `json:"id"`
    Name       string   `json:"name"`
    CategoryId int      `json:"category_id"`
    Category   Category `json:"category"`
    URL        string   `json:"url"`
    Rating     int      `json:"rating"`
    Notes      string   `json:"notes"`
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := os.Getenv("DATABASE_USERNAME")
    dbPass := os.Getenv("DATABASE_PASSWORD")
    dbName := os.Getenv("DATABASE_NAME")
    dbServer := os.Getenv("DATABASE_SERVER")
    dbPort := os.Getenv("DATABASE_PORT")

    log.Printf("Attempting database connection to %s:%s", dbServer, dbPort)

    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        dbUser, dbPass, dbServer, dbPort, dbName)
    
    var err error
    db, err = sql.Open(dbDriver, connectionString)
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return nil
    }

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Printf("Error pinging database: %v", err)
        return nil
    }

    log.Println("Successfully connected to database")
    return db
}

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // Load templates
    r.LoadHTMLGlob("templates/*")
    
    // Tool routes
    r.GET("/", indexHandler)
    r.GET("/show", showHandler)
    r.GET("/new", newHandler)
    r.GET("/edit", editHandler)
    r.POST("/insert", insertHandler)
    r.POST("/update", updateHandler)
    r.GET("/delete", deleteHandler)
    
    // Category routes
    r.GET("/categories", categoriesHandler)
    r.GET("/categories/new", newCategoryHandler)
    r.POST("/categories/insert", insertCategoryHandler)
    r.GET("/categories/edit", editCategoryHandler)
    r.POST("/categories/update", updateCategoryHandler)
    r.GET("/categories/delete", deleteCategoryHandler)
    
    return r
}

// Tool Handlers
func indexHandler(c *gin.Context) {
    db := dbConn()
    if db == nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{
            "error": "Unable to connect to database",
        })
        return
    }
    defer db.Close()

    // Updated query to join with categories
    rows, err := db.Query(`
        SELECT t.id, t.name, t.category_id, c.name, t.url, t.rating, t.notes, c.description 
        FROM tools t 
        JOIN categories c ON t.category_id = c.id 
        ORDER BY t.id DESC
    `)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{
            "error": "Error fetching tools: " + err.Error(),
        })
        return
    }
    defer rows.Close()

    var tools []Tool
    for rows.Next() {
        var tool Tool
        err := rows.Scan(
            &tool.Id, 
            &tool.Name, 
            &tool.CategoryId,
            &tool.Category.Name,
            &tool.URL, 
            &tool.Rating, 
            &tool.Notes,
            &tool.Category.Description,
        )
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }
        tools = append(tools, tool)
    }

    if err = rows.Err(); err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{
            "error": "Error processing results: " + err.Error(),
        })
        return
    }

    c.HTML(http.StatusOK, "Index", gin.H{
        "tools": tools,
    })
}

func newHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    // Get categories for dropdown
    rows, err := db.Query("SELECT id, name FROM categories ORDER BY name")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    var categories []Category
    for rows.Next() {
        var category Category
        err := rows.Scan(&category.Id, &category.Name)
        if err != nil {
            continue
        }
        categories = append(categories, category)
    }

    c.HTML(http.StatusOK, "New", gin.H{"Categories": categories})
}

func insertHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    name := c.PostForm("name")
    categoryId := c.PostForm("category_id")
    url := c.PostForm("url")
    rating := c.PostForm("rating")
    notes := c.PostForm("notes")

    stmt, err := db.Prepare("INSERT INTO tools (name, category_id, url, rating, notes) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(name, categoryId, url, rating, notes)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/")
}

func editHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    nId := c.Query("id")
    row := db.QueryRow(`
        SELECT t.id, t.name, t.category_id, t.url, t.rating, t.notes 
        FROM tools t 
        WHERE t.id = ?
    `, nId)

    var tool Tool
    err := row.Scan(&tool.Id, &tool.Name, &tool.CategoryId, &tool.URL, &tool.Rating, &tool.Notes)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    // Get categories for dropdown
    rows, err := db.Query("SELECT id, name FROM categories ORDER BY name")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    var categories []Category
    for rows.Next() {
        var category Category
        err := rows.Scan(&category.Id, &category.Name)
        if err != nil {
            continue
        }
        categories = append(categories, category)
    }

    c.HTML(http.StatusOK, "Edit", gin.H{
        "Tool":       tool,
        "Categories": categories,
    })
}

func updateHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    id := c.PostForm("uid")
    name := c.PostForm("name")
    categoryId := c.PostForm("category_id")
    url := c.PostForm("url")
    rating := c.PostForm("rating")
    notes := c.PostForm("notes")

    stmt, err := db.Prepare("UPDATE tools SET name=?, category_id=?, url=?, rating=?, notes=? WHERE id=?")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(name, categoryId, url, rating, notes, id)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/")
}

// Category Handlers
func categoriesHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    rows, err := db.Query("SELECT id, name, description FROM categories ORDER BY name")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    var categories []Category
    for rows.Next() {
        var category Category
        err := rows.Scan(&category.Id, &category.Name, &category.Description)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }
        categories = append(categories, category)
    }

    c.HTML(http.StatusOK, "Categories", categories)
}

func newCategoryHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "NewCategory", nil)
}

func insertCategoryHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    name := c.PostForm("name")
    description := c.PostForm("description")

    stmt, err := db.Prepare("INSERT INTO categories (name, description) VALUES (?, ?)")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(name, description)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/categories")
}

func editCategoryHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    nId := c.Query("id")
    row := db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", nId)

    var category Category
    err := row.Scan(&category.Id, &category.Name, &category.Description)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.HTML(http.StatusOK, "EditCategory", category)
}

func updateCategoryHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    id := c.PostForm("id")
    name := c.PostForm("name")
    description := c.PostForm("description")

    stmt, err := db.Prepare("UPDATE categories SET name=?, description=? WHERE id=?")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(name, description, id)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/categories")
}

// Add these functions to your main.go file

func showHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    nId := c.Query("id")
    row := db.QueryRow(`
        SELECT t.id, t.name, t.category_id, c.name as category_name, t.url, t.rating, t.notes 
        FROM tools t 
        JOIN categories c ON t.category_id = c.id 
        WHERE t.id = ?
    `, nId)

    var tool Tool
    var categoryName string
    err := row.Scan(
        &tool.Id, 
        &tool.Name, 
        &tool.CategoryId,
        &categoryName,
        &tool.URL, 
        &tool.Rating, 
        &tool.Notes,
    )
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    tool.Category.Name = categoryName
    c.HTML(http.StatusOK, "Show", tool)
}

func deleteHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    nId := c.Query("id")
    stmt, err := db.Prepare("DELETE FROM tools WHERE id=?")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(nId)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteCategoryHandler(c *gin.Context) {
    db := dbConn()
    defer db.Close()

    id := c.Query("id")
    
    // First check if category is in use
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM tools WHERE category_id = ?", id).Scan(&count)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }
    
    if count > 0 {
        c.HTML(http.StatusBadRequest, "Error", gin.H{
            "error": "Cannot delete category that is in use by tools",
        })
        return
    }

    stmt, err := db.Prepare("DELETE FROM categories WHERE id=?")
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    _, err = stmt.Exec(id)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusMovedPermanently, "/categories")
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    gin.SetMode(gin.DebugMode)
    router := setupRouter()
    
    log.Println("Server started on: http://localhost:8080")
    router.Run(":8080")
}