package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samctang/socoro_api/database"
	"github.com/samctang/socoro_api/models"
	"log"
	"net/http"
)

// GetOperations GET /operation
// Get all operations
func GetOperations(c *gin.Context) {

	rows, err := database.Conn.Query(context.Background(), "select id, operation_no from operations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowSlice []models.Operation
	for rows.Next() {
		var r models.Operation
		err := rows.Scan(&r.ID, &r.OperationNo)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rowSlice)
	c.JSON(http.StatusOK, gin.H{"data": rowSlice})
}