import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'
import { ethers } from 'hardhat'

import { expect } from 'chai'

describe('Framework', function () {
	async function deployFixture() {
		const [owner, otherAccount] = await ethers.getSigners()

		const Contract = await ethers.getContractFactory('Framework')
		const contract = await Contract.deploy()

		return { contract, owner, otherAccount }
	}

	describe('Deployment', function () {
		it('pass: set the right time', async function () {
			const { contract } = await loadFixture(deployFixture)

			expect.fail('Not implemented')

			contract
		})
	})
})
