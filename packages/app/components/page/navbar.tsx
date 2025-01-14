import { useSession } from "next-auth/react"
import Avatar from "boring-avatars"
import { motion } from "framer-motion"
import { Bell, HousePlug, Plus, Search } from "lucide-react"
import { memo } from "react"
import { Image } from "@/components/app/utils/image"
import { cn } from "@/lib"
import { COLUMNS, useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useSocket } from "@/state/authentication"

export const PageNavbar = memo(() => {
    const { data: session } = useSession()
    const { avatar } = useSocket()
    const { column, handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
    const { handle: plugHandle } = usePlugStore()
    const showNavbar = column?.key !== COLUMNS.KEYS.PLUG

    if (!showNavbar) return null

    return (
        <div className="fixed bottom-0 left-0 right-0 z-[10] border-t-[1px] border-plug-green/10 bg-white">
            <div className="relative z-[11] flex flex-row items-center justify-between gap-2 px-8 py-4">
                {/* Home Button */}
                <button
                    className="group flex h-8 w-8 items-center justify-center"
                    onClick={() => handle.navigate({ index: COLUMNS.MOBILE_INDEX, key: COLUMNS.KEYS.HOME })}
                >
                    <HousePlug
                        size={24}
                        className={cn(
                            "text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100",
                            column?.key === COLUMNS.KEYS.HOME && "text-opacity-100"
                        )}
                    />
                </button>

                {/* Search Button */}
                <button
                    className="group flex h-8 w-8 items-center justify-center"
                    onClick={() => handle.navigate({ index: COLUMNS.MOBILE_INDEX, key: COLUMNS.KEYS.HOME })}
                >
                    <Search
                        size={24}
                        className={cn(
                            "text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100",
                            column?.key === COLUMNS.KEYS.HOME && "text-opacity-100"
                        )}
                    />
                </button>

                {/* Create Plug Button */}
                <button
                    className="group flex h-8 w-8 items-center justify-center rounded-md bg-plug-yellow text-plug-green transition-all duration-200 ease-in-out hover:bg-plug-yellow/50"
                    onClick={() => plugHandle.plug.add({ index: COLUMNS.MOBILE_INDEX })}
                >
                    <Plus
                        size={24}
                        className="opacity-80 transition-all duration-200 ease-in-out group-hover:opacity-100"
                    />
                </button>

                {/* Activity Button */}
                <button
                    className="group flex h-8 w-8 items-center justify-center"
                    onClick={() => handle.navigate({ index: COLUMNS.MOBILE_INDEX, key: COLUMNS.KEYS.ACTIVITY })}
                >
                    <Bell
                        size={24}
                        className={cn(
                            "text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100",
                            column?.key === COLUMNS.KEYS.ACTIVITY && "text-opacity-100"
                        )}
                    />
                </button>

                {/* Profile Button */}
                <button
                    className="group h-8 w-8"
                    onClick={() => handle.navigate({ index: COLUMNS.MOBILE_INDEX, key: COLUMNS.KEYS.PROFILE })}
                >
                    {session && (
                        <div className="relative h-8 w-8 rounded-md bg-plug-green/5 transition-all duration-200 ease-in-out">
                            <motion.div
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 1 }}
                                exit={{ opacity: 0 }}
                                transition={{ duration: 0.2 }}
                            >
                                {avatar ? (
                                    <Image
                                        src={avatar}
                                        alt="ENS Avatar"
                                        width={64}
                                        height={64}
                                        className="h-full w-full rounded-sm"
                                    />
                                ) : (
                                    <div className="overflow-hidden rounded-sm">
                                        <Avatar
                                            name={session?.address}
                                            variant="beam"
                                            size={"100%"}
                                            colors={["#00E100", "#A3F700"]}
                                            square
                                        />
                                    </div>
                                )}
                            </motion.div>
                        </div>
                    )}
                </button>
            </div>
        </div>
    )
})
PageNavbar.displayName = 'PageNavbar'
