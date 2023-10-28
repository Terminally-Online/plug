import { defineConfig } from 'tsup'

import packageJson from './package.json'

export default defineConfig({
	entryPoints: [
		'src/index.ts',
		'src/core/framework.ts',
		'src/core/intent.ts'
	],
	outDir: 'dist',
	dts: true,
	sourcemap: true,
	clean: true,
	format: ['cjs', 'esm'],
	minify: true,
	minifyWhitespace: true,
	external: [...Object.keys(packageJson.dependencies || {})]
})
