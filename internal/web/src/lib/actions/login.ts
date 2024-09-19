'use server';

import { AuthError } from 'next-auth';
import { signIn } from '../../auth';
import { redirect } from 'next/navigation';

async function LoginAction(privState: string | undefined, formData: FormData) {
    try {
        await signIn('credentials', Object.fromEntries(formData));
    } catch (error) {
        if (error instanceof AuthError) {
            console.log(error.type);
            switch (error.type) {
                case 'CredentialsSignin':
                    return 'Invalid credentials.';
                default:
                    return 'Something went wrong.';
            }
        }
        if ((error as Error).message === 'NEXT_REDIRECT') {
            redirect('/dashboard');
        }
    }
    return undefined;
}

export default LoginAction;
