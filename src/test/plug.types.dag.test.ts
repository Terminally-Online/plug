import { createWalletClient, getContract, http } from 'viem'

import { PlugSDK } from '@/src/core/sdk'

import { mainnet } from 'viem/chains'

const RUN = false

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
	LiveMail: [
		{ name: 'mail', type: 'Mail' },
		{ name: 'signature', type: 'bytes' }
	]
} as const

export default async function () {
	if (!RUN) throw new Error('This test is not meant to be run.')

	const verifyingContract = getContract({ address: '0x0', abi: [] })
	const owner = createWalletClient({
		chain: mainnet,
		transport: http()
	})

	// * Create the util with the debug types.
	const domain = {
		name: 'PlugMock',
		version: '0.0.0',
		chainId: 1
	}
	const util = new PlugSDK(domain, verifyingContract, DEBUG_TYPES)

	// @ts-expect-error - Should fail because there is no LiveSHOULD_FAIL type.
	await util.sign(owner, 'SHOULD_FAIL', {
		name: 'Bob',
		wallet: '0x0'
	})

	const plug = util.build('Mail', {
		from: { name: 'Bob', wallet: '0x0' },
		to: { name: 'Alice', wallet: '0x0' },
		contents: 'Hello, world!'
	})

	// * Can sign mail.
	const mail = await util.sign(owner, plug)

	// ! The object that is submit onchain.
	// * Can get the signature and the object that was signed.
	const { intent } = mail
	intent?.mail.from
	intent?.mail

	// @ts-expect-error - Should fail because there is no LivePerson type.
	await util.sign(owner, 'Person', {
		name: 'Bob',
		wallet: '0x0'
	})
}
