-- name: GetAllClients :many
SELECT
    *
FROM
    clients;

-- name: GetClientById :one
SELECT
    *
FROM
    clients
WHERE
    id = ?;

-- name: GetClientByName :many
SELECT
    *
FROM
    clients
WHERE
    client_name = ?;

-- name: InsertClient :one
INSERT INTO
    clients (
        id,
        client_name,
        client_phone,
        import_export,
        import_fee,
        export_fee,
        after_hours_charges,
        pet_name,
        species,
        gender,
        breed,
        date_of_birth,
        microchip_number,
        titre,
        last_rabies_date,
        rabies_validity,
        documentation_status,
        rabies_vaccination_valid,
        other_vaccines_completed,
        health_certificate_issues,
        export_permit_approved,
        import_permit_approved,
        airline_approval_received,
        customs_clearance_done,
        origin_country,
        destination_country,
        forwarder_charges,
        departure,
        airline,
        airline_charges,
        crate_cost,
        flight_number,
        type_of_travel,
        etd,
        eta,
        quoted_amount,
        total_cost,
        profit,
        advanced_received,
        balance_pending,
        payment_status,
        remarks
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    ) RETURNING *;

-- name: UpdateClientById :one
UPDATE clients
SET
    client_name = ?,
    client_phone = ?,
    import_export = ?,
    import_fee = ?,
    export_fee = ?,
    after_hours_charges = ?,
    pet_name = ?,
    species = ?,
    gender = ?,
    breed = ?,
    date_of_birth = ?,
    microchip_number = ?,
    titre = ?,
    last_rabies_date = ?,
    rabies_validity = ?,
    documentation_status = ?,
    rabies_vaccination_valid = ?,
    other_vaccines_completed = ?,
    health_certificate_issues = ?,
    export_permit_approved = ?,
    import_permit_approved = ?,
    airline_approval_received = ?,
    customs_clearance_done = ?,
    origin_country = ?,
    destination_country = ?,
    forwarder_charges = ?,
    departure = ?,
    airline = ?,
    airline_charges = ?,
    crate_cost = ?,
    flight_number = ?,
    type_of_travel = ?,
    etd = ?,
    eta = ?,
    quoted_amount = ?,
    total_cost = ?,
    profit = ?,
    advanced_received = ?,
    balance_pending = ?,
    payment_status = ?,
    remarks = ?
WHERE
    id = ? RETURNING *;