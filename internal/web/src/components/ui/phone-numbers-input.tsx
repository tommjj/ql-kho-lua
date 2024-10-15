'use client';

import { parsePhoneNumber, CountryCode, getPhoneCode } from 'libphonenumber-js';
import { useLayoutEffect, useState } from 'react';
import { Input } from '@/components/shadcn-ui/input';
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/shadcn-ui/select';
import { Label } from '@/components/shadcn-ui/label';

function parseE164(
    phoneNumber: string
): [{ countryCode: string; nationalNumber: string }, boolean] {
    try {
        const parsedNumber = parsePhoneNumber(phoneNumber);
        return [
            {
                countryCode: parsedNumber.countryCallingCode,
                nationalNumber: parsedNumber.nationalNumber,
            },
            true,
        ];
    } catch (err) {
        return [{ countryCode: '', nationalNumber: '' }, false];
    }
}

const countryCodeList: CountryCode[] = ['VN', 'AC', 'RU', 'LA'];

type Props = {
    defaultCountryCode?: CountryCode;
    defaultValue?: string;
    onChanged?: (value: string) => void;
    onCountryCodeChanged?: (value: string) => void;
    onNumberChanged?: (value: string) => void;
};

export default function PhoneNumberInput({
    defaultValue,
    defaultCountryCode = 'VN',
    onChanged,
    onCountryCodeChanged,
    onNumberChanged,
}: Props) {
    const [countryCode, setCountryCode] =
        useState<CountryCode>(defaultCountryCode);
    const [phoneNumber, setPhoneNumber] = useState('');

    const handleCountryCodeChange = (value: string) => {
        setCountryCode(value as CountryCode);
        onCountryCodeChanged && onCountryCodeChanged(value as CountryCode);
    };

    const handlePhoneNumberChange = (
        e: React.ChangeEvent<HTMLInputElement>
    ) => {
        const input = e.target.value.replace(/\D/g, '');
        setPhoneNumber(input);

        onNumberChanged && onNumberChanged(input);

        onChanged &&
            onChanged(
                `+${countryCode}${
                    input.startsWith('0') ? input.substring(1) : input
                }`
            );
    };

    useLayoutEffect(() => {
        if (!defaultValue) return;

        const [parsed, ok] = parseE164(defaultValue);
        if (ok) {
            setCountryCode(parsed.countryCode as CountryCode);
            setPhoneNumber(parsed.nationalNumber);
        }
    }, [defaultValue]);

    return (
        <div className="w-full max-w-sm space-y-4">
            <div className="space-y-2">
                <Label htmlFor="phone">Phone Number</Label>
                <div className="flex space-x-2">
                    <Select
                        value={countryCode}
                        onValueChange={handleCountryCodeChange}
                    >
                        <SelectTrigger className="w-[100px]">
                            <SelectValue placeholder="Code" />
                        </SelectTrigger>
                        <SelectContent>
                            {countryCodeList.map((code) => (
                                <SelectItem
                                    key={code}
                                    value={getPhoneCode(code)}
                                >{`${getPhoneCode(
                                    code
                                )} (${code})`}</SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                    <Input
                        id="phone"
                        type="tel"
                        placeholder="Phone number"
                        value={phoneNumber}
                        onChange={handlePhoneNumberChange}
                        className="flex-1"
                    />
                </div>
            </div>
            <div className="text-sm text-muted-foreground">
                Selected: {countryCode} {phoneNumber}
            </div>
        </div>
    );
}
