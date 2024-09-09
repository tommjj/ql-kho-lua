import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

/**
 * withError
 **/
export const withError = <T extends (...args: any[]) => Promise<any>>(
    func: T
): ((
    ...a: Parameters<T>
) => Promise<[ReturnType<T>, undefined] | [undefined, unknown]>) => {
    return async (...a: Parameters<T>) => {
        try {
            const data = await func(...a);
            return [data, undefined];
        } catch (err) {
            return [undefined, err];
        }
    };
};

/**
 * noError is a decorator func to decorated func will return null if there is an error
 **/
export const noError = <T extends (...args: any[]) => Promise<any>>(
    func: T
): T => {
    return (async (...a: Parameters<T>) => {
        try {
            return await func(...a);
        } catch (err) {
            return null;
        }
    }) as T;
};
