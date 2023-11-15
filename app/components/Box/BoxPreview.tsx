import type { CSSProperties, FC } from "react";
import { memo, useEffect, useState } from "react";

import { Box } from "./Box";

const styles: CSSProperties = {
  display: "inline-block",
  transform: "rotate(-7deg)",
  WebkitTransform: "rotate(-7deg)",
};

export interface BoxPreviewProps {
  title: string;
  left: number;
  top: number;
}

export interface BoxPreviewState {
  tickTock: any;
}

export const BoxPreview: FC<BoxPreviewProps> = memo(function BoxPreview({
  title,
}) {
  const [tickTock, setTickTock] = useState(false);

  useEffect(
    function subscribeToIntervalTick() {
      const interval = setInterval(() => setTickTock(!tickTock), 500);
      return () => clearInterval(interval);
    },
    [tickTock]
  );

  return (
    <div style={styles}>
        <Box title={title} yellow={tickTock} preview />
    </div>
  );
});
