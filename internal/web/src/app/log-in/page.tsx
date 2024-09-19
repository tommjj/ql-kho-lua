import LoginForm from '@/components/login/login-form';
import { CubeIcon } from '@radix-ui/react-icons';
import { Metadata } from 'next';

export const metadata: Metadata = {
    title: 'Log in',
    description: 'Log in page',
};

function LoginPage() {
    return (
        <section className="h-screen overflow-hidden">
            <div className="w-full lg:grid lg:grid-cols-2 h-full">
                {/* Left */}
                <div className="hidden bg-primary p-7 lg:block">
                    <CubeIcon className="h-12 w-12 text-primary-foreground" />
                </div>
                {/* Right */}
                <div className="flex items-center h-full justify-center py-12 px-2">
                    <div className="mx-auto grid w-[350px] gap-6">
                        <div className="grid gap-2 text-center">
                            <h1 className="text-4xl font-bold">Login</h1>
                        </div>
                        {/* form */}
                        <LoginForm />
                    </div>
                </div>
            </div>
        </section>
    );
}
export default LoginPage;
