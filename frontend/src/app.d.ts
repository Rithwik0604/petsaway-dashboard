// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        // interface Locals {}
        // interface PageData {}
        // interface PageState {}
        // interface Platform {}
        interface ClientData {
            id: string; // Assuming id is always present and valid
            client_name: { String: string; Valid: boolean };
            client_phone: { String: string; Valid: boolean };
            import_export: { String: "import" | "export"; Valid: boolean };
            import_fee: { Float64: number; Valid: boolean };
            export_fee: { Float64: number; Valid: boolean };
            after_hours_charges: { Float64: number; Valid: boolean };
            pet_name: { String: string; Valid: boolean };
            species: { String: string; Valid: boolean };
            gender: { String: "male" | "female"; Valid: boolean };
            breed: { String: string; Valid: boolean };
            date_of_birth: { Time: string; Valid: boolean };
            microchip_number: { String: string; Valid: boolean };
            titre: { Time: string; Valid: boolean };
            last_rabies_date: { Time: string; Valid: boolean };
            rabies_validity: { Time: string; Valid: boolean };
            documentation_status: { Int64: number; Valid: boolean };
            rabies_vaccination_valid: { Int64: number; Valid: boolean };
            other_vaccines_completed: { String: string; Valid: boolean };
            health_certificate_issues: { Int64: number; Valid: boolean };
            export_permit_approved: { Int64: number; Valid: boolean };
            import_permit_approved: { Int64: number; Valid: boolean };
            airline_approval_received: { Int64: number; Valid: boolean };
            customs_clearance_done: { Int64: number; Valid: boolean };
            origin_country: { String: string; Valid: boolean };
            destination_country: { String: string; Valid: boolean };
            forwarder_charges: { Float64: number; Valid: boolean };
            departure: { Time: string; Valid: boolean };
            airline: { String: string; Valid: boolean };
            airline_charges: { Float64: number; Valid: boolean };
            crate_cost: { Float64: number; Valid: boolean };
            flight_number: { String: string; Valid: boolean };
            type_of_travel: { String: string; Valid: boolean };
            etd: string;
            eta: string;
            quoted_amount: { Float64: number; Valid: boolean };
            total_cost: { Float64: number; Valid: boolean };
            profit: { Float64: number; Valid: boolean };
            advanced_received: { Float64: number; Valid: boolean };
            balance_pending: { Float64: number; Valid: boolean };
            payment_status: { String: "paid" | "pending"; Valid: boolean };
            remarks: { String: string; Valid: boolean };
            created_at: { Time: string; Valid: boolean };
            updated_at: { Time: string; Valid: boolean };
        }

        interface Client {
            id: string; // Assuming id is always present and valid
            client_name: string;
            client_phone: string;
            import_export: "import" | "export";
            import_fee: number;
            export_fee: number;
            after_hours_charges: number;
            pet_name: string;
            species: string;
            gender: "male" | "female";
            breed: string;
            date_of_birth: string;
            microchip_number: string;
            titre: string;
            last_rabies_date: string;
            rabies_validity: string;
            documentation_status: number;
            rabies_vaccination_valid: number;
            other_vaccines_completed: string;
            health_certificate_issues: number;
            export_permit_approved: number;
            import_permit_approved: number;
            airline_approval_received: number;
            customs_clearance_done: number;
            origin_country: string;
            destination_country: string;
            forwarder_charges: number;
            departure: string;
            airline: string;
            airline_charges: number;
            crate_cost: number;
            flight_number: string;
            type_of_travel: string;
            etd: string;
            eta: string;
            quoted_amount: number;
            total_cost: number;
            profit: number;
            advanced_received: number;
            balance_pending: number;
            payment_status: "paid" | "pending";
            remarks: string;
            created_at: string;
            updated_at: string;
        }
    }
}

export {};
