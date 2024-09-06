import NextAuth, { NextAuthOptions } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import { getCsrfToken } from "next-auth/react"

import { SiweMessage } from "siwe"

const authOptions: NextAuthOptions = {
	providers: [
		CredentialsProvider({
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
				const unauthenticatedPairs = ["0x0", "0xdemo"]

				if (
					!credentials ||
					unauthenticatedPairs.includes(credentials.message) ||
					unauthenticatedPairs.includes(credentials.signature)
				) {
					const unixTimestamp = Math.floor(Date.now() / 1000)
					const uuid = crypto.randomUUID()
					const demo = credentials && credentials.message.startsWith("0xdemo")
					const lead = demo ? "demo" : "anonymous"

					return {
						id: `${lead}-${unixTimestamp}-${uuid}`
					}
				}

				try {
					const siwe = new SiweMessage(JSON.parse(credentials?.message || "{}"))
					const nextAuthUrl = new URL(process.env.NEXTAUTH_URL || "")

					const result = await siwe.verify({
						signature: credentials?.signature || "",
						domain: nextAuthUrl.host,
						nonce: await getCsrfToken({
							req: { headers: req.headers }
						})
					})

					if (result.success) return { id: siwe.address }

					return null
				} catch (e) {
					return null
				}
			}
		})
	],
	callbacks: {
		async session({ session, token }: { session: any; token: any }) {
			if (token.sub.startsWith("anonymous") || token.sub.startsWith("demo")) {
				// Create a hot id for the user that is uniquely identifying to the time it was created.
				session.address = token.sub
				session.user = {
					id: "anonymous",
					name: "Anonymous User",
					image: `https://avatar.vercel.sh/anonymous.png`,
					anonymous: true,
					demo: token.sub.startsWith("demo")
				}
			} else if (token.sub) {
				session.address = token.sub
				session.user = {
					id: token.sub,
					name: token.sub,
					image: `https://avatar.vercel.sh/${token.sub}.png`,
					anonymous: false,
					demo: false
				}
			}

			return session
		}
	},
	session: {
		strategy: "jwt"
	},
	secret: process.env.NEXT_AUTH_SECRET
}

// For more information on each option (and a full list of options) go to
// https://next-auth.js.org/configuration/options
export default async function auth(req: any, res: any) {
	const providers = authOptions.providers
	const isDefaultSigninPage = req.method === "GET" && req.query.nextauth.includes("signin")

	if (isDefaultSigninPage) providers.pop()

	return await NextAuth(req, res, {
		...authOptions,
		providers
	})
}
