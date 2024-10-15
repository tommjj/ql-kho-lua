import { z } from 'zod';

import { Rice, RiceSchema } from '../zod.schema';
import fetcher from '../http/fetcher';
import { ResponseOrError } from './type';
import { Res, ResWithPagination } from '@/types/http';
import { getAuthH } from './helper';

export const CreateRiceSchema = RiceSchema.omit({ id: true });
export type CreateRiceRequest = z.infer<typeof CreateRiceSchema>;

/**
 * createRice create a new rice
 *
 * @param key string
 * @param req CreateRiceRequest
 * @returns Promise<ResponseOrError<Res<Rice>>>
 */
export async function createRice(
    key: string,
    req: CreateRiceRequest
): Promise<ResponseOrError<Res<Rice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Rice>>('/rice', req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getRiceByID get rice by rice id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<Rice>>>
 */
export async function getRiceByID(
    key: string,
    id: number
): Promise<ResponseOrError<Res<Rice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Rice>>(`/rice/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListRice = {
    query?: string | undefined;
    skip?: number | undefined;
    limit?: number | undefined;
};

/**
 * getListStorehouse get list rice
 *
 * @param key string
 * @param getListRiceReq GetListRice
 * @returns Promise<ResponseOrError<ResWithPagination<Rice>>>
 */
export async function getListRice(
    key: string,
    { query = '', skip = 1, limit = 5 }: GetListRice
): Promise<ResponseOrError<ResWithPagination<Rice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Rice>>(
            `/rice?q=${query}&skip=${skip}&limit=${limit}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export const UpdateRiceSchema = RiceSchema;
export type UpdateRiceRequest = z.infer<typeof UpdateRiceSchema>;

/**
 * updateRice update a rice
 *
 * @param key string
 * @param req UpdateRiceRequest
 * @returns Promise<ResponseOrError<Res<Rice>>>
 */
export async function updateRice(
    key: string,
    req: UpdateRiceRequest
): Promise<ResponseOrError<Res<Rice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .patch.json<Res<Rice>>(`/rice/${req.id}`, { name: req.name });

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * deleteRice delete a rice
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<undefined>>>
 */
export async function deleteRice(
    key: string,
    id: number
): Promise<ResponseOrError<Res<undefined>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .delete<Res<undefined>>(`/rice/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
