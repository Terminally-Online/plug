import { redirect } from "next/navigation";

import { getServerSession } from "next-auth";
import { authOptions } from "../api/auth/[...nextauth]/route";
import CanvasPreviewGrid from "./components/Blocks/CanvasPreviewGrid";

import { getServerClient } from "../api/trpc/client.server";

export const dynamic = 'force-dynamic'

export default async function Page() {
  const session = await getServerSession(authOptions);
  const t = getServerClient(session);

  const username = session?.user?.name;

  if (!username) redirect(`/connect`);

  const canvases = await t.all();

  return <CanvasPreviewGrid canvases={canvases} />;
}
