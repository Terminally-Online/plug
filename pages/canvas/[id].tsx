import type { ViewportProps } from "@/components/canvas/Viewport";
import Viewport from "@/components/canvas/Viewport";

type PageProps = { params: ViewportProps };

export default async function Page({ params }: PageProps) {
  return <Viewport id={params.id} />;
}
