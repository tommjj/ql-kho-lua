import { z } from 'zod';

export function isEmail(v: string): boolean {
    return z.string().email().safeParse(v).success;
}
