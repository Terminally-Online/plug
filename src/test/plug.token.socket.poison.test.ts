// * This test is in place to make sure that Plug is not exposed to the same issue that
//   was discosed by OpenZeppelin here: https://blog.openzeppelin.com/arbitrary-address-spoofing-vulnerability-erc2771context-multicall-public-disclosure
import { getChainId } from '../lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'
import hre, { network } from 'hardhat'

import { expect } from 'chai'
import { concat, encodeFunctionData, getAddress } from 'viem'

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

describe('Plug Poisoned', function () {
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

	it('fail: claim(): spoofed natively', async function () {
		const { contract, owner, notOwner } = await loadFixture(deploy)

		await contract.write.claim()

		const encodedTransaction = encodeFunctionData({
			abi: contract.abi,
			functionName: 'claim'
		})

		const spoofedTransactionData = concat([
			encodedTransaction,
			getAddress(notOwner.account.address)
		])

		await expect(
			owner.sendTransaction({
				to: contract.address,
				data: spoofedTransactionData as `0x${string}`,
				value: 0n
			})
		).to.be.rejected

		expect(
			await contract.read.balanceOf([
				getAddress(notOwner.account.address)
			])
		).to.eq(0n)
		expect(
			await contract.read.balanceOf([getAddress(owner.account.address)])
		).to.eq(1n)
	})

	it('fail: plug(): claim() spoofed', async function () {})
})
