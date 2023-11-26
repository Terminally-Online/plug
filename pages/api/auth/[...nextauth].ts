import NextAuth, { NextAuthOptions } from 'next-auth'
import { SiweMessage } from 'siwe'

import CredentialsProvider from 'next-auth/providers/credentials'
import { getCsrfToken } from 'next-auth/react'

// https://next-auth.js.org/configuration/providers/oauth
const authOptions: NextAuthOptions = {
	providers: [
		CredentialsProvider({
			name: 'Ethereum',
			credentials: {
				message: {
					label: 'Message',
					type: 'text',
					placeholder: '0x0'
				},
				signature: {
					label: 'Signature',
					type: 'text',
					placeholder: '0x0'
				}
			},
			async authorize(credentials, req) {
				try {
					const siwe = new SiweMessage(
						JSON.parse(credentials?.message || '{}')
					)
					const nextAuthUrl = new URL(process.env.NEXTAUTH_URL || '')

					const result = await siwe.verify({
						signature: credentials?.signature || '',
						domain: nextAuthUrl.host,
						nonce: await getCsrfToken({
							req: { headers: req.headers }
						})
					})

					if (result.success) {
						return {
							id: siwe.address
						}
					}
					return null
				} catch (e) {
					return null
				}
			}
		})
	],
	callbacks: {
		async session({ session, token }: { session: any; token: any }) {
			session.address = token.sub
			session.user.name = token.sub
			session.user.image = `https://avatar.vercel.sh/${token.sub}.png`
			return session
		}
	},
	session: {
		strategy: 'jwt'
	},
	secret: process.env.NEXTAUTH_SECRET
}

// For more information on each option (and a full list of options) go to
// https://next-auth.js.org/configuration/options
export default async function auth(req: any, res: any) {
	const providers = authOptions.providers

	const isDefaultSigninPage =
		req.method === 'GET' && req.query.nextauth.includes('signin')

	// Hide Sign-In with Ethereum from default sign page
	if (isDefaultSigninPage) {
		providers.pop()
	}

	return await NextAuth(req, res, {
		...authOptions,
		providers
	})
}
