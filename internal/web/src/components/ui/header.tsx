import { cn } from '@/lib/utils';
import BackButton from './back-button';

type Props = {
    title: string;
    children?: React.ReactNode;
    className?: string;
};

function Header({ title, className, children }: Props) {
    return (
        <header
            className={cn(
                'flex items-center justify-between px-4 border-b',
                className
            )}
        >
            <div className="text-2xl font-semibold text-opacity-90 my-2 h-9">
                <BackButton className="mr-2" />
                {title}
            </div>
            <div>{children}</div>
        </header>
    );
}

export default Header;
