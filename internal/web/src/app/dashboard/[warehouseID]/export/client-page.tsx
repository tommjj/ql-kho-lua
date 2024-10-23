'use client';

import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';
import DatePicker from '@/components/ui/date-time-picker';
import Pagination from '@/components/ui/pagination';
import { Invoice } from '@/lib/zod.schema';
import { Pagination as PaginationType } from '@/types/http';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { useCallback } from 'react';
import { formatDateString } from '@/lib/utils';
import InvoiceView from './invoice-view';

const umberFormatter = new Intl.NumberFormat();

function ImportClientPage({
    pagination,
    listInvoices,
}: {
    pagination: PaginationType;
    listInvoices: Invoice[];
}) {
    const pathname = usePathname();
    const searchParams = useSearchParams();
    const { push, replace } = useRouter();

    const invoiceID = searchParams.get('invoice');

    const handleEndDayChange = useCallback(
        (v: Date | undefined) => {
            const params = new URLSearchParams(searchParams.toString());
            if (v) {
                params.set('end', v.toISOString());
            } else {
                params.delete('end');
            }

            push(`${pathname}?${params.toString()}`);
        },
        [pathname, push, searchParams]
    );

    const handleStartDayChange = useCallback(
        (v: Date | undefined) => {
            const params = new URLSearchParams(searchParams.toString());
            if (v) {
                params.set('start', v.toISOString());
            } else {
                params.delete('start');
            }

            push(`${pathname}?${params.toString()}`);
        },
        [pathname, push, searchParams]
    );

    return (
        <ResizablePanelGroup
            direction="horizontal"
            className="w-full md:min-w-[450px] h-full pt-[3.25rem]"
        >
            <ResizablePanel defaultSize={65} className="h-full">
                <div className="relative size-full">
                    <div className="flex justify-end items-center w-full h-[3.25rem] px-2 gap-2">
                        <div className="flex-grow">
                            <DatePicker
                                className="w-48 mr-2"
                                placeholder="From"
                                onSelected={handleStartDayChange}
                            />
                            <DatePicker
                                className="w-48"
                                placeholder="To"
                                onSelected={handleEndDayChange}
                            />
                        </div>

                        <Pagination className="" pagination={pagination} />
                    </div>
                    <div className="px-2">
                        <Table>
                            <TableHeader>
                                <TableRow>
                                    <TableHead className="w-[80px] text-lg">
                                        ID
                                    </TableHead>
                                    <TableHead className="text-lg">
                                        Processed by
                                    </TableHead>
                                    <TableHead className="text-lg">
                                        Warehouse id
                                    </TableHead>
                                    <TableHead className="text-lg">
                                        Customer id
                                    </TableHead>
                                    <TableHead className="text-lg">
                                        Created at
                                    </TableHead>
                                    <TableHead className="text-lg">
                                        Total price
                                    </TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {listInvoices.map((v) => (
                                    <TableRow
                                        className="cursor-pointer"
                                        key={v.id}
                                        onClick={() => {
                                            const params = new URLSearchParams(
                                                searchParams.toString()
                                            );
                                            if (v) {
                                                params.set(
                                                    'invoice',
                                                    v.id.toString()
                                                );
                                            } else {
                                                params.delete('invoice');
                                            }

                                            replace(
                                                `${pathname}?${params.toString()}`
                                            );
                                        }}
                                    >
                                        <TableCell className="font-medium text-lg">
                                            {v.id}
                                        </TableCell>
                                        <TableCell className="font-medium text-lg">
                                            {v.user_id}
                                        </TableCell>
                                        <TableCell className="font-medium text-lg">
                                            {v.warehouse_id}
                                        </TableCell>
                                        <TableCell className="font-medium text-lg">
                                            {v.customer_id}
                                        </TableCell>
                                        <TableCell className="font-medium text-lg">
                                            {formatDateString(v.created_at)}
                                        </TableCell>
                                        <TableCell className="font-medium text-lg">
                                            {umberFormatter.format(
                                                v.total_price
                                            )}
                                            <span className="text-sm">
                                                {' '}
                                                VND
                                            </span>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </div>
                </div>
            </ResizablePanel>

            {invoiceID ? (
                <>
                    <ResizableHandle withHandle />
                    <ResizablePanel
                        className="relative min-w-[220px] h-full animate-expand"
                        defaultSize={35}
                        minSize={20}
                        maxSize={50}
                    >
                        <InvoiceView invoiceID={invoiceID} />
                    </ResizablePanel>
                </>
            ) : null}
        </ResizablePanelGroup>
    );
}

export default ImportClientPage;
