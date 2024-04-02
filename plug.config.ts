import { config } from '@nftchance/plug-types'

export default config({
	out: './src/contracts/abstracts/',
	contract: { name: 'PlugTypes', filename: 'Plug.Types', solidity: '0.8.23' }
})
