import { bytesToHex, encodePacked, toBytes } from "viem"
import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { emitter } from "@/server/emitter"

const events = {
	add: "add-vault",
	rename: "rename-vault"
}

export const vault = Prisma.validator<Prisma.VaultDefaultArgs>()({})
export type Vault = Prisma.VaultGetPayload<typeof vault>

const CURRENT_IMPLEMENTATION_VERSION = 0

const FACTORY_ADDRESS = ""
const FACTORY_ABI = []

export default createTRPCRouter({
	// NOTE: Get the next address that a vault will be deployed at.
	address: protectedProcedure.query(async ({ ctx }) => {
		const { nextVaultSalt, nextVaultAddress } = await ctx.db.user.upsert(
			{
				where: {
					id: ctx.session.address
				},
				update: {},
				create: {
					id: ctx.session.address
				}
			}
		)

		if (!nextVaultSalt || !nextVaultAddress) {
			// * Need to cast the type to be viem compatible.
			const salt: string = bytesToHex(
				toBytes(
					encodePacked(
						["address", "uint80", "uint16"],
						[
							ctx.session.address as `0x${string}`,
							BigInt(Date.now() / 1000),
							CURRENT_IMPLEMENTATION_VERSION
						]
					),
					{ size: 32 }
				)
			)

			// TODO: Need to call into the contract to get the next
			//		 address that will be deployed.
			const address = ""

			// TODO: Save the address and salt to the database.
			await ctx.db.user.update(
				{
					where: {
						id: ctx.session.address
					},
					data: {
						id: ctx.session.address,
						nextVaultSalt: salt,
						nextVaultAddress: address
					}
				}
			)

			return { salt, address }
		}

		return {
			salt: nextVaultSalt,
			address: nextVaultAddress
		}
	}),

	all: protectedProcedure.query(async ({ ctx }) => {
		try {
			return await ctx.db.vault.findMany({
				where: {
					ownerAddress: ctx.session.address
				}
			})
		} catch (error) {
			throw new TRPCError({
				code: "BAD_REQUEST"
			})
		}
	}),

	get: protectedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		try {
			const vault = await ctx.db.vault.findUnique({
				where: {
					address_chainId: {
						address: input,
						chainId: 1
					}
				}
			})

			if (!vault) throw new TRPCError({ code: "NOT_FOUND" })

			if (vault.ownerAddress !== ctx.session.address)
				throw new TRPCError({
					code: "UNAUTHORIZED"
				})

			return vault
		} catch (error) {
			throw new TRPCError({
				code: "BAD_REQUEST"
			})
		}
	}),

	// TODO: Get balances.

	// NOTE: Add a Vault to a user's account on specified chains that reflects
	//       a new deployment of the Vault contract.
	// NOTE: Due to the tokenized representation of Vault ownership, a single
	//       Vault may have multiple owners across multiple chains. Though,
	//       a single owner should be able to acquire the Vault on multiple chains
	//       (still with a single click) by using a Merkle signature. The representation
	//       of asset ownership is not multichain itself.
	add: protectedProcedure
		.input(
			z.object({
				address: z.string(),
				chainIds: z.array(z.number()),
				lastBlockIndexed: z.number()
			})
		)
		.mutation(async ({ ctx, input }) => {
			try {
				const { address, chainIds, lastBlockIndexed } = input

				const vaultsPromises = chainIds.map(async chainId => {
					return await ctx.db.vault.create({
						data: {
							address,
							chainId,
							name: "",
							ownerAddress: ctx.session.address,
							lastBlockIndexed
						}
					})
				})

				const vaults = await Promise.all(vaultsPromises)

				emitter.emit(events.add, vaults)

				return vaults
			} catch (error) {
				throw new TRPCError({
					code: "BAD_REQUEST"
				})
			}
		}),

	// NOTE: Rename a Vault across all chain ids that the user owns. If they only
	//		 own it on some chains, then only rename it on those chains. If they
	//		 do not own it on any chains, then throw an error.
	rename: protectedProcedure
		.input(
			z.object({
				address: z.string(),
				name: z.string()
			})
		)
		.mutation(async ({ ctx, input }) => {
			try {
				const { address, name } = input

				const mutation = await ctx.db.vault.updateMany({
					where: {
						address,
						ownerAddress: ctx.session.address
					},
					data: {
						name
					}
				})

				if (mutation.count == 0)
					throw new TRPCError({ code: "NOT_FOUND" })

				const vaults = await ctx.db.vault.findMany({
					where: {
						address,
						ownerAddress: ctx.session.address
					}
				})

				emitter.emit(events.rename, vaults)

				return vaults
			} catch (error) {
				throw new TRPCError({
					code: "BAD_REQUEST"
				})
			}
		}),

	onAdd: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Array<Vault>>(emit => {
			const handleAdd = (data: Array<Vault>) => {
				if (
					!data.some(
						vault => vault.ownerAddress === ctx.session.address
					)
				)
					return

				emit.next(data)
			}

			emitter.on(events.add, handleAdd)

			return () => {
				emitter.off(events.add, handleAdd)
			}
		})
	}),

	onRename: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Array<Vault>>(emit => {
			const handleRename = (data: Array<Vault>) => {
				if (
					!data.some(
						vault => vault.ownerAddress === ctx.session.address
					)
				)
					return

				emit.next(data)
			}

			emitter.on(events.rename, handleRename)

			return () => {
				emitter.off(events.rename, handleRename)
			}
		})
	})
})
