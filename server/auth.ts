import { type GetServerSidePropsContext } from "next"

import {
	type DefaultSession,
	getServerSession,
	type NextAuthOptions
} from "next-auth"
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
				try {
					const siwe = new SiweMessage(
						JSON.parse(credentials?.message || "{}")
					)
					const nextAuthUrl = new URL(`${getBaseUrl()}/api/auth`)

					const result = await siwe.verify({
						signature: credentials?.signature || "",
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
			session.user.id = token.sub
			session.user.name = token.sub
			session.user.image = `https://avatar.vercel.sh/${token.sub}.png`
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
