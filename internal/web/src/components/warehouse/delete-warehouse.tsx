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
import { deleteWarehouse } from '@/lib/services/warehouse.service';

export function DeleteWarehouse({ storeID }: { storeID: number }) {
    const { refresh } = useRouter();
    const user = useSession();

    const handleDelete = useCallback(() => {
        (async () => {
            const [, err] = await deleteWarehouse(user.token, storeID);
            if (!err) {
                refresh();
                return;
            }
        })();
    }, [refresh, storeID, user.token]);

    return (
        <AlertDialog>
            <AlertDialogTrigger asChild>
                <Button
                    className="text-red-600 hover:text-red-600 flex items-center justify-start w-full"
                    variant="ghost"
                >
                    <Trash className="size-5 mr-2 " />
                    Delete
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
