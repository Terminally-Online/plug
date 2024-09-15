import NextAuth from "next-auth"

import { authOptions } from "@/server/auth"

export default async function auth(req: any, res: any) {
	const providers = authOptions.providers
	const isDefaultSigninPage = req.method === "GET" && req.query.nextauth.includes("signin")

	if (isDefaultSigninPage) providers.pop()

	return await NextAuth(req, res, {
		...authOptions,
		providers
	})
}
