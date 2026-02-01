<script lang="ts">
    import { TableHandler } from "@vincjo/datatables";

    const { data } = $props();
    console.log("CLIENTS:", data);

    const table = new TableHandler([], { rowsPerPage: 20 });
</script>
