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

export function formatTimeForInput(timeValue: any): string {
    if (!timeValue) return "";

    // If Go sends "17:12:00", we need to trim it to "17:12"
    if (typeof timeValue === "string") {
        return timeValue.slice(0, 5);
    }

    return "";
}

export function formatDateTimeForInput(dateTimeString: string): string {
    if (!dateTimeString) return "";
    const date = new Date(dateTimeString);
    // Format to "YYYY-MM-DDTHH:mm"
    return date.toISOString().slice(0, 16);
}

export function formatTo12Hour(timeStr: string): string {
    if (!timeStr) return "";

    // Split "17:12" into [17, 12]
    const [hours, minutes] = timeStr.split(":").map(Number);
    if (isNaN(hours) || isNaN(minutes)) return "";

    const suffix = hours >= 12 ? "PM" : "AM";
    const adjustedHour = hours % 12 || 12; // Converts 0 to 12 and 13 to 1

    // Return formatted string: "05:12 PM"
    return `${adjustedHour.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")} ${suffix}`;
}
