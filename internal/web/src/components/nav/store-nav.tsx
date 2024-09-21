'use client';

import NavItem from './nav-items';
import { useParams, usePathname } from 'next/navigation';
import { Home, LucideProps, Map } from 'lucide-react';
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
        title: 'Home',
        icon: Home,
    },
    {
        path: '/dashboard/root/map',
        title: 'Map',
        icon: Map,
    },
];

function getStorePageNavItems(storeID: string): NavItem[] {
    return storePageNavItems.map((item) => ({
        path: `/dashboard/${storeID}${item.path}`,
        title: item.title,
        icon: Map,
    }));
}

function StorehousePageNavBar() {
    const { storeID } = useParams<{ storeID: string }>();
    const pathname = usePathname();
    const navList = useMemo(() => getStorePageNavItems(storeID), [storeID]);

    return (
        <nav className="px-2 pt-2 grid gap-y-1.5">
            {navList.map((Item) => (
                <NavItem
                    key={Item.path}
                    active={pathname === Item.path}
                    title={Item.title}
                    href={Item.path}
                    icon={
                        <Item.icon
                            className={cn('size-5 mr-2', {
                                'opacity-80': pathname !== Item.path,
                            })}
                        />
                    }
                />
            ))}
        </nav>
    );
}

export default StorehousePageNavBar;
