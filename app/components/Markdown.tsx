import { default as MarkdownJSX } from 'markdown-to-jsx';
import { CanvasPosition, Position } from "./Position";

interface TextBlockProps extends CanvasPosition {
  text: string;
  left: number;
  top: number;
  width: number;
  height: number;
}

const Markdown = ({
  text,
  left,
  top,
  width,
  height
}: TextBlockProps) => {
  return (
    <Position left={left} top={top} width={width} height={height}>
      <div
        className="flex items-center justify-center"
        style={{
          width: `${width}px`,
          height: `${height}px`,
        }}
      >
        <MarkdownJSX options={{
          overrides: { 
            h1: { component: 'h1', props: { className: 'text-2xl font-bold' } },
            h2: { component: 'h2', props: { className: 'text-xl font-bold' } },
            h3: { component: 'h3', props: { className: 'text-lg font-bold' } },
            h4: { component: 'h4', props: { className: 'text-base font-bold' } },
            h5: { component: 'h5', props: { className: 'text-sm font-bold' } },
            h6: { component: 'h6', props: { className: 'text-xs font-bold' } },
            p: { component: 'p', props: { className: 'text-base font-normal' } },
            a: { component: 'a', props: { className: 'text-base font-normal' } },
          }
        }}>
          {text}
        </MarkdownJSX>
      </div>
    </Position>
  );
};

export default Markdown;
