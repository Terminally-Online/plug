import { createWalletClient, getContract, http } from 'viem'

import { PlugSDK } from '@/src/core/sdk'

import { mainnet } from 'viem/chains'

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
		Email: [
			{ name: 'from', type: 'Person' },
			{ name: 'to', type: 'Person[]' },
			{ name: 'mail', type: 'LiveMail[]' }
		],
		LiveMail: [
			{ name: 'mail', type: 'Mail' },
			{ name: 'signature', type: 'bytes' }
		],
		LiveEmail: [
			{ name: 'email', type: 'Email' },
			{ name: 'signature', type: 'bytes' }
		]
	} as const

	const verifyingContract = getContract({ address: '0x0', abi: [] })
	const owner = createWalletClient({
		chain: mainnet,
		transport: http()
	})

	// * Create the util with the debug types.
	const domain = { name, version, chainId: 1 }
	const util = new PlugSDK(domain, verifyingContract, DEBUG_TYPES)

	// * Should be able to build Person to get the typehash even
	//   though we are cannot sign it.
	util.build('Person', {
		name: 'Bob',
		wallet: '0x0'
	})

	// @ts-expect-error - SHOULD_FAIL is not a valid type.
	await util.sign(owner, 'SHOULD_FAIL', {
		name: 'Bob',
		wallet: '0x0'
	})

	// * Can sign mail.
	const mail = await util.sign(owner, 'Mail', {
		from: { name: 'Bob', wallet: '0x0' },
		to: { name: 'Alice', wallet: '0x0' },
		contents: 'Hello, world!'
	})

	// * Cannot create an intent with fields that do not belong.
	await util.sign(owner, 'Mail', {
		from: { name: 'Bob', wallet: '0x0' },
		to: { name: 'Alice', wallet: '0x0' },
		contents: 'Hello, world!',
		// @ts-expect-error - Doesnt exist on Mail.
		mail: []
	})

	if (!mail.intent) throw new Error('Mail intent not initialized.')

	// * Can create another type that references the mail type.
	const email = await util.sign(owner, 'Email', {
		from: { name: 'Bob', wallet: '0x0' },
		to: [
			{ name: 'Alice', wallet: '0x0' },
			{ name: 'Charlie', wallet: '0x0' }
		],
		mail: [mail.intent]
	})

	if (!mail.intent) throw new Error('Mail intent not initialized.')

	const Mail = mail.intent

	if (!email.intent) throw new Error('Email intent not initialized.')

	const Email = email.intent

	// * These should both work.
	Mail.mail.from
	Mail.signature

	Email.email
	Email.email.from
	Email.signature

	// * None of these should work.
	// @ts-expect-error - Doesnt exist.
	Mail.email
	// @ts-expect-error - Doesnt exist.
	Mail.Mail
	// @ts-expect-error - Doesnt exist.
	Mail.Email
}
