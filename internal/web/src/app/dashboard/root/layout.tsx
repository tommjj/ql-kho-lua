import BottomBar from '@/components/sidenav/bottom-bar';
import RootPageNavBar from '@/components/sidenav/root-nav';
import StoreSelector from '@/components/sidenav/store-selector';
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import { Separator } from '@/components/shadcn-ui/separator';
import React from 'react';

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
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
                        <StoreSelector
                            storehouses={[{ id: 1, name: 'store 01' }]}
                        />
                    </div>
                    <Separator />
                    <div>
                        <RootPageNavBar />
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
