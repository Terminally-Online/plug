import { db } from "@/server/db"

import axios from "axios"

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR
const FARCASTER_CACHE_TIME = 7 * DAY

type Following = {
	followingAddress: {
		identity: string
		addresses: Array<string>
	}
	followerAddress: {
		identity: string
		addresses: Array<string>
	}
	followingProfileId: string
}

/**
 * Retrieves the Farcaster users that the provided address is following.
 * @param address The address of the user to retrieve the following for.
 * @deprecated This function is deprecated and should no longer be used due to the erosion
 *             of Airstack's API being gated behind Moxie protocol. One must acquire a
 *             Moxie Pass, then buy their shitcoin, then hope that they don't change the
 *             setup of the plans again. Under no circumstances should anyone use this.
 */
export const getFarcasterFollowing = async (address: string) => {
	if (!address) return

	const socket = await db.userSocket.findUnique({
		where: { id: address },
		select: { identity: { select: { farcaster: true } } }
	})

	const cache =
		socket &&
		socket.identity?.farcaster &&
		socket.identity?.farcaster?.updatedAt > new Date(Date.now() - FARCASTER_CACHE_TIME)

	if (cache) return

	const endpoint = "https://api.airstack.xyz/gql"
	const headers = {
		"Content-Type": "application/json",
		"X-API-Key": process.env.AIRSTACK_API_KEY
	}

	const query = `
	        query GetFarcasterFollowing($cursor: String) {
	            SocialFollowings(
	                input: {
	                    filter: {
	                        dappName: { _eq: farcaster }
	                        identity: { _eq: "0x62180042606624f02D8A130dA8A3171e9b33894d" }
	                    }
	                    blockchain: ALL
	                    limit: 200
	                    cursor: $cursor
	                }
	            ) {
	                Following {
	                    followingAddress {
	                        identity
	                        addresses
	                    }
	                    followerAddress {
	                        identity
	                        addresses
	                    }
	                    followingProfileId
	                }
	                pageInfo {
	                    hasNextPage
	                    nextCursor
	                }
	            }
	        }
	    `

	let hasNextPage = true
	let cursor: string | null = null
	const allFollowing: Array<Following> = []

	while (hasNextPage) {
		const variables: { cursor: string | null } = { cursor }
		const response = await axios.post(endpoint, { query, variables }, { headers })

		if (response.data.errors) {
			console.error("GraphQL Errors:", response.data.errors)
		}

		const data = response.data.data.SocialFollowings

		// The Farcaster API returns null when there are no followings or the user does not exist.
		if (data === null) return

		allFollowing.push(...data.Following)

		hasNextPage = data.pageInfo.hasNextPage
		cursor = data.pageInfo.nextCursor
	}

	const followerAddress = allFollowing?.[0]?.followerAddress

	if (!followerAddress) return

	const addresses = {
		connectOrCreate: followerAddress.addresses.map(address => ({
			where: { id: address },
			create: { id: address }
		}))
	}

	const following = {
		connectOrCreate: allFollowing.map(following => ({
			where: { id: following.followingAddress.identity },
			create: {
				id: following.followingAddress.identity,
				addresses: {
					connectOrCreate: following.followingAddress.addresses.map(address => ({
						where: { id: address },
						create: { id: address }
					}))
				}
			}
		}))
	}

	await db.farcasterUser.upsert({
		where: { id: followerAddress.identity },
		create: {
			id: followerAddress.identity,
			addresses,
			following,
			identity: {
				connectOrCreate: {
					where: { socketId: address },
					create: { socketId: address }
				}
			}
		},
		update: {
			addresses,
			following,
			identity: {
				connectOrCreate: {
					where: { socketId: address },
					create: { socketId: address }
				}
			},
			updatedAt: new Date()
		}
	})
}
