import BottomBar from '@/components/sidenav/bottom-bar';
import StorehousePageNavBar from '@/components/sidenav/store-nav';
import StoreSelector from '@/components/sidenav/store-selector';
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import { Separator } from '@/components/shadcn-ui/separator';
import React from 'react';
import { getListStorehouse } from '@/lib/services/storehouse.service';
import { authz } from '@/auth';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';

const DashboardLayout = async ({ children }: { children: React.ReactNode }) => {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const [res, err] = await getListStorehouse(user.token, { limit: 99999 });
    if (err) {
        if (!(err instanceof Response)) {
            handleErr(err);
        }
        if (err.status !== 400) {
            handleErr(err);
        }
    }

    const stores = res?.data;

    return (
        <section className="relative w-screen h-screen overflow-x-hidden">
            <ResizablePanelGroup
                direction="horizontal"
                className="w-full rounded-lg border md:min-w-[450px]"
            >
                <ResizablePanel
                    className="relative min-w-[220px]"
                    defaultSize={18}
                    minSize={16}
                    maxSize={23}
                >
                    <div className="p-2">
                        <StoreSelector storehouses={stores ? stores : []} />
                    </div>
                    <Separator />
                    <div>
                        <StorehousePageNavBar />
                    </div>
                    <BottomBar />
                </ResizablePanel>
                <ResizableHandle />
                <ResizablePanel defaultSize={82}>
                    <div className="size-full">{children}</div>
                </ResizablePanel>
            </ResizablePanelGroup>
        </section>
    );
};

export default DashboardLayout;
