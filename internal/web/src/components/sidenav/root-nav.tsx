'use client';

import NavItem from './nav-items';
import { usePathname } from 'next/navigation';
import {
    BriefcaseBusiness,
    Home,
    Leaf,
    LucideProps,
    Users,
} from 'lucide-react';
import { cn } from '@/lib/utils';
import { ForwardRefExoticComponent, RefAttributes } from 'react';

type NavItem = {
    path: string;
    title: string;
    icon: ForwardRefExoticComponent<
        Omit<LucideProps, 'ref'> & RefAttributes<SVGSVGElement>
    >;
};

const rootPageNavItems: NavItem[] = [
    {
        path: '/dashboard/root',
        title: 'Storehouses',
        icon: Home,
    },
    {
        path: '/dashboard/root/rice',
        title: 'Rice',
        icon: Leaf,
    },
    {
        path: '/dashboard/root/customers',
        title: 'Customers',
        icon: Users,
    },
    {
        path: '/dashboard/root/root/staff',
        title: 'Staff',
        icon: BriefcaseBusiness,
    },
];

function RootPageNavBar() {
    const pathname = usePathname();
    const navList = rootPageNavItems;

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
                                'opacity-90': pathname !== Item.path,
                            })}
                        />
                    }
                />
            ))}
        </nav>
    );
}

export default RootPageNavBar;
