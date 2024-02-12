
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	// register routes
	router.POST("/ping", test)
	// router.GET("/main", getmainHandler)
	// router.POST("/main", createmainHandler)
	// router.PUT("/main/:main-id", updatemainByIDHandler)
	// router.DELETE("/main/:main-id", deletemainByIDHandler)
	
	router.Run(":8060")
	// gin.SetMode(gin.ReleaseMode)
}

func test(c * gin.Context) {
	c.JSON(http.StatusOK, gin.H{"kuy": "test"})
	
}