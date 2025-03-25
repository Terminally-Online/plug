import { Marquee } from "../../utils/marquee"

export const ScrollingError = ({ error }: { error: string | undefined }) => {
	if (!error) return null

	return (
		<div className="relative min-h-6 overflow-x-hidden">
			<div className="absolute bottom-0 left-0 top-0 z-[20] w-12 bg-gradient-to-r from-plug-white to-plug-white/0" />
			<div className="absolute bottom-0 right-0 top-0 z-[20] w-12 bg-gradient-to-l from-plug-white to-plug-white/0" />
			<Marquee className="-z-1 relative max-w-full whitespace-nowrap font-bold text-plug-red">{error}</Marquee>
		</div>
	)
}
