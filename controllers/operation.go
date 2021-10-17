package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/samctang/socoro_api/database"
	"github.com/samctang/socoro_api/models"
	"log"
	"net/http"
)

// swagger:operation GET /operation operation GetOperations
// Gets all operations
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Success
//     '400':
//         description: Bad request
//     '500':
//         description: Internal server error
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
			log.Fatal(err) //add conditions for httputil
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": rowSlice})
}