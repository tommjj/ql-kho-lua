import { authz } from '@/auth';

const DashboardPage = async () => {
    const session = await authz();

    return (
        <div className="size-full">
            <h1>Dashboard</h1>
            <p>{JSON.stringify(session)}</p>
        </div>
    );
};

export default DashboardPage;
