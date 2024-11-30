import axios from "axios"
import fs from "fs"
import path from "path"

type TokenList = {
	tokens: Array<{
		address: string
		name: string
		symbol: string
		logoURI: string
		chainId: number
		decimals: number
	}>
}

async function fetchTokenList(chainId: number): Promise<TokenList["tokens"]> {
	const url = `https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/${chainId}.json`
	const response = await axios.get<TokenList>(url)

	if (response.status !== 200) throw new Error("Could not fetch token list.")

	return response.data.tokens
}

async function generateTokenList() {
	const chainIds = [1, 10, 8453]
	let allTokens: TokenList["tokens"] = []

	for (const chainId of chainIds) {
		const tokens = await fetchTokenList(chainId)
		allTokens = [...allTokens, ...tokens]
	}

	const outputPath = path.join(__dirname, "../constants/tokens.ts")
	const native = 'export const NATIVE_TOKEN_ADDRESS = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"'
	const tokens = `export const TOKENS = ${JSON.stringify(allTokens, null, 4)}`
	const content = [native, tokens].join("\n\n")

	fs.writeFileSync(outputPath, content)
	console.log(`✔︎ Token list saved to ${outputPath}`)
}

generateTokenList().catch(console.error)
