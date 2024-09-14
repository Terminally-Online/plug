import { type GetServerSidePropsContext } from "next"
import { type DefaultSession, getServerSession, type NextAuthOptions } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import { getCsrfToken } from "next-auth/react"

import { SiweMessage } from "siwe"

import { getBaseUrl } from "@/server/client/links"

declare module "next-auth" {
	interface Session extends DefaultSession {
		user: DefaultSession["user"] & {
			id: string
			name: string
			image: string
			anonymous: boolean
			demo: boolean
		}
		address: string
	}
}

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
				if (!credentials || credentials.message === "0x0" || credentials.signature === "0x0") {
					const unixTimestamp = Math.floor(Date.now() / 1000)
					const uuid = crypto.randomUUID()

					return {
						id: `anonymous-${unixTimestamp}-${uuid}`
					}
				}

				try {
					const siwe = new SiweMessage(JSON.parse(credentials.message))
					const nextAuthUrl = new URL(`${getBaseUrl()}/api/auth`)

					const result = await siwe.verify({
						signature: credentials.signature,
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
			} else {
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

export const getServerAuthSession = (ctx: {
	req: GetServerSidePropsContext["req"]
	res: GetServerSidePropsContext["res"]
}) => {
	return getServerSession(ctx.req, ctx.res, authOptions)
}
