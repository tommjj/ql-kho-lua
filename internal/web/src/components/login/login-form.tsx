'use client';

import LoginAction from '@/lib/actions/login';
import { useState } from 'react';
import { useFormState, useFormStatus } from 'react-dom';
import { Button } from '../shadcn-ui/button';
import { Input } from '../shadcn-ui/input';
import Link from 'next/link';
import { Label } from '../shadcn-ui/label';

function LoginForm() {
    const [changed, setChanged] = useState(false);
    const [code, action] = useFormState(LoginAction, undefined);

    return (
        <form
            className="grid gap-4"
            onSubmit={() => {
                setTimeout(() => setChanged(false), 500);
            }}
            action={action}
        >
            <div id="error" className="text-red-400 text-center">
                {!changed && code && code}
            </div>
            <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                    id="email"
                    type="email"
                    name="email"
                    placeholder="m@example.com"
                    required
                    onChange={() => {
                        setChanged(true);
                    }}
                />
            </div>
            <div className="grid gap-2">
                <div className="flex items-center">
                    <Label htmlFor="password">Password</Label>
                    <Link
                        href="/forgot-password"
                        className="ml-auto inline-block text-sm underline"
                    >
                        Forgot your password?
                    </Link>
                </div>
                <Input
                    id="password"
                    type="password"
                    name="password"
                    required
                    onChange={() => {
                        setChanged(true);
                    }}
                />
            </div>
            <SignInButton />
        </form>
    );
}

function SignInButton() {
    const { pending } = useFormStatus();
    return (
        <Button type="submit" className="w-full" aria-disabled={pending}>
            Login
        </Button>
    );
}

export default LoginForm;
