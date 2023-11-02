import { defineConfig } from 'tsup'

import { dependencies } from './package.json'
import { getConfig } from './src/lib/functions/tsup'

export default defineConfig(
	getConfig({
		entry: ['src/index.ts', 'src/core/framework.ts', 'src/core/intent.ts'],
		external: [...Object.keys(dependencies)]
	})
)
