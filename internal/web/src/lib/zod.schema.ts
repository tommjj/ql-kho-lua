import { z } from 'zod';
import { validateE164 } from '@/lib/validator/e164';
import { Role } from '@/types/role';

// phoneNumber validator
const phoneNumber = z.string().refine((v) => validateE164(v), {
    message: 'phone numbers is not valid',
});

// WarehouseSchema
export const WarehouseSchema = z.object({
    id: z.number().int(),
    name: z.string().min(3).max(255),
    location: z.array(z.number()).length(2),
    capacity: z.number().min(1),
    image: z.string().min(4),
});
export type Warehouse = z.infer<typeof WarehouseSchema>;

export const WarehouseItemSchema = z.object({
    id: z.number().int(),
    rice_name: z.string(),
    capacity: z.number().min(1),
});
export type WarehouseItem = z.infer<typeof WarehouseItemSchema>;

// UserSchema
export const UserSchema = z.object({
    id: z.number().int(),
    name: z.string().min(3).max(32),
    email: z.string().email(),
    phone: phoneNumber,
    role: z.enum([Role.ROOT, Role.MEMBER]),
});
export type User = z.infer<typeof UserSchema>;

//RiceSchema
export const RiceSchema = z.object({
    id: z.number().int(),
    name: z.string().min(3).max(50),
});
export type Rice = z.infer<typeof RiceSchema>;

// CustomerSchema
export const CustomerSchema = z.object({
    id: z.number().int(),
    name: z.string().min(3).max(255),
    email: z.string().email(),
    phone: phoneNumber,
    address: z
        .string()
        .min(1, 'Address must contain at least 1 character(s)')
        .max(255),
});
export type Customer = z.infer<typeof CustomerSchema>;
