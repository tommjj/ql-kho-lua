import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type asyncFunc = (...args: any[]) => Promise<any>;
/**
 * withError is a decorator func to decorated func will return with error \
 * ! unsafe types if use for func that infer return types
 *
 * @param func asyncFunc
 **/
export function withError<T extends asyncFunc>(
    func: T
): (
    ...a: Parameters<T>
) => Promise<[Awaited<ReturnType<T>>, undefined] | [undefined, unknown]> {
    return async (...a: Parameters<T>) => {
        try {
            const data = await func(...a);
            return [data, undefined];
        } catch (err) {
            return [undefined, err];
        }
    };
}

/**
 * catchErr is a helper func catch err and return error as second value
 *
 * @param promise Promise<T>
 * @returns Promise<[T, undefined] | [undefined, Error]>
 */
export async function catchErr<T>(
    promise: Promise<T>
): Promise<[T, undefined] | [undefined, Error]> {
    try {
        const data = await promise;
        return [data, undefined];
    } catch (error) {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const wrappedError =
            error instanceof Error
                ? error
                : new Error((error as object).toString());
        return [undefined, wrappedError];
    }
}

/**
 * noError is a decorator func to decorated func will return null if there is an error
 **/
export function noError<T extends asyncFunc>(func: T): T {
    return (async (...a: Parameters<T>) => {
        try {
            return await func(...a);
        } catch (err) {
            return null;
        }
    }) as T;
}

export function formatDateString(dateString: string): string {
    // Parse the date string
    const date = new Date(dateString);

    // Format the date to dd/mm/yy
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-based
    const year = String(date.getFullYear()).slice(-2); // Get last two digits of the year

    return `${day}/${month}/${year}`;
}

export function formatDateTimeString(dateString: string): string {
    // Parse the date string
    const date = new Date(dateString);

    // Format the date to dd/mm/yy
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-based
    const year = String(date.getFullYear()).slice(-2); // Get last two digits of the year

    // Format the time to hh:mm
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${day}/${month}/${year} - ${hours}:${minutes}`;
}
