import { defineConfig } from 'tsup'

import { dependencies } from './package.json'
import { getConfig } from './src/lib/functions/tsup'

export default defineConfig(
	getConfig({
		entry: [
			'src/index.ts',
			'src/core/plug.ts',
			'src/core/sdk.ts',
			'src/lib/types/index.ts',
			'src/lib/constants.ts'
		],
		external: [...Object.keys(dependencies)]
	})
)
