'use client';

import { useSession } from '@/components/session-context';
import { Button } from '@/components/shadcn-ui/button';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuTrigger,
} from '@/components/shadcn-ui/dropdown-menu';
import { Input } from '@/components/shadcn-ui/input';
import { Progress } from '@/components/shadcn-ui/progress';
import { CreateStorehouse } from '@/components/storehouse/create-storehouse';
import { DeleteStorehouse } from '@/components/storehouse/delete-storehouse';
import { UpdateStorehouse } from '@/components/storehouse/update-storehouse';
import { getUsedCapacity } from '@/lib/services/storehouse.service';
import { Storehouse } from '@/lib/zod.schema';
import { Box, Ellipsis, Search } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { ChangeEvent, useCallback, useEffect, useState } from 'react';

type Props = {
    storehouses: Storehouse[];
    mapLocationControl(longitude: number, latitude: number): void;
};

function MoreOption({ storehouse }: { storehouse: Storehouse }) {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button className="h-8 px-2" variant="ghost">
                    <Ellipsis className="size-4 " />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent
                className="w-56"
                onClick={(e) => {
                    e.stopPropagation();
                }}
                onDoubleClick={(e) => {
                    e.stopPropagation();
                }}
            >
                <DropdownMenuGroup>
                    <UpdateStorehouse storehouse={storehouse} />
                    <DeleteStorehouse storeID={storehouse.id} />
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

function StorehouseItem({
    storehouse,
    mapLocationControl,
}: {
    storehouse: Storehouse;
    mapLocationControl(longitude: number, latitude: number): void;
}) {
    const { push } = useRouter();
    const user = useSession();
    const [used, setUsed] = useState(0);

    useEffect(() => {
        (async () => {
            const [res, err] = await getUsedCapacity(user.token, storehouse.id);
            if (!err) {
                setUsed(res.data.used_capacity);
            }
        })();

        const intervalID = setInterval(async () => {
            const [res, err] = await getUsedCapacity(user.token, storehouse.id);
            if (!err) {
                setUsed(res.data.used_capacity);
            }
        }, 1000 * 60 * 5);

        return () => clearInterval(intervalID);
    }, [storehouse.id, user.token]);

    const handleDoubleClick = useCallback(() => {
        push(`/dashboard/${storehouse.id}`);
    }, [push, storehouse.id]);

    const handleClick = useCallback(() => {
        mapLocationControl(storehouse.location[1], storehouse.location[0]);
    }, [mapLocationControl, storehouse.location]);

    return (
        <Button
            className="w-full h-auto p-0 cursor-pointer select-none"
            onDoubleClick={handleDoubleClick}
            onClick={handleClick}
            variant="ghost"
            asChild
        >
            <div>
                <div className="w-full mb-1 rounded p-1.5">
                    <div>
                        <div className="flex items-center text-lg pt-1">
                            <div className="flex items-center flex-grow">
                                <Box className="size-5 mr-2 opacity-80"></Box>
                                {storehouse.name}
                            </div>
                            <MoreOption storehouse={storehouse}></MoreOption>
                        </div>

                        <Progress
                            value={
                                (used / storehouse.capacity) * 100 > 100
                                    ? 100
                                    : (used / storehouse.capacity) * 100
                            }
                            className="w-full my-1"
                        />
                        <div className="flex">
                            <div className="flex-grow text-sm opacity-80">{`${
                                storehouse.capacity - used
                            } free of ${storehouse.capacity}`}</div>
                        </div>
                    </div>
                </div>
            </div>
        </Button>
    );
}

function StorehouseList({ storehouses, mapLocationControl }: Props) {
    const [search, setSearch] = useState('');

    const handleInputChange = useCallback(
        (e: ChangeEvent<HTMLInputElement>) => {
            setSearch(e.target.value);
        },
        []
    );

    return (
        <div className="flex flex-col size-full max-h-screen">
            <div className="group relative w-full p-2 border-b flex">
                <Search className="absolute p-2 size-9 opacity-80 -mt-[1px]" />
                <Input
                    onChange={handleInputChange}
                    value={search}
                    className=" pl-9 focus-visible:ring-none focus-visible:ring-0 mr-1"
                    placeholder="Search..."
                ></Input>

                <CreateStorehouse />
            </div>

            <div className="relative flex-grow w-full h-full  p-2 max-h-screen">
                <div className="absolute inset-0 overflow-y-auto px-2 py-1.5">
                    {storehouses.map((storehouse) =>
                        storehouse.name
                            .toLocaleLowerCase()
                            .includes(search.trim().toLocaleLowerCase()) ? (
                            <StorehouseItem
                                mapLocationControl={mapLocationControl}
                                key={storehouse.id}
                                storehouse={storehouse}
                            />
                        ) : (
                            <></>
                        )
                    )}
                </div>
            </div>
        </div>
    );
}

export default StorehouseList;
