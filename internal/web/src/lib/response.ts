import { notFound, redirect } from 'next/navigation';
import { ErrDataNotFound, ErrInternal, ErrUnauthorized } from './errors';

/**
 * handleErr is a helper func to handing err \
 * NOTE: just use for server components
 * @param err
 * @returns never
 */
export function handleErr(err: unknown): never {
    if (err instanceof Response) {
        handleHTTPErr(err);
    }
    switch (err) {
        case ErrDataNotFound:
            notFound();
        case ErrUnauthorized:
            redirect('/log-out');
        default:
            throw err;
    }
}

function handleHTTPErr(res: Response): never {
    switch (res.status) {
        case 404:
            notFound();
        case 401:
            redirect('/log-out');
        case 403:
            notFound();
        default:
            throw ErrInternal;
    }
}
