import { PUBLIC_API } from "$env/static/public";

// We pass in the fetch instance as an argument
export async function apiFetch(
    endpoint: string,
    options: RequestInit = {},
    fetcher: typeof fetch = fetch,
): Promise<Response> {
    const res = await fetcher(`${PUBLIC_API}${endpoint}`, options);
    if (!res.ok) {
        throw new Error(`API Error: ${res.status}`);
    }
    return res;
}
