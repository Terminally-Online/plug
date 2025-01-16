import { spawn } from "child_process"

import { env } from "@/env"

const outputString = `   ◍ Plug Anvil Fork Network`

function startAnvilFork() {
	const forkUrl = `https://eth-mainnet.g.alchemy.com/v2/${env.NEXT_PUBLIC_ALCHEMY_KEY}`
	const blockNumber = "19250000" // We can make this configurable later if needed

	const anvil = spawn("anvil", ["--fork-url", forkUrl, "--fork-block-number", blockNumber])

	anvil.stdout.on("data", data => {
		console.log(data.toString())
	})

	anvil.stderr.on("data", data => {
		console.error(data.toString())
	})

	anvil.on("close", code => {
		if (code !== 0) {
			console.error(`Anvil process exited with code ${code}`)
		}
	})

	process.on("SIGINT", () => {
		anvil.kill()
		process.exit()
	})

	process.on("SIGTERM", () => {
		anvil.kill()
		process.exit()
	})

	console.log(`${outputString}
   - URL: http://127.0.0.1:8545
   - Block: ${blockNumber}
   - Chain ID: 1
`)
}

startAnvilFork()
