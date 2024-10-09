export const API_HOST =
    process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

type ReturnType<T> = [undefined, Error | Response] | [T, undefined];

type Path = string | URL | Request;

function createMethods(dfPath: Path = '', init: RequestInit = {}) {
    const method = async function <T>(
        path: string = '',
        requestInit: RequestInit = {}
    ): Promise<ReturnType<T>> {
        try {
            const res = await fetch(`${dfPath}${path}`, {
                ...init,
                ...requestInit,
            });

            if (res.ok) {
                return [(await res.json()) as T, undefined];
            }

            return [undefined, res];
        } catch (error) {
            return [undefined, error as Error];
        }
    };

    method.json = async function <T>(
        path: Path = '',
        body: object,
        requestInit: RequestInit = init
    ): Promise<Promise<ReturnType<T>>> {
        try {
            const res = await fetch(`${dfPath}${path}`, {
                ...init,
                ...requestInit,
                headers: {
                    ...init.headers,
                    ...requestInit.headers,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            });

            if (res.ok) {
                return [(await res.json()) as T, undefined];
            }

            return [undefined, res];
        } catch (error) {
            return [undefined, error as Error];
        }
    };

    method.formData = async function <T>(
        path: Path = '',
        body: FormData,
        requestInit: RequestInit = init
    ): Promise<Promise<ReturnType<T>>> {
        try {
            const res = await fetch(`${dfPath}${path}`, {
                ...init,
                ...requestInit,
                headers: {
                    ...init.headers,
                    ...requestInit.headers,
                    'Content-Type': 'multipart/form-data',
                },
                body: body,
            });

            if (res.ok) {
                return [(await res.json()) as T, undefined];
            }

            return [undefined, res];
        } catch (error) {
            return [undefined, error as Error];
        }
    };

    method.set = (key: string, value: string) =>
        new Fetcher(dfPath, {
            ...init,
            headers: {
                ...init.headers,
                [key]: value,
            },
        });

    return method;
}

class Fetcher {
    get;
    post;
    delete;
    put;
    patch;
    set;
    constructor(input: string | URL | Request = '', init: RequestInit = {}) {
        this.get = createMethods(input, { ...init, method: 'GET' });
        this.post = createMethods(input, { ...init, method: 'POST' });
        this.put = createMethods(input, { ...init, method: 'PUT' });
        this.patch = createMethods(input, { ...init, method: 'PATCH' });
        this.delete = createMethods(input, { ...init, method: 'DELETE' });
        this.set = (key: string, value: string) =>
            new Fetcher(input, {
                ...init,
                headers: {
                    ...init.headers,
                    [key]: value,
                },
            });
    }
}

const fetcher = new Fetcher(API_HOST);

export default fetcher;
