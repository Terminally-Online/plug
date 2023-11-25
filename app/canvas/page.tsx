import { redirect } from "next/navigation";

import { getServerSession } from "next-auth";

import { authOptions } from "@/server/auth";
import { api } from "@/trpc/server";

export const dynamic = "force-dynamic";

export default async function Page() {
  const session = await getServerSession(authOptions);
  //   const t = getServerClient(session);

  const username = session?.user?.name;

  if (!username) redirect(`/connect`);

  const canvases = await api.healthcheck.query();

  return <>{canvases}</>;

  //   return <CanvasPreviewGrid canvases={canvases} />;
}
