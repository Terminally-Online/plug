import { ZerionPosition } from "@/lib/functions"

type CreatePlaceholderTokenProps = { symbol: string, name: string, balance: number, price: number, icon: string }

const createPlaceholderToken = ({ symbol, name, balance, price = 0, icon }: CreatePlaceholderTokenProps): ZerionPosition => ({
	id: Math.random().toString(),
	type: "positions",
	attributes: {
		parent: null,
		protocol: null,
		name,
		position_type: "wallet",
		quantity: {
			int: "0",
			decimals: 18,
			float: balance,
			numeric: "0"
		},
		value: balance * price,
		price,
		changes: {
			absolute_1d: 0,
			percent_1d: Math.random() * Math.max(10, Math.random() * 10)
		},
		flags: {
			displayable: true,
			is_trash: false
		},
		updated_at: "",
		updated_at_block: 0,
		fungible_info: {
			name,
			symbol,
			icon: { url: icon.startsWith('http') ? icon : `https://cdn.zerion.io/${icon}.png` },
			flags: {
				verified: true,
			},
			implementations: [{
				chain_id: "base",
				address: "0x",
				decimals: 18,
				balance,
				percentage: Math.random() * Math.max(10, Math.random() * 10),
				value: balance * price,
			}]
		}
	},
	relationships: {
		fungible: {
			data: {
				type: "fungibles",
				id: Math.random().toString()
			},
			links: {
				related: ""
			}
		},
		chain: {
			data: {
				type: "chains",
				id: "base"
			},
			links: {
				related: ""
			}
		}
	}

})

const PLACEHOLDER_AERODROME = createPlaceholderToken({ name: "Aerodrome", symbol: "AERO", balance: 302194, price: 0.318, icon: "430f1d3d-9a4b-4a56-b804-896b34843ac0" })
const PLACEHOLDER_MORPHO = createPlaceholderToken({ name: "Morpho", symbol: "MORPHO", balance: 103729, price: 0.38, icon: "38f1c334-bbe0-4d99-aef4-e6d0a3a4207b" })
const PLACEHOLDER_ETH = createPlaceholderToken({ name: "Ethereum", symbol: "ETH", balance: 7.26, price: 1650.42, icon: "eth" })
const PLACEHOLDER_DAI = createPlaceholderToken({ name: "DAI", symbol: "DAI", balance: 4192, price: 1.01, icon: "https://cdn.zerion.io/0x6b175474e89094c44da98b954eedeac495271d0f.png" })
const PLACEHOLDER_WETH = createPlaceholderToken({ name: "Wrapped Ethereum", symbol: "WETH", balance: 1.349, price: 1650.42, icon: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2" })
const PLACEHOLDER_USDC = createPlaceholderToken({ name: "USD Coin", symbol: "USDC", balance: 56.23, price: 0.99, icon: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48" })

export const PLACEHOLDER_TOKENS = [
	PLACEHOLDER_AERODROME,
	PLACEHOLDER_MORPHO,
	PLACEHOLDER_ETH,
	PLACEHOLDER_DAI,
	PLACEHOLDER_WETH,
	PLACEHOLDER_USDC,
]
