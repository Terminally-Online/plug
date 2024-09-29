#!/usr/bin/env node
import { Command } from 'commander'

import { schema, zod } from '@/src/lib'

const command = new Command()

command
	.name('plug')
	.description('The Plug toolkit powering full-suite intent management.')

command
	.command('schema')
	.description('Generate the full-suite schema.')
	.option('-c --config <config>', 'Path to config file.')
	.option('-r --root <root>', 'Path to root directory.')
	.action(schema)

command
	.command('zod')
	.description('Generate the accompanying Zod schema.')
	.option('-c --config <config>', 'Path to config file.')
	.option('-r --root <root>', 'Path to root directory.')
	.action(zod)

// Run commander with async/await.
;(async function () {
	await command.parseAsync(process.argv)
})()
