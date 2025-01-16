import { useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"

export enum Flag {
	SHOW_PWA,
	SHOW_DEVELOPER
}

const DEFAULT_FLAGS: Partial<Record<keyof typeof Flag, true>> = {
	SHOW_PWA: true,
	SHOW_DEVELOPER: true
}

const setBit = (bitmap: number, position: number): number => bitmap | (1 << position)
const clearBit = (bitmap: number, position: number): number => bitmap & ~(1 << position)
const getBit = (bitmap: number, position: number): boolean => !!(bitmap & (1 << position))

const createInitialBitmap = (): number => {
	return Object.entries(Flag)
		.filter(([key]) => isNaN(Number(key)))
		.reduce((bitmap, [flag]) => {
			const position = Flag[flag as keyof typeof Flag]
			return flag in DEFAULT_FLAGS ? setBit(bitmap, position) : bitmap
		}, 0)
}

const flagsAtom = atomWithStorage<number>("plug.flags", createInitialBitmap())

export const useFlags = () => {
	const [flagsBitmap, setFlagsBitmap] = useAtom(flagsAtom)

	const getFlag = (flagType: Flag): boolean => {
		return getBit(flagsBitmap, flagType)
	}

	const handleFlag = (flagType: Flag, value: boolean | ((currentValue: boolean) => boolean)) => {
		setFlagsBitmap(prevBitmap => {
			const currentValue = getBit(prevBitmap, flagType)
			const newValue = typeof value === "function" ? value(currentValue) : value
			return newValue ? setBit(prevBitmap, flagType) : clearBit(prevBitmap, flagType)
		})
	}

	const flags = Object.fromEntries(
		Object.entries(Flag)
			.filter(([key]) => isNaN(Number(key)))
			.map(([key]) => [key.toLowerCase().replace(/_/g, "-"), getFlag(Flag[key as keyof typeof Flag])])
	) as Record<string, boolean>

	return { flags, getFlag, handleFlag }
}
