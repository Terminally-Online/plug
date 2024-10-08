import { type GetServerSidePropsContext } from "next"
import { type DefaultSession, getServerSession, type NextAuthOptions } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import { getCsrfToken } from "next-auth/react"

import { parseSiweMessage } from "viem/siwe"

import { env } from "@/env"
import { createClient } from "@/lib"

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

// For more information on each option (and a full list of options) go to
// https://next-auth.js.org/configuration/options
export const authOptions: NextAuthOptions = {
	providers: [
		CredentialsProvider({
			name: "Ethereum",
			credentials: {
				message: {
					label: "Message",
					type: "text"
				},
				signature: {
					label: "Signature",
					type: "text"
				},
				chainId: {
					label: "Chain ID",
					type: "text"
				}
			},
			async authorize(credentials, req) {
				if (!credentials?.message || !credentials?.signature || !credentials?.chainId) return null

				console.log("Verifying credentials")

				const unauthenticatedPairs = ["0x0", "0xdemo"]
				if (
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
					const nextAuthUrl = new URL(env.NEXT_AUTH_URL || "")

					console.log("Creating client")
					const client = createClient(Number(credentials.chainId))
					console.log("Created client")
					const valid = await client.verifySiweMessage({
						message: credentials.message,
						signature: credentials.signature as `0x${string}`,
						domain: nextAuthUrl.host,
						nonce: await getCsrfToken({
							req: { headers: req.headers }
						})
					})
					console.log("Verified message")
					if (valid) {
						const address = parseSiweMessage(credentials.message).address as string
						return { id: address }
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
			if (token.sub.startsWith("anonymous") || token.sub.startsWith("demo")) {
				// Create a hot id for the user that is uniquely identifying to the time it was created.
				session.address = token.sub
				session.user = {
					id: token.sub,
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
	secret: env.NEXT_AUTH_SECRET
}

export const getServerAuthSession = (ctx: {
	req: GetServerSidePropsContext["req"]
	res: GetServerSidePropsContext["res"]
}) => {
	return getServerSession(ctx.req, ctx.res, authOptions)
}
