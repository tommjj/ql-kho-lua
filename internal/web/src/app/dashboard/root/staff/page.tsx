import { authz } from '@/auth';
import { NotFound } from '@/components/pages/not-found';
import { Button } from '@/components/shadcn-ui/button';
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/shadcn-ui/table';
import CreateStaffModal from '@/components/staff/create-staff';
import { DeleteStaff } from '@/components/staff/delete-staff';
import UpdateStaffModal from '@/components/staff/update-staff';
import Header from '@/components/ui/header';
import PaginationBar from '@/components/ui/pagination';
import SearchBar from '@/components/ui/search';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListUser } from '@/lib/services/user.service';
import { Role } from '@/types/role';
import { Metadata } from 'next/types';

export const metadata: Metadata = {
    title: 'Staffs',
};

type Props = {
    searchParams: {
        q?: string;
        page?: string;
    };
};

async function StaffsPage({ searchParams: { page = '1', q = '' } }: Props) {
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

    const [res, err] = await getListUser(user.token, {
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
            <Header title="Staffs">
                <CreateStaffModal />
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
                                <TableHead className="w-[120px] text-lg">
                                    ID
                                </TableHead>
                                <TableHead className="text-lg">Name</TableHead>
                                <TableHead className="text-lg">Email</TableHead>
                                <TableHead className=" text-lg">
                                    Phone
                                </TableHead>
                                <TableHead className="text-lg">Role</TableHead>
                                <TableHead className="text-right"></TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {res.data.map((staff) => (
                                <TableRow key={staff.id}>
                                    <TableCell className="font-medium text-lg">
                                        {staff.id}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {staff.name}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {staff.email}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg">
                                        {staff.phone}
                                    </TableCell>
                                    <TableCell className="font-medium text-lg text-nowrap truncate ">
                                        {staff.role}
                                    </TableCell>
                                    <TableCell className="text-right">
                                        <UpdateStaffModal staff={staff} />
                                        {staff.role === Role.MEMBER ? (
                                            <DeleteStaff staff={staff} />
                                        ) : (
                                            <Button className="mr-2 px-2 opacity-0">
                                                <div className="size-4"></div>
                                            </Button>
                                        )}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                )}
            </div>
        </section>
    );
}

export default StaffsPage;
