<script lang="ts">
    import favicon from "$lib/assets/favicon.svg";
    import { fly } from "svelte/transition";
    import { notifications } from "../lib/notifications.svelte";
    import "./layout.css";
    import { apiFetch } from "$lib/api";
    import { goto } from "$app/navigation";

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

<nav class="bg-white shadow-md mb-4">
    <div class="container flex flex-row justify-between items-center mx-auto px-4">
        <div class="flex justify-between items-center py-4">
            <a href="/" class="text-2xl font-bold text-gray-800">PetsAway</a>
        </div>
        {#if tokenExists}
            <div class="flex justify-between w-min h-min items-center py-4 bg-red-500 p-2 rounded-xl">
                <button onclick={logout} class=" font-bold text-gray-800 cursor-pointer">Logout</button>
            </div>
        {/if}
    </div>
</nav>

{@render children()}

{#if notifications.message}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div transition:fly={{ y: 20, duration: 300 }} onclick={() => notifications.clear()} class="toast bg-red-400">
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
