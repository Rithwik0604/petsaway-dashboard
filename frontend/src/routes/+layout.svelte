<script lang="ts">
    import { apiFetch } from "$lib/api";
    import favicon from "$lib/assets/favicon.svg";
    import { fly } from "svelte/transition";
    import { notifications } from "../lib/notifications.svelte";
    import "./layout.css";

    let { children } = $props();

    let tokenExists = true;

    function logout() {
        apiFetch("/auth/logout", { method: "POST" });
        location.reload();
    }
</script>

<svelte:head>
    <title>PetsAway Dashboard</title>
    <link rel="icon" href={favicon} /></svelte:head
>

<nav class="text-white navbar bg-blue-950">
    <div class="container flex flex-row items-center justify-between px-4 mx-auto">
        <div class="flex items-center justify-between py-4">
            <a href="/" class="text-2xl font-bold">PetsAway</a>
        </div>
        {#if tokenExists}
            <div class="flex items-center justify-between p-2 py-4 bg-red-500 w-min h-min rounded-xl">
                <button onclick={logout} class="cursor-pointer btn-error">Logout</button>
            </div>
        {/if}
    </div>
</nav>
<br /><br />

{@render children()}

{#if notifications.message}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div transition:fly={{ y: 20, duration: 300 }} onclick={() => notifications.clear()} class="bg-red-400 toast">
        <span>{notifications.message}</span>
    </div>
{/if}

<style>
    .toast {
        position: fixed;
        bottom: 20px;
        left: 50%;
        transform: translateX(-50%);
        background: #333;
        color: white;
        padding: 12px 24px;
        border-radius: 99px;
        cursor: pointer;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        z-index: 9999;
    }
</style>
