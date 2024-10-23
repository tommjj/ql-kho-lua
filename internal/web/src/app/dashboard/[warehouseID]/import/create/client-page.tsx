'use client';

import CustomerCard from '@/components/invoices/customer-card';
import RiceCard from '@/components/invoices/rice-card';
import { useSession } from '@/components/session-context';
import { Button } from '@/components/shadcn-ui/button';
import { Input } from '@/components/shadcn-ui/input';
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
import {
    CardContent,
    CardHeader,
    CardTitle,
    CardFooter,
} from '@/components/ui/card';
import Header from '@/components/ui/header';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import {
    createImportInvoice,
    CreateImportInvoiceSchema,
} from '@/lib/services/import_invoice.service';
import { cn } from '@/lib/utils';
import { Customer, InvoiceDetail, Rice, Warehouse } from '@/lib/zod.schema';
import { AlertCircle, Leaf, Plus, Search, Users, X } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { ChangeEvent, useCallback, useRef, useState } from 'react';

const umberFormatter = new Intl.NumberFormat();

type Props = {
    customers: Customer[];
    rice: Rice[];
    warehouse: Warehouse;
};

function CreateImInvoiceClientPage({ customers, rice, warehouse }: Props) {
    const { push } = useRouter();
    const user = useSession();
    const lock = useRef(false);

    const [selectedCustomer, setSelectedCustomer] = useState<
        undefined | Customer
    >(undefined);
    const [selectedRice, setSelectedRice] = useState<InvoiceDetail[]>([]);

    const [tabValue, setTabValue] = useState('customer');

    const [error, setError] = useState<undefined | string>();

    const createOnQuantityChange = useCallback(
        (id: number, key: 'quantity' | 'price') => {
            return (e: ChangeEvent<HTMLInputElement>) => {
                const v = Number(e.target.value);
                if (Number.isNaN(v)) {
                    return;
                }
                setSelectedRice((priv) =>
                    priv.map((v) => {
                        if (v.rice_id !== id) return v;

                        return {
                            ...v,
                            [key]: Number(e.target.value),
                        };
                    })
                );
            };
        },
        []
    );

    const createRemove = useCallback((id: number) => {
        return () => {
            setSelectedRice((priv) => priv.filter((v) => v.rice_id !== id));
        };
    }, []);

    const handleSubmit = () => {
        if (lock.current) {
            return;
        }
        lock.current = true;

        if (!selectedCustomer) {
            setError('select customer please!');
            return;
        }

        if (selectedRice.length === 0) {
            setError('select rice please!');
            return;
        }

        const parsed = CreateImportInvoiceSchema.safeParse({
            customer_id: selectedCustomer.id,
            warehouse_id: warehouse.id,
            details: selectedRice.map((v) => ({
                price: v.price,
                quantity: v.quantity,
                rice_id: v.rice_id,
            })),
        });

        if (!parsed.success) {
            console.log('err');
            setError(parsed.error.errors[0].message);
            return;
        }

        (async function () {
            const [res, err] = await createImportInvoice(
                user.token,
                parsed.data
            );
            if (res) {
                push(`/dashboard/${warehouse.id}/import`);
                return;
            }

            if (!(err instanceof Response)) return;
            switch (err.status) {
                case 409:
                    setError('Conflicting data error');
                    break;
                case 400:
                    const data = await err.json();
                    setError(data.messages[0] as string);
                    break;
            }
        })();
    };

    return (
        <ResizablePanelGroup direction="horizontal" className="w-full h-full ">
            <ResizablePanel defaultSize={70} className="flex flex-col h-full">
                <Header title="Create invoice"></Header>
                <div className="w-full flex-grow overflow-y-auto custom-scrollbar">
                    <CardHeader>
                        <CardTitle className="text-lg">Invoice </CardTitle>
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
                    </CardHeader>
                    <CardContent>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                            <div>
                                <h3 className="text-lg font-semibold">
                                    Customer
                                </h3>
                                <p>
                                    {selectedCustomer
                                        ? selectedCustomer.name
                                        : ''}
                                </p>
                                <p>
                                    {selectedCustomer
                                        ? `Customer ID:${selectedCustomer.id}`
                                        : ''}
                                </p>
                            </div>
                            <div>
                                <h3 className="text-lg font-semibold">
                                    Warehouse
                                </h3>
                                <p>{warehouse.name}</p>
                                <p>Warehouse ID: {warehouse.id}</p>
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
                                        <TableHead></TableHead>
                                    </TableRow>
                                </TableHeader>
                                <TableBody>
                                    {selectedRice.map((rice) => (
                                        <TableRow key={rice.rice_id}>
                                            <TableCell>{rice.name}</TableCell>
                                            <TableCell>
                                                <Input
                                                    onChange={createOnQuantityChange(
                                                        rice.rice_id,
                                                        'quantity'
                                                    )}
                                                    className={cn('w-40', {
                                                        'focus-visible:ring-red-700':
                                                            rice.price <= 0,
                                                    })}
                                                    autoFocus
                                                    value={rice.quantity}
                                                ></Input>
                                            </TableCell>
                                            <TableCell>
                                                <Input
                                                    onChange={createOnQuantityChange(
                                                        rice.rice_id,
                                                        'price'
                                                    )}
                                                    className={cn('w-40', {
                                                        'focus-visible:ring-red-700':
                                                            rice.price <= 0,
                                                    })}
                                                    value={rice.price}
                                                ></Input>
                                            </TableCell>
                                            <TableCell>
                                                {umberFormatter.format(
                                                    rice.quantity * rice.price
                                                )}{' '}
                                                <span className="text-[0.5rem]">
                                                    {' '}
                                                    VND
                                                </span>
                                            </TableCell>

                                            <TableCell>
                                                <Button
                                                    onClick={createRemove(
                                                        rice.rice_id
                                                    )}
                                                    className="bg-red-400 px-2.5 hover:bg-red-400/90"
                                                >
                                                    <X className="size-4" />
                                                </Button>
                                            </TableCell>
                                        </TableRow>
                                    ))}
                                </TableBody>
                            </Table>
                        </div>
                        <div className="flex justify-between items-center">
                            <div>
                                <p>Processed by: {user.name}</p>
                                <p>User ID: {user.id}</p>
                            </div>
                            <div className="text-right">
                                <h3 className="text-lg font-semibold">
                                    Total Price
                                </h3>
                                <p className="text-2xl font-bold">
                                    {umberFormatter.format(
                                        selectedRice.reduce(
                                            (p, v) => p + v.price * v.quantity,
                                            0
                                        )
                                    )}
                                    <span className="text-sm">VND</span>
                                </p>
                            </div>
                        </div>
                    </CardContent>
                    <CardFooter className="justify-end">
                        <Button onClick={handleSubmit}>
                            Create invoice <Plus className="size-4 ml-2"></Plus>
                        </Button>
                    </CardFooter>
                </div>
            </ResizablePanel>

            <ResizableHandle withHandle />
            <ResizablePanel
                className="relative min-w-[220px] h-full animate-expand"
                defaultSize={30}
                minSize={20}
                maxSize={50}
            >
                <Tabs
                    defaultValue="customer"
                    className="flex flex-col w-full h-screen"
                    value={tabValue}
                >
                    <div className="flex py-2 px-2 border-b">
                        <div className="relative flex-grow">
                            <Search className="absolute p-2 size-9 opacity-80 -mt-[1px]" />
                            <Input
                                className="w-full pl-9"
                                placeholder="Search..."
                            ></Input>
                        </div>
                        <TabsList className="bg-foreground">
                            <TabsTrigger
                                onClick={() => setTabValue('customer')}
                                value="customer"
                                className="py-2 data-[state=active]:bg-primary data-[state=active]:text-white"
                            >
                                <Users className="size-4" />
                            </TabsTrigger>
                            <TabsTrigger
                                onClick={() => setTabValue('rice')}
                                value="rice"
                                className="py-2 data-[state=active]:bg-primary data-[state=active]:text-white"
                            >
                                <Leaf className="size-4" />
                            </TabsTrigger>
                        </TabsList>
                    </div>

                    <TabsContent
                        className="px-2 overflow-y-auto flex-grow custom-scrollbar animate-left-to-right"
                        value="customer"
                    >
                        {customers.map((customer) => (
                            <CustomerCard
                                key={customer.id}
                                customer={customer}
                                className={cn('mb-1', {
                                    'border-primary':
                                        customer === selectedCustomer,
                                })}
                                onSelect={() => {
                                    setTabValue('rice');
                                    setSelectedCustomer(customer);
                                }}
                            />
                        ))}
                    </TabsContent>
                    <TabsContent
                        className="px-2 overflow-y-auto flex-grow custom-scrollbar animate-right-to-left"
                        value="rice"
                    >
                        {rice.map((r) =>
                            selectedRice.some(
                                (v) => v.rice_id === r.id
                            ) ? null : (
                                <RiceCard
                                    key={r.id}
                                    rice={r}
                                    className={cn('mb-1')}
                                    onSelect={() => {
                                        setSelectedRice((priv) => [
                                            ...priv,
                                            {
                                                rice_id: r.id,
                                                name: r.name,
                                                price: 0,
                                                quantity: 0,
                                            },
                                        ]);
                                    }}
                                />
                            )
                        )}
                    </TabsContent>
                </Tabs>
            </ResizablePanel>
        </ResizablePanelGroup>
    );
}

export default CreateImInvoiceClientPage;
