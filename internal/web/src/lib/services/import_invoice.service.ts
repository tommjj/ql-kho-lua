import { Res, ResWithPagination } from '@/types/http';
import { ResponseOrError } from './type';
import fetcher from '../http/fetcher';
import { Invoice, InvoiceDetailSchema, InvoiceSchema } from '../zod.schema';
import { getAuthH } from './helper';
import { z } from 'zod';

export const CreateImportInvoiceSchema = InvoiceSchema.omit({
    id: true,
    total_price: true,
    user_name: true,
    warehouse_name: true,
    customer_name: true,
    details: true,
}).merge(
    z.object({
        details: z.array(InvoiceDetailSchema.omit({ name: true })),
    })
);
export type CreateImportInvoiceRequest = z.infer<
    typeof CreateImportInvoiceSchema
>;

/**
 * createImportInvoice create a new import invoice
 *
 * @param key string
 * @param req CreateImportInvoiceRequest
 * @returns Promise<ResponseOrError<Res<Invoice>>>
 */
export async function createImportInvoice(
    key: string,
    req: CreateImportInvoiceRequest
): Promise<ResponseOrError<Res<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Invoice>>(`/import_invoices`, req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getImportInvoiceByID get a invoice by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<Invoice>>>
 */
export async function getImportInvoiceByID(
    key: string,
    id: number
): Promise<ResponseOrError<Res<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Invoice>>(`/import_invoices/${id}`);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

export type GetListInvoicesRequest = {
    skip?: number | undefined;
    limit?: number | undefined;
    start?: Date | undefined;
    end?: Date | undefined;
};

export async function getListImportInvoices(
    key: string,
    { end, limit, skip, start }: GetListInvoicesRequest
): Promise<ResponseOrError<ResWithPagination<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Invoice>>(
            `/import_invoices?skip=${skip}&limit=${limit}${
                start ? `&start=${start.toISOString()}` : ''
            }${end ? `&end=${end.toISOString()}` : ''}`
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
