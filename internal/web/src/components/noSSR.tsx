'use client';

import { useEffect, useLayoutEffect, useState } from 'react';

const useEnhancedEffect =
    typeof window !== 'undefined' && process.env.NODE_ENV !== 'test'
        ? useLayoutEffect
        : useEffect;

const NoSSR = ({
    children,
    defer = false,
    fallback = null,
}: {
    children: React.ReactNode;
    defer?: boolean;
    fallback?: React.ReactNode;
}) => {
    const [isMounted, setMountedState] = useState(false);

    useEnhancedEffect(() => {
        if (!defer) setMountedState(true);
    }, [defer]);

    useEffect(() => {
        if (defer) setMountedState(true);
    }, [defer]);

    return isMounted ? children : fallback;
};

export default NoSSR;
