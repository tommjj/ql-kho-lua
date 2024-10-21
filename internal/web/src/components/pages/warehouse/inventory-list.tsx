'use client';

import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';
import { WarehouseItem } from '@/lib/zod.schema';
import { useSearchParams } from 'next/navigation';

function InventoryList({
    warehouseItems,
}: {
    warehouseItems: WarehouseItem[];
}) {
    const searchParams = useSearchParams();

    const query = searchParams.get('q');

    return (
        <div className="relative w-full h-full ">
            <div className="absolute overflow-y-auto inset-0">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-[120px] text-lg">
                                Rice id
                            </TableHead>
                            <TableHead className="text-lg">Name</TableHead>
                            <TableHead className="text-lg">Capacity</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {warehouseItems.map((v) =>
                            query &&
                            !v.rice_name
                                .toLowerCase()
                                .includes(query.toLowerCase()) ? null : (
                                <TableRow key={v.id}>
                                    <TableCell className="font-medium text-lg">
                                        {v.id}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {v.rice_name}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {v.capacity}
                                    </TableCell>
                                </TableRow>
                            )
                        )}
                    </TableBody>
                </Table>
            </div>
        </div>
    );
}

export default InventoryList;
