import { getChainId } from '@/lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox-viem/network-helpers'
import hre, { network } from 'hardhat'

import { expect } from 'chai'
import { bytesToHex, hexToBytes } from 'viem'

import { PlugSDK } from '@/core/sdk'

import { constants } from '@nftchance/plug-types'

const [vaultName] = ['PlugVaultSocket', '0.0.0']
const [name, version] = ['PlugFactorySocket', '0.0.0']

async function deployImplementation() {
	const chainId = await getChainId(network)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const implementation = await hre.viem.deployContract(vaultName, [])

	const publicClient = await hre.viem.getPublicClient()

	return {
		chainId,
		name,
		version,
		implementation,
		owner,
		notOwner,
		publicClient
	}
}

async function deploy() {
	const { chainId, publicClient, implementation } =
		await loadFixture(deployImplementation)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const contract = await hre.viem.deployContract(name, [name, version])

	const util = new PlugSDK(name, version, chainId, constants.types, contract)

	return {
		chainId,
		name,
		version,
		implementation,
		contract,
		util,
		owner,
		notOwner,
		publicClient
	}
}

describe('Plug Factory', function () {
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

	it('pass: deploy(notOwner)', async function () {
		const { implementation, contract, owner, notOwner, publicClient } =
			await loadFixture(deploy)

		const salt = bytesToHex(
			hexToBytes(notOwner.account.address, {
				size: 32
			})
		)

		const hash = await contract.write.deploy([
			implementation.address,
			notOwner.account.address,
			salt
		])
		await publicClient.waitForTransactionReceipt({ hash })

		await expect(
			contract.write.deploy([
				implementation.address,
				owner.account.address,
				salt
			])
		).to.be.rejected
	})
})
