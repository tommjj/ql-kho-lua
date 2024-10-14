'use client';

import { useSearchParams } from 'next/navigation';

import {
    Pagination,
    PaginationContent,
    PaginationEllipsis,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
} from '@/components/shadcn-ui/pagination';
import { cn } from '@/lib/utils';

import { Pagination as PaginationType } from '@/types/http';
import { useMemo } from 'react';

type Props = {
    pagination: PaginationType;
    className?: string;
};

function PaginationBar({ className, pagination }: Props) {
    const searchParams = useSearchParams();

    const paramsStr = useMemo(() => {
        const params = new URLSearchParams(searchParams.toString());
        if (params.has('page')) {
            params.delete('page');
        }

        return params.toString();
    }, [searchParams]);

    return (
        <Pagination className={cn('w-fit', className)}>
            <PaginationContent>
                <PaginationItem>
                    <PaginationPrevious
                        className="aria-disabled:cursor-not-allowed aria-disabled:opacity-70 "
                        {...(pagination.prev_page
                            ? {
                                  href: `?page=${pagination.prev_page}&${paramsStr}`,
                              }
                            : { 'aria-disabled': true })}
                    />
                </PaginationItem>
                {pagination.prev_page ? (
                    <PaginationItem>
                        <PaginationLink href={`?page=1&${paramsStr}`}>
                            {1}
                        </PaginationLink>
                    </PaginationItem>
                ) : null}
                <PaginationItem>
                    <PaginationLink
                        className="bg-primary hover:bg-primary/90 text-white hover:text-white"
                        href={`?page=${pagination.current_page}&${paramsStr}`}
                    >
                        {pagination.current_page}
                    </PaginationLink>
                </PaginationItem>
                {pagination.next_page ? (
                    <PaginationItem>
                        <PaginationLink
                            href={`?page=${pagination.next_page}&${paramsStr}`}
                        >
                            {pagination.next_page}
                        </PaginationLink>
                    </PaginationItem>
                ) : null}
                {pagination.current_page + 1 < pagination.total_pages ? (
                    <PaginationItem>
                        <PaginationLink
                            href={`?page=${
                                pagination.current_page + 2
                            }&${paramsStr}`}
                        >
                            {pagination.current_page + 2}
                        </PaginationLink>
                    </PaginationItem>
                ) : null}

                {pagination.current_page + 3 < pagination.total_pages ? (
                    <PaginationItem>
                        <PaginationItem>
                            <PaginationEllipsis />
                        </PaginationItem>
                    </PaginationItem>
                ) : null}
                {pagination.current_page + 2 < pagination.total_pages ? (
                    <PaginationItem>
                        <PaginationLink
                            href={`?page=${pagination.total_pages}&${paramsStr}`}
                        >
                            {pagination.total_pages}
                        </PaginationLink>
                    </PaginationItem>
                ) : null}
                <PaginationItem>
                    <PaginationNext
                        className="aria-disabled:cursor-not-allowed aria-disabled:opacity-70"
                        {...(pagination.next_page
                            ? {
                                  href: `?page=${pagination.next_page}&${paramsStr}`,
                              }
                            : { 'aria-disabled': true })}
                    />
                </PaginationItem>
            </PaginationContent>
        </Pagination>
    );
}

export default PaginationBar;
