'use client';

import { Trash } from 'lucide-react';
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from '../shadcn-ui/alert-dialog';
import { Button } from '../shadcn-ui/button';
import { useSession } from '../session-context';
import { useRouter } from 'next/navigation';
import { useCallback } from 'react';
import { Customer } from '@/lib/zod.schema';
import { deleteCustomer } from '@/lib/services/customer.service';

export function DeleteCustomer({ customer }: { customer: Customer }) {
    const { refresh } = useRouter();
    const user = useSession();

    const handleDelete = useCallback(() => {
        (async () => {
            const [, err] = await deleteCustomer(user.token, customer.id);
            if (!err) {
                refresh();
                return;
            }
        })();
    }, [refresh, customer.id, user.token]);

    return (
        <AlertDialog>
            <AlertDialogTrigger asChild>
                <Button
                    className="mr-2 px-2 opacity-85 text-red-700 hover:text-red-700"
                    variant={'outline'}
                >
                    <Trash className="size-4" />
                </Button>
            </AlertDialogTrigger>
            <AlertDialogContent className="bg-white">
                <AlertDialogHeader>
                    <AlertDialogTitle>
                        Are you absolutely sure?
                    </AlertDialogTitle>
                    <AlertDialogDescription>
                        This action cannot be undone. This will permanently
                        delete your account and remove your data from our
                        servers.
                    </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <AlertDialogAction className="px-0">
                        <Button
                            variant={'destructive'}
                            onClick={handleDelete}
                            asChild
                        >
                            <div>Continue</div>
                        </Button>
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    );
}
