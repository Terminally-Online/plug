import { getServerSession } from "next-auth";

import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { getServerClient } from "@/app/api/trpc/client.server";
import Viewport from "@/components/canvas/Viewport";

export default async function Page({ params }: { params: { id: string } }) {
  const session = await getServerSession(authOptions);
  const t = getServerClient(session);

  const canvas = await t.canvas.get(params.id);

  return <Viewport canvas={canvas} />;
}
