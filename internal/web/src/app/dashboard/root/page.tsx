import { authz } from '@/auth';
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from '@/components/shadcn-ui/resizable';
import { ErrUnauthorized } from '@/lib/errors';
import { handleErr } from '@/lib/response';
import { getListStorehouse } from '@/lib/services/storehouse.service';

async function RootPage() {
    const user = await authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    const [store, err] = await getListStorehouse(user.token, { limit: 99999 });
    if (err) {
        if (!(err instanceof Response)) {
            handleErr(err);
        }
        if (err.status !== 400) {
            handleErr(err);
        }
    }

    return (
        <section className="flex w-full h-screen">
            <ResizablePanelGroup
                direction="horizontal"
                className="w-full rounded-lg border md:min-w-[450px]"
            >
                <ResizablePanel
                    className="relative min-w-[220px]"
                    defaultSize={70}
                    minSize={50}
                    maxSize={80}
                >
                    <div>{JSON.stringify(store)}</div>
                </ResizablePanel>
                <ResizableHandle withHandle />
                <ResizablePanel defaultSize={30}>
                    <div className="size-full">{}</div>
                </ResizablePanel>
            </ResizablePanelGroup>
        </section>
    );
}

export default RootPage;
