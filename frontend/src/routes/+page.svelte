<script lang="ts">
    import { formatDate, formatTo12Hour } from "$lib/format";
    import StatusToggle from "$lib/StatusToggle.svelte";
    import { check, Datatable, TableHandler, ThFilter, ThSort, type AdvancedFilterBuilder } from "@vincjo/datatables";
    import { ArrowUp, Check, EditIcon, X } from "lucide-svelte";
    import { _toggleCheck } from "./+page.js";

    let filtersVisibility = $state(true);

    let viewFilterExpanded = $state(true);
    let statusFilterExpanded = $state(true);

    const { data } = $props();
    let clients = $derived(data.data as App.Client[]);
    // svelte-ignore non_reactive_update
    let view: any,
        paymentFilter: AdvancedFilterBuilder<App.Client>,
        typeFilter: AdvancedFilterBuilder<App.Client>,
        petFilter: AdvancedFilterBuilder<App.Client>;
    // svelte-ignore state_referenced_locally
    const table = new TableHandler(clients, { rowsPerPage: 100, selectBy: "id" });

    // svelte-ignore state_referenced_locally
    if (clients) {
        view = table.createView([
            { index: 0, name: "Edit" },
            { index: 1, name: "Client Name" },
            { index: 2, name: "Client Phone" },
            { index: 3, name: "Import/Export" },
            { index: 4, name: "Import Fee" },
            { index: 5, name: "Export Fee" },
            { index: 6, name: "After Hours Charges" },
            { index: 7, name: "Pet Name" },
            { index: 8, name: "Species" },
            { index: 9, name: "Gender" },
            { index: 10, name: "Breed" },
            { index: 11, name: "Date of Birth" },
            { index: 12, name: "Microchip Number" },
            { index: 13, name: "Titre" },
            { index: 14, name: "Last Rabies Date" },
            { index: 15, name: "Rabies Validity" },
            { index: 16, name: "Documentation Status" },
            { index: 17, name: "Rabies Vaccination Valid" },
            { index: 18, name: "Other Vaccines Completed" },
            { index: 19, name: "Health Certificate Issues" },
            { index: 20, name: "Export Permit Approved" },
            { index: 21, name: "Import Permit Approved" },
            { index: 22, name: "Airline Approval Received" },
            { index: 23, name: "Customs Clearance Done" },
            { index: 24, name: "Origin Country" },
            { index: 25, name: "Destination Country" },
            { index: 26, name: "Forwarder Charges" },
            { index: 27, name: "Departure" },
            { index: 28, name: "Airline" },
            { index: 29, name: "Airline Charges" },
            { index: 30, name: "Crate Cost" },
            { index: 31, name: "Flight Number" },
            { index: 32, name: "Type of Travel" },
            { index: 33, name: "ETD" },
            { index: 34, name: "ETA" },
            { index: 35, name: "Quoted Amount" },
            { index: 36, name: "Total Cost" },
            { index: 37, name: "Profit" },
            { index: 38, name: "Advanced Received" },
            { index: 39, name: "Balance Pending" },
            { index: 40, name: "Payment Status" },
            { index: 41, name: "Remarks" },
            { index: 42, name: "Created At" },
            { index: 43, name: "Updated At" },
        ]);

        paymentFilter = table.createAdvancedFilter("payment_status", check.isEqualTo);

        typeFilter = table.createAdvancedFilter("import_export", check.isEqualTo);

        petFilter = table.createAdvancedFilter("species", check.isEqualTo);
        clients.forEach((row) => {
            row.date_of_birth.Time = formatDate(row.date_of_birth.Time);
            row.titre.Time = formatDate(row.titre.Time);
            row.last_rabies_date.Time = formatDate(row.last_rabies_date.Time);
            row.rabies_validity.Time = formatDate(row.rabies_validity.Time);
            row.departure.Time = formatDate(row.departure.Time);
            row.created_at.Time = formatDate(row.created_at.Time);
            row.updated_at.Time = formatDate(row.updated_at.Time);
            row.etd = formatTo12Hour(row.etd);
            row.eta = formatTo12Hour(row.eta);
        });
    }

    async function handleToggle(id: string, field: string, isChecked: boolean) {
        const body = { id, [field]: isChecked };
        console.log(body);
        try {
            const res = await _toggleCheck(id, body);
            if (!res!.ok) {
                return;
            }

            // 1. Find the specific client in the source array
            // We use data.data because 'clients' is derived from it.
            const index = data.data.findIndex((c: App.Client) => c.id === id);

            if (index !== -1) {
                // 2. Update the nested value
                data.data[index][field].Int64 = isChecked ? 1 : 0;

                // 3. TRIGGER REACTIVITY: Re-assign the array to itself
                data.data = [...data.data];

                // // 4. SYNC THE TABLE: Most datatable handlers need to be told the source changed
                table.setRows(data.data);
            }
        } catch (err) {
            // Error handling already in _toggleCheck
        }
    }
</script>

<div class="container p-0 mx-auto mt-5">
    <a href="/add" class=" btn btn-info">Add client</a>
    <br />
    <br />
    <br />

    {#if !clients}
        <h1>No clients</h1>
    {:else}
        <div class="flex flex-row items-start gap-4">
            <div>
                <button
                    class="btn btn-primary min-w-max"
                    onclick={() => (filtersVisibility = !filtersVisibility.valueOf())}>Show Filters</button
                >
                {#if filtersVisibility}
                    <!-- svelte-ignore a11y_click_events_have_key_events -->
                    <!-- svelte-ignore a11y_no_static_element_interactions -->
                    <div
                        class="p-4 mt-4 overflow-y-auto rounded-lg shadow-lg bg-base-200 h-96 w-max"
                        class:h-96={viewFilterExpanded}
                        class:h-10={!viewFilterExpanded}
                    >
                        <div
                            class="flex items-center justify-between mb-4 cursor-pointer"
                            onclick={() => (viewFilterExpanded = !viewFilterExpanded)}
                        >
                            <span class="font-bold">View Filters</span>
                            <div class="transition-transform duration-300" class:rotate-180={viewFilterExpanded}>
                                <ArrowUp size={18} />
                            </div>
                        </div>
                        <table class="table table-compact">
                            <tbody>
                                {#each view.columns as column}
                                    <tr>
                                        <td>
                                            <label class="flex items-center gap-2 cursor-pointer">
                                                <input
                                                    type="checkbox"
                                                    checked={column.isVisible}
                                                    onchange={() => column.toggle()}
                                                    class="checkbox checkbox-primary"
                                                />
                                                <span class="text-left label-text">{column.name}</span>
                                            </label>
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </div>

                    <div
                        class="p-4 mt-4 overflow-y-auto rounded-lg shadow-lg bg-base-200 h-96 w-max"
                        class:h-96={statusFilterExpanded}
                        class:h-12={!statusFilterExpanded}
                    >
                        <!-- svelte-ignore a11y_click_events_have_key_events -->
                        <!-- svelte-ignore a11y_no_static_element_interactions -->
                        <div
                            class="flex items-center justify-between mb-4 cursor-pointer"
                            onclick={() => (statusFilterExpanded = !statusFilterExpanded)}
                        >
                            <span class="font-bold">Status Filters</span>
                            <div class="transition-transform duration-300" class:rotate-180={statusFilterExpanded}>
                                <ArrowUp size={18} />
                            </div>
                        </div>
                        <table class="table table-compact">
                            <tbody>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    paymentFilter.set("pending");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("pending")}
                                            />
                                            <span class="text-left label-text">Payment: Pending</span>
                                        </label>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    paymentFilter.set("paid");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("paid")}
                                            />
                                            <span class="text-left label-text">Payment: Paid</span>
                                        </label>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    typeFilter.set("import");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("import")}
                                            />
                                            <span class="text-left label-text">Type: Import</span>
                                        </label>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    typeFilter.set("export");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("export")}
                                            />
                                            <span class="text-left label-text">Type: Export</span>
                                        </label>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    petFilter.set("cat");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("cat")}
                                            />
                                            <span class="text-left label-text">Species: Cat</span>
                                        </label>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <label class="flex items-center gap-2 cursor-pointer">
                                            <input
                                                type="checkbox"
                                                checked={false}
                                                onchange={() => {
                                                    petFilter.set("dog");
                                                }}
                                                class="checkbox checkbox-primary"
                                                class:active={paymentFilter.criteria.includes("dog")}
                                            />
                                            <span class="text-left label-text">Species: Dog</span>
                                        </label>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                {/if}
            </div>

            <div class="overflow-x-auto grow">
                <Datatable {table}>
                    <table class="table w-full">
                        <thead>
                            <tr>
                                <ThSort {table} field="id">Edit</ThSort>
                                <ThSort {table} field="client_name">Client Name</ThSort>
                                <ThSort {table} field="client_phone">Client Phone</ThSort>
                                <ThSort {table} field="import_export">Import/Export</ThSort>
                                <ThSort {table} field="import_fee">Import Fee</ThSort>
                                <ThSort {table} field="export_fee">Export Fee</ThSort>
                                <ThSort {table} field="after_hours_charges">After Hours Charges</ThSort>
                                <ThSort {table} field="pet_name">Pet Name</ThSort>
                                <ThSort {table} field="species">Species</ThSort>
                                <ThSort {table} field="gender">Gender</ThSort>
                                <ThSort {table} field="breed">Breed</ThSort>
                                <ThSort {table} field="date_of_birth">Date of Birth</ThSort>
                                <ThSort {table} field="microchip_number">Microchip Number</ThSort>
                                <ThSort {table} field="titre">Titre</ThSort>
                                <ThSort {table} field="last_rabies_date">Last Rabies Date</ThSort>
                                <ThSort {table} field="rabies_validity">Rabies Validity</ThSort>
                                <ThSort {table} field="documentation_status">Documentation Status</ThSort>
                                <ThSort {table} field="rabies_vaccination_valid">Rabies Vaccination Valid</ThSort>
                                <ThSort {table} field="other_vaccines_completed">Other Vaccines Completed</ThSort>
                                <ThSort {table} field="health_certificate_issues">Health Certificate Issues</ThSort>
                                <ThSort {table} field="export_permit_approved">Export Permit Approved</ThSort>
                                <ThSort {table} field="import_permit_approved">Import Permit Approved</ThSort>
                                <ThSort {table} field="airline_approval_received">Airline Approval Received</ThSort>
                                <ThSort {table} field="customs_clearance_done">Customs Clearance Done</ThSort>
                                <ThSort {table} field="origin_country">Origin Country</ThSort>
                                <ThSort {table} field="destination_country">Destination Country</ThSort>
                                <ThSort {table} field="forwarder_charges">Forwarder Charges</ThSort>
                                <ThSort {table} field="departure">Departure</ThSort>
                                <ThSort {table} field="airline">Airline</ThSort>
                                <ThSort {table} field="airline_charges">Airline Charges</ThSort>
                                <ThSort {table} field="crate_cost">Crate Cost</ThSort>
                                <ThSort {table} field="flight_number">Flight Number</ThSort>
                                <ThSort {table} field="type_of_travel">Type of Travel</ThSort>
                                <ThSort {table} field="etd">ETD</ThSort>
                                <ThSort {table} field="eta">ETA</ThSort>
                                <ThSort {table} field="quoted_amount">Quoted Amount</ThSort>
                                <ThSort {table} field="total_cost">Total Cost</ThSort>
                                <ThSort {table} field="profit">Profit</ThSort>
                                <ThSort {table} field="advanced_received">Advanced Received</ThSort>
                                <ThSort {table} field="balance_pending">Balance Pending</ThSort>
                                <ThSort {table} field="payment_status">Payment Status</ThSort>
                                <ThSort {table} field="remarks">Remarks</ThSort>
                                <ThSort {table} field="created_at">Created At</ThSort>
                                <ThSort {table} field="updated_at">Updated At</ThSort>
                            </tr>
                            <tr>
                                <ThFilter {table} field="id" />
                                <ThFilter {table} field="client_name" />
                                <ThFilter {table} field="client_phone" />
                                <ThFilter {table} field="import_export" />
                                <ThFilter {table} field="import_fee" />
                                <ThFilter {table} field="export_fee" />
                                <ThFilter {table} field="after_hours_charges" />
                                <ThFilter {table} field="pet_name" />
                                <ThFilter {table} field="species" />
                                <ThFilter {table} field="gender" />
                                <ThFilter {table} field="breed" />
                                <ThFilter {table} field="date_of_birth" />
                                <ThFilter {table} field="microchip_number" />
                                <ThFilter {table} field="titre" />
                                <ThFilter {table} field="last_rabies_date" />
                                <ThFilter {table} field="rabies_validity" />
                                <ThFilter {table} field="documentation_status" />
                                <ThFilter {table} field="rabies_vaccination_valid" />
                                <ThFilter {table} field="other_vaccines_completed" />
                                <ThFilter {table} field="health_certificate_issues" />
                                <ThFilter {table} field="export_permit_approved" />
                                <ThFilter {table} field="import_permit_approved" />
                                <ThFilter {table} field="airline_approval_received" />
                                <ThFilter {table} field="customs_clearance_done" />
                                <ThFilter {table} field="origin_country" />
                                <ThFilter {table} field="destination_country" />
                                <ThFilter {table} field="forwarder_charges" />
                                <ThFilter {table} field="departure" />
                                <ThFilter {table} field="airline" />
                                <ThFilter {table} field="airline_charges" />
                                <ThFilter {table} field="crate_cost" />
                                <ThFilter {table} field="flight_number" />
                                <ThFilter {table} field="type_of_travel" />
                                <ThFilter {table} field="etd" />
                                <ThFilter {table} field="eta" />
                                <ThFilter {table} field="quoted_amount" />
                                <ThFilter {table} field="total_cost" />
                                <ThFilter {table} field="profit" />
                                <ThFilter {table} field="advanced_received" />
                                <ThFilter {table} field="balance_pending" />
                                <ThFilter {table} field="payment_status" />
                                <ThFilter {table} field="remarks" />
                                <ThFilter {table} field="created_at" />
                                <ThFilter {table} field="updated_at" />
                            </tr>
                        </thead>
                        <tbody>
                            {#each table.rows as row}
                                <tr class="hover:bg-transparent">
                                    <!-- <td>{row.id}</td> -->
                                    <td class="whitespace-nowrap"
                                        ><div>
                                            <a href={`/edit/${row.id}`}><EditIcon /></a>
                                        </div></td
                                    >
                                    <td class="whitespace-nowrap">{row.client_name.String}</td>
                                    <td class="whitespace-nowrap">{row.client_phone.String}</td>
                                    <td class="whitespace-nowrap">{row.import_export.String}</td>
                                    <td class="whitespace-nowrap">{row.import_fee.Float64}</td>
                                    <td class="whitespace-nowrap">{row.export_fee.Float64}</td>
                                    <td class="whitespace-nowrap">{row.after_hours_charges.Float64}</td>
                                    <td class="whitespace-nowrap">{row.pet_name.String}</td>
                                    <td class="whitespace-nowrap">{row.species.String}</td>
                                    <td class="whitespace-nowrap">{row.gender.String}</td>
                                    <td class="whitespace-nowrap">{row.breed.String}</td>
                                    <td class="whitespace-nowrap">{row.date_of_birth.Time}</td>
                                    <td class="whitespace-nowrap">{row.microchip_number.String}</td>
                                    <td class="whitespace-nowrap">{row.titre.Time}</td>
                                    <td class="whitespace-nowrap">{row.last_rabies_date.Time}</td>
                                    <td class="whitespace-nowrap">{row.rabies_validity.Time}</td>
                                    <td class="text-center whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.documentation_status.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "documentation_status", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.rabies_vaccination_valid.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "rabies_vaccination_valid", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.other_vaccines_completed.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "other_vaccines_completed", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.health_certificate_issues.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "health_certificate_issues", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.export_permit_approved.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "export_permit_approved", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.import_permit_approved.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "import_permit_approved", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.airline_approval_received.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "airline_approval_received", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">
                                        <StatusToggle
                                            checked={row.customs_clearance_done.Int64 === 1}
                                            onchange={(val: boolean) =>
                                                handleToggle(row.id, "customs_clearance_done", val)}
                                        />
                                    </td>
                                    <td class="whitespace-nowrap">{row.origin_country.String}</td>
                                    <td class="whitespace-nowrap">{row.destination_country.String}</td>
                                    <td class="whitespace-nowrap">{row.forwarder_charges.Float64}</td>
                                    <td class="whitespace-nowrap">{row.departure.Time}</td>
                                    <td class="whitespace-nowrap">{row.airline.String}</td>
                                    <td class="whitespace-nowrap">{row.airline_charges.Float64}</td>
                                    <td class="whitespace-nowrap">{row.crate_cost.Float64}</td>
                                    <td class="whitespace-nowrap">{row.flight_number.String}</td>
                                    <td class="whitespace-nowrap">{row.type_of_travel.String}</td>
                                    <td class="whitespace-nowrap">{row.etd}</td>
                                    <td class="whitespace-nowrap">{row.eta}</td>
                                    <td class="whitespace-nowrap">{row.quoted_amount.Float64}</td>
                                    <td class="whitespace-nowrap">{row.total_cost.Float64}</td>
                                    <td class="whitespace-nowrap">{row.profit.Float64}</td>
                                    <td class="whitespace-nowrap">{row.advanced_received.Float64}</td>
                                    <td class="whitespace-nowrap">{row.balance_pending.Float64}</td>
                                    <td
                                        class="whitespace-nowrap"
                                        class:text-emerald-500={row.payment_status.String === "paid"}
                                        class:text-rose-500={row.payment_status.String === "pending"}
                                    >
                                        {row.payment_status.String}
                                    </td>
                                    <td class="whitespace-nowrap">{row.remarks.String}</td>
                                    <td class="whitespace-nowrap">{row.created_at.Time}</td>
                                    <td class="whitespace-nowrap">{row.updated_at.Time}</td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </Datatable>
            </div>
        </div>
    {/if}
</div>

<style>
    /* Targeting the hover state specifically to prevent the background change */
    .table tr:hover {
        background-color: rgb(60, 60, 60) !important;
        /* If there is a "glow" caused by opacity or other effects: */
        --tw-bg-opacity: 0 !important;
    }
</style>
