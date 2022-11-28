package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Heyya",
	})

}
func main() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(false)
	log.SetLevel(log.InfoLevel)
	log.Trace("Starting the go gin application")
	defer log.Warn("Shutting down the gin application server")

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/home", HomeHandler)
	log.Fatal(r.Run(":8080")) //listen & serve on 0.0.0.0:8080
}
