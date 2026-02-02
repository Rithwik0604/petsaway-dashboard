<script lang="ts">
    import { goto } from "$app/navigation";
    import { apiFetch } from "$lib/api";
    import { page } from "$app/stores";
    import { formatDateForInput, formatTimeForInput, formatDateTimeForInput } from "$lib/format";

    const { data } = $props();
    const clientData = data.data;

    let client = $state({
        client_name: clientData.client_name.String,
        client_phone: clientData.client_phone.String,
        import_export: clientData.import_export.String,
        import_fee: clientData.import_fee.Float64,
        export_fee: clientData.export_fee.Float64,
        after_hours_charges: clientData.after_hours_charges.Float64,
        pet_name: clientData.pet_name.String,
        species: clientData.species.String,
        gender: clientData.gender.String,
        breed: clientData.breed.String,
        date_of_birth: formatDateForInput(clientData.date_of_birth.Time),
        microchip_number: clientData.microchip_number.String,
        titre: formatDateForInput(clientData.titre.Time),
        last_rabies_date: formatDateForInput(clientData.last_rabies_date.Time),
        rabies_validity: formatDateForInput(clientData.rabies_validity.Time),
        documentation_status: clientData.documentation_status.Int64 === 1,
        rabies_vaccination_valid: clientData.rabies_vaccination_valid.Int64 === 1,
        other_vaccines_completed: clientData.other_vaccines_completed.String,
        health_certificate_issues: clientData.health_certificate_issues.Int64 === 1,
        export_permit_approved: clientData.export_permit_approved.Int64 === 1,
        import_permit_approved: clientData.import_permit_approved.Int64 === 1,
        airline_approval_received: clientData.airline_approval_received.Int64 === 1,
        customs_clearance_done: clientData.customs_clearance_done.Int64 === 1,
        origin_country: clientData.origin_country.String,
        destination_country: clientData.destination_country.String,
        forwarder_charges: clientData.forwarder_charges.Float64,
        departure: formatDateTimeForInput(clientData.departure.Time),
        airline: clientData.airline.String,
        airline_charges: clientData.airline_charges.Float64,
        crate_cost: clientData.crate_cost.Float64,
        flight_number: clientData.flight_number.String,
        type_of_travel: clientData.type_of_travel.String,
        etd: formatTimeForInput(clientData.etd),
        eta: formatTimeForInput(clientData.eta),
        quoted_amount: clientData.quoted_amount.Float64,
        total_cost: clientData.total_cost.Float64,
        profit: clientData.profit.Float64,
        advanced_received: clientData.advanced_received.Float64,
        balance_pending: clientData.balance_pending.Float64,
        payment_status: clientData.payment_status.String,
        remarks: clientData.remarks.String,
    });

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        try {
            const res = await apiFetch(`/clients/${$page.params.id}`, {
                method: "PATCH",
                body: JSON.stringify(client),
            });
            if (res.ok) {
                goto("/");
            }
        } catch (err) {
            console.log(err);
        }
    }
</script>

<div class="container p-8 mx-auto">
    <h1 class="mb-6 text-3xl font-bold">Edit Client</h1>
    <form onsubmit={handleSubmit} class="space-y-8">
        <!-- Client Details -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Client Details</h2>
                <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Client Name</span>
                        </div>
                        <input
                            type="text"
                            required
                            bind:value={client.client_name}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Client Phone</span>
                        </div>
                        <input type="text" bind:value={client.client_phone} class="w-full input input-bordered" />
                    </label>
                </div>
            </div>
        </div>

        <!-- Pet Information -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Pet Information</h2>
                <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Pet Name</span>
                        </div>
                        <input type="text" bind:value={client.pet_name} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Species</span>
                        </div>
                        <input type="text" bind:value={client.species} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Gender</span>
                        </div>
                        <select bind:value={client.gender} class="w-full select select-bordered">
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                        </select>
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Breed</span>
                        </div>
                        <input type="text" bind:value={client.breed} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Date of Birth</span>
                        </div>
                        <input type="date" bind:value={client.date_of_birth} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Microchip Number</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.microchip_number}
                            class="w-full input input-bordered"
                        />
                    </label>
                </div>
            </div>
        </div>

        <!-- Health & Documentation -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Health & Documentation</h2>
                <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Titre Date</span>
                        </div>
                        <input type="date" bind:value={client.titre} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Last Rabies Date</span>
                        </div>
                        <input
                            type="date"
                            bind:value={client.last_rabies_date}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Rabies Validity</span>
                        </div>
                        <input
                            type="date"
                            bind:value={client.rabies_validity}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Other Vaccines</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.other_vaccines_completed}
                            class="w-full input input-bordered"
                            placeholder="e.g., DHPPI, Lepto"
                        />
                    </label>
                </div>
                <div class="grid grid-cols-2 gap-4 mt-6 md:grid-cols-3 lg:grid-cols-4">
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Documentation Complete</span>
                            <input
                                type="checkbox"
                                bind:checked={client.documentation_status}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Rabies Valid</span>
                            <input
                                type="checkbox"
                                bind:checked={client.rabies_vaccination_valid}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Health Certificate Issues</span>
                            <input
                                type="checkbox"
                                bind:checked={client.health_certificate_issues}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Export Permit Approved</span>
                            <input
                                type="checkbox"
                                bind:checked={client.export_permit_approved}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Import Permit Approved</span>
                            <input
                                type="checkbox"
                                bind:checked={client.import_permit_approved}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Airline Approval Received</span>
                            <input
                                type="checkbox"
                                bind:checked={client.airline_approval_received}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                    <div class="form-control">
                        <label class="gap-2 cursor-pointer label">
                            <span class="label-text">Customs Clearance Done</span>
                            <input
                                type="checkbox"
                                bind:checked={client.customs_clearance_done}
                                class="checkbox checkbox-primary"
                            />
                        </label>
                    </div>
                </div>
            </div>
        </div>

        <!-- Travel Details -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Travel Details</h2>
                <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Type</span>
                        </div>
                        <select bind:value={client.import_export} class="w-full select select-bordered">
                            <option value="import">Import</option>
                            <option value="export">Export</option>
                        </select>
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Origin Country</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.origin_country}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Destination Country</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.destination_country}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Departure Date/Time</span>
                        </div>
                        <input
                            type="datetime-local"
                            bind:value={client.departure}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Airline</span>
                        </div>
                        <input type="text" bind:value={client.airline} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Flight Number</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.flight_number}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Type of Travel</span>
                        </div>
                        <input
                            type="text"
                            bind:value={client.type_of_travel}
                            class="w-full input input-bordered"
                            placeholder="e.g., Cargo, Accompanied"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">ETD</span>
                        </div>
                        <input type="time" bind:value={client.etd} class="w-full input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">ETA</span>
                        </div>
                        <input type="time" bind:value={client.eta} class="w-full input input-bordered" />
                    </label>
                </div>
            </div>
        </div>

        <!-- Financials -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Financials</h2>
                <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-4">
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Import Fee</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.import_fee}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Export Fee</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.export_fee}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">After Hours Charges</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.after_hours_charges}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Forwarder Charges</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.forwarder_charges}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Airline Charges</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.airline_charges}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Crate Cost</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.crate_cost}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Quoted Amount</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.quoted_amount}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Total Cost</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.total_cost}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Profit</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.profit}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Advanced Received</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.advanced_received}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Balance Pending</span>
                        </div>
                        <input
                            type="number"
                            step="0.01"
                            bind:value={client.balance_pending}
                            class="w-full input input-bordered"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="label-text">Payment Status</span>
                        </div>
                        <select bind:value={client.payment_status} class="w-full select select-bordered">
                            <option value="pending">Pending</option>
                            <option value="paid">Paid</option>
                        </select>
                    </label>
                </div>
            </div>
        </div>

        <!-- Remarks -->
        <div class="shadow-xl card bg-base-100">
            <div class="card-body">
                <h2 class="card-title">Remarks</h2>
                <label class="w-full form-control">
                    <div class="label">
                        <span class="label-text">Remarks</span>
                    </div>
                    <textarea
                        bind:value={client.remarks}
                        rows="4"
                        class="w-full textarea textarea-bordered"
                    ></textarea>
                </label>
            </div>
        </div>

        <div class="flex justify-end">
            <button type="submit" class="btn btn-primary"> Update Client </button>
        </div>
    </form>
</div>