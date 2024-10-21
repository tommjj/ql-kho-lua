import { authz } from '@/auth';
import WarehousePage from '@/components/pages/warehouses/warehouse';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListWarehouse } from '@/lib/services/warehouse.service';
import { Metadata } from 'next/types';

export const metadata: Metadata = {
    title: 'Warehouse',
};

async function Page() {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const [res, err] = await getListWarehouse(user.token, { limit: 99999 });
    if (err) {
        if (!(err instanceof Response) || err.status !== 404) {
            handleErr(err);
        }
    }

    const stores = res?.data;

    return <WarehousePage stores={stores ? stores : []} />;
}

export default Page;
