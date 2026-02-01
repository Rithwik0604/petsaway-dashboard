import { goto } from "$app/navigation";
import { PUBLIC_API } from "$env/static/public";
import { APIError } from "./errors";

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
            goto("/login");
        }
        throw new APIError(message, res.status);
    }
    return res;
}
