package models

import (
	"time"
)

type Operation struct {
	ID                 int       `json:"id"`
	CompanyId          int       `json:"company_id"`
	EmployeeId         int       `json:"employee_id"`
	OperationNo        string    `json:"operation_no"`
	OperationTypeId    int       `json:"operation_type_id"`
	AgentId            int       `json:"agent_id"`
	ShipperId          int       `json:"shipper_id"`
	ConsigneeId        int       `json:"consignee_id"`
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
	Progress           int       `json:"progress"`
	Status             int       `json:"status"`
	SubmittedDate      time.Time `json:"submitted_date"`
	CompletedDate      time.Time `json:"completed_date"`
	LastModifiedBy     string    `json:"last_modified_by"`
	LastModifiedOn     time.Time `json:"last_modified_on"`
}
