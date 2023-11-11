import deploy, { name, version } from '../lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'

import { expect } from 'chai'
import { encodeFunctionData, getAddress } from 'viem'

const BASE_AUTH =
	'0x0000000000000000000000000000000000000000000000000000000000000000'

describe('Plug', function () {
	it('pass: instantiate a PlugUtil class instance', async function () {
		const { chainId, contract, util } = await loadFixture(deploy)

		expect(util).to.not.be.null.and.not.be.undefined
		expect(util.info).to.not.be.null

		expect(util.info?.domain).to.eql({
			chainId: chainId,
			verifyingContract: contract.address,
			name,
			version
		})
	})

	// * Run a test from start to finish for a SignedPermission.
	it('pass: getSignedPermissionSigner', async function () {
		const { util, contract, owner } = await loadFixture(deploy)

		// * Create a Permission.
		const intent = {
			delegate: getAddress(owner.account.address),
			authority: BASE_AUTH,
			caveats: [],
			salt: BASE_AUTH
		} as const

		// * Sign the permission to make it executable.
		const signedIntent = await util.sign(owner, 'Permission', intent)

		if (!signedIntent) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedPermission = signedIntent.intent

		if (!SignedPermission) expect.fail('Intent could not be signed.')

		// * Make sure the intent signer matched the recovered signer.
		expect(getAddress(owner.account.address)).to.eq(
			await signedIntent.address({})
		)
		expect(
			await signedIntent.verify({
				address: getAddress(owner.account.address)
			})
		).to.be.true
		expect(getAddress(owner.account.address)).to.eq(
			await contract.read.getSignedPermissionSigner([SignedPermission])
		)
	})

	// * Run a test from start to finish for a SignedIntents.
	it('pass: getSignedIntentsSigner(SignedIntent memory signedIntent)', async function () {
		const { util, contract, owner, notOwner } = await loadFixture(deploy)

		// * Create a Permission.
		const intent = {
			delegate: getAddress(owner.account.address),
			authority: BASE_AUTH,
			caveats: [],
			salt: BASE_AUTH
		} as const

		// * Sign the permission to make it executable.
		const signedPermission = await util.sign(owner, 'Permission', intent)

		if (!signedPermission) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedPermission = signedPermission.intent

		if (!SignedPermission) expect.fail('Intent could not be signed.')

		expect(
			await contract.read.getSignedPermissionSigner([SignedPermission])
		).to.eq(getAddress(owner.account.address))

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'mutedEcho'
		})

		const signedIntent = await util.sign(notOwner, 'Intents', {
			replayProtection: {
				nonce: 1n,
				queue: 0n
			},
			batch: [
				{
					authority: [SignedPermission],
					transaction: {
						to: getAddress(owner.account.address),
						gasLimit: 21000n,
						data: encodedTransaction
					}
				}
			]
		})

		if (!signedIntent) expect.fail('Signed intent does not exist.')
		if (!signedIntent.intent) expect.fail('Intent does not exist.')

		const SignedIntent = signedIntent.intent

		expect(
			await contract.read.getSignedIntentsSigner([SignedIntent])
		).to.eq(getAddress(notOwner.account.address))
	})

	it('fail: signedIntents: mutedEcho()', async function () {
		const { util, contract, owner } = await loadFixture(deploy)

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'mutedEcho'
		})

		const signedIntents = await util.sign(owner, 'Intents', {
			replayProtection: {
				nonce: 1n,
				queue: 0n
			},
			batch: [
				{
					authority: [],
					transaction: {
						to: contract.address,
						gasLimit: 21000n,
						data: encodedTransaction
					}
				}
			]
		})

		if (!signedIntents) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedIntents = signedIntents.intent

		if (!SignedIntents) expect.fail('Intent could not be signed.')

		await expect(
			contract.write.invoke([[SignedIntents]])
		).to.be.rejectedWith('EchoMuted')
	})
})
