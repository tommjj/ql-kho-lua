import { authz } from '@/auth';
import { NotFound } from '@/components/pages/not-found';
import { Button } from '@/components/shadcn-ui/button';
import Header from '@/components/ui/header';
import PaginationBar from '@/components/ui/pagination';
import SearchBar from '@/components/ui/search';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListRice } from '@/lib/services/rice.service';
import { Plus } from 'lucide-react';

import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';

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

    const [res, err] = await getListRice(user.token, {
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
            <Header title="Rice">
                <Button>
                    Create rice <Plus className="size-5 ml-1"></Plus>
                </Button>
            </Header>
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
                                <TableHead className="w-[180px] text-lg">
                                    ID
                                </TableHead>
                                <TableHead className="text-lg">Name</TableHead>
                                <TableHead className="text-right"></TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {res.data.map((rice) => (
                                <TableRow key={rice.id}>
                                    <TableCell className="font-medium text-lg">
                                        {rice.id}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {rice.name}
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
