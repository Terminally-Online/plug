import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'
import { ethers } from 'hardhat'

// import { expect } from 'chai'

describe('Lock', function () {
	async function deployMock() {
		const [owner, other] = await ethers.getSigners()

		const name = 'FrameworkMock'
		const version = '0.0.0'

		const Contract = await ethers.getContractFactory(name)
		const contract = await Contract.deploy(name, version)

		return { contract, name, version, owner, other }
	}

	describe('Deployment', function () {
		it('Has the right domain hash', async function () {
			await loadFixture(deployMock)
		})
	})
})
