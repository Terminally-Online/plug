import { exec } from "child_process"

import { version } from "../package.json"

const outputString = `   ◍ Plug PostgreSQL Database ${version}`
const envString = `- Environments: .env`

function startDatabase(
	containerName: string,
	databaseName: string,
	databasePassword: string,
	port: string
) {
	exec(
		`docker start ${containerName} || docker run --name ${containerName} -e POSTGRES_PASSWORD=${databasePassword} -p ${port}:${port} -d ${databaseName}`,
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
   ${envString}\n
			`)
		}
	)
}

function stopDatabase(containerName: string) {
	exec(`docker stop ${containerName}`, err => {
		if (err) {
			console.error("Error stopping the PostgreSQL container:", err)

			return
		}

		console.log("PostgreSQL container stopped.")
	})
}

const containerName = process.env.CONTAINER_NAME ?? "emporium"
const databaseName = process.env.DATABASE_NAME ?? "postgres"
const databasePort = process.env.DATABASE_PORT ?? "5432"
const databasePassword = process.env.DATABASE_PASSWORD ?? "postgres"

startDatabase(containerName, databaseName, databasePassword, databasePort)

process.on("SIGTERM", () => {
	stopDatabase(containerName)

	setTimeout(() => process.exit(0), 1000)
})
