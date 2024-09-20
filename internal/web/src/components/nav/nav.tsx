'use client';

import NavItem from './nav-items';
import { usePathname } from 'next/navigation';
import { Home, LucideProps, Map } from 'lucide-react';
import { cn } from '@/lib/utils';
import { ForwardRefExoticComponent, RefAttributes } from 'react';

type NavItem = {
    path: string;
    title: string;
    icon: ForwardRefExoticComponent<
        Omit<LucideProps, 'ref'> & RefAttributes<SVGSVGElement>
    >;
};

const navList: NavItem[] = [
    {
        path: '/dashboard',
        title: 'Home',
        icon: Home,
    },
    {
        path: '/dashboard/map',
        title: 'Map',
        icon: Map,
    },
];

function NavBar() {
    const pathname = usePathname();

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

export default NavBar;
