package api

import (
	"strconv"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gin-gonic/gin"
)

// API returns handler for Gin routes
type API struct {
	Repository repository.Interface
}

// GetUser handles request to get all users
func (api API) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{"error": "User ID is not an integer."})
		return
	}

	user, err := api.Repository.GetUser(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "User not found."})
		return
	}

	c.JSON(200, user)
}
