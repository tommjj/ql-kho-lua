'use client';

import { cn } from '@/lib/utils';
import { Button } from '../shadcn-ui/button';
import { useRouter } from 'next/navigation';
import { useCallback } from 'react';
import { ChevronLeft } from 'lucide-react';

export default function BackButton({ className }: { className?: string }) {
    const { back } = useRouter();

    const handleClick = useCallback(() => {
        back();
    }, [back]);

    return (
        <Button
            className={cn('px-2', className)}
            variant="ghost"
            onClick={handleClick}
        >
            <ChevronLeft className="size-5" />
        </Button>
    );
}
