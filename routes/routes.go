package routes

import (
	"github.com/PedroALeo/crudRedis/handlers"
	"github.com/gin-gonic/gin"
)

func RequestContollers() {
	r := gin.Default()
	r.GET("/blocks", handlers.GetAll)
	r.GET("/blocks/:id")
	r.GET("/blocks/tree/:id")
	r.POST("/blocks")
	r.PUT("/blocks/:id")
	r.DELETE("/blocks/:id")
	r.Run()
}
