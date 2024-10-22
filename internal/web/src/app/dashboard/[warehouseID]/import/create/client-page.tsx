'use client';

import { Input } from '@/components/shadcn-ui/input';
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import Header from '@/components/ui/header';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Leaf, Search, Users } from 'lucide-react';

function CreateImInvoiceClientPage() {
    return (
        <ResizablePanelGroup direction="horizontal" className="w-full h-full ">
            <ResizablePanel defaultSize={70} className="flex flex-col h-full">
                <Header title="Create invoice"></Header>
                <div className="w-full flex-grow overflow-y-auto"></div>
            </ResizablePanel>

            <ResizableHandle withHandle />
            <ResizablePanel
                className="relative min-w-[220px] h-full animate-right-to-left "
                defaultSize={30}
                minSize={20}
                maxSize={50}
            >
                <Tabs defaultValue="customer" className="w-full">
                    <div className="flex py-2 px-2 border-b">
                        <div className="relative flex-grow">
                            <Search className="absolute p-2 size-9 opacity-80 -mt-[1px]" />
                            <Input
                                className="w-full pl-9"
                                placeholder="Search..."
                            ></Input>
                        </div>
                        <TabsList className="bg-foreground">
                            <TabsTrigger
                                value="customer"
                                className="py-2 data-[state=active]:bg-primary data-[state=active]:text-white"
                            >
                                <Users className="size-4" />
                            </TabsTrigger>
                            <TabsTrigger
                                value="rice"
                                className="py-2 data-[state=active]:bg-primary data-[state=active]:text-white"
                            >
                                <Leaf className="size-4" />
                            </TabsTrigger>
                        </TabsList>
                    </div>

                    <TabsContent value="customer">
                        Make changes to your account here.
                    </TabsContent>
                    <TabsContent value="rice">
                        Change your password here.
                    </TabsContent>
                </Tabs>
            </ResizablePanel>
        </ResizablePanelGroup>
    );
}
export default CreateImInvoiceClientPage;
