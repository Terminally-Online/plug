import '@nomicfoundation/hardhat-toolbox-viem'
import { HardhatUserConfig } from 'hardhat/config'

import 'tsconfig-paths/register'

const config: HardhatUserConfig = {
	solidity: {
		version: '0.8.19',
		settings: {
			viaIR: true,
			optimizer: {
				enabled: true,
				details: {
					yulDetails: {
						optimizerSteps: 'u'
					}
				}
			}
		}
	},
	paths: {
		sources: './src/contracts',
		tests: './src/test'
	}
}

export default config
