package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dslipak/pdf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dataSourceName := "goUser:12345678@tcp(127.0.0.1:3306)/pdfreader"
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	fmt.Println("Connected to the database successfully")
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Serve static files from the "frontend/build" directory
	router.StaticFile("/", "./frontend/build/index.html")
	router.Static("/static", "./frontend/build/static")

	// Set a route for uploading PDF files
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		// Check if the uploaded file is a PDF
		if filepath.Ext(file.Filename) != ".pdf" {
			c.String(http.StatusBadRequest, "only PDF files are allowed")
			return
		}

		// Save the uploaded file to disk
		dst, err := saveUploadedFile(file)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		content, err := readPdf(dst) // Read local pdf file
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, content)
	})

	router.Run(":12345")
}

func readPdf(path string) (string, error) {
	src, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer src.Close()

	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func saveUploadedFile(file *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create a destination file on disk
	newpath := filepath.Join(".", "uploads")
	dirErr := os.MkdirAll(newpath, os.ModePerm)

	if dirErr != nil {
		return "", dirErr
	}

	dst, err := os.Create("./uploads/" + file.Filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return dst.Name(), nil
}
