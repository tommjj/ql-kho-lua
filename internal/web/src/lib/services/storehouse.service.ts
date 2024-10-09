import { z } from 'zod';

import { StorehouseSchema, Storehouse } from '../zod.schema';
import fetcher from '../http/fetcher';
import { ResponseOrError } from './type';
import { Res, ResWithPagination } from '@/types/http';
import { getAuthH } from './helper';

export const CreateStorehouseSchema = StorehouseSchema.omit({ id: true });
export type CreateStorehouseRequest = z.infer<typeof CreateStorehouseSchema>;

/**
 * createStorehouse create a new storehouse
 *
 * @param key string
 * @param req CreateStorehouseRequest
 * @returns Promise<ResponseOrError<Res<Storehouse>>>
 */
export async function createStorehouse(
    key: string,
    req: CreateStorehouseRequest
): Promise<ResponseOrError<Res<Storehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Storehouse>>('/storehouses', req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getStorehouseByID get storehouse by id
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<Storehouse>>>
 */
export async function getStorehouseByID(
    key: string,
    id: number
): Promise<ResponseOrError<Res<Storehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Storehouse>>(`/storehouses/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListStorehouse = {
    query?: string | undefined;
    skip?: number | undefined;
    limit?: number | undefined;
};

/**
 * getListStorehouse get list storehouse
 * @param key string
 * @param req GetListStorehouse
 * @returns Promise<ResponseOrError<ResWithPagination<Storehouse>>>
 */
export async function getListStorehouse(
    key: string,
    { query = '', skip = 1, limit = 5 }: GetListStorehouse
): Promise<ResponseOrError<ResWithPagination<Storehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Storehouse>>(
            `/storehouses?query=${query}&skip=${skip}&limit=${limit}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getUsedCapacity get use capacity of storehouse by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<{ used_capacity: number }>>>
 */
export async function getUsedCapacity(
    key: string,
    id: number
): Promise<ResponseOrError<Res<{ used_capacity: number }>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<{ used_capacity: number }>>(
            `/storehouses/${id}/used_capacity`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export const UpdateStorehouseSchema = StorehouseSchema.partial({
    image: true,
    capacity: true,
    location: true,
    name: true,
});
export type UpdateStorehouseRequest = z.infer<typeof UpdateStorehouseSchema>;

/**
 * updateStorehouse update storehouse, only update non-zero fields
 *
 * @param key string
 * @param req UpdateStorehouseRequest
 * @returns Promise<ResponseOrError<Res<Storehouse>>>
 */
export async function updateStorehouse(
    key: string,
    req: UpdateStorehouseRequest
): Promise<ResponseOrError<Res<Storehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .patch.json<Res<Storehouse>>(`/storehouses`, req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * deleteStorehouse delete storehouse by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<undefined>>>
 */
export async function deleteStorehouse(
    key: string,
    id: number
): Promise<ResponseOrError<Res<undefined>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .delete<Res<undefined>>(`/storehouses/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
