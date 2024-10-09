export type ResponseOrError<T> = [undefined, Error | Response] | [T, undefined];
