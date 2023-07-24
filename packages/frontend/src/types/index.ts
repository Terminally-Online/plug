const ZeroAddress = `${'0x'}${'0'.repeat(40)}` as const

export type AddressLike = typeof ZeroAddress
