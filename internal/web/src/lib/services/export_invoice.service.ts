import { Res, ResWithPagination } from '@/types/http';
import { ResponseOrError } from './type';
import fetcher from '../http/fetcher';
import { Invoice, InvoiceDetailSchema, InvoiceSchema } from '../zod.schema';
import { getAuthH } from './helper';
import { z } from 'zod';

export const CreateExportInvoiceSchema = InvoiceSchema.omit({
    id: true,
    total_price: true,
    user_name: true,
    user_id: true,
    warehouse_name: true,
    customer_name: true,
    details: true,
    created_at: true,
}).merge(
    z.object({
        details: z.array(InvoiceDetailSchema.omit({ name: true })),
    })
);
export type CreateExportInvoiceRequest = z.infer<
    typeof CreateExportInvoiceSchema
>;

/**
 * createExportInvoice create a new export invoice
 *
 * @param key string
 * @param req CreateExportInvoiceRequest
 * @returns Promise<ResponseOrError<Res<Invoice>>>
 */
export async function createExportInvoice(
    key: string,
    req: CreateExportInvoiceRequest
): Promise<ResponseOrError<Res<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.json<Res<Invoice>>(`/export_invoices`, req);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}

/**
 * getExportInvoiceByID get a invoice by id
 *
 * @param key string
 * @param id number
 * @returns Promise<ResponseOrError<Res<Invoice>>>
 */
export async function getExportInvoiceByID(
    key: string,
    id: number
): Promise<ResponseOrError<Res<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<Res<Invoice>>(`/export_invoices/${id}`);

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
    warehouseID?: number;
};

export async function getListExportInvoices(
    key: string,
    { end, limit, skip, start, warehouseID }: GetListInvoicesRequest
): Promise<ResponseOrError<ResWithPagination<Invoice>>> {
    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .get<ResWithPagination<Invoice>>(
            `/export_invoices?skip=${skip}&limit=${limit}${
                start ? `&start=${start.toISOString()}` : ''
            }${end ? `&end=${end.toISOString()}` : ''}${
                warehouseID ? `&warehouse_id=${warehouseID}` : ''
            }`,
            {
                cache: 'no-store',
            }
        );

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
