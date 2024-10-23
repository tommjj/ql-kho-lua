'use client';

import { useSession } from '@/components/session-context';
import {
    Card,
    CardContent,
    CardHeader,
    CardTitle,
} from '@/components/shadcn-ui/card';
import { Skeleton } from '@/components/shadcn-ui/skeleton';
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';
import { getImportInvoiceByID } from '@/lib/services/import_invoice.service';
import { Invoice } from '@/lib/zod.schema';
import { useLayoutEffect, useState } from 'react';

const umberFormatter = new Intl.NumberFormat();

function InvoiceView({ invoiceID }: { invoiceID: string | number }) {
    const user = useSession();
    const [invoice, setInvoice] = useState<undefined | Invoice>(undefined);

    useLayoutEffect(() => {
        let numID = 0;
        if (!Number.isInteger(Number(invoiceID))) {
            return;
        }

        numID = Number(invoiceID);

        (async () => {
            const [res, err] = await getImportInvoiceByID(user.token, numID);
            if (err) {
                return;
            }
            setInvoice(res.data);
        })();
    }, [invoiceID, user.token]);

    if (!invoice)
        return <Skeleton className="w-full h-full m-2 rounded"></Skeleton>;

    return (
        <Card className="w-full max-w-4xl mx-auto border-none shadow-none mt-2">
            <CardHeader>
                <CardTitle>Invoice #{invoice.id}</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                    <div>
                        <h3 className="text-lg font-semibold">Customer</h3>
                        <p>{invoice.customer_name}</p>
                        <p>Customer ID: {invoice.customer_id}</p>
                    </div>
                    <div>
                        <h3 className="text-lg font-semibold">Warehouse</h3>
                        <p>{invoice.warehouse_name}</p>
                        <p>Warehouse ID: {invoice.warehouse_id}</p>
                    </div>
                </div>
                <div className="mb-6">
                    <h3 className="text-lg font-semibold mb-2">
                        Invoice Details
                    </h3>
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead>Item</TableHead>
                                <TableHead>Quantity</TableHead>
                                <TableHead>Price</TableHead>
                                <TableHead>Total</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {invoice.details.map((detail) => (
                                <TableRow key={detail.rice_id}>
                                    <TableCell>{detail.name}</TableCell>
                                    <TableCell>{detail.quantity}</TableCell>
                                    <TableCell>
                                        {umberFormatter.format(detail.price)}
                                        <span className="text-[0.5rem]">
                                            {' '}
                                            VND
                                        </span>
                                    </TableCell>
                                    <TableCell>
                                        {umberFormatter.format(
                                            detail.quantity * detail.price
                                        )}{' '}
                                        <span className="text-[0.5rem]">
                                            {' '}
                                            VND
                                        </span>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </div>
                <div className="flex justify-between items-center">
                    <div>
                        <p>Processed by: {invoice.user_name}</p>
                        <p>User ID: {invoice.user_id}</p>
                    </div>
                    <div className="text-right">
                        <h3 className="text-lg font-semibold">Total Price</h3>
                        <p className="text-2xl font-bold">
                            {umberFormatter.format(invoice.total_price)}{' '}
                            <span className="text-sm">VND</span>
                        </p>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}

export default InvoiceView;
