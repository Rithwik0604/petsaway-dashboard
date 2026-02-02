<script lang="ts">
    import { Check, X } from "lucide-svelte";

    let { checked, onchange } = $props();
</script>

<button
    type="button"
    onclick={() => onchange(!checked)}
    class="flex items-center justify-center w-8 h-8 transition-colors border rounded-md"
    class:bg-green-500={checked}
    class:bg-red-500={!checked}
    class:border-green-600={checked}
    class:border-red-600={!checked}
>
    {#if checked}
        <Check size={18} color="white" strokeWidth={3} />
    {:else}
        <X size={18} color="white" strokeWidth={3} />
    {/if}
</button>

<style>
    button {
        cursor: pointer;
    }
    button:active {
        transform: scale(0.9);
    }
</style>
