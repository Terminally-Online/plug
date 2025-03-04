import { ensureDir, writeFile } from 'fs-extra'
import { dirname, resolve } from 'pathe'

import { configs, generateZodSchema } from '@/src/lib'

export const zod = async () => {
	const cwd = process.cwd()
	const configurations = await configs()

	for await (const config of configurations) {
		if (!config.out.zod) continue

		const { schemas } = await generateZodSchema(config)

		const zodPath = resolve(cwd, config.out.zod, 'types.ts')

		await ensureDir(dirname(zodPath))
		await writeFile(zodPath, schemas)
	}

	process.exit()
}
