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
import { Warehouse } from '@/lib/zod.schema';
import Link from 'next/link';

type Props = {
    warehouses: Warehouse[];
};

function StoreSelector({ warehouses }: Props) {
    const user = useSession();
    const { push } = useRouter();
    const { warehouseID } = useParams<{ warehouseID: string }>();
    const [open, setOpen] = useState(false);

    const items = useMemo(
        () =>
            user.role === Role.ROOT
                ? [{ id: 0, name: 'Root' }, ...warehouses]
                : [...warehouses],
        [warehouses, user.role]
    );

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <div className="flex ">
                {user.role === Role.ROOT && warehouseID ? (
                    <Button className="px-3 mr-1" asChild>
                        <Link href="/dashboard/root">
                            <Boxes className="size-[18px]  opacity-80" />
                        </Link>
                    </Button>
                ) : (
                    <></>
                )}

                <PopoverTrigger asChild>
                    <Button
                        variant="outline"
                        role="combobox"
                        aria-expanded={open}
                        className="w-full justify-between"
                    >
                        <div className="flex items-start justify-center">
                            {warehouseID ? (
                                <Box className="size-[18px] mr-2 opacity-80" />
                            ) : (
                                <Boxes className="size-[18px] mr-2 opacity-80" />
                            )}
                            {warehouseID
                                ? warehouses.find(
                                      (warehouse) =>
                                          warehouse.id.toString() ===
                                          warehouseID
                                  )?.name
                                : 'Root'}
                        </div>
                        <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                    </Button>
                </PopoverTrigger>
            </div>

            <PopoverContent className="w-[260px] p-0">
                <Command>
                    <CommandInput placeholder="Search warehouse..." />
                    <CommandList>
                        <CommandEmpty>No warehouse found.</CommandEmpty>
                        <CommandGroup>
                            {items.map((warehouse) => (
                                <CommandItem
                                    className="cursor-pointer"
                                    key={warehouse.id}
                                    value={warehouse.name}
                                    onSelect={() => {
                                        if (warehouse.id === 0) {
                                            push(`/dashboard/root`);
                                        } else {
                                            push(`/dashboard/${warehouse.id}`);
                                        }
                                        setOpen(false);
                                    }}
                                >
                                    {warehouse.id ? (
                                        <Box className="size-[18px] mr-2 opacity-80" />
                                    ) : (
                                        <Boxes className="size-[18px] mr-2 opacity-80" />
                                    )}
                                    {warehouse.name}
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
