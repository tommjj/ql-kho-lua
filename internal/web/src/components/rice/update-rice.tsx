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
import { updateRice, UpdateRiceSchema } from '@/lib/services/rice.service';
import { useSession } from '../session-context';
import { useRouter } from 'next/navigation';
import { Rice } from '@/lib/zod.schema';

export default function UpdateRiceModal({ rice }: { rice: Rice }) {
    const { refresh } = useRouter();
    const user = useSession();
    const [isOpen, setIsOpen] = useState(false);
    const [name, setName] = useState(rice.name);
    const [error, setError] = useState<string | null>();

    const handleToggle = useCallback(() => {
        setIsOpen((priv) => !priv);
    }, []);

    const handleSubmit = useCallback(() => {
        const parse = UpdateRiceSchema.safeParse({
            id: rice.id,
            name: name,
        });

        if (!parse.success) {
            setError('name must be longer than 3 characters');
            return;
        }

        (async function () {
            const [res, err] = await updateRice(user.token, parse.data);
            if (res) {
                refresh();
                handleToggle();
                return;
            }

            if (!(err instanceof Response)) return;
            switch (err.status) {
                case 409:
                    setError('rice name is exist');
                    break;
                case 400:
                    const data = await err.json();
                    setError(data.messages[0] as string);
                    break;
            }
        })();
    }, [handleToggle, name, refresh, rice.id, user.token]);

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
                    <DialogTitle>Edit rice</DialogTitle>
                    <DialogDescription>
                        Edit rice here. Click save when you are done.
                    </DialogDescription>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="">
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
                                'focus-visible:ring-red-700': name.length < 3,
                            })}
                            autoFocus
                        />
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
