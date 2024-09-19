import { auth } from '@/auth';

const DashboardPage = async () => {
    const session = await auth();

    return (
        <div className="size-full">
            <h1>Dashboard</h1>
            <p>{'adaw ' + JSON.stringify(session?.user)}</p>
        </div>
    );
};

export default DashboardPage;
