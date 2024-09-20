'use client';

import { useState } from 'react';
import { Popover, PopoverContent, PopoverTrigger } from '../shadcn-ui/popover';
import { Button } from '../shadcn-ui/button';
import { ChevronsUpDown } from 'lucide-react';
import {
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    Command,
} from '../shadcn-ui/command';
import { CubeIcon } from '@radix-ui/react-icons';

type Storehouse = {
    id: number;
    name: string;
};

type Props = {
    storehouses: Storehouse[];
};

function StoreSelector({ storehouses }: Props) {
    const [open, setOpen] = useState(false);
    const [value, setValue] = useState(0);

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
                        <CubeIcon className="size-[18px] mr-2 opacity-80" />
                        {value
                            ? storehouses.find(
                                  (storehouse) => storehouse.id === value
                              )?.name
                            : 'Select storehouse...'}
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
                            {storehouses.map((storehouse) => (
                                <CommandItem
                                    key={storehouse.id}
                                    value={storehouse.id.toString()}
                                    onSelect={(currentValue) => {
                                        setValue(
                                            Number(currentValue) === value
                                                ? 0
                                                : storehouse.id
                                        );
                                        setOpen(false);
                                    }}
                                >
                                    <CubeIcon className="size-5 mr-2 opacity-80" />
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
