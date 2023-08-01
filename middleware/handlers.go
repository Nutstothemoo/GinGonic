package middleware

import (
	"database/sql"
	"fmt"
	"ginapi/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() gin.HandlerFunc{

	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err =db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetAlbums(c *gin.Context) {
	// Obtenir la connexion à la base de données à partir du contexte.
	db, exists := c.Get("db")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get database connection"})
		return
	}

	// Cast db en *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid database connection"})
		return
	}

	albums, err := models.GetAllAlbums(dbConn)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	// Obtenir la connexion à la base de données à partir du contexte.
	db, exists := c.Get("db")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get database connection"})
		return
	}

	// Cast db en *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid database connection"})
		return
	}

	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON"})
		return
	}

	err := newAlbum.Save(dbConn)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to save album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Obtenir la connexion à la base de données à partir du contexte.
	db, exists := c.Get("db")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get database connection"})
		return
	}

	// Cast db en *sql.DB
	dbConn, ok := db.(*sql.DB)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid database connection"})
		return
	}

	album, err := models.GetAlbumByID(dbConn, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
