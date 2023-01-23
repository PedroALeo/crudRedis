package routes

import (
	"github.com/PedroALeo/crudRedis/handlers"
	"github.com/gin-gonic/gin"
)

func RequestContollers() {
	r := gin.Default()
	r.GET("/blocks", handlers.GetAll)
	r.GET("/blocks/:id", handlers.GetById)
	r.GET("/blocks/tree/:id", handlers.GetTree)
	r.DELETE("/blocks/:id", handlers.Delete)
	r.POST("/blocks")
	r.PUT("/blocks/:id")
	r.Run()
}
