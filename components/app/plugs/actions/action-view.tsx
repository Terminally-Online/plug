import { Sentence } from "@/components/app/sentences"
import { usePlugs } from "@/contexts"

export const ActionView = () => {
	const { plug, version } = usePlugs()

	if (!plug) return null

	return (
		<div className="mb-72">
			<div className="flex flex-col">
				{plug.versions[plug.versions.length - version] &&
					plug.versions[plug.versions.length - version].actions.map(
						(action, index) => (
							<>
								<Sentence action={action} />

								{index <
									plug.versions[
										plug.versions.length - version
									].actions.length -
										1 && (
									<div className="mx-auto h-4 w-[2px] bg-grayscale-100" />
								)}
							</>
						)
					)}
			</div>

			<div className="mt-12">
				<h4 className="font-bold opacity-40">
					Next Action Suggestions
				</h4>
			</div>
		</div>
	)
}
