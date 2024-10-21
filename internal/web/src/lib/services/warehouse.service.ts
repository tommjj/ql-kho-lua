import { z } from 'zod';

import { WarehouseSchema, Warehouse, WarehouseItem } from '../zod.schema';
import fetcher from '../http/fetcher';
import { ResponseOrError } from './type';
import { Res, ResWithPagination } from '@/types/http';
import { getAuthH } from './helper';

export const CreateWarehouseSchema = WarehouseSchema.omit({ id: true });
export type CreateWarehouseRequest = z.infer<typeof CreateWarehouseSchema>;

/**
 * createWarehouse create a new warehouse
 *
 * @param key string
 * @param req CreateWarehouseRequest
 * @returns Promise<ResponseOrError<Res<Warehouse>>>
 */
export async function createWarehouse(
    key: string,
    req: CreateWarehouseRequest
): Promise<ResponseOrError<Res<Warehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Warehouse>>('/warehouses', req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getWarehouseByID get warehouse by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<Warehouse>>>
 */
export async function getWarehouseByID(
    key: string,
    id: number
): Promise<ResponseOrError<Res<Warehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Warehouse>>(`/warehouses/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListWarehouse = {
    query?: string | undefined;
    skip?: number | undefined;
    limit?: number | undefined;
};

/**
 * getListWarehouse get list warehouse
 *
 * @param key string
 * @param req GetListWarehouse
 * @returns Promise<ResponseOrError<ResWithPagination<Warehouse>>>
 */
export async function getListWarehouse(
    key: string,
    { query = '', skip = 1, limit = 5 }: GetListWarehouse
): Promise<ResponseOrError<ResWithPagination<Warehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Warehouse>>(
            `/warehouses?q=${query}&skip=${skip}&limit=${limit}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getUsedCapacity get use capacity of warehouse by id
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
        .get<Res<{ used_capacity: number }>>(`/warehouses/${id}/used_capacity`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export async function getWarehouseInventory(
    key: string,
    id: number
): Promise<ResponseOrError<Res<WarehouseItem[]>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<WarehouseItem[]>>(`/warehouses/${id}/inventory`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export const UpdateWarehouseSchema = WarehouseSchema.partial({
    image: true,
    capacity: true,
    location: true,
    name: true,
});
export type UpdateWarehouseRequest = z.infer<typeof UpdateWarehouseSchema>;

/**
 * updateWarehouse update warehouse, only update non-zero fields
 *
 * @param key string
 * @param req UpdateWarehouseRequest
 * @returns Promise<ResponseOrError<Res<Warehouse>>>
 */
export async function updateWarehouse(
    key: string,
    req: UpdateWarehouseRequest
): Promise<ResponseOrError<Res<Warehouse>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .patch.json<Res<Warehouse>>(`/warehouses/${req.id}`, req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * deleteWarehouse delete warehouse by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<undefined>>>
 */
export async function deleteWarehouse(
    key: string,
    id: number
): Promise<ResponseOrError<Res<undefined>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .delete<Res<undefined>>(`/warehouses/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
