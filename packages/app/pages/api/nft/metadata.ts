import { type NextApiRequest, type NextApiResponse } from "next"

import { db } from "@/server/db"

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
	const { address } = req.query

	if (!address || typeof address !== "string") {
		return res.status(400).json({ error: "Invalid address parameter" })
	}

	const identity = await db.socketIdentity.findFirst({
		where: {
			socketId: {
				equals: address.toLowerCase(),
				mode: "insensitive"
			}
		}
	})
	if (!identity) return res.status(404).json({ error: "User not found" })

	res.status(200).json({
		name: `Plug Founding Ticket #${identity.onboardingCount}`,
		description: "A founding ticket for early Plug users",
		image: `${process.env.NEXT_PUBLIC_APP_URL}/api/nft/image?color=${identity.onboardingColor?.replace("#", "")}&number=${identity.onboardingCount}`,
		attributes: [
			{
				trait_type: "Ticket Number",
				value: identity.onboardingCount
			},
			{
				trait_type: "Color",
				value: identity.onboardingColor
			}
		]
	})
}
