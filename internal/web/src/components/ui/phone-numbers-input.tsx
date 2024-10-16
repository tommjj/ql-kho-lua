'use client';

import { CountryCode, getPhoneCode } from 'libphonenumber-js';
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
import { parseE164 } from '@/lib/validator/e164';

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
                `+${getPhoneCode(countryCode)}${
                    input.startsWith('0') ? input.substring(1) : input
                }`
            );
    };

    useLayoutEffect(() => {
        if (!defaultValue) return;

        const [parsed, ok] = parseE164(defaultValue);
        if (ok) {
            setCountryCode(parsed.countryCode || 'VN');
            setPhoneNumber(parsed.nationalNumber);
        }
    }, [defaultValue]);

    return (
        <div className="w-full  space-y-4">
            <div className="">
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
                                    value={code}
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
                        className="flex-1 w-full"
                    />
                </div>
            </div>
        </div>
    );
}
