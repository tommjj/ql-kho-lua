'use client';

import { AlertCircle, FilePenLine } from 'lucide-react';
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
import { useSession } from '../session-context';
import { useRouter } from 'next/navigation';
import { Customer } from '@/lib/zod.schema';
import PhoneNumberInput from '../ui/phone-numbers-input';
import { isEmail } from '@/lib/validator/email';
import {
    updateCustomer,
    UpdateCustomerSchema,
} from '@/lib/services/customer.service';

export default function UpdateCustomerModal({
    customer,
}: {
    customer: Customer;
}) {
    const { refresh } = useRouter();
    const user = useSession();
    const [isOpen, setIsOpen] = useState(false);
    const [name, setName] = useState(customer.name);
    const [email, setEmail] = useState(customer.email);
    const [phone, setPhone] = useState(customer.phone);
    const [address, setAddress] = useState(customer.address);
    const [error, setError] = useState<string | null>();

    const handleToggle = useCallback(() => {
        setIsOpen((priv) => !priv);
    }, []);

    const handleSubmit = useCallback(() => {
        const parse = UpdateCustomerSchema.safeParse({
            id: customer.id,
            name: name,
            email: email,
            phone: phone,
            address: address,
        });

        if (!parse.success) {
            setError(parse.error.errors[0].message);
            return;
        }

        (async function () {
            const [res, err] = await updateCustomer(user.token, parse.data);
            if (res) {
                refresh();
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
        address,
        customer.id,
        email,
        handleToggle,
        name,
        phone,
        refresh,
        user.token,
    ]);

    return (
        <Dialog open={isOpen}>
            <DialogTrigger asChild onClick={handleToggle}>
                <Button className="mr-2 px-2 opacity-85" variant={'outline'}>
                    <FilePenLine className="size-4 text-amber-700 hover:text-amber-700" />
                </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px] bg-white">
                <DialogClose
                    onClick={handleToggle}
                    className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                >
                    <Cross2Icon className="h-4 w-4" />
                    <span className="sr-only">Close</span>
                </DialogClose>
                <DialogHeader>
                    <DialogTitle>Edit customer</DialogTitle>
                    <DialogDescription>
                        Edit customer here. Click save when you are done.
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
                                    'focus-visible:ring-red-700':
                                        name.length < 3,
                                })}
                                placeholder="Customer name"
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
                                    'focus-visible:ring-red-700':
                                        !isEmail(email),
                                })}
                                placeholder="Customer email"
                            />
                        </div>

                        <PhoneNumberInput
                            onChanged={(v) => {
                                setError('');
                                setPhone(v);
                            }}
                            defaultValue={customer.phone}
                            defaultCountryCode="VN"
                        />
                        <div>
                            <Label htmlFor="address" className="text-right">
                                Address
                            </Label>
                            <Input
                                id="name"
                                value={address}
                                onChange={(e) => {
                                    setError('');
                                    setAddress(e.target.value);
                                }}
                                className={cn('', {
                                    'focus-visible:ring-red-700':
                                        address.length < 3,
                                })}
                                placeholder="Customer address"
                            />
                        </div>
                    </div>
                </div>
                <DialogFooter>
                    <Button type="submit" onClick={handleSubmit}>
                        Save change
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
