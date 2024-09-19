import { Role } from './role';

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
