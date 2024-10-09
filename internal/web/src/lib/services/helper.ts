import {
    AUTHORIZATION_HEADER_KEY,
    AUTHORIZATION_TYPE,
} from '../constants/auth.constants';

/**
 * getAuthH is helper get authorization header for http request
 *
 * @param key string
 * @returns [string, string]
 */
export function getAuthH(key: string): [string, string] {
    return [AUTHORIZATION_HEADER_KEY, `${AUTHORIZATION_TYPE} ${key}`];
}
