import { exec } from "child_process"

import { env } from "@/env"

import { version } from "../../package.json"

const outputString = `   ◍ Plug PostgreSQL Database ${version}`
const envString = `- Environments: .env`

function startDatabase(containerName: string, databaseName: string, databasePassword: string, port: string) {
	exec(
		`docker start ${containerName} || docker run --name ${containerName} -e POSTGRES_PASSWORD=${databasePassword} -p ${port}:${
			parseInt(port) - 2
		} -d ${databaseName}`,
		err => {
			if (err) {
				console.error("Error starting the PostgreSQL container in Docker:", err)

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

startDatabase(
	env.DOCKER_CONTAINER_NAME,
	env.DOCKER_DATABASE_NAME,
	env.DOCKER_DATABASE_PASSWORD,
	env.DOCKER_DATABASE_PORT
)
