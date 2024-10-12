import { authz } from '@/auth';
import StorehousePage from '@/components/pages/storehouse/storehouse';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListStorehouse } from '@/lib/services/storehouse.service';

async function Page() {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const [res, err] = await getListStorehouse(user.token, { limit: 99999 });
    if (err) {
        if (!(err instanceof Response)) {
            handleErr(err);
        }
        if (err.status !== 400) {
            handleErr(err);
        }
    }

    const stores = res?.data;

    return <StorehousePage stores={stores ? stores : []} />;
}

export default Page;
