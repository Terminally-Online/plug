import { OpenseaCollection } from "@prisma/client"

export type OpenseaCollectible = {
	identifier: string
	collection: OpenseaCollection
	contract: string
	token_standard: string
	name: string
	description: string
	image_url: string
	display_image_url: string
	display_animation_url: string
	metadata_url: string
	opensea_url: string
	updated_at: string
	is_disabled: boolean
	is_nsfw: boolean
}
