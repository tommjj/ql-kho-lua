import { authz } from '@/auth';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListWarehouse } from '@/lib/services/warehouse.service';
import { Role } from '@/types/role';
import { redirect } from 'next/navigation';

const DashboardPage = async () => {
    const session = await authz();
    if (!session) {
        handleErr(ErrUnauthorized);
    }

    if (session.role === Role.ROOT) {
        redirect('/dashboard/root');
    }

    const [res, err] = await getListWarehouse(session.token, { limit: 10000 });
    if (err) {
        handleErr(err);
    }
    redirect(`/dashboard/${res.data[0]}`);

    return <div className="size-full"></div>;
};

export default DashboardPage;
