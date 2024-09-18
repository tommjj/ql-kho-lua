'use client';

import React from 'react';
import { Button } from '../shadcn-ui/button';
import Link from 'next/link';
import { GearIcon } from '@radix-ui/react-icons';

type Props = {
    href: string;
    active: boolean;
};

function NavItem({ active, href }: Props) {
    return (
        <Button
            asChild
            className="w-full justify-start px-3"
            size="lg"
            variant={active ? 'default' : 'secondary'}
        >
            <Link href={href}>
                <GearIcon className="size-5 mr-2" />
                LINK
            </Link>
        </Button>
    );
}

export default NavItem;
