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
import { CreateWarehouse } from '@/components/warehouse/create-warehouse';
import { DeleteWarehouse } from '@/components/warehouse/delete-warehouse';
import { UpdateWarehouse } from '@/components/warehouse/update-warehouse';
import { getUsedCapacity } from '@/lib/services/warehouse.service';
import { Warehouse } from '@/lib/zod.schema';
import { Box, Ellipsis, Search } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { ChangeEvent, useCallback, useEffect, useState } from 'react';

type Props = {
    warehouses: Warehouse[];
    mapLocationControl(longitude: number, latitude: number): void;
};

function MoreOption({ warehouse }: { warehouse: Warehouse }) {
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
                    <UpdateWarehouse warehouse={warehouse} />
                    <DeleteWarehouse storeID={warehouse.id} />
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

function WarehouseItem({
    warehouse,
    mapLocationControl,
}: {
    warehouse: Warehouse;
    mapLocationControl(longitude: number, latitude: number): void;
}) {
    const { push } = useRouter();
    const user = useSession();
    const [used, setUsed] = useState(0);

    useEffect(() => {
        (async () => {
            const [res, err] = await getUsedCapacity(user.token, warehouse.id);
            if (!err) {
                setUsed(res.data.used_capacity);
            }
        })();

        const intervalID = setInterval(async () => {
            const [res, err] = await getUsedCapacity(user.token, warehouse.id);
            if (!err) {
                setUsed(res.data.used_capacity);
            }
        }, 1000 * 60 * 5);

        return () => clearInterval(intervalID);
    }, [warehouse.id, user.token]);

    const handleDoubleClick = useCallback(() => {
        push(`/dashboard/${warehouse.id}`);
    }, [push, warehouse.id]);

    const handleClick = useCallback(() => {
        mapLocationControl(warehouse.location[1], warehouse.location[0]);
    }, [mapLocationControl, warehouse.location]);

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
                                {warehouse.name}
                            </div>
                            <MoreOption warehouse={warehouse}></MoreOption>
                        </div>

                        <Progress
                            value={
                                (used / warehouse.capacity) * 100 > 100
                                    ? 100
                                    : (used / warehouse.capacity) * 100
                            }
                            className="w-full my-1"
                        />
                        <div className="flex">
                            <div className="flex-grow text-sm opacity-80">{`${
                                warehouse.capacity - used
                            } free of ${warehouse.capacity}`}</div>
                        </div>
                    </div>
                </div>
            </div>
        </Button>
    );
}

function WarehouseList({ warehouses, mapLocationControl }: Props) {
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

                <CreateWarehouse />
            </div>

            <div className="relative flex-grow w-full h-full  p-2 max-h-screen">
                <div className="absolute inset-0 overflow-y-auto px-2 py-1.5">
                    {warehouses.map((warehouse) =>
                        warehouse.name
                            .toLocaleLowerCase()
                            .includes(search.trim().toLocaleLowerCase()) ? (
                            <WarehouseItem
                                mapLocationControl={mapLocationControl}
                                key={warehouse.id}
                                warehouse={warehouse}
                            />
                        ) : null
                    )}
                </div>
            </div>
        </div>
    );
}

export default WarehouseList;
