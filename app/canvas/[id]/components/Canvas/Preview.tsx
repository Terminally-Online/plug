import type { CSSProperties, FC, PropsWithChildren } from "react";
import { memo } from "react";

const styles: CSSProperties = {
  display: "inline-block",
  opacity: 0.8,
};

export const Preview: FC<PropsWithChildren> = memo(function Preview({
  children,
}) {
  return <div style={styles}>{children}</div>;
});
