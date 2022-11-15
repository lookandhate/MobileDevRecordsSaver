package controllers

import (
	"MobileDevRecordsSaver/db"
	"MobileDevRecordsSaver/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRecords
// GET /records
// Get all records
func GetRecords(c *gin.Context) {
	var records []db.Record
	db.DB.Find(&records)
	c.JSON(200, records)
}

// PostRecord
// POST /record
// Add new record
func PostRecord(c *gin.Context) {
	var input models.CreateRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		println(err.Error())
		println(input.Points, input.Timestamp)
		return
	}

	record := db.Record{Timestamp: input.Timestamp, Points: input.Points}
	db.DB.Create(&record)

	c.JSON(http.StatusOK, record)
}
