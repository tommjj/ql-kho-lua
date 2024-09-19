'use client';

import { SessionPayload } from '@/types/jwt-payload';
import { createContext, useContext } from 'react';

export const SessionContext = createContext<SessionPayload | undefined>(
    undefined
);

export const useSession = () => {
    const session = useContext(SessionContext);
    if (!session) throw new Error('session');
    return session;
};

export const SessionProvider = ({
    children,
    user,
}: {
    children: React.ReactNode;
    user?: SessionPayload;
}) => {
    return (
        <SessionContext.Provider value={user}>
            {children}
        </SessionContext.Provider>
    );
};
