package main

import (
	"MobileDevRecordsSaver/controllers"
	"MobileDevRecordsSaver/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/records", controllers.GetRecords)

	r.POST("/record", controllers.PostRecord)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
