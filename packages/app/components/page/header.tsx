import { useSession } from "next-auth/react"
import { COLUMNS, useColumnData } from "@/state/columns"
import { Container } from "@/components/shared/layout/container"

export const PageHeader = () => {
    const { column } = useColumnData(COLUMNS.MOBILE_INDEX)
    const { data: session } = useSession()

    if (!column) return null

    return (
        <Container className="sticky top-0 z-10 bg-white border-b border-plug-green/10">
            <div className="flex items-center justify-between py-4">
                <h1 className="font-bold">
                    {column.key === COLUMNS.KEYS.HOME && "Home"}
                    {column.key === COLUMNS.KEYS.DISCOVER && "Discover"}
                    {column.key === COLUMNS.KEYS.MY_PLUGS && "My Plugs"}
                    {column.key === COLUMNS.KEYS.ACTIVITY && "Activity"}
                    {column.key === COLUMNS.KEYS.PROFILE && "Profile"}
                    {column.key === COLUMNS.KEYS.AUTHENTICATE && "Login"}
                    {column.key === COLUMNS.KEYS.PLUG && "Plug"}
                </h1>
            </div>
        </Container>
    )
}