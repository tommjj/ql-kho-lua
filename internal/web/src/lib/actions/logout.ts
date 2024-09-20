'use server';

import { signOut } from '@/auth';

async function logout() {
    await signOut();
}

export default logout;
