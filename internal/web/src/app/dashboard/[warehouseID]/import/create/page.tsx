import { authz } from '@/auth';
import CreateImInvoiceClientPage from './client-page';
import { handleErr } from '@/lib/response';
import { ErrUnauthorized } from '@/lib/errors';

function CreateImportInvoicePage() {
    const user = authz();
    if (!user) {
        handleErr(ErrUnauthorized);
    }

    return (
        <section className="relative w-full h-full">
            <CreateImInvoiceClientPage />
        </section>
    );
}

export default CreateImportInvoicePage;
