import { useSession } from "next-auth/react"
import { ColumnAuthenticate } from "@/components/app/columns/utils/column-authenticate"
import { Container } from "@/components/app/layout/container"
import { PlugsDiscover } from "@/components/app/plugs/discover"
import { PlugsMine } from "@/components/app/plugs/mine"
import { Plug } from "@/components/app/plugs/plug"
import { SocketActivity } from "@/components/app/sockets/activity/activity-list"
import { SocketAssets } from "@/components/app/sockets/assets"
import { SocketProfile } from "@/components/app/sockets/profile"
import { Plugs } from "@/components/shared/framework/plugs"
import { COLUMNS, useColumnStore } from "@/state/columns"
import { useSocket } from "@/state/authentication"
import { ReferralRequired } from "@/components/app/utils/referral-required"
import { Button } from "@/components/shared/buttons/button"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"

const ProfileContent = () => {
    const { data: session } = useSession()
    const { socket } = useSocket()
    
    if (!socket || !session?.user.id) return null
    
    return (
        <div className="flex flex-col gap-4 p-4">
            {/* Show user's assets using existing SocketAssets component */}
            <SocketAssets 
                index={COLUMNS.MOBILE_INDEX} 
                address={session.user.id} 
                hasTokens 
                hasCollectibles 
            />
        </div>
    )
}

export const PageContent = () => {
    const { data: session } = useSession()
    const { socket } = useSocket()
    const { column, handle } = useColumnStore(COLUMNS.MOBILE_INDEX)

    const isAuthenticated = session?.user.id?.startsWith("0x")
    const isReferred = Boolean(socket && socket.identity?.referrerId)

    console.log("[PageContent] Render", {
        hasSession: !!session,
        isAuthenticated,
        hasSocket: !!socket,
        isReferred,
        hasColumn: !!column,
        columnKey: column?.key,
        hasHandle: !!handle
    })

    // Show auth content if not logged in
    if (!isAuthenticated) {
        return (
            <Container>
                <ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
            </Container>
        )
    }

    // Show referral screen if not referred
    if (!isReferred) {
        return (
            <Container>
                <ReferralRequired />
            </Container>
        )
    }

    // Add null check for handle
    if (!column) {
        console.log("[PageContent] Redirecting to HOME - No column")
        if (!handle) {
            console.warn("[PageContent] Handle is undefined")
            return null
        }
        handle.navigate({ 
            index: COLUMNS.MOBILE_INDEX, 
            key: COLUMNS.KEYS.HOME 
        })
        return null
    }

    if (column.key === COLUMNS.KEYS.PROFILE) {
        return (
            <Container>
                <ProfileContent />
            </Container>
        )
    }

    // Render the appropriate content based on column key
    switch (column.key) {
        case COLUMNS.KEYS.PLUG:
            return <Plug index={COLUMNS.MOBILE_INDEX} item={column.item} from={column.from} />
        case COLUMNS.KEYS.HOME:
            return (
                <Container className="mb-24">
                    <Plugs hideEmpty={true} />
                </Container>
            )
        case COLUMNS.KEYS.DISCOVER:
            return <PlugsDiscover className="pt-4" />
        case COLUMNS.KEYS.MY_PLUGS:
            return <PlugsMine className="pt-4" />
        case COLUMNS.KEYS.ACTIVITY:
            return (
                <Container className="pt-4">
                    <SocketActivity />
                </Container>
            )
        case COLUMNS.KEYS.AUTHENTICATE:
            return (
                <Container className="pt-4">
                    <ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
                </Container>
            )
        default:
            console.warn("[PageContent] No matching column key:", column?.key)
            return null
    }
}