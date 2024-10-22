'use client';

import NavItem from './nav-items';
import { useParams, usePathname } from 'next/navigation';
import { ClipboardCopy, ClipboardPaste, Home, LucideProps } from 'lucide-react';
import { cn } from '@/lib/utils';
import { ForwardRefExoticComponent, RefAttributes, useMemo } from 'react';

type NavItem = {
    path: string;
    title: string;
    icon: ForwardRefExoticComponent<
        Omit<LucideProps, 'ref'> & RefAttributes<SVGSVGElement>
    >;
};

const storePageNavItems: NavItem[] = [
    {
        path: '',
        title: 'Warehouse',
        icon: Home,
    },
    {
        path: '/import',
        title: 'Import',
        icon: ClipboardCopy,
    },
    {
        path: '/export',
        title: 'Export',
        icon: ClipboardPaste,
    },
];

function getStorePageNavItems(storeID: string): NavItem[] {
    return storePageNavItems.map((item) => ({
        path: `/dashboard/${storeID}${item.path}`,
        title: item.title,
        icon: item.icon,
    }));
}

function WarehousePageNavBar() {
    const { warehouseID } = useParams<{ warehouseID: string }>();
    const pathname = usePathname();
    const navList = useMemo(
        () => getStorePageNavItems(warehouseID),
        [warehouseID]
    );

    return (
        <nav className="px-2 pt-2 grid gap-y-1.5">
            {navList.map((item) => (
                <NavItem
                    key={item.path}
                    active={
                        item.path !== `/dashboard/${warehouseID}`
                            ? pathname.startsWith(item.path)
                            : item.path === pathname
                    }
                    title={item.title}
                    href={item.path}
                    icon={
                        <item.icon
                            className={cn('size-5 mr-2', {
                                'opacity-80': pathname !== item.path,
                            })}
                        />
                    }
                />
            ))}
        </nav>
    );
}

export default WarehousePageNavBar;
