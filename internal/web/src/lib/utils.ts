import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type asyncFunc = (...args: any[]) => Promise<any>;
/**
 * withError is a decorator func to decorated func will return with error\
 * ! unsafe type if use for func that infer return types
 **/
export function withError<T extends asyncFunc>(
    func: T
): (
    ...a: Parameters<T>
) => Promise<[ReturnType<T>, undefined] | [undefined, unknown]> {
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
