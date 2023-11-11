import hre from 'hardhat'

import { Plug } from '@/core/framework'

const RUN = false

export default async function () {
	if (!RUN) throw new Error('This test is not meant to be run.')

	const [name, version] = ['PlugMock', '0.0.0']

	const DEBUG_TYPES = {
		Mail: [
			{ name: 'from', type: 'Person' },
			{ name: 'to', type: 'Person' },
			{ name: 'contents', type: 'string' }
		],
		Person: [
			{ name: 'name', type: 'string' },
			{ name: 'wallet', type: 'address' }
		],
		SignedMail: [
			{ name: 'mail', type: 'Mail' },
			{ name: 'signature', type: 'bytes' }
		]
	} as const

	const contract = await hre.viem.deployContract(name)
	const [owner] = await hre.viem.getWalletClients()

	// * Create the util with the debug types.
	const util = new Plug(name, version, 1, DEBUG_TYPES, contract)

	// @ts-expect-error - Should fail because there is no SignedSHOULD_FAIL type.
	await util.sign(owner, 'SHOULD_FAIL', {
		name: 'Bob',
		wallet: owner.account.address
	})

	// * Can sign mail.
	const mail = await util.sign(owner, 'Mail', {
		from: { name: 'Bob', wallet: owner.account.address },
		to: { name: 'Alice', wallet: owner.account.address },
		contents: 'Hello, world!'
	})

	// ! The object that is submit onchain.
	// * Can get the signature and the object that was signed.
	const { intent } = mail
	intent?.mail.from
	intent?.mail

	// @ts-expect-error - Should fail because there is no SignedPerson type.
	await util.sign(owner, 'Person', {
		name: 'Bob',
		wallet: owner.account.address
	})
}
