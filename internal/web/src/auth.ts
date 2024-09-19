import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
import fetcher from './lib/http/fetcher';
import { jwtDecode } from 'jwt-decode';
import { TokenPayload } from './types/jwt-payload';

export const { handlers, signIn, signOut, auth } = NextAuth({
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
            const isLoggedIn = !!auth?.user;

            const isOnDashboard = nextUrl.pathname.startsWith('/dashboard');
            const isOnHomePage = nextUrl.pathname === '/';

            if (isOnHomePage) {
                if (isLoggedIn) {
                    return Response.redirect(new URL('/dashboard', nextUrl));
                }
                return true;
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
                const [res, err] = await fetcher.post.json(
                    '/v1/api/auth/login',
                    {
                        email: credentials.email,
                        password: credentials.password,
                    }
                );
                if (!res?.ok || err) {
                    return null;
                }

                const data = (await res.json()).data;

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
