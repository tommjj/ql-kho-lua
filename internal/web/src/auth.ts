import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
import fetcher from './lib/http/fetcher';
import { jwtDecode } from 'jwt-decode';
import { SessionPayload, TokenPayload } from './types/jwt-payload';
import { z } from 'zod';
import { Role } from './types/role';

const { handlers, signIn, signOut, auth } = NextAuth({
    trustHost: true,
    pages: {
        signIn: '/log-in',
        newUser: '/sign-up',
    },
    callbacks: {
        jwt: async ({ token, user }) => {
            if (user) {
                token.id = user.id;
                token.token = user.token;
                token.role = user.role;
            }
            return token;
        },
        session: async ({ session, token }) => {
            if (session.user) {
                session.user.id = token.sub || '';
                session.user.token = token.token as string;
                session.user.role = token.role as string;
            }
            return session;
        },
        authorized({ auth, request: { nextUrl } }) {
            const user = auth?.user;
            const isLoggedIn = !!user;
            const isOnDashboard = nextUrl.pathname.startsWith('/dashboard');
            const isOnRoot = nextUrl.pathname.startsWith('/dashboard/root');
            const isOnHomePage = nextUrl.pathname === '/';

            if (isLoggedIn) {
                const isOnLogout = nextUrl.pathname.startsWith('/log-out');
                if (isOnLogout) {
                    return true;
                }

                const dec = jwtDecode(user.token);
                if (dec.exp && Date.now() >= dec.exp * 1000) {
                    return false;
                }
            }

            if (isOnHomePage) {
                if (isLoggedIn) {
                    return Response.redirect(new URL('/dashboard', nextUrl));
                }
                return true;
            }

            if (isOnRoot) {
                if (isLoggedIn) {
                    if (user.role === Role.ROOT) {
                        return true;
                    }
                    return Response.redirect(new URL('/dashboard', nextUrl));
                }
                return false;
            }

            if (isOnDashboard) {
                if (isLoggedIn) {
                    return true;
                }
                return false;
            } else if (isLoggedIn) {
                return Response.redirect(new URL('/dashboard', nextUrl));
            }
            return true;
        },
    },
    providers: [
        Credentials({
            credentials: {
                email: {},
                password: {},
            },

            authorize: async (credentials) => {
                const parsedCredentials = z
                    .object({
                        email: z.string().email(),
                        password: z.string().min(8),
                    })
                    .safeParse(credentials);

                if (parsedCredentials.error) {
                    return null;
                }

                const [res, err] = await fetcher.post.json<{
                    data: { token: string };
                }>('/auth/login', {
                    email: parsedCredentials.data.email,
                    password: parsedCredentials.data.password,
                });
                if (err) {
                    return null;
                }

                const data = res.data;

                try {
                    const dec = jwtDecode<TokenPayload>(data.token);

                    return {
                        id: dec.id.toString(),
                        name: dec.name,
                        email: dec.email,
                        role: dec.role,
                        token: data.token as string,
                    };
                } catch (err) {}

                return null;
            },
        }),
    ],
});

/**
 * authz is a helper func convert return of auth func
 * @returns Promise<SessionPayload | undefined>
 */
async function authz(): Promise<SessionPayload | undefined> {
    const s = await auth();

    if (!s?.user) {
        return undefined;
    }
    return {
        id: Number(s.user.id),
        name: s.user.name,
        email: s.user.email,
        role: s.user.role as Role,
        token: s.user.token,
    };
}

export { handlers, signIn, signOut, auth, authz };
