// $lib/notifications.svelte.ts

// We use an object to wrap our state so it stays reactive when imported
class NotificationManager {
    #message = $state<string | null>(null);

    get message() {
        return this.#message;
    }

    show(msg: string) {
        this.#message = msg;
        // Auto-clear after 5 seconds
        setTimeout(() => {
            this.#message = null;
        }, 5000);
    }

    clear() {
        this.#message = null;
    }
}

export const notifications = new NotificationManager();
