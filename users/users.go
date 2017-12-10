package users

import (
	"log"
	"net/http"
	"strconv"
	"testApi/storage"

	"github.com/gin-gonic/gin"
)

var db *storage.Storage

// SetDatabase loads selected storage into package
func SetDatabase(s *storage.Storage) {
	db = s
}

// GetAll ...
func GetAll(c *gin.Context) {
	orgs := db.GetUsers()
	c.JSON(http.StatusOK, orgs)
}

// Get ...
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Could not parse %s to int", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param not valid"})
		return
	}
	orgs := db.GetUser(id)
	c.JSON(http.StatusOK, orgs)
}

// Create ...
func Create(c *gin.Context) {

	c.JSON(200, "File saved")
}
