import { Role } from './role';

/**
 * Token payload is the data that is stored in the JWT token.
 */
export type TokenPayload = {
    id: number;
    name: string;
    email: string;
    role: Role;
};

export type SessionPayload = {
    id: number;
    name: string;
    email: string;
    role: Role;
    token: string;
};
