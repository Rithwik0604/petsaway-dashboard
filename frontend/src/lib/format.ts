export function formatDate(dateString: string) {
    if (!dateString) return "";
    const date = new Date(dateString);

    return new Intl.DateTimeFormat("en-US", {
        day: "numeric",
        month: "short",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    }).format(date);
}

export function formatDateForInput(dateString: string): string {
    if (!dateString) return "";
    const date = new Date(dateString);
    return date.toISOString().split("T")[0];
}

export function formatTimeForInput(timeValue: any): string {
    if (!timeValue) return "";

    // If Go sends "17:12:00", we need to trim it to "17:12"
    if (typeof timeValue === "string") {
        return timeValue.slice(0, 5);
    }

    return "";
}

export function formatDateTimeForInput(dateTimeString: string): string {
    if (!dateTimeString) return "";
    const date = new Date(dateTimeString);
    // Format to "YYYY-MM-DDTHH:mm"
    return date.toISOString().slice(0, 16);
}

export function formatTo12Hour(timeStr: string): string {
    if (!timeStr) return "";

    // Split "17:12" into [17, 12]
    const [hours, minutes] = timeStr.split(":").map(Number);
    if (isNaN(hours) || isNaN(minutes)) return "";

    const suffix = hours >= 12 ? "PM" : "AM";
    const adjustedHour = hours % 12 || 12; // Converts 0 to 12 and 13 to 1

    // Return formatted string: "05:12 PM"
    return `${adjustedHour.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")} ${suffix}`;
}

export function formatClientData(data: App.ClientData[]): App.Client[] {
    return data.map((client) => ({
        id: client.id,
        client_name: client.client_name.Valid ? client.client_name.String : "",
        client_phone: client.client_phone.Valid ? client.client_phone.String : "",
        import_export: client.import_export.Valid ? client.import_export.String : "import",
        import_fee: client.import_fee.Valid ? client.import_fee.Float64 : 0,
        export_fee: client.export_fee.Valid ? client.export_fee.Float64 : 0,
        after_hours_charges: client.after_hours_charges.Valid ? client.after_hours_charges.Float64 : 0,
        pet_name: client.pet_name.Valid ? client.pet_name.String : "",
        species: client.species.Valid ? client.species.String : "",
        gender: client.gender.Valid ? client.gender.String : "male",
        breed: client.breed.Valid ? client.breed.String : "",
        date_of_birth: client.date_of_birth.Valid ? client.date_of_birth.Time : "",
        microchip_number: client.microchip_number.Valid ? client.microchip_number.String : "",
        titre: client.titre.Valid ? client.titre.Time : "",
        last_rabies_date: client.last_rabies_date.Valid ? client.last_rabies_date.Time : "",
        rabies_validity: client.rabies_validity.Valid ? client.rabies_validity.Time : "",
        documentation_status: client.documentation_status.Valid ? client.documentation_status.Int64 : 0,
        rabies_vaccination_valid: client.rabies_vaccination_valid.Valid ? client.rabies_vaccination_valid.Int64 : 0,
        other_vaccines_completed: client.other_vaccines_completed.Valid ? client.other_vaccines_completed.String : "",
        health_certificate_issues: client.health_certificate_issues.Valid ? client.health_certificate_issues.Int64 : 0,
        export_permit_approved: client.export_permit_approved.Valid ? client.export_permit_approved.Int64 : 0,
        import_permit_approved: client.import_permit_approved.Valid ? client.import_permit_approved.Int64 : 0,
        airline_approval_received: client.airline_approval_received.Valid ? client.airline_approval_received.Int64 : 0,
        customs_clearance_done: client.customs_clearance_done.Valid ? client.customs_clearance_done.Int64 : 0,
        origin_country: client.origin_country.Valid ? client.origin_country.String : "",
        destination_country: client.destination_country.Valid ? client.destination_country.String : "",
        forwarder_charges: client.forwarder_charges.Valid ? client.forwarder_charges.Float64 : 0,
        departure: client.departure.Valid ? client.departure.Time : "",
        airline: client.airline.Valid ? client.airline.String : "",
        airline_charges: client.airline_charges.Valid ? client.airline_charges.Float64 : 0,
        crate_cost: client.crate_cost.Valid ? client.crate_cost.Float64 : 0,
        flight_number: client.flight_number.Valid ? client.flight_number.String : "",
        type_of_travel: client.type_of_travel.Valid ? client.type_of_travel.String : "",
        etd: client.etd,
        eta: client.eta,
        quoted_amount: client.quoted_amount.Valid ? client.quoted_amount.Float64 : 0,
        total_cost: client.total_cost.Valid ? client.total_cost.Float64 : 0,
        profit: client.profit.Valid ? client.profit.Float64 : 0,
        advanced_received: client.advanced_received.Valid ? client.advanced_received.Float64 : 0,
        balance_pending: client.balance_pending.Valid ? client.balance_pending.Float64 : 0,
        payment_status: client.payment_status.Valid ? client.payment_status.String : "paid",
        remarks: client.remarks.Valid ? client.remarks.String : "",
        created_at: client.created_at.Valid ? client.created_at.Time : "",
        updated_at: client.updated_at.Valid ? client.updated_at.Time : "",
    }));
}
