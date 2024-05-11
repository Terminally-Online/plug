import { exec } from "child_process"

import { version } from "../../package.json"

const outputString = `   ◍ Plug PostgreSQL Database ${version}`
const envString = `- Environments: .env`

function startDatabase(
	containerName: string,
	databaseName: string,
	databasePassword: string,
	port: string
) {
	exec(
		`docker start ${containerName} || docker run --name ${containerName} -e POSTGRES_PASSWORD=${databasePassword} -p ${port}:${
			parseInt(port) - 2
		} -d ${databaseName}`,
		err => {
			if (err) {
				console.error(
					"Error starting the PostgreSQL container in Docker:",
					err
				)

				return
			}

			console.log(`${outputString}
   - Local: postgres://postgres:${databasePassword}@localhost:${port}/${databaseName}
   - Container Name: ${containerName}
   ${envString}
`)
		}
	)
}

const containerName = process.env.CONTAINER_NAME ?? "postgres"
const databaseName = process.env.DATABASE_NAME ?? "postgres"
const databasePort = process.env.DATABASE_PORT ?? "5434"
const databasePassword = process.env.DATABASE_PASSWORD ?? "postgres"

startDatabase(containerName, databaseName, databasePassword, databasePort)
