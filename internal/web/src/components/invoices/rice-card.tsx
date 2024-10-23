import { Card, CardContent, CardTitle, CardDescription } from '../ui/card';
import { Rice } from '@/lib/zod.schema';
import { cn } from '@/lib/utils';

interface RiceProps {
    rice: Rice;
    className?: string;
    onSelect?: (rice: Rice) => void;
}

export default function RiceCard({
    rice,
    className,
    onSelect = function () {},
}: RiceProps) {
    return (
        <Card
            onClick={() => onSelect(rice)}
            className={cn(
                'flex w-full min-h-14 p-0 shadow-none rounded overflow-hidden hover:bg-slate-100 cursor-pointer',
                className
            )}
        >
            <CardContent className="p-3 py-1 flex-grow">
                <CardTitle className="text-xl font-medium">
                    {rice.name}
                </CardTitle>
                <CardDescription># {rice.id}</CardDescription>
            </CardContent>
        </Card>
    );
}
