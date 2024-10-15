import { z } from 'zod';
import { Customer, CustomerSchema } from '../zod.schema';
import { ResponseOrError } from './type';
import { Res, ResWithPagination } from '@/types/http';
import fetcher from '../http/fetcher';
import { getAuthH } from './helper';

export const CreateCustomerSchema = CustomerSchema.omit({ id: true });
export type CreateCustomerRequest = z.infer<typeof CreateCustomerSchema>;

/**
 * createCustomer create a new customer
 *
 * @param key string
 * @param req CreateCustomerRequest
 * @returns Promise<ResponseOrError<Res<Customer>>>
 */
export async function createCustomer(
    key: string,
    req: CreateCustomerRequest
): Promise<ResponseOrError<Res<Customer>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Customer>>('/customers', req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getCustomerByID get customer by id
 *
 * @param key string
 * @param id number
 * @returns
 */
export async function getCustomerByID(key: string, id: number) {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Customer>>(`/customers/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListCustomer = {
    query?: string | undefined;
    skip?: number | undefined;
    limit?: number | undefined;
};

/**
 * getListCustomers get list customer
 *
 * @param key string
 * @param query GetListCustomer
 * @returns Promise<ResponseOrError<ResWithPagination<Customer>>>
 */
export async function getListCustomers(
    key: string,
    { query = '', skip = 1, limit = 5 }: GetListCustomer
): Promise<ResponseOrError<ResWithPagination<Customer>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Customer>>(
            `/customers?q=${query}&skip=${skip}&limit=${limit}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export const UpdateCustomerSchema = CustomerSchema.partial({
    name: true,
    address: true,
    email: true,
    phone: true,
});
export type UpdateCustomerRequest = z.infer<typeof UpdateCustomerSchema>;

/**
 * updateCustomer update customer info
 *
 * @param key string
 * @param req UpdateCustomerRequest
 * @returns Promise<ResponseOrError<Res<Customer>>>
 */
export async function updateCustomer(
    key: string,
    req: UpdateCustomerRequest
): Promise<ResponseOrError<Res<Customer>>> {
    const { id, ...body } = req;

    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .patch.json<Res<Customer>>(`/customers/${id}`, body);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * deleteCustomer delete customer by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<undefined>>>
 */
export async function deleteCustomer(
    key: string,
    id: number
): Promise<ResponseOrError<Res<undefined>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .delete<Res<undefined>>(`/customers/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
