import adapter from "@sveltejs/adapter-static";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: vitePreprocess(),
    kit: {
        adapter: adapter({
            // fallback is essential for SPA mode
            fallback: "index.html", // You can also use 'index.html'
            strict: false,
        }),
    },
};

export default config;
