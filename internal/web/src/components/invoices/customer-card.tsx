import { Mail } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { Customer } from '@/lib/zod.schema';
import { cn } from '@/lib/utils';

interface CustomerProps {
    customer: Customer;
    className?: string;
    onSelect?: (id: number) => void;
}

export default function CustomerCard({
    customer: { id, name, email },
    className,
    onSelect = function () {},
}: CustomerProps) {
    return (
        <Card
            className={cn(
                ' w-full py-2 shadow-none rounded cursor-pointer',
                className
            )}
            onClick={() => {
                onSelect(id);
            }}
        >
            <CardHeader className="p-3 py-1">
                <CardTitle className="text-xl font-bold">
                    #{id} {name}
                </CardTitle>
            </CardHeader>
            <CardContent className="px-3 py-1">
                <div className="">
                    <div className="flex items-center space-x-2">
                        <Mail className="h-4 w-4 text-muted-foreground" />
                        <p className="text-sm">{email}</p>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}
