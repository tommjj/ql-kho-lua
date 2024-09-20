'use client';

import { Settings } from 'lucide-react';
import { Button } from '../shadcn-ui/button';

function SettingButton() {
    return (
        <Button className="px-2" variant="ghost">
            <Settings className="size-5" />
        </Button>
    );
}

export default SettingButton;
