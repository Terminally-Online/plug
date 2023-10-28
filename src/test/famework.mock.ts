import deploy, { name, version } from '@/lib/functions/hardhat'
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers'

import { expect } from 'chai'

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
		const { contract } = await loadFixture(deploy)
		await contract.write.echo(['hi'])
	})

	it('pass: pureEcho()', async function () {
		const { contract } = await loadFixture(deploy)
		await contract.read.pureEcho()
	})

	it('fail: mutedEcho()', async function () {
		const { contract } = await loadFixture(deploy)
		await expect(contract.read.mutedEcho()).to.be.rejectedWith('EchoMuted')
	})
})
