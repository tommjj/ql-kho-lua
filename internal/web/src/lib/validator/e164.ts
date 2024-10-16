import { CountryCode, parsePhoneNumber } from 'libphonenumber-js';

/**
 * validateE164 is a validator for e164
 *
 * @param phoneNumber
 * @returns bool
 */
export function validateE164(phoneNumber: string) {
    const regEx = /^\+[1-9]\d{10,14}$/;

    return regEx.test(phoneNumber);
}

export function parseE164(phoneNumber: string): [
    {
        countryCode: CountryCode | undefined;
        nationalNumber: string;
        countryCallingCode: string;
    },
    boolean
] {
    try {
        const parsedNumber = parsePhoneNumber(phoneNumber);
        return [
            {
                countryCallingCode: parsedNumber.countryCallingCode,
                nationalNumber: parsedNumber.nationalNumber,
                countryCode: parsedNumber.country,
            },
            true,
        ];
    } catch (err) {
        return [
            {
                countryCode: undefined,
                nationalNumber: '',
                countryCallingCode: '',
            },
            false,
        ];
    }
}
