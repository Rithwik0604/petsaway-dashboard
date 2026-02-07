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
        microchip_validity: "",
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

    /**
     * Manual Trigger for Financials
     */
    function runCalculations() {
        client.total_cost =
            Number(client.import_fee) +
            Number(client.export_fee) +
            Number(client.forwarder_charges) +
            Number(client.after_hours_charges) +
            Number(client.airline_charges) +
            Number(client.crate_cost);

        client.profit = Number(client.quoted_amount) - client.total_cost;
        client.balance_pending = Number(client.quoted_amount) - Number(client.advanced_received);
    }

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        try {
            const res = await apiFetch("/clients", {
                method: "POST",
                body: JSON.stringify(client),
            });
            if (res.ok) {
                goto("/");
            }
        } catch (err) {
            console.error(err);
        }
    }
</script>

<div class="container max-w-6xl p-8 mx-auto">
    <div class="flex items-center justify-between mb-8">
        <h1 class="text-3xl font-bold">Add New Client</h1>
        <button type="button" class="btn btn-ghost" onclick={() => goto("/")}>Cancel</button>
    </div>

    <form onsubmit={handleSubmit} class="pb-20 space-y-8">
        <section class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div class="md:col-span-1">
                <h2 class="text-xl font-semibold">Client Identity</h2>
                <p class="text-sm opacity-60">Primary contact and service type.</p>
            </div>
            <div class="shadow-xl md:col-span-2 card bg-base-100">
                <div class="grid grid-cols-1 gap-4 card-body sm:grid-cols-2">
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Full Name</span></div>
                        <input
                            type="text"
                            required
                            bind:value={client.client_name}
                            class="input input-bordered focus:input-primary"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Phone Number</span></div>
                        <input type="text" bind:value={client.client_phone} class="input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="font-medium label-text text-secondary">Service Type</span>
                        </div>
                        <select bind:value={client.import_export} class="select select-bordered border-secondary/40">
                            <option value="import">Import</option>
                            <option value="export">Export</option>
                        </select>
                    </label>
                </div>
            </div>
        </section>

        <hr class="opacity-10" />

        <section class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div class="md:col-span-1">
                <h2 class="text-xl font-semibold">Pet Information</h2>
                <p class="text-sm opacity-60">Specify pet details and age.</p>
            </div>
            <div class="shadow-xl md:col-span-2 card bg-base-100">
                <div class="grid grid-cols-1 gap-4 card-body sm:grid-cols-2 lg:grid-cols-3">
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Pet Name</span></div>
                        <input type="text" bind:value={client.pet_name} class="input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Species</span></div>
                        <input type="text" bind:value={client.species} class="input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Gender</span></div>
                        <select bind:value={client.gender} class="select select-bordered">
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                        </select>
                    </label>
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Breed</span></div>
                        <input type="text" bind:value={client.breed} class="input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="font-medium label-text text-primary">Date of Birth</span>
                        </div>
                        <input
                            type="date"
                            bind:value={client.date_of_birth}
                            onclick={(e) => e.currentTarget.showPicker()}
                            class="input input-bordered border-primary/30"
                        />
                    </label>
                    <label class="w-full form-control">
                        <div class="label"><span class="font-medium label-text">Microchip #</span></div>
                        <input type="text" bind:value={client.microchip_number} class="input input-bordered" />
                    </label>
                    <label class="w-full form-control">
                        <div class="label">
                            <span class="font-medium label-text text-primary">Microchip Validity</span>
                        </div>
                        <input
                            type="date"
                            bind:value={client.microchip_validity}
                            onclick={(e) => e.currentTarget.showPicker()}
                            class="input input-bordered border-primary/30"
                        />
                    </label>
                </div>
            </div>
        </section>

        <hr class="opacity-10" />

        <section class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div class="md:col-span-1">
                <h2 class="text-xl font-semibold">Medical & Documentation</h2>
                <p class="text-sm opacity-60">Vaccination history and check-list.</p>
            </div>
            <div class="border-l-4 shadow-xl md:col-span-2 card bg-base-100 border-success">
                <div class="card-body">
                    <div class="grid grid-cols-1 gap-4 mb-4 sm:grid-cols-2">
                        <label class="form-control">
                            <div class="label"><span class="font-medium label-text">Titre Date</span></div>
                            <input
                                type="date"
                                bind:value={client.titre}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered"
                            />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="font-medium label-text">Last Rabies Date</span></div>
                            <input
                                type="date"
                                bind:value={client.last_rabies_date}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered"
                            />
                        </label>
                        <label class="form-control">
                            <div class="label">
                                <span class="font-medium label-text text-success">Rabies Validity</span>
                            </div>
                            <input
                                type="date"
                                bind:value={client.rabies_validity}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered border-success/30"
                            />
                        </label>
                        <label class="form-control">
                            <div class="label">
                                <span class="font-medium label-text">Other Vaccines Completed</span>
                            </div>
                            <input
                                type="text"
                                bind:value={client.other_vaccines_completed}
                                class="input input-bordered"
                                placeholder="e.g., DHPPI, Lepto"
                            />
                        </label>
                    </div>

                    <div class="text-xs uppercase opacity-50 divider">Checklist</div>

                    <div class="grid grid-cols-2 gap-y-4 sm:grid-cols-3">
                        {#each [{ id: "documentation_status", label: "Docs Status" }, { id: "rabies_vaccination_valid", label: "Rabies Valid" }, { id: "health_certificate_issues", label: "Health Certificate" }, { id: "export_permit_approved", label: "Export Permit" }, { id: "import_permit_approved", label: "Import Permit" }, { id: "airline_approval_received", label: "Airline Approval" }, { id: "customs_clearance_done", label: "Customs Clear" }] as item}
                            <label class="flex items-center gap-3 cursor-pointer group">
                                <input
                                    type="checkbox"
                                    bind:checked={client[item.id]}
                                    class="checkbox checkbox-success checkbox-sm"
                                />
                                <span class="transition-colors label-text group-hover:text-success">{item.label}</span>
                            </label>
                        {/each}
                    </div>
                </div>
            </div>
        </section>

        <hr class="opacity-10" />

        <section class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div class="md:col-span-1">
                <h2 class="text-xl font-semibold">Travel Logistics</h2>
                <p class="text-sm opacity-60">Origin, destination, and flight times.</p>
            </div>
            <div class="shadow-xl md:col-span-2 card bg-base-100">
                <div class="grid grid-cols-1 gap-4 card-body sm:grid-cols-2">
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Origin Country</span></div>
                        <input type="text" bind:value={client.origin_country} class="input input-bordered" />
                    </label>
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Destination Country</span></div>
                        <input type="text" bind:value={client.destination_country} class="input input-bordered" />
                    </label>
                    <label class="col-span-1 form-control sm:col-span-2">
                        <div class="label">
                            <span class="font-medium label-text text-warning">Departure Date & Time</span>
                        </div>
                        <input
                            type="datetime-local"
                            bind:value={client.departure}
                            onclick={(e) => e.currentTarget.showPicker()}
                            class="w-full input input-bordered border-warning/40"
                        />
                    </label>
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Airline</span></div>
                        <input type="text" bind:value={client.airline} class="input input-bordered" />
                    </label>
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Flight Number</span></div>
                        <input type="text" bind:value={client.flight_number} class="input input-bordered" />
                    </label>
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Travel Type</span></div>
                        <input
                            type="text"
                            bind:value={client.type_of_travel}
                            placeholder="Cargo / Accompanied"
                            class="input input-bordered"
                        />
                    </label>
                    <div class="grid grid-cols-2 col-span-1 gap-2 sm:col-span-2">
                        <label class="form-control">
                            <div class="label"><span class="font-medium label-text">ETD</span></div>
                            <input
                                type="time"
                                bind:value={client.etd}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered"
                            />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="font-medium label-text">ETA</span></div>
                            <input
                                type="time"
                                bind:value={client.eta}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered"
                            />
                        </label>
                    </div>
                </div>
            </div>
        </section>

        <hr class="opacity-10" />

        <section class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div class="md:col-span-1">
                <h2 class="text-xl font-semibold">Financials & Remarks</h2>
                <p class="text-sm opacity-60">Invoicing details and custom notes.</p>
                <button type="button" class="mt-4 btn btn-info btn-outline btn-sm" onclick={runCalculations}>
                    Auto Calculate
                </button>
            </div>
            <div class="border-t-4 shadow-xl md:col-span-2 card bg-base-100 border-info">
                <div class="card-body">
                    <div class="grid grid-cols-2 gap-4 mb-6 lg:grid-cols-3">
                        {#each [{ id: "import_fee", label: "Import Fee" }, { id: "export_fee", label: "Export Fee" }, { id: "after_hours_charges", label: "After Hours" }, { id: "forwarder_charges", label: "Forwarder" }, { id: "airline_charges", label: "Air Charges" }, { id: "crate_cost", label: "Crate Cost" }, { id: "quoted_amount", label: "Quoted" }, { id: "total_cost", label: "Total Cost" }, { id: "profit", label: "Profit" }, { id: "advanced_received", label: "Advanced" }, { id: "balance_pending", label: "Balance" }] as money}
                            <label class="form-control">
                                <div class="label">
                                    <span class="text-xs uppercase label-text opacity-70">{money.label}</span>
                                </div>
                                <div class="relative">
                                    <span class="absolute -translate-y-1/2 opacity-50 left-3 top-1/2">$</span>
                                    <input
                                        type="number"
                                        step="0.01"
                                        bind:value={client[money.id]}
                                        class="w-full input input-bordered pl-7"
                                        class:bg-cyan-900={money.id.startsWith("total") ||
                                            money.id.startsWith("balance") ||
                                            money.id.startsWith("profit")}
                                    />
                                </div>
                            </label>
                        {/each}
                        <label class="form-control">
                            <div class="label">
                                <span class="font-bold label-text text-info">Payment Status</span>
                            </div>
                            <select
                                bind:value={client.payment_status}
                                class="font-bold select select-bordered border-info/40"
                            >
                                <option value="pending">Pending</option>
                                <option value="paid">Paid</option>
                            </select>
                        </label>
                    </div>
                    <label class="form-control">
                        <div class="label"><span class="font-medium label-text">Remarks / Private Notes</span></div>
                        <textarea
                            bind:value={client.remarks}
                            rows="3"
                            class="textarea textarea-bordered focus:textarea-info"
                            placeholder="Any additional travel details..."
                        ></textarea>
                    </label>
                </div>
            </div>
        </section>

        <div class="flex justify-end gap-4 mt-12">
            <button type="submit" class="px-12 text-lg shadow-lg btn btn-primary">Create Client</button>
        </div>
    </form>
</div>

<style>
    input[type="date"]::-webkit-calendar-picker-indicator,
    input[type="datetime-local"]::-webkit-calendar-picker-indicator,
    input[type="time"]::-webkit-calendar-picker-indicator {
        background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="%233b82f6" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>');
        background-position: center;
        background-size: contain;
        width: 1.2rem;
        height: 1.2rem;
        cursor: pointer;
    }

    .container {
        animation: slideUp 0.5s ease-out;
    }

    @keyframes slideUp {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
</style>
