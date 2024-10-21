'use client';

import { cn } from '@/lib/utils';
import { Input } from '../shadcn-ui/input';
import { Search } from 'lucide-react';
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import { useDebouncedCallback } from 'use-debounce';

type Props = {
    className?: string;
    shallow?: boolean;
};

function SearchBar({ className, shallow = false }: Props) {
    const searchParams = useSearchParams();
    const pathname = usePathname();
    const { replace } = useRouter();

    const setSearch = useDebouncedCallback((q: string) => {
        const params = new URLSearchParams(searchParams.toString());
        params.append('q', q);
        params.toString();
        params.delete('page');
        if (q) {
            params.set('q', q);
        } else {
            params.delete('q');
        }

        if (shallow) {
            window.history.replaceState(null, '', `?${params.toString()}`);
        } else {
            replace(`${pathname}?${params.toString()}`);
        }
    }, 300);

    return (
        <div className={cn('relative ', className)}>
            <Search className="absolute p-2 size-9 opacity-80 -mt-[1px]" />
            <Input
                defaultValue={searchParams.get('query')?.toString()}
                onChange={(e) => setSearch(e.target.value)}
                className="w-full pl-9"
                placeholder="Search..."
            ></Input>
        </div>
    );
}

export default SearchBar;
