'use client';

import { Role } from '@/types/role';
import { useSession } from '../session-context';
import SettingButton from '../setting/setting-buttom';
import { Button } from '../shadcn-ui/button';
import { UsersRound, UserRoundCog, LogOut } from 'lucide-react';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from '../shadcn-ui/dropdown-menu';
import logout from '@/lib/actions/logout';

function BottomBar() {
    const user = useSession();

    return (
        <div className="absolute bottom-0 left-0 flex border-t w-full p-2 gap-x-1.5">
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button
                        className="flex-grow justify-start font-medium text-base"
                        variant="ghost"
                    >
                        {user.role === Role.ROOT ? (
                            <UserRoundCog className="size-5 mr-2" />
                        ) : (
                            <UsersRound className="size-5 mr-2" />
                        )}
                        {user.name}
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-56">
                    <DropdownMenuGroup>
                        <DropdownMenuItem asChild>
                            <form className="px-0 py-0" action={logout}>
                                <Button
                                    type="submit"
                                    className="flex-grow justify-start font-medium text-base px-2"
                                    variant="ghost"
                                >
                                    <LogOut className="size-5 mr-2" />
                                    sign out
                                </Button>
                            </form>
                        </DropdownMenuItem>
                    </DropdownMenuGroup>
                </DropdownMenuContent>
            </DropdownMenu>

            <SettingButton />
        </div>
    );
}

export default BottomBar;
