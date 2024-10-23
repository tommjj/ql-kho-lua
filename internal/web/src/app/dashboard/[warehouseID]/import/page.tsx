import Header from '@/components/ui/header';
import ImportClientPage from './client-page';
import { authz } from '@/auth';
import { handleErr } from '@/lib/response';
import { getListImportInvoices } from '@/lib/services/import_invoice.service';
import { z } from 'zod';
import { ErrDataNotFound } from '@/lib/errors';
import { Button } from '@/components/shadcn-ui/button';
import Link from 'next/link';
import { Plus } from 'lucide-react';
import { Metadata } from 'next/types';

export const metadata: Metadata = {
    title: 'Import invoices',
};

type Props = {
    searchParams: {
        page?: string;
        start?: string;
        end?: string;
    };
    params: {
        warehouseID: string;
    };
};

const searchParamsParser = z.object({
    start: z.coerce.date().optional(),
    end: z.coerce.date().optional(),
    page: z.coerce.number().int().min(1).optional().default(1),
});

async function ImportPage({ searchParams, params: { warehouseID } }: Props) {
    const user = await authz();
    if (!user) {
        handleErr(user);
    }

    const parsed = searchParamsParser.safeParse(searchParams);
    if (!parsed.success) {
        handleErr(ErrDataNotFound);
    }

    const [res, err] = await getListImportInvoices(user.token, {
        end: parsed.data.end,
        start: parsed.data.start,
        limit: 10,
        skip: parsed.data.page,
        warehouseID: Number.isInteger(Number(warehouseID))
            ? Number(warehouseID)
            : 0,
    });

    if (err) {
        if (!(err instanceof Response) || err.status !== 404) {
            console.log(err);
            handleErr(err);
        }
    }

    const data = res?.data;
    const pagination = res?.pagination;

    return (
        <section className="relative w-full h-screen">
            <Header className="absolute top-0 left-0 right-0" title="Import">
                <Button asChild>
                    <Link href={`/dashboard/${warehouseID}/import/create`}>
                        Create invoice <Plus className="size-4" />
                    </Link>
                </Button>
            </Header>
            <ImportClientPage
                listInvoices={data ? data : []}
                pagination={
                    pagination
                        ? pagination
                        : {
                              current_page: 1,
                              limit_records: 0,
                              next_page: null,
                              prev_page: null,
                              total_pages: 1,
                              total_records: 0,
                          }
                }
            />
        </section>
    );
}

export default ImportPage;
