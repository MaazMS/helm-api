package routes

import (
	"github.com/gin-gonic/gin"
)

var (

	router = gin.Default()
)

func Run(){
	router.Use(CORS())
	getRoutes()
	if err := router.Run(); err != nil {
		panic(err)
	}
}

func getRoutes() {
	v1 := router.Group("/v1")
	clusterOnboarding(v1)
}

func CORS() gin.HandlerFunc {
	return func (c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}