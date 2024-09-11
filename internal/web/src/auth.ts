import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
// Your own logic for dealing with plaintext password strings; be careful!

export const { handlers, signIn, signOut, auth } = NextAuth({
    pages: {
        signIn: '/log-in',
        newUser: '/sign-up',
    },
    callbacks: {
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
            // You can specify which fields should be submitted, by adding keys to the `credentials` object.
            // e.g. domain, username, password, 2FA token, etc.
            credentials: {
                email: {},
                password: {},
            },
            authorize: async () => {
                const user = {
                    id: 'test',
                    email: 'adw',
                    name: 'a',
                };

                if (!user) {
                    throw new Error('User not found.');
                }
                // return user object with their profile data
                return user;
            },
        }),
    ],
});
