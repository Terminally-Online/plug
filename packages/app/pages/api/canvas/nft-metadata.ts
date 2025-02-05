import { type NextApiRequest, type NextApiResponse } from "next"

import { db } from "@/server/db"
import { env } from "@/env"

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
	const { address } = req.query

	if (!address || typeof address !== "string") {
		return res.status(400).json({ error: "Invalid address parameter" })
	}

	try {
		const normalizedAddress = address.toLowerCase()
		console.log("Querying with address:", normalizedAddress)

		const identity = await db.socketIdentity.findFirst({
			where: {
				socketId: {
					equals: normalizedAddress,
					mode: "insensitive"
				}
			}
		})

		console.log("Query result:", identity)

		if (!identity) {
			return res.status(404).json({ error: "User not found" })
		}

		// Return NFT metadata
		const metadata = {
			name: `Plug Founding Ticket #${identity.onboardingCount}`,
			description: "A founding ticket for early Plug users",
			image: `${env.NEXT_PUBLIC_APP_URL}/api/nft/image?color=${identity.onboardingColor?.replace("#", "")}&number=${identity.onboardingCount}`,
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
		}

		res.status(200).json(metadata)
	} catch (error) {
		console.error("Detailed error:", error)
		return res.status(500).json({
			error: "Internal server error",
			details: error instanceof Error ? error.message : String(error)
		})
	}
}
