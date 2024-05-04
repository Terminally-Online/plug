import { redirect } from "next/navigation"

import { routes } from "@/lib/routes"

const Page = () => {
	redirect(routes.earlyAccess)
}

export default Page
