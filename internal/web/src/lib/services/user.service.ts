import { z } from 'zod';
import { User, UserSchema } from '../zod.schema';
import { ResponseOrError } from './type';
import { Res, ResWithPagination } from '@/types/http';
import fetcher from '../http/fetcher';
import { getAuthH } from './helper';

export const CreateUserSchema = UserSchema.omit({ id: true, role: true });
export type CreateUserRequest = z.infer<typeof CreateUserSchema>;

/**
 * createUser create a new user
 *
 * @param key string
 * @param req CreateUserRequest
 * @returns Promise<ResponseOrError<Res<User>>>
 */
export async function createUser(
    key: string,
    req: CreateUserRequest
): Promise<ResponseOrError<Res<User>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<User>>('/users', req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getUserByID get user by id
 *
 * @param key string
 * @param id
 * @returns
 */
export async function getUserByID(
    key: string,
    id: string
): Promise<ResponseOrError<Res<User>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<User>>(`/users/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListUsers = {
    query?: string | undefined;
    skip?: number | undefined;
    limit?: number | undefined;
};

/**
 * getListUser get a list users
 *
 * @param key strings
 * @param req GetListUsers
 * @returns
 */
export async function getListUser(
    key: string,
    { query = '', skip = 1, limit = 5 }: GetListUsers
): Promise<ResponseOrError<ResWithPagination<User>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<User>>(
            `/users?q=${query}&skip=${skip}&limit=${limit}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export const UpdateUserSchema = UserSchema.omit({ role: true }).partial({
    email: true,
    name: true,
    phone: true,
});
export type UpdateUserRequest = z.infer<typeof UpdateUserSchema>;

/**
 * updateCustomer update user info
 *
 * @param key string
 * @param req UpdateUserRequest
 * @returns Promise<ResponseOrError<Res<User>>>
 */
export async function updateCustomer(
    key: string,
    req: UpdateUserRequest
): Promise<ResponseOrError<Res<User>>> {
    const { id, ...body } = req;

    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .patch.json<Res<User>>(`/users/${id}`, body);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * deleteCustomer delete user by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<undefined>>>
 */
export async function deleteUser(
    key: string,
    id: number
): Promise<ResponseOrError<Res<undefined>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .delete<Res<undefined>>(`/users/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
