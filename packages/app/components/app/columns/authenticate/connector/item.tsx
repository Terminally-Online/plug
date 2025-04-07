import { FC } from "react"

import { useConnect, Connector as wagmiConnector } from "wagmi"

import { Loader2 } from "lucide-react"

import { useSetAtom } from "jotai"

import { ConnectorImage } from "@/components/app/columns/authenticate/connector/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { cn, CONNECTOR_ICON_OVERRIDE_MAP, greenGradientStyle, recentConnectorIdAtom, useRecentConnectorId } from "@/lib"
import { useAuthenticate } from "@/lib/hooks/account/useAuthenticate"
import { useColumnActions } from "@/state/columns"

type Props = { connector: wagmiConnector; index: number; from?: string }
export const ConnectorItem: FC<Props> = ({ connector, index, from }) => {
	const connection = useConnect()
	const { authenticate } = useAuthenticate()

	const { navigate } = useColumnActions()

	const updateRecentConnectorId = useSetAtom(recentConnectorIdAtom)

	const isLoading = connection.isPending && connection.variables?.connector === connector
	const isRecent = connector.id === useRecentConnectorId()
	const isDetected = connector.isInjected as boolean
	const isDisabled = Boolean(connection?.isPending)
	const icon = CONNECTOR_ICON_OVERRIDE_MAP[connector.id] ?? connector.icon

	const Badge = () => {
		if (isLoading) return <Loader2 className="animate-spin opacity-40" size={14} />
		if (isRecent) return <span style={{ ...greenGradientStyle }}>Recent</span>
		if (isDetected) return <span className="opacity-40">Detected</span>
		return null
	}

	const handleConnect = () => {
		if (isDisabled) return

		connection.connect(
			{ connector },
			{
				onSuccess: data => {
					updateRecentConnectorId(connector.id)
					authenticate({ address: data.accounts[0] }, { onSuccess: () => navigate({ index, from }) })
				}
			}
		)
	}

	return (
		<Accordion className={cn(isDisabled && "cursor-not-allowed bg-plug-green/5")} onExpand={handleConnect}>
			<div className="flex flex-row items-center gap-4">
				<ConnectorImage icon={icon} name={connector.name} />
				<p className="font-bold">{connector.name}</p>
				<p className="ml-auto text-sm font-bold">
					<Badge />
				</p>
			</div>
		</Accordion>
	)
}
