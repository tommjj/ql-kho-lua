import { authz } from '@/auth';
import { NotFound } from '@/components/pages/not-found';
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';
import Header from '@/components/ui/header';
import PaginationBar from '@/components/ui/pagination';
import SearchBar from '@/components/ui/search';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListCustomers } from '@/lib/services/customer.service';
import { Metadata } from 'next/types';

export const metadata: Metadata = {
    title: 'Customers',
};

type Props = {
    searchParams: {
        q?: string;
        page?: string;
    };
};

async function RicePage({ searchParams: { page = '1', q = '' } }: Props) {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    let skip = 1;
    if (page) {
        skip = Number(page);
        if (!Number.isInteger(skip)) {
            skip = 1;
        }
    }

    const [res, err] = await getListCustomers(user.token, {
        limit: 10,
        query: q,
        skip: skip,
    });

    if (err) {
        if (!(err instanceof Response) || err.status !== 404) {
            handleErr(err);
        }
    }

    return (
        <section className="relative w-full h-screen max-h-screen">
            <Header title="Customer">{/* < /> */}</Header>
            <div className="flex py-2 px-3">
                <SearchBar className="flex-grow mr-2 w-full"></SearchBar>
                <PaginationBar
                    pagination={
                        err
                            ? {
                                  current_page: 1,
                                  limit_records: 0,
                                  next_page: null,
                                  prev_page: null,
                                  total_pages: 1,
                                  total_records: 0,
                              }
                            : res.pagination
                    }
                ></PaginationBar>
            </div>
            <div className="px-3 ">
                {err ? (
                    <NotFound />
                ) : (
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead className="w-[120px] text-lg">
                                    ID
                                </TableHead>
                                <TableHead className="text-lg">Name</TableHead>
                                <TableHead className="text-lg">Email</TableHead>
                                <TableHead className=" text-lg">
                                    Phone
                                </TableHead>
                                <TableHead className="text-lg">
                                    Address
                                </TableHead>
                                <TableHead className="text-right"></TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {res.data.map((customer) => (
                                <TableRow key={customer.id}>
                                    <TableCell className="font-medium text-lg">
                                        {customer.id}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {customer.name}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {customer.email}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {customer.phone}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg text-nowrap truncate ">
                                        {customer.address}
                                    </TableCell>
                                    <TableCell className="text-right"></TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                )}
            </div>
        </section>
    );
}

export default RicePage;
