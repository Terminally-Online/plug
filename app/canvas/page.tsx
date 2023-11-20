import { redirect } from "next/navigation";

import { getServerSession } from "next-auth";

import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { getServerClient } from "@/app/api/trpc/client.server";
import CanvasPreviewGrid from "@/components/canvas/blocks/CanvasPreviewGrid";

export const dynamic = "force-dynamic";

export default async function Page() {
  const session = await getServerSession(authOptions);
  const t = getServerClient(session);

  const username = session?.user?.name;

  if (!username) redirect(`/connect`);

  const canvases = await t.canvas.all();

  return <CanvasPreviewGrid canvases={canvases} />;
}
