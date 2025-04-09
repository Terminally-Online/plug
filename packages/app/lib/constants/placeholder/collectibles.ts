import { RouterOutputs } from "@/server/client"

type CreatePlaceholderTokenProps = Pick<
	NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number],
	"name"
> & { icon: string }
const createPlaceholderCollection = ({ name, icon }: CreatePlaceholderTokenProps) => {
	return {
		iconUrl: icon.startsWith('http') ? icon : `https://cdn.zerion.io/${icon}.png`,
		name,
		chain: "base",
		collectibles: [{ collectionAddress: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", tokenId: 1 }]
	}
}

const PLACEHOLDER_COOLCATS = createPlaceholderCollection({
	name: "Cool Cats",
	icon: "https://i2.seadn.io/ethereum/aaa51fcc7e44421b8a9f54bab4269fd1/2c31f59ad9d5f4565d33e0351fa8e9/872c31f59ad9d5f4565d33e0351fa8e9.png?h=250&w=250"
})
const PLACEHOLDER_CHONKS = createPlaceholderCollection({
	name: "Chonks",
	icon: "https://i2.seadn.io/base/418db73be72f4b208e054529221f79fc/51f67e1576a7c070996a7d42be80b3/5651f67e1576a7c070996a7d42be80b3.png?h=250&w=250"
})
const PLACEHOLDER_CHROMIE = createPlaceholderCollection({
	name: "Chromie Squiggles",
	icon: "https://i2.seadn.io/ethereum/f171aeee4d65341191a1de574affab53/33026283124b2200c938c7a3395107/f233026283124b2200c938c7a3395107.png?h=250&w=250"
})
const PLACEHOLDER_PENGUIN = createPlaceholderCollection({
	name: "Pudgy Penguins",
	icon: "https://i.seadn.io/s/raw/files/cdf489fb69fd11886b468c0f7ff1376c.png?h=250&w=250"
})
const PLACEHOLDER_MILADY = createPlaceholderCollection({
	name: "Milady",
	icon: "https://i.seadn.io/gae/a_frplnavZA9g4vN3SexO5rrtaBX_cBTaJYcgrPtwQIqPhzgzUendQxiwUdr51CGPE2QyPEa1DHnkW1wLrHAv5DgfC3BP-CWpFq6BA?h=250&w=25"
})

export const PLACEHOLDER_COLLECTIONS = [
	PLACEHOLDER_COOLCATS,
	PLACEHOLDER_CHONKS,
	PLACEHOLDER_CHROMIE,
	PLACEHOLDER_PENGUIN,
	PLACEHOLDER_MILADY,
]
