package operation

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/samctang/socoro_api/database"
	"log"
	"net/http"
	"time"
)

func GetOperations(c echo.Context) error {

	rows, err := database.Conn.Query(context.Background(), "select id, operation_no, operation_type_id, progress, status, agent_id, shipper_id, submitted_date from operations;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowSlice []Operation
	for rows.Next() {
		var r Operation

		err := rows.Scan(&r.Id, &r.OperationNo,&r.OperationTypeId,&r.Progress,&r.Status, &r.AgentId, &r.ShipperId,&r.SubmittedDate)
		if err != nil {
			log.Fatal(err) //add conditions for httputil
		}

		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, rowSlice)
}

type Operation struct {
	Id                 int       `json:"id"`
	CompanyId          int       `json:"company_id" validate:"required"`
	UserId		       int       `json:"user_id" validate:"required"`
	OperationNo        string    `json:"operation_no" validate:"required"`
	OperationTypeId    int       `json:"operation_type_id" validate:"required"`
	AgentId            int       `json:"agent_id" validate:"required"`
	ShipperId          int       `json:"shipper_id" validate:"required"`
	ConsigneeId        int       `json:"consignee_id" validate:"required"`
	CarrierId          int       `json:"carrier_id"`
	AgentRefNo         string    `json:"agent_ref_no"`
	ShipperRefNo       string    `json:"shipper_ref_no"`
	ConsigneeRefNo     string    `json:"consignee_ref_no"`
	PortOfLoading      string    `json:"port_of_loading"`
	PortOfDeparture    string    `json:"port_of_departure"`
	OriginAddr1        string    `json:"origin_addr1"`
	OriginAddr2        string    `json:"origin_addr2"`
	OriginCity         string    `json:"origin_city"`
	OriginState        string    `json:"origin_state"`
	OriginZip          string    `json:"origin_zip"`
	OriginCountry      string    `json:"origin_country"`
	DestinationAddr1   string    `json:"destination_addr1"`
	DestinationAddr2   string    `json:"destination_addr2"`
	DestinationCity    string    `json:"destination_city"`
	DestinationState   string    `json:"destination_state"`
	DestinationZip     string    `json:"destination_zip"`
	DestinationCountry string    `json:"destination_country"`
	Progress           int       `json:"progress" validate:"required"`
	Status             int       `json:"status" validate:"required"`
	SubmittedDate      time.Time `json:"submitted_date" validate:"required"`
	CompletedDate      time.Time `json:"completed_date"`
	LastModifiedBy     string    `json:"last_modified_by" validate:"required"`
	LastModifiedOn     time.Time `json:"last_modified_on" validate:"required"`
}
