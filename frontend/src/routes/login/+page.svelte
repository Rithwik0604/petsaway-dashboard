<script lang="ts">
    import { goto } from "$app/navigation";
    import { apiFetch } from "$lib/api";

    let username = $state("");
    let password = $state("");
    let errorMessage: string | null = $state(null);

    async function handleSubmit(event: SubmitEvent) {
        // We prevent default to handle the API call manually
        event.preventDefault();

        try {
            await apiFetch("/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, password }),
            });

            // Redirecting to a different URL is the strongest signal
            // to the browser to trigger the "Save Password" popup.
            goto("/");
        } catch (err) {
            errorMessage = "Invalid username or password.";
        }
    }
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
            Sign in to your account
        </h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form class="space-y-6" onsubmit={handleSubmit}>
            <div>
                <label for="username" class="block text-sm font-medium leading-6 text-gray-900"> Username </label>
                <div class="mt-2">
                    <input
                        id="username"
                        name="username"
                        type="text"
                        autocomplete="username"
                        required
                        bind:value={username}
                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                    />
                </div>
            </div>

            <div>
                <div class="flex items-center justify-between">
                    <label for="password" class="block text-sm font-medium leading-6 text-gray-900"> Password </label>
                </div>
                <div class="mt-2">
                    <input
                        id="password"
                        name="password"
                        type="password"
                        autocomplete="current-password"
                        required
                        bind:value={password}
                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                    />
                </div>
            </div>

            <div>
                <button
                    type="submit"
                    class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                    Sign in
                </button>
            </div>
            <i class="text-rose-500 m-auto">{errorMessage}</i>
        </form>
    </div>
</div>
