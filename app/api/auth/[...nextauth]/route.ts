import NextAuth from 'next-auth';
import { AppProviders } from 'next-auth/providers/index';
import CredentialsProvider from 'next-auth/providers/credentials';
import { getCsrfToken } from 'next-auth/react';
import { SiweMessage } from 'siwe';
import type { NextApiRequest, NextApiResponse } from 'next';

const providers = [CredentialsProvider({
    name: "Ethereum",
    credentials: {
        message: { 
            label: "Message", 
            type: "text", 
            placeholder: "0x0" 
        },
        signature: { 
            label: "Signature",
            type: "text",
            placeholder: "0x0"
        }
    },
    async authorize(credentials, req) { 
        try { 
            const siwe = new SiweMessage(JSON.parse(credentials?.message || "{}"));
            const nextAuthUrl = new URL(process.env.NEXTAUTH_URL || "http://localhost:3000/api/auth");
            const result = await siwe.verify({ 
                signature: credentials?.signature || '',
                domain: nextAuthUrl.host,
                nonce: await getCsrfToken({ req: { headers: req.headers } })
            })

            if(result.success) { return { id: siwe.address } }
            return null
        } catch (e) { 
            return null
        }
    }
}) ]

const callbacks = { 
    async session({ session, token }: { session: any, token: any }) {
        session.address = token.sub
        session.user.name = token.sub
        session.user.image = `https://avatar.vercel.sh/${pin.value}.png`
        return session
    }
}

const handler = NextAuth({ 
    providers,
    session: { strategy: "jwt" },
    secret: process.env.NEXTAUTH_SECRET || "secret",
    callbacks
})

export { handler as GET, handler as POST }
