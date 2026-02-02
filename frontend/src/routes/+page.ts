import { apiFetch } from "$lib/api";
import { APIError } from "$lib/errors";
import { notifications } from "$lib/notifications.svelte";

export const ssr = false;

export async function load({ fetch }) {
    try {
        const res = await apiFetch("/clients", {}, fetch);
        const data = await res.json();
        return { data };
    } catch (err) {
        if (err instanceof APIError) {
            notifications.show(`API Error (${err.status}): ${err.message}`);
        } else {
            notifications.show(`Failed to connect to the API. ${err}`);
        }
        return;
    }
}

export async function _toggleCheck(id: string, body: object) {
    try {
        return await apiFetch(`/clients/${id}`, { method: "PATCH", body: JSON.stringify(body) });
    } catch (err) {
        if (err instanceof APIError) {
            notifications.show(`API Error (${err.status}): ${err.message}`);
        } else {
            notifications.show(`Failed to connect to the API. ${err}`);
        }
        return;
    }
}
