export function formatDate(dateString: string) {
    if (!dateString) return "";
    const date = new Date(dateString);

    return new Intl.DateTimeFormat("en-US", {
        day: "numeric",
        month: "short",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    }).format(date);
}

export function formatDateForInput(dateString: string): string {
    if (!dateString) return "";
    const date = new Date(dateString);
    return date.toISOString().split("T")[0];
}

export function formatTimeForInput(time: any): string {
    if (!time || !time.Valid) return "";
    // Assuming time.Time is in "HH:MM:SS" format
    return time.Time.substring(0, 5);
}

export function formatDateTimeForInput(dateTimeString: string): string {
    if (!dateTimeString) return "";
    const date = new Date(dateTimeString);
    // Format to "YYYY-MM-DDTHH:mm"
    return date.toISOString().slice(0, 16);
}
