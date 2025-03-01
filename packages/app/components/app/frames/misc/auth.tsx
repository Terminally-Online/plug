import { useSession } from "next-auth/react"

import { User } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { AuthButton } from "@/components/shared/buttons/auth"
import { columnByIndexAtom, COLUMNS, isFrameAtom} from "@/state/columns"
import { useAtom, useAtomValue } from "jotai"

export const AuthFrame = () => {
	const { data: session } = useSession()

	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))
	const frameKey = "auth"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	return (
		<Frame
			index={COLUMNS.MOBILE_INDEX}
			icon={<User size={18} />}
			label={session?.address ? "Account" : "Login"}
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-4 px-6 pb-4">
				<AuthButton />
			</div>
		</Frame>
	)
}
