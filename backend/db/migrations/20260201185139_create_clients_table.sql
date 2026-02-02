-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    clients (
        id TEXT PRIMARY KEY,
        client_name TEXT,
        client_phone TEXT,
        import_export TEXT CHECK (import_export IN ('import', 'export')),
        import_fee REAL DEFAULT 0.0,
        export_fee REAL DEFAULT 0.0,
        after_hours_charges REAL DEFAULT 0.0,
        pet_name TEXT,
        species TEXT,
        gender TEXT CHECK (gender IN ('male', 'female')),
        breed TEXT,
        date_of_birth DATE,
        microchip_number TEXT,
        titre DATE,
        last_rabies_date DATE,
        rabies_validity DATE,
        documentation_status INTEGER DEFAULT 0,
        rabies_vaccination_valid INTEGER DEFAULT 0,
        other_vaccines_completed TEXT,
        health_certificate_issues INTEGER DEFAULT 0,
        export_permit_approved INTEGER DEFAULT 0,
        import_permit_approved INTEGER DEFAULT 0,
        airline_approval_received INTEGER DEFAULT 0,
        customs_clearance_done INTEGER DEFAULT 0,
        origin_country TEXT,
        destination_country TEXT,
        forwarder_charges REAL DEFAULT 0.0,
        departure DATETIME,
        airline TEXT,
        airline_charges REAL DEFAULT 0.0,
        crate_cost REAL DEFAULT 0.0,
        flight_number TEXT,
        type_of_travel TEXT,
        etd TIME,
        eta TIME,
        quoted_amount REAL DEFAULT 0.0,
        total_cost REAL DEFAULT 0.0,
        profit REAL DEFAULT 0.0,
        advanced_received REAL DEFAULT 0.0,
        balance_pending REAL DEFAULT 0.0,
        payment_status TEXT CHECK (payment_status IN ('paid', 'pending')) DEFAULT 'pending',
        remarks TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

CREATE TRIGGER update_client_timestamp AFTER
UPDATE ON clients FOR EACH ROW WHEN NEW.updated_at <= OLD.updated_at BEGIN
UPDATE clients
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = NEW.id;

END;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;

-- +goose StatementEnd