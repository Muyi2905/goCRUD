package main

import (

 "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
 "log"
)

func init()  {
	err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hjsfkjbfd",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}