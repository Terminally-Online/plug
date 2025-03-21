import { getMetadataForToken } from "@/lib/opensea/metadata"
import { db } from "@/server"

const work = async () => {
	const collectiblesWithoutMetadata = await db.collectible.findMany({
		where: {
			collectibleMetadata: null
		},
		take: 100,
		include: {
			collection: true
		}
	})

	await Promise.all(
		collectiblesWithoutMetadata.map(async collectible => {
			const token = {
				address: collectible.collectionAddress,
				chain: collectible.collectionChain,
				tokenId: collectible.tokenId
			}

			try {
				await getMetadataForToken(token)
				return { success: true, token }
			} catch (error) {
				return { success: false, token }
			}
		})
	)
}

work()
