<script lang="ts">
    import { TableHandler } from "@vincjo/datatables";

    const { data } = $props();
    const clients = data.data;

    const table = new TableHandler([], { rowsPerPage: 20 });
</script>

<a href="/add" class="m-10 p-4 bg-blue-400 rounded-xl w-full">Add client</a>
<br />
<br />
<br />

{#if !clients}
    <h1>No clients</h1>
{/if}
