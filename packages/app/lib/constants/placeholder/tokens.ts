import { RouterOutputs } from "@/server/client"

type CreatePlaceholderTokenProps = Pick<
	NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number],
	"name" | "symbol" | "balance" | "price"
> & { icon: string }
const createPlaceholderToken = ({ symbol, name, balance, price = 0, icon }: CreatePlaceholderTokenProps) => {
	return {
		icon: icon.startsWith('http') ? icon : `https://cdn.zerion.io/${icon}.png`,
		name,
		symbol,
		balance,
		change: Math.random() * Math.max(10, Math.random() * 10),
		value: balance * price,
		price: price,
		verified: false,
		implementations: [{
			contract: icon,
			chain: "base",
			balance: 1234,
			percentage: 100,
			createdAt: new Date(),
			updatedAt: new Date(),
			decimals: 6,
			fungibleName: name,
			fungibleSymbol: symbol,
			balances: [{ balance: 12345 }]
		}]
	}
}

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
