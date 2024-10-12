import Image from 'next/image';
import { Storehouse } from '@/lib/zod.schema';
import { Button } from '../shadcn-ui/button';
import { Box } from 'lucide-react';
import Link from 'next/link';
import React from 'react';

const StorehousePopup = ({ storehouse }: { storehouse: Storehouse }) => {
    return (
        <div className="w-[240px] overflow-hidden rounded">
            <div className="flex justify-center w-full h-32 overflow-hidden p-1 ">
                <Image
                    className="max-h-36 w-auto overflow-hidden"
                    width={256}
                    height={256}
                    src={storehouse.image}
                    alt={storehouse.name}
                />
            </div>
            <div className="flex justify-between px-2 pb-2 pt-1">
                <div>
                    <samp>name</samp>
                    <p className="text-lg font-semibold leading-4">
                        {storehouse.name}
                    </p>
                </div>

                <Button className="h-9 " asChild>
                    <Link href={`/dashboard/${storehouse.id}`}>
                        <Box className="size-4 mr-2" /> View
                    </Link>
                </Button>
            </div>
        </div>
    );
};

export default StorehousePopup;
