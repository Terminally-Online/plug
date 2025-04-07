import { FC } from "react"


import { Animate } from "@/components/app/utils/animate"
import { useOrderedConnections } from "@/lib/hooks/account/useConnections"
import { Connector } from "./connector";
import { ConnectorQrCode } from "./connector-qr-code";

export const Connectors: FC<{ index: number; from?: string }> = ({ index, from }) => {
	const connectors = useOrderedConnections(true)

	return (
		<div className="mb-auto w-full pt-2">
			<ConnectorQrCode />
			<div className="h-[1px] w-full bg-plug-green/10" />
			<div className="px-4 pt-4">
				<Animate.List>
					{connectors.map(connector => (
						<Animate.ListItem key={connector.id}>
							<Connector connector={connector} index={index} from={from} />
						</Animate.ListItem>
					))}
				</Animate.List>
			</div>
		</div>
	)
}


