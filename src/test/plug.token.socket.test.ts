import { getChainId } from '../lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'
import hre, { network } from 'hardhat'

import { expect } from 'chai'
import { encodeFunctionData, getAddress } from 'viem'

import { PlugSDK } from '@/core/sdk'

import { constants } from '@nftchance/plug-types'

const [name, version] = ['PlugERC20Socket', '0.0.0']

async function deploy() {
	const chainId = await getChainId(network)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const contract = await hre.viem.deployContract(name, [name, version])

	const publicClient = await hre.viem.getPublicClient()

	const util = new PlugSDK(name, version, chainId, constants.types, contract)

	return {
		chainId,
		name,
		version,
		contract,
		util,
		owner,
		notOwner,
		publicClient
	}
}

describe('Plug Token', function () {
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

	it('pass: name()', async function () {
		const { contract } = await loadFixture(deploy)

		await contract.read.name()
	})

	it('pass: symbol()', async function () {
		const { contract } = await loadFixture(deploy)

		await contract.read.symbol()
	})

	it('pass: claim()', async function () {
		const { contract } = await loadFixture(deploy)

		await contract.write.claim()
	})

	it('pass: plug(): claim()', async function () {
		const { util, contract, owner, notOwner } = await loadFixture(deploy)

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'claim'
		})

		const estimate = await contract.estimateGas.claim()

		const signedPlugs = await util.sign(notOwner, 'Plugs', {
			breaker: {
				nonce: 1n,
				queue: 0n
			},
			plugs: [
				{
					pins: [],
					current: {
						ground: contract.address,
						voltage: estimate,
						data: encodedTransaction
					},
					forced: true
				}
			]
		})

		if (!signedPlugs) expect.fail('Live intent does not exist.')

		const LivePlugs = signedPlugs.intent

		if (!LivePlugs) expect.fail('Plug could not be signed.')

		expect(
			await contract.read.balanceOf([getAddress(owner.account.address)])
		).to.eq(0n)
		expect(await contract.write.plug([[LivePlugs]]))
		expect(
			await contract.read.balanceOf([
				getAddress(notOwner.account.address)
			])
		).to.eq(1n)
	})

	it('fail: claim() twice', async function () {
		const { contract } = await loadFixture(deploy)

		await contract.write.claim()

		await expect(contract.write.claim()).to.be.rejectedWith(
			'PlugERC20Socket:claimed'
		)
	})
})
