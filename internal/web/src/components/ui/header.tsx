import { cn } from '@/lib/utils';
import { Button } from '../shadcn-ui/button';
import Link from 'next/link';
import { ChevronLeft } from 'lucide-react';

type Props = {
    title: string;
    children?: React.ReactNode;
    className?: string;
    backLink?: string;
};

function Header({ title, className, children, backLink }: Props) {
    return (
        <header
            className={cn(
                'flex items-center justify-between px-4 border-b',
                className
            )}
        >
            <div className="text-2xl font-semibold text-opacity-90 my-2 h-9">
                {backLink ? (
                    <Button className="mr-2 px-2 max-h-9" variant={'ghost'}>
                        <Link href={backLink}>
                            <ChevronLeft className="size-5" />
                        </Link>
                    </Button>
                ) : null}
                {title}
            </div>
            <div>{children}</div>
        </header>
    );
}

export default Header;
