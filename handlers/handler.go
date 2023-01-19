package handlers

import (
	"net/http"

	"github.com/PedroALeo/crudRedis/models"
	"github.com/gin-gonic/gin"
)

// GetAll handles the route to get all blocks on db
//
//	GET /blocks
//
// Returns a list of blocks as JSON or an empty list if the db is empty and OK Status
func GetAll(c *gin.Context) {
	blocks, err := models.GetAllBlocks()
	switch err {
	case models.ErrStatusInternalServerError:
		c.JSON(http.StatusInternalServerError, models.Block{})
	}
	c.JSON(http.StatusOK, blocks)
}
