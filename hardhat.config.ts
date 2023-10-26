import '@nomicfoundation/hardhat-toolbox-viem'
import { HardhatUserConfig } from 'hardhat/config'

import 'tsconfig-paths/register'

const config: HardhatUserConfig = {
	solidity: '0.8.19',
	paths: {
		sources: './src/contracts',
		tests: './src/test'
	}
}

export default config
