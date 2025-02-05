import { type NextApiRequest, type NextApiResponse } from "next"
import { prisma } from "@/server/db"

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { address } = req.query
  
  if (!address || typeof address !== "string") {
    return res.status(400).json({ error: "Invalid address parameter" })
  }

  try {
    const identity = await prisma.socketIdentity.findFirst({
      where: {
        socket: {
          socketAddress: address.toLowerCase()
        }
      },
      select: {
        onboardingColor: true,
        onboardingCount: true
      }
    })

    if (!identity) {
      return res.status(404).json({ error: "User not found" })
    }

    // Return NFT metadata
    const metadata = {
      name: `Plug Founding Ticket #${identity.onboardingCount}`,
      description: "A founding ticket for early Plug users",
      // Use existing image generation endpoint
      image: `${process.env.NEXT_PUBLIC_APP_URL}/api/canvas/nft?color=${identity.onboardingColor?.replace("#", "")}&number=${identity.onboardingCount}`,
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
    console.error("Error fetching NFT metadata:", error)
    res.status(500).json({ error: "Internal server error" })
  }
}