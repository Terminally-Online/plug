import { memo } from "react";
import type { FC, PropsWithChildren } from "react";

import MarkdownJSX from "markdown-to-jsx";

export type BoxProps = {
  preview?: boolean;
};

export const Markdown: FC<PropsWithChildren<BoxProps>> = memo(function Markdown({
  children,
  preview,
}) {
  return (
    <div
      className="bg-white cursor-move p-2 px-4 border-[1px] border-gray-200"
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
        {children as string}
      </MarkdownJSX>
    </div>
  );
});
