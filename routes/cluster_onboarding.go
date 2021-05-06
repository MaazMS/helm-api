package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tanalam2411/gripo_rest/pkg/controllers"
)

func clusterOnboarding(rg *gin.RouterGroup) {
	configureOperatorsRG := rg.Group("/onboard/clusters")
	configureOperatorsRG.POST("/", controllers.ClusterOnboarding)
}
