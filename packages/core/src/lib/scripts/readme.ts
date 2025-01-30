import fs from 'fs'
import path from 'path'

interface DeploymentInfo {
	initCode: string
	initCodeHash: string
	deployment: {
		salt: string
		address: string
		rarity: string
	}
}

interface MiningConfig {
	seconds: number
	leading: number
	total: number
	addresses: number
	factory: string
	caller: string
	quick: boolean
}

interface VersionInfo {
	config: MiningConfig
	contracts: Record<string, DeploymentInfo>
}

interface VersionedAddressMap {
	[version: string]: VersionInfo
}

function generateReadme(templateString: string): string {
	try {
		const addressesPath = path.join(__dirname, '../addresses.json')
		const addressesContent = fs.readFileSync(addressesPath, 'utf8')
		const addresses: VersionedAddressMap = JSON.parse(addressesContent)

		let deploymentTree = ''

		// Sort versions in descending order (newest first)
		const versions = Object.keys(addresses).sort().reverse()

		versions.forEach(ver => {
			const versionInfo = addresses[ver]
			deploymentTree += `\`\`\`ml\n[v${ver}]:\n`

			// Add mining configuration
			deploymentTree += `├─ Mining Configuration:\n`
			deploymentTree += `│  ├─ Leading Zeros: ${versionInfo.config.leading}\n`
			deploymentTree += `│  ├─ Total Zeros: ${versionInfo.config.total}\n`
			deploymentTree += `│  ├─ Factory: ${versionInfo.config.factory}\n`
			deploymentTree += `│  └─ Quick Mode: ${versionInfo.config.quick ? 'Yes' : 'No'}\n`
			deploymentTree += `│\n`
			deploymentTree += `└─ Contracts:\n`

			const contracts = Object.entries(versionInfo.contracts).sort()
			contracts.forEach(([contract, info], contractIndex) => {
				const isLastContract = contractIndex === contracts.length - 1
				const symbol = isLastContract ? '   └' : '   ├'

				// Safely access the deployment data
				const deployment = info?.deployment
				if (!deployment) {
					console.error(
						`Missing deployment data for ${contract} in version ${ver}`
					)
					deploymentTree += `${symbol}─ ${contract} — "No deployment data"\n`
					return
				}

				deploymentTree += `${symbol}─ ${contract} [${deployment.rarity}] — "${deployment.address}"\n`
			})

			deploymentTree += '```\n\n'
		})

		return templateString.replace(
			'<__DEPLOYMENT_ADDRESSES__>',
			deploymentTree.trim()
		)
	} catch (error) {
		console.error('Error generating README:', error)
		return templateString
	}
}

// Read template and generate new README
const templatePath = path.join(process.cwd(), './README.Template.md')
const template = fs.readFileSync(templatePath, 'utf8')
const newReadme = generateReadme(template)
const readmePath = path.join(process.cwd(), './README.md')
fs.writeFileSync(readmePath, newReadme)

console.log('README.md has been updated with current deployment addresses')
