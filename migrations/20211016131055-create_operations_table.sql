
-- +migrate Up
CREATE TABLE IF NOT EXISTS operations (
    id serial PRIMARY KEY,
    company_id INT NOT NULL,
    employee_id INT NOT NULL,
    operation_no text UNIQUE NOT NULL,
    operation_type_id INT NOT NULL,
    agent_id INT,
    shipper_id INT,
    consignee_id INT,
    carrier_id INT,
    agent_ref_no text,
    shipper_ref_no text,
    consignee_ref_no text,
    port_of_loading text,
    port_of_departure text,
    origin_addr1 text,
    origin_addr2 text,
    origin_city text,
    origin_state text,
    origin_zip text,
    origin_country text,
    destination_addr1 text,
    destination_addr2 text,
    destination_city text,
    destination_state text,
    destination_zip text,
    destination_country text,
    progress INT NOT NULL DEFAULT 1,
    status INT NOT NULL DEFAULT 1,
    submitted_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_date TIMESTAMP,
    last_modified_by text NOT NULL,
    last_modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- +migrate Down
DROP TABLE operations;