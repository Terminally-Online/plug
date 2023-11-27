import { type FC, memo, type PropsWithChildren } from 'react'

import { useDroppable } from '@dnd-kit/core'

import { cn } from '@/lib/utils'

// Next, let's set up your first Droppable component.
// To do so, we'll be using the useDroppable hook.
//
// The useDroppable hook isn't opinionated about how your app should be structured.
// At minimum though, it requires you pass a ref to the DOM element that you would like
// to become droppable.
//
// You'll also need to provide a unique id attribute to all your droppable components.
// When a draggable element is moved over your droppable element, the isOver property will
// become true.
//
// Note: This is the area that a component can be dropped into, not the component that is
// being dragged. The components are referred to as Draggable.
export const Droppable: FC<PropsWithChildren> = ({ children }) => {
	const { isOver, setNodeRef } = useDroppable({
		id: 'droppable'
	})

	return (
		<div
			ref={setNodeRef}
			className={cn(
				isOver ? 'border-green-100' : '',
				'flex items-center justify-center border-[1px] border-stone-950/10 bg-red-300 h-full'
			)}
		>
			{children}
		</div>
	)
}

export default memo(Droppable)
