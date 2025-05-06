type CreatePlaceholderTokenProps = {
	name: string
	icon: string
	values: Array<number>
}

const createPlaceholderPosition = ({ name, icon, values }: CreatePlaceholderTokenProps) => {
	return values.map((value, idx) => {
		const percentChange = Math.round((Math.random() * 20 - 10) * 100) / 100 // -10% to +10%, 2 decimals
		const absoluteChange = Math.round(((value * percentChange) / 100) * 100) / 100
		return {
			id: `placeholder-${name}-${idx}`,
			type: "position",
			attributes: {
				name: name,
				price: 1,
				value: value,
				changes: { absolute_1d: absoluteChange, percent_1d: percentChange },
				position_type: "placeholder",
				quantity: { int: "0", decimals: 18, float: value, numeric: value.toString() },
				fungible_info: {
					name: name,
					symbol: "PH",
					icon: { url: icon.startsWith("http") ? icon : `https://cdn.zerion.io/${icon}.png` },
					flags: { verified: false },
					implementations: [
						{
							chain_id: "base",
							address: icon,
							decimals: 18,
							balance: value,
							value: value,
							percentage: 100
						}
					]
				},
				flags: { displayable: true, is_trash: false },
				updated_at: new Date().toISOString(),
				updated_at_block: null,
				parent: null,
				protocol: null,
				poolAddress: null,
				groupId: null,
				application_metadata: {
					name: name,
					icon: { url: icon.startsWith("http") ? icon : `https://cdn.zerion.io/${icon}.png` },
					url: null
				}
			},
			relationships: {
				chain: {
					links: { related: "" },
					data: { type: "chain", id: "base" }
				},
				fungible: {
					links: { related: "" },
					data: { type: "fungible", id: `placeholder-fungible-${name}` }
				},
				dapp: {
					data: { type: "dapp", id: `placeholder-dapp-${name}` }
				}
			}
		}
	})
}

// Group placeholder positions by protocol name for UI compatibility
const placeholderGroups: Record<string, { name: string; positions: typeof allPositions }> = {}
const allPositions = [
	...createPlaceholderPosition({
		name: "Yearn",
		icon: "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
		values: [9300, 1521, 24.34]
	}),
	...createPlaceholderPosition({
		name: "Aave",
		icon: "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9",
		values: [150, 21050, 350.21]
	}),
	...createPlaceholderPosition({
		name: "Morpho",
		icon: "38f1c334-bbe0-4d99-aef4-e6d0a3a4207b",
		values: [1650.4, 3300.82]
	}),
	...createPlaceholderPosition({
		name: "Euler",
		icon: "0xd9fcd98c322942075a5c3860693e9f4f03aae07b",
		values: [56321.96]
	})
]
allPositions.forEach(pos => {
	const name = pos.attributes.application_metadata?.name || "Placeholder"
	if (!placeholderGroups[name]) {
		placeholderGroups[name] = { name, positions: [] }
	}
	placeholderGroups[name].positions.push(pos)
})

export const PLACEHOLDER_POSITIONS = placeholderGroups
