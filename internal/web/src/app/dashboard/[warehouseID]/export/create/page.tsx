import { authz } from '@/auth';
import CreateImInvoiceClientPage from './client-page';
import { handleErr } from '@/lib/response';
import { ErrDataNotFound, ErrUnauthorized } from '@/lib/errors';
import { getListCustomers } from '@/lib/services/customer.service';
import {
    getWarehouseByID,
    getWarehouseInventory,
} from '@/lib/services/warehouse.service';
import { Metadata } from 'next/types';

export const fetchCache = 'default-no-store';

export const metadata: Metadata = {
    title: 'Create - import invoice',
};

type Props = {
    params: {
        warehouseID: string;
    };
};

async function CreateImportInvoicePage({ params: { warehouseID } }: Props) {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const numWarehouseID = Number(warehouseID);
    if (!Number.isInteger(numWarehouseID)) {
        handleErr(ErrDataNotFound);
    }

    const riceReq = getWarehouseInventory(user.token, numWarehouseID);
    const customerReq = getListCustomers(user.token, { limit: 1000 });
    const wareHouseReq = getWarehouseByID(user.token, numWarehouseID);

    const [
        [riceRes, riceErr],
        [customerRes, customerErr],
        [warehouseRes, warehouseErr],
    ] = await Promise.all([riceReq, customerReq, wareHouseReq]);
    if (riceErr || customerErr || warehouseErr) {
        handleErr(ErrDataNotFound);
    }

    return (
        <section className="relative w-full h-full">
            <CreateImInvoiceClientPage
                customers={customerRes.data}
                rice={riceRes.data}
                warehouse={warehouseRes.data}
            />
        </section>
    );
}

export default CreateImportInvoicePage;
