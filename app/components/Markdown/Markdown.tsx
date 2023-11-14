import type { CSSProperties, FC } from "react";
import { memo } from "react";

import { default as MarkdownJSX } from "markdown-to-jsx";

const styles: CSSProperties = {
  border: "1px dashed gray",
  padding: "0.5rem 1rem",
  cursor: "move",
};

export interface MarkdownProps {
  title: string;
  yellow?: boolean;
  preview?: boolean;
}

export const Markdown: FC<MarkdownProps> = memo(function Markdown({
  title,
  yellow,
  preview,
}) {
  const backgroundColor = yellow ? "yellow" : "white";

  return (
    <div
      style={{ ...styles, backgroundColor }}
      role={preview ? "MarkdownPreview" : "Markdown"}
    >
      <MarkdownJSX
        options={{
          overrides: {
            h1: { component: "h1", props: { className: "text-2xl font-bold" } },
            h2: { component: "h2", props: { className: "text-xl font-bold" } },
            h3: { component: "h3", props: { className: "text-lg font-bold" } },
            h4: {
              component: "h4",
              props: { className: "text-base font-bold" },
            },
            h5: { component: "h5", props: { className: "text-sm font-bold" } },
            h6: { component: "h6", props: { className: "text-xs font-bold" } },
            p: {
              component: "p",
              props: { className: "text-base font-normal" },
            },
            a: {
              component: "a",
              props: { className: "text-base font-normal" },
            },
          },
        }}
      >
        {title}
      </MarkdownJSX>
    </div>
  );
});
