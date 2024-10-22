'use client';

import { useState } from 'react';
import { CalendarIcon } from 'lucide-react';
import { format } from 'date-fns';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Calendar } from '@/components/ui/calendar';
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover';

export default function DatePicker({
    className,
    placeholder,
    onSelected,
}: {
    className?: string;
    placeholder?: string;
    onSelected?: (date: Date | undefined) => void;
}) {
    const [isOpen, setIsOpen] = useState(false);
    const [date, setDate] = useState<Date>();

    const formatDate = (date: Date | undefined) => {
        if (!date) return placeholder || 'Pick a date';
        return format(date, 'PPP');
    };

    return (
        <Popover open={isOpen}>
            <PopoverTrigger asChild onClick={() => setIsOpen(true)}>
                <Button
                    variant={'outline'}
                    className={cn(
                        'w-[240px] justify-start text-left font-normal',
                        !date && 'text-muted-foreground',
                        className
                    )}
                >
                    <CalendarIcon className="mr-2 h-4 w-4 mb-0.5" />
                    {formatDate(date)}
                </Button>
            </PopoverTrigger>
            <PopoverContent
                onInteractOutside={() => setIsOpen(false)}
                className="w-auto p-0"
            >
                <Calendar
                    mode="single"
                    selected={date}
                    onSelect={(v) => {
                        setDate(v);
                        onSelected && onSelected(v);
                        setIsOpen(false);
                    }}
                    initialFocus
                />
            </PopoverContent>
        </Popover>
    );
}
