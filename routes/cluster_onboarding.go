package routes

import (
	"github.com/gin-gonic/gin"
)

func clusterOnboarding(rg *gin.RouterGroup) {
	configureOperatorsRG := rg.Group("/onboard/clusters")
	configureOperatorsRG.POST("/", controllers.ClusterOnboarding)
}
