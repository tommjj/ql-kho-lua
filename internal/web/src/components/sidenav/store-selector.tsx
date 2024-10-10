'use client';

import { useMemo, useState } from 'react';
import { Popover, PopoverContent, PopoverTrigger } from '../shadcn-ui/popover';
import { Button } from '../shadcn-ui/button';
import { ChevronsUpDown, Box, Boxes } from 'lucide-react';
import {
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    Command,
} from '../shadcn-ui/command';
import { useParams, useRouter } from 'next/navigation';
import { useSession } from '../session-context';
import { Role } from '@/types/role';
import { Storehouse } from '@/lib/zod.schema';

type Props = {
    storehouses: Storehouse[];
};

function StoreSelector({ storehouses }: Props) {
    const user = useSession();
    const { push } = useRouter();
    const { storeID } = useParams<{ storeID: string }>();
    const [open, setOpen] = useState(false);

    const items = useMemo(
        () =>
            user.role === Role.ROOT
                ? [{ id: 0, name: 'Root' }, ...storehouses]
                : [...storehouses],
        [storehouses, user.role]
    );

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className="w-full justify-between"
                >
                    <div className="flex items-start justify-center">
                        {storeID ? (
                            <Box className="size-[18px] mr-2 opacity-80" />
                        ) : (
                            <Boxes className="size-[18px] mr-2 opacity-80" />
                        )}
                        {storeID
                            ? storehouses.find(
                                  (storehouse) =>
                                      storehouse.id.toString() === storeID
                              )?.name
                            : 'Root'}
                    </div>
                    <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="w-[260px] p-0">
                <Command>
                    <CommandInput placeholder="Search storehouse..." />
                    <CommandList>
                        <CommandEmpty>No storehouse found.</CommandEmpty>
                        <CommandGroup>
                            {items.map((storehouse) => (
                                <CommandItem
                                    key={storehouse.id}
                                    value={storehouse.name}
                                    onSelect={() => {
                                        if (storehouse.id === 0) {
                                            push(`/dashboard/root`);
                                        } else {
                                            push(`/dashboard/${storehouse.id}`);
                                        }
                                        setOpen(false);
                                    }}
                                >
                                    {storehouse.id ? (
                                        <Box className="size-[18px] mr-2 opacity-80" />
                                    ) : (
                                        <Boxes className="size-[18px] mr-2 opacity-80" />
                                    )}
                                    {storehouse.name}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    );
}

export default StoreSelector;
