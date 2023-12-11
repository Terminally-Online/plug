import { defineConfig } from 'tsup'

import { getConfig } from './lib/functions/tsup'
import { dependencies } from './package.json'

export default defineConfig(
	getConfig({
		entry: ['src/index.ts', 'src/core/plug.ts', 'src/core/sdk.ts'],
		external: [...Object.keys(dependencies)]
	})
)
