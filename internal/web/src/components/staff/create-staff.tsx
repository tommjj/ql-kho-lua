'use client';

import { AlertCircle, Plus } from 'lucide-react';
import { Button } from '../shadcn-ui/button';
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from '../shadcn-ui/dialog';
import { Input } from '../shadcn-ui/input';
import { Label } from '../shadcn-ui/label';
import { useCallback, useState } from 'react';
import { cn } from '@/lib/utils';
import { Cross2Icon } from '@radix-ui/react-icons';
import PhoneNumberInput from '../ui/phone-numbers-input';
import { isEmail } from '@/lib/validator/email';
import { useRouter } from 'next/navigation';
import { useSession } from '../session-context';
import { createUser, CreateUserSchema } from '@/lib/services/user.service';

export default function CreateStaffModal() {
    const { refresh } = useRouter();
    const user = useSession();
    const [isOpen, setIsOpen] = useState(false);
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [phone, setPhone] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState<string | null>();

    const handleToggle = useCallback(() => {
        setIsOpen((priv) => !priv);
    }, []);

    const handleSubmit = useCallback(() => {
        if (password !== confirmPassword) {
            setError('password is not match');
            return;
        }

        const parse = CreateUserSchema.safeParse({
            name: name,
            email: email,
            phone: phone,
            password: password,
        });

        if (!parse.success) {
            console.log(parse.error);
            setError(parse.error.errors[0].message);
            return;
        }

        (async function () {
            const [res, err] = await createUser(user.token, parse.data);
            if (res) {
                refresh();
                setName('');
                handleToggle();
                return;
            }

            if (!(err instanceof Response)) return;
            switch (err.status) {
                case 409:
                    setError('Conflicting data error');
                    break;
                case 400:
                    const data = await err.json();
                    setError(data.messages[0] as string);
                    break;
            }
        })();
    }, [
        password,
        confirmPassword,
        name,
        email,
        phone,
        user.token,
        refresh,
        handleToggle,
    ]);

    return (
        <Dialog open={isOpen}>
            <DialogTrigger asChild onClick={handleToggle}>
                <Button>
                    Create staff <Plus className="size-5 ml-1"></Plus>
                </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[525px] bg-white">
                <DialogClose
                    onClick={handleToggle}
                    className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                >
                    <Cross2Icon className="h-4 w-4" />
                    <span className="sr-only">Close</span>
                </DialogClose>
                <DialogHeader>
                    <DialogTitle>Create staff</DialogTitle>
                    <DialogDescription>
                        Create a new staff here. Click create when you are done.
                    </DialogDescription>
                    {error && (
                        <div
                            className="flex items-center text-destructive"
                            id="file-upload-error"
                            role="alert"
                        >
                            <AlertCircle className="h-4 w-4 mr-2" />
                            <span className="text-sm">{error}</span>
                        </div>
                    )}
                </DialogHeader>
                <div className="grid gap-4 py-4 ">
                    <div className="space-y-2">
                        <div>
                            <Label htmlFor="name" className="text-right">
                                Name
                            </Label>
                            <Input
                                id="name"
                                onChange={(e) => {
                                    setError('');
                                    setName(e.target.value);
                                }}
                                value={name}
                                className={cn('', {
                                    'focus-visible:ring-red-700 ring-red-700':
                                        name.length < 3,
                                })}
                                placeholder="Staff name"
                                autoFocus
                            />
                        </div>
                        <div>
                            <Label htmlFor="email" className="text-right">
                                Email
                            </Label>
                            <Input
                                id="email"
                                type="email"
                                value={email}
                                onChange={(e) => {
                                    setError('');
                                    setEmail(e.target.value);
                                }}
                                className={cn('', {
                                    'focus-visible:ring-red-700 ring-red-700':
                                        !isEmail(email),
                                })}
                                placeholder="Staff email"
                            />
                        </div>

                        <PhoneNumberInput
                            onChanged={(v) => {
                                setError('');
                                setPhone(v);
                            }}
                            defaultCountryCode="VN"
                        />
                        <div>
                            <Label htmlFor="password" className="text-right">
                                Password
                            </Label>
                            <Input
                                id="password"
                                type="password"
                                value={password}
                                onChange={(e) => {
                                    setError('');
                                    setPassword(e.target.value);
                                }}
                                className={cn('', {
                                    'focus-visible:ring-red-700 ring-red-700':
                                        password.length < 3,
                                })}
                                placeholder="Password"
                            />
                        </div>
                        <div>
                            <Label
                                htmlFor="confirm_password"
                                className="text-right"
                            >
                                Confirm password
                            </Label>
                            <Input
                                type="password"
                                id="confirm_password"
                                value={confirmPassword}
                                onChange={(e) => {
                                    setError('');
                                    setConfirmPassword(e.target.value);
                                }}
                                className={cn('', {
                                    'focus-visible:ring-red-700 ring-red-700 ring-1':
                                        password.length < 3,
                                })}
                                placeholder="Confirm password"
                            />
                        </div>
                    </div>
                </div>
                <DialogFooter>
                    <Button type="submit" onClick={handleSubmit}>
                        Create
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
