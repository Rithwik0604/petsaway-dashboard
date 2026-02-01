import { apiFetch } from "$lib/api";
import { notifications } from "$lib/notifications.svelte";

export const ssr = false;

export async function load({ fetch }) {
    try {
        const res = await apiFetch("/api/clients", {}, fetch);
        const data = await res.json();
        return { data };
    } catch (err) {
        notifications.show(`Failed to connect to the API. ${err}`);
        return;
    }
}
