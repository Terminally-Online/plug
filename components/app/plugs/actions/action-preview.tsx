import { Sentence } from "@/components/app/sentences"
import { usePlugs } from "@/contexts"

export const ActionPreview = () => {
	const { plug, version } = usePlugs()

	if (!plug) return null

	const versionData = plug.versions[plug.versions.length - version]

	return (
		<div className="mb-8 flex flex-col gap-2">
			{versionData &&
				versionData.actions.map((action, index) => (
					<div key={index} className="relative">
						{index < versionData.actions.length - 1 && (
							<div className="absolute bottom-[-12px] top-2 z-[3] ml-[11px] w-[2px] bg-grayscale-100" />
						)}

						<div className="relative z-[4] flex flex-col gap-1">
							<div className="flex flex-row items-center gap-4">
								<p className="flex h-6 w-6 items-center justify-center rounded-full bg-grayscale-0 text-xs text-opacity-60">
									{index + 1}
								</p>

								<Sentence action={action} preview={true} />
							</div>

							<p className="ml-10 text-sm opacity-60">Ready</p>
						</div>
					</div>
				))}
		</div>
	)
}
