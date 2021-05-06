package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ClusterOnboarding(c *gin.Context) {

	configFile, _ := c.FormFile("file")
	log.Println(configFile.Filename)

	fileObj, err := configFile.Open()
	if err != nil {
		log.Printf("Failed to open file: %s", err)
	}

	configData, err := file.ReadAll(fileObj)
	if err != nil {
		log.Printf("Failed to read file: %s", err)
	}

	err = file.Write("./cluster_configs/config", configData)
	if err != nil {
		log.Printf("Failed to write file: %s", err)
	}

	c.JSON(http.StatusOK, "Cluster Onboarded")
}
