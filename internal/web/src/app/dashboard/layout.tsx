import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import React from 'react';

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <section className="relative w-screen h-screen overflow-x-hidden">
            <ResizablePanelGroup
                direction="horizontal"
                className="w-full rounded-lg border md:min-w-[450px]"
            >
                <ResizablePanel defaultSize={20} minSize={18} maxSize={23}>
                    <div className="">
                        <span className="font-semibold">One</span>
                    </div>
                </ResizablePanel>
                <ResizableHandle />
                <ResizablePanel defaultSize={80}>
                    <div className="size-full">{children}</div>
                </ResizablePanel>
            </ResizablePanelGroup>
        </section>
    );
};

export default DashboardLayout;
