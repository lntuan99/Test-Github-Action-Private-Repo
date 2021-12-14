package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
)

var version = "dev"

func main() {
	fmt.Printf("Version: %s\n", version)

	fmt.Println(hello())

	r := gin.New()

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowHeaders = []string{"authorization", "Authorization", "content-type",
		"accept", "referer", "user-agent"}
	r.Use(cors.New(corConfig))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "CICD newbie.AAAAAAAAA NICE ALO 1234"})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9696"
	}

	go func() {
		r.Run(":" + PORT)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

}

func hello() string {
	return "Hello Golang"
}
