import deploy, { name, version } from '@/lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'

import { expect } from 'chai'
import { getAddress } from 'viem'

const BASE_AUTH =
	'0x0000000000000000000000000000000000000000000000000000000000000000'

describe('Framework', function () {
	it('pass: instantiate a FrameworkUtil class instance', async function () {
		const { chainId, contract, util } = await loadFixture(deploy)

		expect(util).to.not.be.null.and.not.be.undefined
		expect(util.signedIntents).to.be.empty
		expect(util.info).to.not.be.null

		expect(util.info?.domain).to.eql({
			chainId: chainId,
			verifyingContract: contract.address,
			name,
			version
		})
	})

	it("pass: echo('hi')", async function () {
		const { contract, publicClient } = await loadFixture(deploy)

		const hash = await contract.write.echo(['hi'])

		await publicClient.waitForTransactionReceipt({ hash })
	})

	it('pass: pure echo()', async function () {
		const { contract } = await loadFixture(deploy)

		await contract.read.pureEcho()
	})

	it('pass: getDelegationPacketHash(Delegation memory $input)', async function () {
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
			await signedIntent.address()
		)
		expect(await signedIntent.verify(getAddress(owner.account.address))).to
			.be.true
		expect(getAddress(owner.account.address)).to.eq(
			await contract.read.getSignedDelegationSigner([SignedDelegation])
		)

		// * Validate the digest.
		// const delegationPacketDigest = await contract.read.getDelegationDigest([intent])
		// expect(delegationPacketDigest).to.eq(signedIntent.hash())

		// const typeHash = hashTypedData({
		// 	domain: signedIntent.domain,
		// 	types: signedIntent.types,
		// 	primaryType: 'SignedDelegation',
		//     // @ts-expect-error - test
		// 	message: signedIntent.intent
		// })

		// console.log(typeHash, typedDataHash, signedIntent.hash(intent))

		// expect(typedDataHash).to.eq(typeHash)
		// * Make sure the offchain and onchain hashes match.
		// expect(typedDataHash).to.eq(signedIntent.hash(intent))
	})

	// it.only('pass: getInvocationsTypedDataHash(Invocations memory invocations)', async function () {
	// 	const { util, contract, owner } = await loadFixture(deploy)

	// 	const signedDelegation = await util.sign(owner, {
	// 		delegate: (await owner.getAddress()) as `0x${string}`,
	// 		authority: BASE_AUTH,
	// 		caveats: [],
	// 		salt: BASE_AUTH
	// 	})

	// 	if (signedDelegation.signedMessage === null)
	// 		expect.fail('Signed delegation is null')

	// 	const intent = util.build({
	// 		replayProtection: {
	// 			nonce: 1n,
	// 			queue: 0n
	// 		},
	// 		batch: [
	// 			{
	// 				authority: [signedDelegation.signedMessage],
	// 				transaction: {
	// 					to: (await contract.getAddress()) as `0x${string}`,
	// 					gasLimit: 21000n,
	// 					data: '0x'
	// 				}
	// 			}
	// 		]
	// 	})

	// 	const typedDataHash = await contract.getInvocationsTypedDataHash({
	// 		replayProtection: {
	// 			nonce: 1n,
	// 			queue: 0n
	// 		},
	// 		batch: [
	// 			{
	// 				authority: [signedDelegation.signedMessage],
	// 				transaction: {
	// 					to: (await contract.getAddress()) as `0x${string}`,
	// 					gasLimit: 21000n,
	// 					data: '0x'
	// 				}
	// 			}
	// 		]
	// 	})

	// 	console.table({
	// 		typedDataHash,
	// 		hash: intent.hash()
	// 	})

	// 	expect(typedDataHash).to.eq(intent.hash())
	// })

	// it('fail: alwaysFail()', async function () {
	// 	const { util, contract, owner } = await loadFixture(deploy)

	// 	const signedInvocation = await util.sign(owner, {
	// 		replayProtection: {
	// 			nonce: 1n,
	// 			queue: 0n
	// 		},
	// 		batch: [
	// 			{
	// 				authority: [],
	// 				transaction: {
	// 					to: (await contract.getAddress()) as `0x${string}`,
	// 					gasLimit: 21000n,
	// 					data: (await contract.alwaysFail.populateTransaction())
	// 						.data as `0x${string}`
	// 				}
	// 			}
	// 		]
	// 	})

	// 	if (signedInvocation.signedMessage === null)
	// 		expect.fail('Signed invocation is null')

	// 	await expect(
	// 		contract.invoke([signedInvocation.signedMessage])
	// 	).to.be.revertedWith('I always fail')
	// })

	// it('pass: getEIP712DomainHash(string,string,uint256,address)', async function () {
	// 	expect.fail('Not implemented')
	// })

	// it('pass: verifyDelegationSignature(SignedDelegation memory signedDelegation)`', async function () {
	// 	const { util, contract, owner } = await loadFixture(deploy)

	// 	const signedDelegation = await util.sign(owner, {
	// 		delegate: (await owner.getAddress()) as `0x${string}`,
	// 		authority: BASE_AUTH,
	// 		caveats: [],
	// 		salt: BASE_AUTH
	// 	})

	// 	if (signedDelegation.signedMessage === null)
	// 		expect.fail('Signed delegation is null')

	// 	expect(
	// 		await contract.verifyDelegationSignature(
	// 			signedDelegation.signedMessage
	// 		)
	// 	).to.eq(await owner.getAddress())
	// })

	// it('pass: verifyInvocationSignature(SignedInvocation memory signedInvocation)', async function () {
	// 	const { util, contract, owner, notOwner } = await loadFixture(deploy)

	// 	const signedDelegation = await util.sign(owner, {
	// 		delegate: (await notOwner.getAddress()) as `0x${string}`,
	// 		authority: BASE_AUTH,
	// 		caveats: [],
	// 		salt: BASE_AUTH
	// 	})

	// 	if (signedDelegation.signedMessage === null)
	// 		expect.fail('Signed delegation is null')

	// 	expect(
	// 		await contract.verifyDelegationSignature(
	// 			signedDelegation.signedMessage
	// 		)
	// 	).to.eq(await owner.getAddress())

	// 	const signedInvocation = await util.sign(owner, {
	// 		replayProtection: {
	// 			nonce: 1n,
	// 			queue: 0n
	// 		},
	// 		batch: [
	// 			{
	// 				authority: [signedDelegation.signedMessage],
	// 				transaction: {
	// 					to: (await contract.getAddress()) as `0x${string}`,
	// 					gasLimit: 21000n,
	// 					data: (await contract.alwaysFail.populateTransaction())
	// 						.data as `0x${string}`
	// 				}
	// 			}
	// 		]
	// 	})

	// 	console.log(signedInvocation.signedMessage, signedInvocation.hash())

	// 	if (signedInvocation.signedMessage === null)
	// 		expect.fail('Signed invocation is null')

	// 	console.log(await owner.getAddress(), await notOwner.getAddress())

	// 	expect(
	// 		await contract.verifyInvocationSignature(
	// 			signedInvocation.signedMessage
	// 		)
	// 	).to.eq(await owner.getAddress())
	// })
})
