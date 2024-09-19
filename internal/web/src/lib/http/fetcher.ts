export const API_HOST =
    process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const createMethods = (
    dfPath: string | URL | Request = '',
    init: RequestInit = {}
) => {
    const method = async (
        path: string = '',
        requestInit: RequestInit = {}
    ): Promise<[undefined, unknown] | [Response, undefined]> => {
        try {
            const res = await fetch(`${dfPath}${path}`, {
                ...init,
                ...requestInit,
            });
            return [res, undefined];
        } catch (error) {
            return [undefined, error];
        }
    };

    method.json = async (
        path: string = '',
        body: object,
        requestInit: RequestInit = init
    ): Promise<[undefined, unknown] | [Response, undefined]> => {
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
            return [res, undefined];
        } catch (error) {
            return [undefined, error];
        }
    };

    method.formData = async (
        path: string = '',
        body: FormData,
        requestInit: RequestInit = init
    ): Promise<[undefined, unknown] | [Response, undefined]> => {
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
            return [res, undefined];
        } catch (error) {
            return [undefined, error];
        }
    };
    return method;
};

class Fetcher {
    get;
    post;
    delete;
    put;
    patch;
    constructor(input: string | URL | Request = '', init: RequestInit = {}) {
        this.get = createMethods(input, { ...init, method: 'GET' });
        this.post = createMethods(input, { ...init, method: 'POST' });
        this.put = createMethods(input, { ...init, method: 'PUT' });
        this.patch = createMethods(input, { ...init, method: 'PATCH' });
        this.delete = createMethods(input, { ...init, method: 'DELETE' });
    }
}

const fetcher = new Fetcher(API_HOST);

export default fetcher;
