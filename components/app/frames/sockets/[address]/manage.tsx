import { useEffect } from "react"

import { PencilLine } from "lucide-react"

import { Button } from "@/components/buttons"
import { Search } from "@/components/inputs"
import { useFrame, useSockets } from "@/contexts"
import { useClipboard, useDebounce } from "@/lib"

import { Frame } from "../../base"

export const ManageFrame = () => {
	const { socket, handleRename } = useSockets()
	const { frameVisible } = useFrame()
	const { copied, handleCopied } = useClipboard(socket?.socketAddress ?? "")

	const [name, debouncedName, handleDebounce, debouncedRef] = useDebounce(
		socket?.name ?? ""
	)

	useEffect(() => {
		if (socket === undefined || debouncedRef.current === debouncedName)
			return
		debouncedRef.current = debouncedName
		handleRename(debouncedName)
	}, [socket, debouncedName, debouncedRef, handleRename])

	return (
		<Frame
			className="z-[2]"
			label="Manage Socket"
			visible={frameVisible === "manage"}
		>
			<Search
				icon={<PencilLine size={14} className="opacity-60" />}
				placeholder="Socket name"
				search={name}
				handleSearch={handleDebounce}
			/>

			<div className="mt-[20px] flex flex-row gap-2">
				<Button
					variant="secondary"
					className="w-max"
					onClick={handleCopied}
				>
					{copied ? "Copied" : "Copy"}
				</Button>
				<Button variant="primary" className="w-full" onClick={() => {}}>
					Deploy Onchain
				</Button>
			</div>
		</Frame>
	)
}
