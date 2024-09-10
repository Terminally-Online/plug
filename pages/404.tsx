import { Callout, StaticLayout } from "@/components"

const NotFound = () => {
	return (
		<StaticLayout title="Page Not Found">
			<div className="flex h-screen flex-col items-center justify-center">
				<Callout.EmptyPage />
			</div>
		</StaticLayout>
	)
}

export default NotFound
