import { goto } from "$app/navigation";
import { PUBLIC_API } from "$env/static/public";
import { APIError } from "./errors";
import { notifications } from "./notifications.svelte";

export let AUTHENTICATED = false;

// We pass in the fetch instance as an argument
export async function apiFetch(
    endpoint: string,
    options: RequestInit = {},
    fetcher: typeof fetch = fetch,
): Promise<Response> {
    const res = await fetcher(`${PUBLIC_API}${endpoint}`, { ...options, credentials: "include" });
    if (!res.ok) {
        const message = (await res.text()) || `API Error`;
        if (res.status == 401) {
            AUTHENTICATED = false;
            goto("/login");
            return new Response();
        } else {
            AUTHENTICATED = true;
            const err = new APIError(message, res.status);
            notifications.show(`API Error (${err.status}): ${err.message}`);
            throw err;
        }
    }
    AUTHENTICATED = true;
    return res;
}
