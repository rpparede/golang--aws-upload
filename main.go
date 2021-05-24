  
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rpparede/golang--aws-upload/handler"
)

func main() {
	//LoadEnv()

	sess := handler.ConnectAws()
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	router.POST("/upload", handler.UploadImage)
	_ = router.Run(":4000")
}