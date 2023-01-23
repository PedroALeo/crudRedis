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

// GetById handles the route to get a block from the db using id
//
// GET /blocks/id
//
// Returns a block as JSON or an empty block if the db is empty and OK Status
func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	block, err := models.GetBlockByID(id)

	switch err {
	case models.ErrStatusInternalServerError:
		c.JSON(http.StatusInternalServerError, models.Block{})
	case models.ErrStatusNotFound:
		c.JSON(http.StatusNotFound, models.Block{})
	}

	c.JSON(http.StatusOK, block)
}

// GetTree handles the route to get a tree using a block as root specified by id
//
// GET /blocks/tree/id
//
// Returns a tree as JSON or an empty block if the db is empty and OK Status
func GetTree(c *gin.Context) {
	id := c.Params.ByName("id")
	tree := models.GetTreeByID(id)
	c.JSON(http.StatusOK, tree)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	err := models.DeleteBlock(id)

	switch err {
	case models.ErrStatusNotFound:
		c.JSON(http.StatusNotFound, nil)
	default:
		c.JSON(http.StatusOK, nil)
	}
}

func Put(c *gin.Context) {
	var block models.Block
	id := c.Params.ByName("id")
	block.ID = id
	if err := c.ShouldBindJSON(&block); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error()})
		return
	}
	block, err := models.PutBlock(&block)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "nonExistant key on database",
		})
	}
	c.JSON(http.StatusOK, block)

}
