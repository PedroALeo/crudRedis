package routes

import (
	"github.com/PedroALeo/crudRedis/handlers"
	"github.com/gin-gonic/gin"
)

// HandleRoutes handles the application routes
func RequestContollers() {
	r := gin.Default()
	r.GET("/blocks", handlers.GetAll)
	r.GET("/blocks/:id", handlers.GetById)
	r.GET("/blocks/tree/:id", handlers.GetTree)
	r.DELETE("/blocks/:id", handlers.Delete)
	r.PUT("/blocks/:id", handlers.Put)
	r.POST("/blocks", handlers.PostBlock)
	r.Run()
}
