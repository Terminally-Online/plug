import type { CSSProperties, FC } from "react";
import { memo, useEffect, useState } from "react";

import { Markdown } from "./Markdown";

const styles: CSSProperties = {
  display: "inline-block",
  transform: "rotate(-7deg)",
  WebkitTransform: "rotate(-7deg)",
};

export type MarkdownPreviewProps = {
  title: string;
};

export type MarkdownPreviewState = {
  tickTock: any;
};

export const MarkdownPreview: FC<MarkdownPreviewProps> = memo(
  function MarkdownPreview({ title }) {
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
        <Markdown title={title} yellow={tickTock} preview />
      </div>
    );
  }
);
