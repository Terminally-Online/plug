import { contracts } from './src/lib/contracts'
import { defineConfig } from '@wagmi/cli'
import { react } from '@wagmi/cli/plugins'

export default defineConfig({
	out: 'src/lib/hooks.ts',
	// @ts-ignore -- Not sure why this is broken, but not worth figuring out.
	contracts,
	plugins: [react()]
})
