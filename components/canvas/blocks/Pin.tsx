import { memo, useState } from 'react'

import { CaretSortIcon, CheckIcon } from '@radix-ui/react-icons'

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import {
	Command,
	CommandEmpty,
	CommandGroup,
	CommandInput,
	CommandItem,
	CommandList
} from '@/components/ui/command'
import { Input } from '@/components/ui/input'
import {
	Popover,
	PopoverContent,
	PopoverTrigger
} from '@/components/ui/popover'
import CanvasStore from '@/lib/store'
import { Pins, Pin as PinType } from '@/lib/types'
import { cn } from '@/lib/utils'

type PopoverTriggerProps = React.ComponentPropsWithoutRef<typeof PopoverTrigger>

interface PinProps extends PopoverTriggerProps {
	selectedPin: PinType
	pins: Pins
	onPinChange: (pin: PinType) => void
}

export const Pin = ({
	className,
	selectedPin,
	pins,
	onPinChange
}: PinProps) => {
	const [open, setOpen] = useState(false)

	// * Lock the camera from moving when a Popover is open so that the user can scroll without moving the camera.
	CanvasStore.lockCamera(open)

	return (
		<Popover open={open} onOpenChange={setOpen}>
			<div className="relative bg-stone-900 w-full">
				<p className="absolute top-[-10px] left-4 bg-stone-900 border-[1px] border-stone-950 rounded-full text-xs text-white/60 p-[1px] px-[8px]">
					{selectedPin.type.slice(0, 1).toUpperCase() +
						selectedPin.type.slice(1)}
				</p>
				<p className="absolute top-[-10px] right-4 text-right bg-stone-900 border-[1px] border-stone-950 rounded-full text-xs text-white/60 p-[1px] px-[8px]">
					Pending
				</p>

				<PopoverTrigger asChild>
					<Button
						variant="outline"
						role="combobox"
						aria-expanded={open}
						aria-label="Select a Pin"
						className={cn(
							'w-full justify-between border-b-[1px] border-transparent border-b-stone-950 p-4 py-6',
							className
						)}
					>
						<Avatar className="mr-2 h-4 w-4">
							<AvatarImage
								src={`https://avatar.vercel.sh/${selectedPin.value}.png`}
								alt={selectedPin.label}
							/>
							<AvatarFallback>SC</AvatarFallback>
						</Avatar>

						{selectedPin.label}

						<CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />
					</Button>
				</PopoverTrigger>

				{Object.keys(selectedPin.schema.shape).map(key => (
					<div
						key={key}
						className="w-full flex flex-col bg-stone-800"
					>
						<Input
							type="text"
							id={key}
							placeholder={key
								.replace(/-/g, ' ')
								.replace(/\w\S*/g, w =>
									w.replace(/^\w/, c => c.toUpperCase())
								)}
							autoComplete="off"
							className="border-transparent"
						/>
					</div>
				))}
			</div>

			<PopoverContent className="w-[400px] p-0">
				<Command>
					<CommandList>
						<CommandInput placeholder="Search Pin..." />
						<CommandEmpty>No Pin found.</CommandEmpty>
						{pins.map(pinSource =>
							pinSource ? (
								<CommandGroup
									key={pinSource.label}
									heading={pinSource.label}
								>
									{pinSource.pins.map(pin => (
										<CommandItem
											key={pin.value}
											onSelect={() => {
												onPinChange(pin)
												setOpen(false)
											}}
											className="text-sm"
										>
											<Avatar className="mr-2 h-5 w-5">
												<AvatarImage
													src={`https://avatar.vercel.sh/${pin.value}.png`}
													alt={pin.label}
													className="grayscale"
												/>
												<AvatarFallback>
													SC
												</AvatarFallback>
											</Avatar>

											{pin.label}

											<CheckIcon
												className={cn(
													'ml-auto h-4 w-4',
													selectedPin.value ===
														pin.value
														? 'opacity-100'
														: 'opacity-0'
												)}
											/>
										</CommandItem>
									))}
								</CommandGroup>
							) : null
						)}
					</CommandList>
				</Command>
			</PopoverContent>
		</Popover>
	)
}

export default memo(Pin)
