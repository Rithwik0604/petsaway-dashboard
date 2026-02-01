<script lang="ts">
    import { goto } from "$app/navigation";
    import { apiFetch } from "$lib/api";

    let client = $state({
        client_name: "",
        client_phone: "",
        import_export: "import",
        import_fee: 0,
        export_fee: 0,
        after_hours_charges: 0,
        pet_name: "",
        species: "",
        gender: "male",
        breed: "",
        date_of_birth: "",
        microchip_number: "",
        titre: "",
        last_rabies_date: "",
        rabies_validity: "",
        documentation_status: false,
        rabies_vaccination_valid: false,
        other_vaccines_completed: "",
        health_certificate_issues: false,
        export_permit_approved: false,
        import_permit_approved: false,
        airline_approval_received: false,
        customs_clearance_done: false,
        origin_country: "",
        destination_country: "",
        forwarder_charges: 0,
        departure: "",
        airline: "",
        airline_charges: 0,
        crate_cost: 0,
        flight_number: "",
        type_of_travel: "",
        etd: "",
        eta: "",
        quoted_amount: 0,
        total_cost: 0,
        profit: 0,
        advanced_received: 0,
        balance_pending: 0,
        payment_status: "pending",
        remarks: "",
    });

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        try {
            const res = await apiFetch("/clients", { method: "POST", body: JSON.stringify(client) });
            if (res.ok) {
                goto("/");
            }
        } catch (err) {
            console.log(err);
        }
    }
</script>

<div class="container mx-auto p-8">
    <h1 class="text-3xl font-bold mb-6">Add New Client</h1>
    <form onsubmit={handleSubmit} class="space-y-8">
        <!-- Client Details -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Client Details</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <label for="client_name" class="block text-sm font-medium text-gray-700">Client Name</label>
                    <input
                        type="text"
                        id="client_name"
                        required
                        bind:value={client.client_name}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="client_phone" class="block text-sm font-medium text-gray-700">Client Phone</label>
                    <input
                        type="text"
                        id="client_phone"
                        bind:value={client.client_phone}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
            </div>
        </div>

        <!-- Pet Information -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Pet Information</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div>
                    <label for="pet_name" class="block text-sm font-medium text-gray-700">Pet Name</label>
                    <input
                        type="text"
                        id="pet_name"
                        bind:value={client.pet_name}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="species" class="block text-sm font-medium text-gray-700">Species</label>
                    <input
                        type="text"
                        id="species"
                        bind:value={client.species}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="gender" class="block text-sm font-medium text-gray-700">Gender</label>
                    <select
                        id="gender"
                        bind:value={client.gender}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    >
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                    </select>
                </div>
                <div>
                    <label for="breed" class="block text-sm font-medium text-gray-700">Breed</label>
                    <input
                        type="text"
                        id="breed"
                        bind:value={client.breed}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="date_of_birth" class="block text-sm font-medium text-gray-700">Date of Birth</label>
                    <input
                        type="date"
                        id="date_of_birth"
                        bind:value={client.date_of_birth}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="microchip_number" class="block text-sm font-medium text-gray-700"
                        >Microchip Number</label
                    >
                    <input
                        type="text"
                        id="microchip_number"
                        bind:value={client.microchip_number}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
            </div>
        </div>

        <!-- Health & Documentation -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Health & Documentation</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div>
                    <label for="titre" class="block text-sm font-medium text-gray-700">Titre Date</label>
                    <input
                        type="date"
                        id="titre"
                        bind:value={client.titre}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="last_rabies_date" class="block text-sm font-medium text-gray-700"
                        >Last Rabies Date</label
                    >
                    <input
                        type="date"
                        id="last_rabies_date"
                        bind:value={client.last_rabies_date}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="rabies_validity" class="block text-sm font-medium text-gray-700">Rabies Validity</label>
                    <input
                        type="date"
                        id="rabies_validity"
                        bind:value={client.rabies_validity}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="other_vaccines_completed" class="block text-sm font-medium text-gray-700"
                        >Other Vaccines</label
                    >
                    <input
                        type="text"
                        id="other_vaccines_completed"
                        bind:value={client.other_vaccines_completed}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                        placeholder="e.g., DHPPI, Lepto"
                    />
                </div>
            </div>
            <div class="mt-6 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
                <div class="flex items-center">
                    <input
                        id="documentation_status"
                        type="checkbox"
                        bind:checked={client.documentation_status}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="documentation_status" class="ml-2 block text-sm text-gray-900"
                        >Documentation Complete</label
                    >
                </div>
                <div class="flex items-center">
                    <input
                        id="rabies_vaccination_valid"
                        type="checkbox"
                        bind:checked={client.rabies_vaccination_valid}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="rabies_vaccination_valid" class="ml-2 block text-sm text-gray-900">Rabies Valid</label>
                </div>
                <div class="flex items-center">
                    <input
                        id="health_certificate_issues"
                        type="checkbox"
                        bind:checked={client.health_certificate_issues}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="health_certificate_issues" class="ml-2 block text-sm text-gray-900"
                        >Health Certificate Issues</label
                    >
                </div>
                <div class="flex items-center">
                    <input
                        id="export_permit_approved"
                        type="checkbox"
                        bind:checked={client.export_permit_approved}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="export_permit_approved" class="ml-2 block text-sm text-gray-900"
                        >Export Permit Approved</label
                    >
                </div>
                <div class="flex items-center">
                    <input
                        id="import_permit_approved"
                        type="checkbox"
                        bind:checked={client.import_permit_approved}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="import_permit_approved" class="ml-2 block text-sm text-gray-900"
                        >Import Permit Approved</label
                    >
                </div>
                <div class="flex items-center">
                    <input
                        id="airline_approval_received"
                        type="checkbox"
                        bind:checked={client.airline_approval_received}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="airline_approval_received" class="ml-2 block text-sm text-gray-900"
                        >Airline Approval Received</label
                    >
                </div>
                <div class="flex items-center">
                    <input
                        id="customs_clearance_done"
                        type="checkbox"
                        bind:checked={client.customs_clearance_done}
                        class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label for="customs_clearance_done" class="ml-2 block text-sm text-gray-900"
                        >Customs Clearance Done</label
                    >
                </div>
            </div>
        </div>

        <!-- Travel Details -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Travel Details</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div>
                    <label for="import_export" class="block text-sm font-medium text-gray-700">Type</label>
                    <select
                        id="import_export"
                        bind:value={client.import_export}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    >
                        <option value="import">Import</option>
                        <option value="export">Export</option>
                    </select>
                </div>
                <div>
                    <label for="origin_country" class="block text-sm font-medium text-gray-700">Origin Country</label>
                    <input
                        type="text"
                        id="origin_country"
                        bind:value={client.origin_country}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="destination_country" class="block text-sm font-medium text-gray-700"
                        >Destination Country</label
                    >
                    <input
                        type="text"
                        id="destination_country"
                        bind:value={client.destination_country}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="departure" class="block text-sm font-medium text-gray-700">Departure Date/Time</label>
                    <input
                        type="datetime-local"
                        id="departure"
                        bind:value={client.departure}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="airline" class="block text-sm font-medium text-gray-700">Airline</label>
                    <input
                        type="text"
                        id="airline"
                        bind:value={client.airline}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="flight_number" class="block text-sm font-medium text-gray-700">Flight Number</label>
                    <input
                        type="text"
                        id="flight_number"
                        bind:value={client.flight_number}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="type_of_travel" class="block text-sm font-medium text-gray-700">Type of Travel</label>
                    <input
                        type="text"
                        id="type_of_travel"
                        bind:value={client.type_of_travel}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                        placeholder="e.g., Cargo, Accompanied"
                    />
                </div>
                <div>
                    <label for="etd" class="block text-sm font-medium text-gray-700">ETD</label>
                    <input
                        type="time"
                        id="etd"
                        bind:value={client.etd}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="eta" class="block text-sm font-medium text-gray-700">ETA</label>
                    <input
                        type="time"
                        id="eta"
                        bind:value={client.eta}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
            </div>
        </div>

        <!-- Financials -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Financials</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <div>
                    <label for="import_fee" class="block text-sm font-medium text-gray-700">Import Fee</label>
                    <input
                        type="number"
                        step="0.01"
                        id="import_fee"
                        bind:value={client.import_fee}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="export_fee" class="block text-sm font-medium text-gray-700">Export Fee</label>
                    <input
                        type="number"
                        step="0.01"
                        id="export_fee"
                        bind:value={client.export_fee}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="after_hours_charges" class="block text-sm font-medium text-gray-700"
                        >After Hours Charges</label
                    >
                    <input
                        type="number"
                        step="0.01"
                        id="after_hours_charges"
                        bind:value={client.after_hours_charges}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="forwarder_charges" class="block text-sm font-medium text-gray-700"
                        >Forwarder Charges</label
                    >
                    <input
                        type="number"
                        step="0.01"
                        id="forwarder_charges"
                        bind:value={client.forwarder_charges}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="airline_charges" class="block text-sm font-medium text-gray-700">Airline Charges</label>
                    <input
                        type="number"
                        step="0.01"
                        id="airline_charges"
                        bind:value={client.airline_charges}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="crate_cost" class="block text-sm font-medium text-gray-700">Crate Cost</label>
                    <input
                        type="number"
                        step="0.01"
                        id="crate_cost"
                        bind:value={client.crate_cost}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="quoted_amount" class="block text-sm font-medium text-gray-700">Quoted Amount</label>
                    <input
                        type="number"
                        step="0.01"
                        id="quoted_amount"
                        bind:value={client.quoted_amount}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="total_cost" class="block text-sm font-medium text-gray-700">Total Cost</label>
                    <input
                        type="number"
                        step="0.01"
                        id="total_cost"
                        bind:value={client.total_cost}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="profit" class="block text-sm font-medium text-gray-700">Profit</label>
                    <input
                        type="number"
                        step="0.01"
                        id="profit"
                        bind:value={client.profit}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="advanced_received" class="block text-sm font-medium text-gray-700"
                        >Advanced Received</label
                    >
                    <input
                        type="number"
                        step="0.01"
                        id="advanced_received"
                        bind:value={client.advanced_received}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="balance_pending" class="block text-sm font-medium text-gray-700">Balance Pending</label>
                    <input
                        type="number"
                        step="0.01"
                        id="balance_pending"
                        bind:value={client.balance_pending}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                </div>
                <div>
                    <label for="payment_status" class="block text-sm font-medium text-gray-700">Payment Status</label>
                    <select
                        id="payment_status"
                        bind:value={client.payment_status}
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    >
                        <option value="pending">Pending</option>
                        <option value="paid">Paid</option>
                    </select>
                </div>
            </div>
        </div>

        <!-- Remarks -->
        <div class="p-6 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-semibold mb-4">Remarks</h2>
            <div>
                <label for="remarks" class="block text-sm font-medium text-gray-700">Remarks</label>
                <textarea
                    id="remarks"
                    bind:value={client.remarks}
                    rows="4"
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                ></textarea>
            </div>
        </div>

        <div class="flex justify-end">
            <button
                type="submit"
                class="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
                Add Client
            </button>
        </div>
    </form>
</div>
