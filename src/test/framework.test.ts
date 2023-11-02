import deploy, { name, version } from '../lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'

import { expect } from 'chai'
import { encodeFunctionData, getAddress } from 'viem'

const BASE_AUTH =
	'0x0000000000000000000000000000000000000000000000000000000000000000'

describe('Framework', function () {
	it('pass: instantiate a FrameworkUtil class instance', async function () {
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

	// * Run a test from start to finish for a SignedDelegation.
	it('pass: getSignedDelegationSigner', async function () {
		const { util, contract, owner } = await loadFixture(deploy)

		// * Create a Delegation.
		const intent = {
			delegate: getAddress(owner.account.address),
			authority: BASE_AUTH,
			caveats: [],
			salt: BASE_AUTH
		} as const

		// * Sign the delegation to make it executable.
		const signedIntent = await util.sign(owner, 'Delegation', intent)

		if (!signedIntent) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedDelegation = signedIntent.intent

		if (!SignedDelegation) expect.fail('Intent could not be signed.')

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
			await contract.read.getSignedDelegationSigner([SignedDelegation])
		)
	})

	// * Run a test from start to finish for a SignedInvocations.
	it('pass: getSignedInvocationsSigner(SignedInvocation memory signedInvocation)', async function () {
		const { util, contract, owner, notOwner } = await loadFixture(deploy)

		// * Create a Delegation.
		const intent = {
			delegate: getAddress(owner.account.address),
			authority: BASE_AUTH,
			caveats: [],
			salt: BASE_AUTH
		} as const

		// * Sign the delegation to make it executable.
		const signedDelegation = await util.sign(owner, 'Delegation', intent)

		if (!signedDelegation) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedDelegation = signedDelegation.intent

		if (!SignedDelegation) expect.fail('Intent could not be signed.')

		expect(
			await contract.read.getSignedDelegationSigner([SignedDelegation])
		).to.eq(getAddress(owner.account.address))

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'mutedEcho'
		})

		const signedInvocation = await util.sign(notOwner, 'Invocations', {
			replayProtection: {
				nonce: 1n,
				queue: 0n
			},
			batch: [
				{
					authority: [SignedDelegation],
					transaction: {
						to: getAddress(owner.account.address),
						gasLimit: 21000n,
						data: encodedTransaction
					}
				}
			]
		})

		if (!signedInvocation) expect.fail('Signed invocation does not exist.')
		if (!signedInvocation.intent) expect.fail('Intent does not exist.')

		const SignedInvocation = signedInvocation.intent

		expect(
			await contract.read.getSignedInvocationsSigner([SignedInvocation])
		).to.eq(getAddress(notOwner.account.address))
	})

	it('fail: signedInvocations: mutedEcho()', async function () {
		const { util, contract, owner } = await loadFixture(deploy)

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'mutedEcho'
		})

		const signedInvocations = await util.sign(owner, 'Invocations', {
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

		if (!signedInvocations) expect.fail('Signed intent does not exist.')

		// * Retrieve the object that will be passed onchain.
		const SignedInvocations = signedInvocations.intent

		if (!SignedInvocations) expect.fail('Intent could not be signed.')

		await expect(
			contract.write.invoke([[SignedInvocations]])
		).to.be.rejectedWith('EchoMuted')
	})
})
