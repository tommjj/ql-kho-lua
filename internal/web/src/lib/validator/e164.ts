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
