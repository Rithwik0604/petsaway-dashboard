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

<div class="min-h-screen hero bg-base-200">
    <div class="flex-col hero-content">
        <div class="text-center">
            <h1 class="text-5xl font-bold">Login now!</h1>
        </div>
        <div class="w-full max-w-sm shadow-2xl card shrink-0 bg-base-100">
            <form class="card-body" onsubmit={handleSubmit}>
                <div class="form-control">
                    <label for="username" class="label">
                        <span class="label-text">Username</span>
                    </label>
                    <input
                        id="username"
                        bind:value={username}
                        type="text"
                        placeholder="username"
                        class="input input-bordered"
                        required
                    />
                </div>
                <div class="form-control">
                    <label for="password" class="label">
                        <span class="label-text">Password</span>
                    </label>
                    <input
                        id="password"
                        bind:value={password}
                        type="password"
                        placeholder="password"
                        class="input input-bordered"
                        required
                    />
                </div>
                {#if errorMessage}
                    <div class="text-sm text-red-500">
                        {errorMessage}
                    </div>
                {/if}
                <div class="mt-6 form-control">
                    <button type="submit" class="btn btn-primary">Login</button>
                </div>
            </form>
        </div>
    </div>
</div>
