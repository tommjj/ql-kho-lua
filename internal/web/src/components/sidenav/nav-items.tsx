'use client';

import React from 'react';
import { Button } from '../shadcn-ui/button';
import Link from 'next/link';

type Props = {
    href: string;
    active?: boolean;
    title: string;
    icon: JSX.Element;
};

function NavItem({ icon, active = false, href, title }: Props) {
    return (
        <Button
            asChild
            className="w-full justify-start px-3 font-medium text-base"
            size="default"
            variant={active ? 'default' : 'ghost'}
        >
            <Link href={href}>
                {icon}
                {title}
            </Link>
        </Button>
    );
}

export default NavItem;
