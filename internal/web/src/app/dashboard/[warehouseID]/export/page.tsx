import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import DatePicker from '@/components/ui/date-time-picker';
import Header from '@/components/ui/header';

function ExportPage() {
    return (
        <section className="relative w-full h-screen">
            <Header
                className="absolute top-0 left-0 right-0"
                title="Export"
            ></Header>
            <ResizablePanelGroup
                direction="horizontal"
                className="w-full md:min-w-[450px] h-full pt-[3.25rem]"
            >
                <ResizablePanel defaultSize={65} className="h-full">
                    <div>
                        <DatePicker placeholder="Start day" />
                    </div>
                </ResizablePanel>

                <ResizableHandle withHandle />

                <ResizablePanel
                    className="relative min-w-[220px] h-full"
                    defaultSize={35}
                    minSize={20}
                    maxSize={50}
                >
                    <div></div>
                </ResizablePanel>
            </ResizablePanelGroup>
        </section>
    );
}

export default ExportPage;
