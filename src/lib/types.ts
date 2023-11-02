import { TypedData } from 'abitype'

// ! Turn a dictionary of EIP-712 types into a union of the keys
//   that have a signed pair.
export type TypedDataToKeysWithSignedPair<T extends TypedData> =
	// * Remove all the 'never' types from the union.
	string &
		Exclude<
			{
				// * For each key in the typed data, if the capitalized version of the
				//   key prefixed with 'Signed' is a key in the typed data, then return
				//   the key, otherwise return never.
				[K in keyof T]: `Signed${Capitalize<
					string & K
				>}` extends keyof T
					? K
					: never
				// * Then, get the union of all the keys.
			}[keyof T],
			never
		>

// * Turn an intent object into the expected onchain shape of the
//   SignedIntent pair output.
export type TypedDataToSignedIntent<K, U> = Record<
	'signature',
	`0x${string}`
> & {
	[TK in K as Lowercase<string & TK>]: U
}
