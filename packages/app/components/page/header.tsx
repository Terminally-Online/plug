import { useSession } from "next-auth/react"

export const PageHeader = () => {
    const { column } = useColumnData(COLUMNS.MOBILE_INDEX)
    const { data: session } = useSession()

    if (!column) return null

    if (!session?.user.id?.startsWith("0x")) {
        return (
            <Container className="sticky top-0 z-10 bg-white border-b border-plug-green/10">
                <AuthenticateHeader />
            </Container>
        )
    }

    return (
        <Container className="sticky top-0 z-10 bg-white border-b border-plug-green/10">
            {column.key === COLUMNS.KEYS.PLUG && <PlugHeader />}
            {column.key === COLUMNS.KEYS.DISCOVER && <DiscoverHeader />}
            {column.key === COLUMNS.KEYS.MY_PLUGS && <MyPlugsHeader />}
            {column.key === COLUMNS.KEYS.PROFILE && <ProfileHeader />}
        </Container>
    )
}
