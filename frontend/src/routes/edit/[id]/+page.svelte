<script lang="ts">
    import { goto } from "$app/navigation";
    import { apiFetch } from "$lib/api";
    import { page } from "$app/stores";
    import { formatDateForInput, formatTimeForInput, formatDateTimeForInput } from "$lib/format";

    let { data } = $props();

    function getInitialClient() {
        const d = data?.data;
        if (!d) return {};

        return {
            client_name: d.client_name?.String ?? "",
            client_phone: d.client_phone?.String ?? "",
            import_export: d.import_export?.String ?? "import",
            import_fee: d.import_fee?.Float64 ?? 0,
            export_fee: d.export_fee?.Float64 ?? 0,
            after_hours_charges: d.after_hours_charges?.Float64 ?? 0,
            pet_name: d.pet_name?.String ?? "",
            species: d.species?.String ?? "",
            gender: d.gender?.String ?? "male",
            breed: d.breed?.String ?? "",
            date_of_birth: formatDateForInput(d.date_of_birth?.Time),
            microchip_number: d.microchip_number?.String ?? "",
            titre: formatDateForInput(d.titre?.Time),
            last_rabies_date: formatDateForInput(d.last_rabies_date?.Time),
            rabies_validity: formatDateForInput(d.rabies_validity?.Time),
            documentation_status: d.documentation_status?.Int64 === 1,
            rabies_vaccination_valid: d.rabies_vaccination_valid?.Int64 === 1,
            other_vaccines_completed: d.other_vaccines_completed?.String ?? "",
            health_certificate_issues: d.health_certificate_issues?.Int64 === 1,
            export_permit_approved: d.export_permit_approved?.Int64 === 1,
            import_permit_approved: d.import_permit_approved?.Int64 === 1,
            airline_approval_received: d.airline_approval_received?.Int64 === 1,
            customs_clearance_done: d.customs_clearance_done?.Int64 === 1,
            origin_country: d.origin_country?.String ?? "",
            destination_country: d.destination_country?.String ?? "",
            forwarder_charges: d.forwarder_charges?.Float64 ?? 0,
            departure: formatDateTimeForInput(d.departure?.Time),
            airline: d.airline?.String ?? "",
            airline_charges: d.airline_charges?.Float64 ?? 0,
            crate_cost: d.crate_cost?.Float64 ?? 0,
            flight_number: d.flight_number?.String ?? "",
            type_of_travel: d.type_of_travel?.String ?? "",
            etd: formatTimeForInput(d.etd),
            eta: formatTimeForInput(d.eta),
            quoted_amount: d.quoted_amount?.Float64 ?? 0,
            total_cost: d.total_cost?.Float64 ?? 0,
            profit: d.profit?.Float64 ?? 0,
            advanced_received: d.advanced_received?.Float64 ?? 0,
            balance_pending: d.balance_pending?.Float64 ?? 0,
            payment_status: d.payment_status?.String ?? "pending",
            remarks: d.remarks?.String ?? "",
        };
    }

    let client = $state(getInitialClient());

    $effect(() => {
        if (data?.data) {
            Object.assign(client, getInitialClient());
        }
    });

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        try {
            const payload = Object.fromEntries(
                Object.entries(client).map(([key, value]) => [key, value === "" ? null : value]),
            );

            const res = await apiFetch(`/clients/${$page.params.id}`, {
                method: "PATCH",
                body: JSON.stringify(payload),
            });
            if (res.ok) goto("/");
        } catch (err) {
            console.error(err);
        }
    }
</script>

<div class="container p-8 mx-auto max-w-6xl">
    <div class="flex justify-between items-center mb-8">
        <h1 class="text-3xl font-bold">Edit Client</h1>
        <button type="button" class="btn btn-ghost" onclick={() => goto("/")}>Cancel</button>
    </div>

    {#if client.client_name !== undefined}
        <form onsubmit={handleSubmit} class="space-y-8 pb-20">
            <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="md:col-span-1">
                    <h2 class="text-xl font-semibold">Client Identity</h2>
                    <p class="text-sm opacity-60">Primary contact and service type.</p>
                </div>
                <div class="md:col-span-2 shadow-xl card bg-base-100">
                    <div class="card-body grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Full Name</span></div>
                            <input
                                type="text"
                                required
                                bind:value={client.client_name}
                                class="input input-bordered focus:input-primary"
                            />
                        </label>
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Phone Number</span></div>
                            <input type="text" bind:value={client.client_phone} class="input input-bordered" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text font-medium text-secondary">Service Type</span>
                            </div>
                            <select
                                bind:value={client.import_export}
                                class="select select-bordered border-secondary/40"
                            >
                                <option value="import">Import</option>
                                <option value="export">Export</option>
                            </select>
                        </label>
                    </div>
                </div>
            </section>

            <hr class="opacity-10" />

            <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="md:col-span-1">
                    <h2 class="text-xl font-semibold">Pet Information</h2>
                    <p class="text-sm opacity-60">Specify pet details and age.</p>
                </div>
                <div class="md:col-span-2 shadow-xl card bg-base-100">
                    <div class="card-body grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Pet Name</span></div>
                            <input type="text" bind:value={client.pet_name} class="input input-bordered" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Species</span></div>
                            <input type="text" bind:value={client.species} class="input input-bordered" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Gender</span></div>
                            <select bind:value={client.gender} class="select select-bordered">
                                <option value="male">Male</option>
                                <option value="female">Female</option>
                            </select>
                        </label>
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Breed</span></div>
                            <input type="text" bind:value={client.breed} class="input input-bordered" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text font-medium text-primary">Date of Birth</span>
                            </div>
                            <input
                                type="date"
                                bind:value={client.date_of_birth}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered border-primary/30"
                            />
                        </label>
                        <label class="form-control w-full">
                            <div class="label"><span class="label-text font-medium">Microchip #</span></div>
                            <input type="text" bind:value={client.microchip_number} class="input input-bordered" />
                        </label>
                    </div>
                </div>
            </section>

            <hr class="opacity-10" />

            <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="md:col-span-1">
                    <h2 class="text-xl font-semibold">Medical & Documentation</h2>
                    <p class="text-sm opacity-60">Vaccination history and check-list.</p>
                </div>
                <div class="md:col-span-2 shadow-xl card bg-base-100 border-l-4 border-success">
                    <div class="card-body">
                        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
                            <label class="form-control">
                                <div class="label"><span class="label-text font-medium">Titre Date</span></div>
                                <input
                                    type="date"
                                    bind:value={client.titre}
                                    onclick={(e) => e.currentTarget.showPicker()}
                                    class="input input-bordered"
                                />
                            </label>
                            <label class="form-control">
                                <div class="label"><span class="label-text font-medium">Last Rabies Date</span></div>
                                <input
                                    type="date"
                                    bind:value={client.last_rabies_date}
                                    onclick={(e) => e.currentTarget.showPicker()}
                                    class="input input-bordered"
                                />
                            </label>
                            <label class="form-control">
                                <div class="label">
                                    <span class="label-text font-medium text-success">Rabies Validity</span>
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
                                    <span class="label-text font-medium">Other Vaccines Completed</span>
                                </div>
                                <input
                                    type="text"
                                    bind:value={client.other_vaccines_completed}
                                    class="input input-bordered"
                                    placeholder="e.g., DHPPI, Lepto"
                                />
                            </label>
                        </div>

                        <div class="divider text-xs opacity-50 uppercase">Checklist</div>

                        <div class="grid grid-cols-2 sm:grid-cols-3 gap-y-4">
                            {#each [{ id: "documentation_status", label: "Docs Status" }, { id: "rabies_vaccination_valid", label: "Rabies Valid" }, { id: "health_certificate_issues", label: "Health Certificate" }, { id: "export_permit_approved", label: "Export Permit" }, { id: "import_permit_approved", label: "Import Permit" }, { id: "airline_approval_received", label: "Airline Approval" }, { id: "customs_clearance_done", label: "Customs Clear" }] as item}
                                <label class="flex items-center gap-3 cursor-pointer group">
                                    <input
                                        type="checkbox"
                                        bind:checked={client[item.id]}
                                        class="checkbox checkbox-success checkbox-sm"
                                    />
                                    <span class="label-text group-hover:text-success transition-colors"
                                        >{item.label}</span
                                    >
                                </label>
                            {/each}
                        </div>
                    </div>
                </div>
            </section>

            <hr class="opacity-10" />

            <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="md:col-span-1">
                    <h2 class="text-xl font-semibold">Travel Logistics</h2>
                    <p class="text-sm opacity-60">Origin, destination, and flight times.</p>
                </div>
                <div class="md:col-span-2 shadow-xl card bg-base-100">
                    <div class="card-body grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Origin Country</span></div>
                            <input type="text" bind:value={client.origin_country} class="input input-bordered" />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Destination Country</span></div>
                            <input type="text" bind:value={client.destination_country} class="input input-bordered" />
                        </label>
                        <label class="form-control col-span-1 sm:col-span-2">
                            <div class="label">
                                <span class="label-text font-medium text-warning">Departure Date & Time</span>
                            </div>
                            <input
                                type="datetime-local"
                                bind:value={client.departure}
                                onclick={(e) => e.currentTarget.showPicker()}
                                class="input input-bordered border-warning/40 w-full"
                            />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Airline</span></div>
                            <input type="text" bind:value={client.airline} class="input input-bordered" />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Flight Number</span></div>
                            <input type="text" bind:value={client.flight_number} class="input input-bordered" />
                        </label>
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Travel Type</span></div>
                            <input
                                type="text"
                                bind:value={client.type_of_travel}
                                placeholder="Cargo / Accompanied"
                                class="input input-bordered"
                            />
                        </label>
                        <div class="grid grid-cols-2 gap-2 col-span-1 sm:col-span-2">
                            <label class="form-control">
                                <div class="label"><span class="label-text font-medium">ETD</span></div>
                                <input
                                    type="time"
                                    bind:value={client.etd}
                                    onclick={(e) => e.currentTarget.showPicker()}
                                    class="input input-bordered"
                                />
                            </label>
                            <label class="form-control">
                                <div class="label"><span class="label-text font-medium">ETA</span></div>
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

            <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="md:col-span-1">
                    <h2 class="text-xl font-semibold">Financials & Remarks</h2>
                    <p class="text-sm opacity-60">Invoicing details and custom notes.</p>
                </div>
                <div class="md:col-span-2 shadow-xl card bg-base-100 border-t-4 border-info">
                    <div class="card-body">
                        <div class="grid grid-cols-2 lg:grid-cols-3 gap-4 mb-6">
                            {#each [{ id: "import_fee", label: "Import Fee" }, { id: "export_fee", label: "Export Fee" }, { id: "after_hours_charges", label: "After Hours" }, { id: "forwarder_charges", label: "Forwarder" }, { id: "airline_charges", label: "Air Charges" }, { id: "crate_cost", label: "Crate Cost" }, { id: "quoted_amount", label: "Quoted" }, { id: "total_cost", label: "Total Cost" }, { id: "profit", label: "Profit" }, { id: "advanced_received", label: "Advanced" }, { id: "balance_pending", label: "Balance" }] as money}
                                <label class="form-control">
                                    <div class="label">
                                        <span class="label-text text-xs opacity-70 uppercase">{money.label}</span>
                                    </div>
                                    <div class="relative">
                                        <span class="absolute left-3 top-1/2 -translate-y-1/2 opacity-50">$</span>
                                        <input
                                            type="number"
                                            step="0.01"
                                            bind:value={client[money.id]}
                                            class="input input-bordered w-full pl-7"
                                        />
                                    </div>
                                </label>
                            {/each}
                            <label class="form-control">
                                <div class="label">
                                    <span class="label-text font-bold text-info">Payment Status</span>
                                </div>
                                <select
                                    bind:value={client.payment_status}
                                    class="select select-bordered font-bold border-info/40"
                                >
                                    <option value="pending">Pending</option>
                                    <option value="paid">Paid</option>
                                </select>
                            </label>
                        </div>
                        <label class="form-control">
                            <div class="label"><span class="label-text font-medium">Remarks / Private Notes</span></div>
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
                <button type="button" class="btn btn-ghost px-8" onclick={() => goto("/")}>Discard</button>
                <button type="submit" class="btn btn-primary px-12 text-lg shadow-lg">Save Changes</button>
            </div>
        </form>
    {:else}
        <div class="flex flex-col justify-center items-center h-96 gap-4">
            <span class="loading loading-spinner loading-lg text-primary"></span>
            <p class="animate-pulse opacity-50">Loading client data...</p>
        </div>
    {/if}
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
