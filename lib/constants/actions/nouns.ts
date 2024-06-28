import { HandCoinsIcon, HashIcon, ShirtIcon } from "lucide-react"
import { keccak256, parseAbi, toHex, zeroAddress } from "viem"

import { abis } from "../abis"

const BACKGROUND_SELECTOR = keccak256(toHex("background"))
const BODY_SELECTOR = keccak256(toHex("body"))
const ACCESSORY_SELECTOR = keccak256(toHex("accessory"))
const HEAD_SELECTOR = keccak256(toHex("head"))
const GLASSES_SELECTOR = keccak256(toHex("glasses"))

const GLASSES_PATH = "/actions/options/nouns/glasses/"
const HEAD_PATH = "/actions/options/nouns/heads/"
const BODY_PATH = "/actions/options/nouns/bodies/"
const BACKGROUND_PATH = "/actions/options/nouns/backgrounds/"
const ACCESSORY_PATH = "/actions/options/nouns/accessories/"

export const nouns = {
	bid: {
		address: zeroAddress,
		abi: abis.nouns.bid,
		inputs: parseAbi([abis.nouns.bid])[0]["inputs"],
		options: undefined,
		sentence: "Bid on Noun with {0} $ETH",
		info: "Bid on a Noun with a specific amount of ETH. This can be used to make an offer on a Noun.",
		icon: HandCoinsIcon,
		primary: true
	},
	hasTrait: {
		address: zeroAddress,
		abi: abis.nouns.hasTrait,
		inputs: parseAbi([abis.nouns.hasTrait])[0]["inputs"],
		options: [
			[
				{
					label: "background",
					value: BACKGROUND_SELECTOR,
					imagePath: `${BACKGROUND_PATH}bg-warm.png`
				},
				{
					label: "body",
					value: BODY_SELECTOR,
					imagePath: `${BODY_PATH}body-blue-sky.png`
				},
				{
					label: "accessory",
					value: ACCESSORY_SELECTOR,
					imagePath: `${ACCESSORY_PATH}accessory-bling-anchor.png`
				},
				{
					label: "head",
					value: HEAD_SELECTOR,
					imagePath: `${HEAD_PATH}head-dino.png`
				},
				{
					label: "glasses",
					value: GLASSES_SELECTOR,
					imagePath: `${GLASSES_PATH}glasses-square-red.png`
				}
			],
			{
				[BACKGROUND_SELECTOR]: [
					{
						label: "cool",
						value: "0x03",
						imagePath: `${BACKGROUND_PATH}bg-cool.png`
					},
					{
						label: "warm",
						value: "0x04",
						imagePath: `${BACKGROUND_PATH}bg-warm.png`
					}
				],
				[BODY_SELECTOR]: [
					{
						label: "Beige Light",
						value: "0x04",
						imagePath: `${BODY_PATH}body-bege-crt.png`
					},
					{
						label: "Blue Sky",
						value: "0x04",
						imagePath: `${BODY_PATH}body-blue-sky.png`
					},
					{
						label: "Bluegrey",
						value: "0x04",
						imagePath: `${BODY_PATH}body-bluegrey.png`
					},
					{
						label: "Cold",
						value: "0x04",
						imagePath: `${BODY_PATH}body-cold.png`
					},
					{
						label: "Computer Blue",
						value: "0x04",
						imagePath: `${BODY_PATH}body-computerblue.png`
					},
					{
						label: "Dark Brown",
						value: "0x04",
						imagePath: `${BODY_PATH}body-darkbrown.png`
					},
					{
						label: "Dark Pink",
						value: "0x04",
						imagePath: `${BODY_PATH}body-darkpink.png`
					},
					{
						label: "Fog Grey",
						value: "0x04",
						imagePath: `${BODY_PATH}body-foggrey.png`
					},
					{
						label: "Gold",
						value: "0x04",
						imagePath: `${BODY_PATH}body-gold.png`
					},
					{
						label: "White",
						value: "0x04",
						imagePath: `${BODY_PATH}body-grayscale-1.png`
					},
					{
						label: "Dark Grey",
						value: "0x04",
						imagePath: `${BODY_PATH}body-grayscale-7.png`
					},
					{
						label: "Darker Grey",
						value: "0x04",
						imagePath: `${BODY_PATH}body-grayscale-8.png`
					},
					{
						label: "Darkest Grey",
						value: "0x04",
						imagePath: `${BODY_PATH}body-grayscale-9.png`
					},
					{
						label: "Green",
						value: "0x04",
						imagePath: `${BODY_PATH}body-green.png`
					},
					{
						label: "Gunk",
						value: "0x04",
						imagePath: `${BODY_PATH}body-gunk.png`
					},
					{
						label: "Hot Brown",
						value: "0x04",
						imagePath: `${BODY_PATH}body-hotbrown.png`
					},
					{
						label: "Magenta",
						value: "0x04",
						imagePath: `${BODY_PATH}body-magenta.png`
					},
					{
						label: "Orange",
						value: "0x04",
						imagePath: `${BODY_PATH}body-orange.png`
					},
					{
						label: "Orange Light",
						value: "0x04",
						imagePath: `${BODY_PATH}body-orange-yellow.png`
					},
					{
						label: "Peach",
						value: "0x04",
						imagePath: `${BODY_PATH}body-peachy-a.png`
					},
					{
						label: "Peachy",
						value: "0x04",
						imagePath: `${BODY_PATH}body-peachy-B.png`
					},
					{
						label: "Purple",
						value: "0x04",
						imagePath: `${BODY_PATH}body-purple.png`
					},
					{
						label: "Red",
						value: "0x04",
						imagePath: `${BODY_PATH}body-red.png`
					},
					{
						label: "Pink",
						value: "0x04",
						imagePath: `${BODY_PATH}body-redpinkish.png`
					},
					{
						label: "Rust",
						value: "0x04",
						imagePath: `${BODY_PATH}body-rust.png`
					},
					{
						label: "Slime Green",
						value: "0x04",
						imagePath: `${BODY_PATH}body-slimegreen.png`
					},
					{
						label: "Teal",
						value: "0x04",
						imagePath: `${BODY_PATH}body-teal.png`
					},
					{
						label: "Teal Light",
						value: "0x04",
						imagePath: `${BODY_PATH}body-teal-light.png`
					},
					{
						label: "Yellow",
						value: "0x04",
						imagePath: `${BODY_PATH}body-yellow.png`
					}
				],
				[ACCESSORY_SELECTOR]: [
					{
						label: "1n",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-1n.png`
					},
					{
						label: "Aardvark",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-aardvark.png`
					},
					{
						label: "Anchor Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-anchor.png`
					},
					{
						label: "Arrow",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-arrow.png`
					},
					{
						label: "Axe",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-axe.png`
					},
					{
						label: "Band Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-marsface.png`
					},
					{
						label: "Beige Grid",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-grid-simple-bege.png`
					},
					{
						label: "Big Red Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-big-red.png`
					},
					{
						label: "Bird Flying",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bird-flying.png`
					},
					{
						label: "Black Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-shirt-black.png`
					},
					{
						label: "Black Tie",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-tie-black-on-white.png`
					},
					{
						label: "Blit Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-blit.png`
					},
					{
						label: "Blood Stains",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stains-blood.png`
					},
					{
						label: "Blue Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-blue-med.png`
					},
					{
						label: "Brown Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-brown.png`
					},
					{
						label: "Cardinal",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bird-side.png`
					},
					{
						label: "Carrot",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-carrot.png`
					},
					{
						label: "CC",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-cc.png`
					},
					{
						label: "CC 2",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-cc2.png`
					},
					{
						label: "Chain Logo",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-chain-logo.png`
					},
					{
						label: "Chain Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-scissors.png`
					},
					{
						label: "Chameleon",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-belly-chameleon.png`
					},
					{
						label: "Checker Big Green",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checkers-big-green.png`
					},
					{
						label: "Checker Big Red",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checkers-big-red-cold.png`
					},
					{
						label: "Checker Blue",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checkers-blue.png`
					},
					{
						label: "Checker Disco Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-checkerdisco.png`
					},
					{
						label: "Checker Grey",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-bigwalk-greylight.png`
					},
					{
						label: "Checker Rainbow",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-bigwalk-rainbow.png`
					},
					{
						label: "Checker RGB",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-RGB.png`
					},
					{
						label: "Checker Small Black",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-spaced-black.png`
					},
					{
						label: "Checker Small Blue",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-spaced-white.png`
					},
					{
						label: "Checker Small Magenta",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checkers-magenta-80.png`
					},
					{
						label: "Checker Vibrant",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-vibrant.png`
					},
					{
						label: "Checker Wide Black",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-bigwalk-blue-prime.png`
					},
					{
						label: "Checker Wide White",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-checker-bigwalk-blue-prime.png`
					},
					{
						label: "Cheese Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-cheese.png`
					},
					{
						label: "Chicken",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-chicken.png`
					},
					{
						label: "Cloud",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-cloud.png`
					},
					{
						label: "Clover",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-clover.png`
					},
					{
						label: "Clown Scarf",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-scarf-clown.png`
					},
					{
						label: "Collar Sunset",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-collar-sunset.png`
					},
					{
						label: "Copy",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-copy.png`
					},
					{
						label: "Cow",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-cow.png`
					},
					{
						label: "DAO",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-dao-black.png`
					},
					{
						label: "Dawn Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-dawn.png`
					},
					{
						label: "Doom",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-doom.png`
					},
					{
						label: "Dope",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-dope-text.png`
					},
					{
						label: "Dusk Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-dusk.png`
					},
					{
						label: "Decay",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-decay-gray-dark.png`
					},
					{
						label: "Decay Pride",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-decay-pride.png`
					},
					{
						label: "Dinosaur",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-dinosaur.png`
					},
					{
						label: "Dollar Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-dollar-bling.png`
					},
					{
						label: "Dragon",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-dragon.png`
					},
					{
						label: "Ducky",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-ducky.png`
					},
					{
						label: "ETH",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-eth.png`
					},
					{
						label: "Eye",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-eye.png`
					},
					{
						label: "Flash",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-flash.png`
					},
					{
						label: "Foo",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-foo-black.png`
					},
					{
						label: "Fries",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-fries.png`
					},
					{
						label: "Glacier Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-glacier.png`
					},
					{
						label: "Glasses",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-glasses.png`
					},
					{
						label: "Gold Bar Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-gold-ingot.png`
					},
					{
						label: "Grease",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-grease.png`
					},
					{
						label: "Heart",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-heart.png`
					},
					{
						label: "Heart Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-love.png`
					},
					{
						label: "Hoodie Strings",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-hoodiestrings-uneven.png`
					},
					{
						label: "Ice Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-ice.png`
					},
					{
						label: "ICO",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-ico.png`
					},
					{
						label: "ID",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-id.png`
					},
					{
						label: "Infinity",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-infinity.png`
					},
					{
						label: "Insignia",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-insignia.png`
					},
					{
						label: "I/O",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-io.png`
					},
					{
						label: "Leaf",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-leaf.png`
					},
					{
						label: "Lightbulb",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-lightbulb.png`
					},
					{
						label: "Lines Green",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-lines-45-greens.png`
					},
					{
						label: "Lines Rose",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-lines-45-rose.png`
					},
					{
						label: "LMAO",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-lmao.png`
					},
					{
						label: "LOL",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-lol.png`
					},
					{
						label: "LP",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-lp.png`
					},
					{
						label: "Mask Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-mask.png`
					},
					{
						label: "Mars Face",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-marsface.png`
					},
					{
						label: "Matrix White",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-matrix-white.png`
					},
					{
						label: "Mint",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-mint.png`
					},
					{
						label: "Moon Block",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-moon-block.png`
					},
					{
						label: "Nil",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-nil-grey-dark.png`
					},
					{
						label: "Noun",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-noun.png`
					},
					{
						label: "Noun Green",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-noun-green.png`
					},
					{
						label: "Nouns Glasses",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-glasses-logo-sun.png`
					},
					{
						label: "Nouns Glasses 2",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-glasses-logo.png`
					},
					{
						label: "Olive Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-olive.png`
					},
					{
						label: "Pi",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-pi.png`
					},
					{
						label: "Pop",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-pop.png`
					},
					{
						label: "Pizza Necklace",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-pizza-bling.png`
					},
					{
						label: "Pocket Pencil",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-pocket-pencil.png`
					},
					{
						label: "Pride Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-pride.png`
					},
					{
						label: "Purple Tie",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-tie-purple-on-white.png`
					},
					{
						label: "Red/Pink Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-body-gradient-redpink.png`
					},
					{
						label: "Rain",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-rain.png`
					},
					{
						label: "Rainbow Steps",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-rainbow-steps.png`
					},
					{
						label: "Red Stripes",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-red-cold.png`
					},
					{
						label: "RGB",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-rgb.png`
					},
					{
						label: "Rings",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-rings.png`
					},
					{
						label: "Robot",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-robot.png`
					},
					{
						label: "ROFL",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-rofl.png`
					},
					{
						label: "Scissors",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-scissors.png`
					},
					{
						label: "Sparkles",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-bling-sparkles.png`
					},
					{
						label: "Sunset Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-sunset.png`
					},
					{
						label: "Safety Vest",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-safety-vest.png`
					},
					{
						label: "Shrimp",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-shrimp.png`
					},
					{
						label: "Slime Splat",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-slimesplat.png`
					},
					{
						label: "Snowflake",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-snowflake.png`
					},
					{
						label: "Stripes & Checks",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stripes-and-checks.png`
					},
					{
						label: "Sunset",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-sunset.png`
					},
					{
						label: "Tatewaku",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-a2+b2.png`
					},
					{
						label: "Taxi",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-taxi-checkers.png`
					},
					{
						label: "Think",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-think.png`
					},
					{
						label: "Tie Dye",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-tie-dye.png`
					},
					{
						label: "a2 + b2",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-a2+b2.png`
					},
					{
						label: "Uroko",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-noun.png`
					},
					{
						label: "We",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-we.png`
					},
					{
						label: "YAY",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-txt-yay.png`
					},
					{
						label: "Wall",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-wall.png`
					},
					{
						label: "Wave",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-wave.png`
					},
					{
						label: "Wet Money",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-wet-money.png`
					},
					{
						label: "Wool Weave",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-woolweave-bicolor.png`
					},
					{
						label: "Wool Weave 2",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-woolweave-dirt.png`
					},
					{
						label: "XCOPY X",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-secret-x.png`
					},
					{
						label: "Yolo",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-text-yolo.png`
					},
					{
						label: "Yin & Yang",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-yingyang.png`
					},
					{
						label: "Yo Shirt",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-tee-yo.png`
					},
					{
						label: "Zombie Stains",
						value: "0x03",
						imagePath: `${ACCESSORY_PATH}accessory-stains-zombie.png`
					}
				],
				[HEAD_SELECTOR]: [
					{
						label: "35mm Film",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-film-35mm.png`
					},
					{
						label: "Aardvark",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-aardvark.png`
					},
					{
						label: "Abstract",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-abstract.png`
					},
					{
						label: "Ape",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ape.png`
					},
					{
						label: "Bag",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bag.png`
					},
					{
						label: "Bagpipe",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bagpipe.png`
					},
					{
						label: "Banana",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-banana.png`
					},
					{
						label: "Bank",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bank.png`
					},
					{
						label: "Bao",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-porkbao.png`
					},
					{
						label: "Baseball",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-baseball-gameball.png`
					},
					{
						label: "Basketball",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-basketball.png`
					},
					{
						label: "Bat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bat.png`
					},
					{
						label: "Beached Whale",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-whale.png`
					},
					{
						label: "Bear",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bear.png`
					},
					{
						label: "Beer",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-beer.png`
					},
					{
						label: "Beet",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-beet.png`
					},
					{
						label: "Bell",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bell.png`
					},
					{
						label: "Bigfoot",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bigfoot.png`
					},
					{
						label: "Blackhole",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-blackhole.png`
					},
					{
						label: "Blueberry",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-blueberry.png`
					},
					{
						label: "Bomb",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bomb.png`
					},
					{
						label: "Bonsai",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bonsai.png`
					},
					{
						label: "Boombox",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-boombox.png`
					},
					{
						label: "Boot",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-boot.png`
					},
					{
						label: "Box",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-box.png`
					},
					{
						label: "Boxing Glove",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-boxingglove.png`
					},
					{
						label: "Brain",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-brain.png`
					},
					{
						label: "Bubble Gum",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bubblegum.png`
					},
					{
						label: "Burger",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-burger-dollarmenu.png`
					},
					{
						label: "Cake",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cake.png`
					},
					{
						label: "Calculator",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-calculator.png`
					},
					{
						label: "Calendar",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-calendar.png`
					},
					{
						label: "Camcorder",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-camcorder.png`
					},
					{
						label: "Canned Ham",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cannedham.png`
					},
					{
						label: "Capybara",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-capybara.png`
					},
					{
						label: "Car",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-car.png`
					},
					{
						label: "Cash Register",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cash-register.png`
					},
					{
						label: "Cassette Tape",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cassettetape.png`
					},
					{
						label: "Cat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cat.png`
					},
					{
						label: "CD",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cd.png`
					},
					{
						label: "Chain",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chain.png`
					},
					{
						label: "Chainsaw",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chainsaw.png`
					},
					{
						label: "Chameleon",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chameleon.png`
					},
					{
						label: "Chart",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chart-bars.png`
					},
					{
						label: "Cheese",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cheese.png`
					},
					{
						label: "Chef Hat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chefhat.png`
					},
					{
						label: "Cherry",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cherry.png`
					},
					{
						label: "Chicken",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chicken.png`
					},
					{
						label: "Chilli",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chilli.png`
					},
					{
						label: "Chipboard",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chipboard.png`
					},
					{
						label: "Chips",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chips.png`
					},
					{
						label: "Chocolate",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-chocolate.png`
					},
					{
						label: "Cloud",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cloud.png`
					},
					{
						label: "Clover",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-clover.png`
					},
					{
						label: "Clutch",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-clutch.png`
					},
					{
						label: "Coffee Bean",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-coffeebean.png`
					},
					{
						label: "Cone",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cone.png`
					},
					{
						label: "Cookie",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cookie.png`
					},
					{
						label: "Cordless Phone",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cordlessphone.png`
					},
					{
						label: "Cotton Ball",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cottonball.png`
					},
					{
						label: "Couch",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-couch.png`
					},
					{
						label: "Cow",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-cow.png`
					},
					{
						label: "Crab",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-crab.png`
					},
					{
						label: "Crane",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-crane.png`
					},
					{
						label: "Crocodile",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-croc-hat.png`
					},
					{
						label: "Crown King",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-crown.png`
					},
					{
						label: "Crown Queen",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-queencrown.png`
					},
					{
						label: "CRT",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-crt-bsod.png`
					},
					{
						label: "Crystal Ball",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-crystalball.png`
					},
					{
						label: "Diamond Blue",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-diamond-blue.png`
					},
					{
						label: "Diamond Red",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-diamond-red.png`
					},
					{
						label: "Dictionary",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-dictionary.png`
					},
					{
						label: "Dinosaur",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-dino.png`
					},
					{
						label: "DNA",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-dna.png`
					},
					{
						label: "Dog",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-dog.png`
					},
					{
						label: "Doughnut",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-doughnut.png`
					},
					{
						label: "Drill",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-drill.png`
					},
					{
						label: "Duck",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-duck.png`
					},
					{
						label: "Egg",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-egg.png`
					},
					{
						label: "Earth",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-earth.png`
					},
					{
						label: "Faberge Egg",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-faberge.png`
					},
					{
						label: "Factory",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-factory-dark.png`
					},
					{
						label: "Fan",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-fan.png`
					},
					{
						label: "Fence",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-fence.png`
					},
					{
						label: "Film Strip",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-film-strip.png`
					},
					{
						label: "Fir",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-fir.png`
					},
					{
						label: "Fire Hydrant",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-firehydrant.png`
					},
					{
						label: "Flamingo",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-flamingo.png`
					},
					{
						label: "Flower",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-flower.png`
					},
					{
						label: "Fox",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-fox.png`
					},
					{
						label: "Frog",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-frog.png`
					},
					{
						label: "Garlic",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-garlic.png`
					},
					{
						label: "Gavel",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-gavel.png`
					},
					{
						label: "Ghost",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ghost-B.png`
					},
					{
						label: "Glasses",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-glasses-big.png`
					},
					{
						label: "Gnome",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-gnome.png`
					},
					{
						label: "Goat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-goat.png`
					},
					{
						label: "Gold Coin",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-goldcoin.png`
					},
					{
						label: "Goldfish",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-goldfish.png`
					},
					{
						label: "Grouper",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-grouper.png`
					},
					{
						label: "Hair",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-hair.png`
					},
					{
						label: "Hanger",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-hanger.png`
					},
					{
						label: "Handheld Console",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-console-handheld.png`
					},
					{
						label: "Hard Hat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-hardhat.png`
					},
					{
						label: "Heart",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-heart.png`
					},
					{
						label: "Helicopter",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-helicopter.png`
					},
					{
						label: "High Heel",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-highheel.png`
					},
					{
						label: "Hockey Punk",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-hockeypuck.png`
					},
					{
						label: "Horse",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-horse-deepfried.png`
					},
					{
						label: "Hotdog",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-hotdog.png`
					},
					{
						label: "House",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-house.png`
					},
					{
						label: "Icepop",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-icepop-b.png`
					},
					{
						label: "Igloo",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-igloo.png`
					},
					{
						label: "Index Card",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-indexcard.png`
					},
					{
						label: "Island",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-island.png`
					},
					{
						label: "Jellyfish",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-jellyfish.png`
					},
					{
						label: "Jupiter",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-jupiter.png`
					},
					{
						label: "Kangaroo",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-kangaroo.png`
					},
					{
						label: "Ketchup",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ketchup.png`
					},
					{
						label: "Laptop",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-laptop.png`
					},
					{
						label: "Lightning Bolt",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-lightning-bolt.png`
					},
					{
						label: "Lint",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-lint.png`
					},
					{
						label: "Lips",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-lips.png`
					},
					{
						label: "Lipstick",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-lipstick2.png`
					},
					{
						label: "Lock",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-lock.png`
					},
					{
						label: "Macaroni",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-macaroni.png`
					},
					{
						label: "Mailbox",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mailbox.png`
					},
					{
						label: "Maze",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-maze.png`
					},
					{
						label: "Microwave",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-microwave.png`
					},
					{
						label: "Milk",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-milk.png`
					},
					{
						label: "Mirror",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mirror.png`
					},
					{
						label: "Mixer",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mixer.png`
					},
					{
						label: "Moon",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-moon.png`
					},
					{
						label: "Moose",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-moose.png`
					},
					{
						label: "Mosquito",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mosquito.png`
					},
					{
						label: "Mountain",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mountain-snowcap.png`
					},
					{
						label: "Mouse",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mouse.png`
					},
					{
						label: "Mug",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mug.png`
					},
					{
						label: "Mushroom",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mushroom.png`
					},
					{
						label: "Mustard",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-mustard.png`
					},
					{
						label: "Nigiri",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-nigiri.png`
					},
					{
						label: "Noodles",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-noodles.png`
					},
					{
						label: "Onion",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-onion.png`
					},
					{
						label: "Orangutan",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-orangutan.png`
					},
					{
						label: "Orca",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-orca.png`
					},
					{
						label: "Otter",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-otter.png`
					},
					{
						label: "Outlet",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-outlet.png`
					},
					{
						label: "Owl",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-owl.png`
					},
					{
						label: "Oyster",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-oyster.png`
					},
					{
						label: "Paintbrush",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-paintbrush.png`
					},
					{
						label: "Panda",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-panda.png`
					},
					{
						label: "Paperclip",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-paperclip.png`
					},
					{
						label: "Peanut",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-peanut.png`
					},
					{
						label: "Pencil",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pencil-tip.png`
					},
					{
						label: "Peyote",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-peyote.png`
					},
					{
						label: "Piano",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-piano.png`
					},
					{
						label: "Pickle",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pickle.png`
					},
					{
						label: "Pie",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pie.png`
					},
					{
						label: "Piggybank",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-piggybank.png`
					},
					{
						label: "Pill",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pill.png`
					},
					{
						label: "Pillow",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pillow.png`
					},
					{
						label: "Pineapple",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pineapple.png`
					},
					{
						label: "Pipe",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pipe.png`
					},
					{
						label: "Pirateship",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pirateship.png`
					},
					{
						label: "Pizza",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pizza.png`
					},
					{
						label: "Plane",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-plane.png`
					},
					{
						label: "Pop",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pop.png`
					},
					{
						label: "Potato",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-potato.png`
					},
					{
						label: "Pufferfish",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pufferfish.png`
					},
					{
						label: "Pumpkin",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pumpkin.png`
					},
					{
						label: "Pyramid",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-pyramid.png`
					},
					{
						label: "Rabbit",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rabbit.png`
					},
					{
						label: "Rainbow",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rainbow.png`
					},
					{
						label: "Rangefinder",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rangefinder.png`
					},
					{
						label: "Raven",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-raven.png`
					},
					{
						label: "Retainer",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-retainer.png`
					},
					{
						label: "RGB",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rgb.png`
					},
					{
						label: "Ring",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ring.png`
					},
					{
						label: "Road",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-road.png`
					},
					{
						label: "Robot",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-robot.png`
					},
					{
						label: "Rock",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rock.png`
					},
					{
						label: "Rosebud",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-rosebud.png`
					},
					{
						label: "Rubber Ducky",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ducky.png`
					},
					{
						label: "Saguaro",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-saguaro.png`
					},
					{
						label: "Sailboat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-sailboat.png`
					},
					{
						label: "Sandwich",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-sandwich.png`
					},
					{
						label: "Saturn",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-saturn.png`
					},
					{
						label: "Saw",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-saw.png`
					},
					{
						label: "Scorpion",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-scorpion.png`
					},
					{
						label: "Shark",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-shark.png`
					},
					{
						label: "Shower",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-shower.png`
					},
					{
						label: "Skateboard",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-skateboard.png`
					},
					{
						label: "Skeleton",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-skeleton-hat.png`
					},
					{
						label: "Skilift",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-skilift.png`
					},
					{
						label: "Smile",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-smile.png`
					},
					{
						label: "Snowglobe",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-snowglobe.png`
					},
					{
						label: "Snowman",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-snowmobile.png`
					},
					{
						label: "Snowmobile",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-snowmobile.png`
					},
					{
						label: "Spaghetti",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-spaghetti.png`
					},
					{
						label: "Speech Bubble",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bubble-speech.png`
					},
					{
						label: "Sponge",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-sponge.png`
					},
					{
						label: "Squid",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-squid.png`
					},
					{
						label: "Stapler",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-stapler.png`
					},
					{
						label: "Star",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-star-sparkles.png`
					},
					{
						label: "Steak",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-steak.png`
					},
					{
						label: "Sunset",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-sunset.png`
					},
					{
						label: "Taco",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-taco-classic.png`
					},
					{
						label: "Taxi",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-taxi.png`
					},
					{
						label: "Thumbs Up",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-thumbsup.png`
					},
					{
						label: "Toaster",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-toaster.png`
					},
					{
						label: "Toilet Paper",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-toiletpaper-full.png`
					},
					{
						label: "Tooth",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-tooth.png`
					},
					{
						label: "Toothbrush",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-toothbrush-fresh.png`
					},
					{
						label: "Tornado",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-tornado.png`
					},
					{
						label: "Trash Can",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-trashcan.png`
					},
					{
						label: "Treasure Chest",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-treasurechest.png`
					},
					{
						label: "Triangular Ruler",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ruler-triangular.png`
					},
					{
						label: "Turing Machine",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-turing.png`
					},
					{
						label: "UFO",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-ufo.png`
					},
					{
						label: "Undead",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-undead.png`
					},
					{
						label: "Unicorn",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-unicorn.png`
					},
					{
						label: "Vending Machine",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-vendingmachine.png`
					},
					{
						label: "Vent",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-vent.png`
					},
					{
						label: "Void",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-void.png`
					},
					{
						label: "Volcano",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-volcano.png`
					},
					{
						label: "Volleyball",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-volleyball.png`
					},
					{
						label: "Wall",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wall.png`
					},
					{
						label: "Wallet",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wallet.png`
					},
					{
						label: "Wallsafe",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wallsafe.png`
					},
					{
						label: "Washing Machine",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-washingmachine.png`
					},
					{
						label: "Watch",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-watch.png`
					},
					{
						label: "Watermelon",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-watermelon.png`
					},
					{
						label: "Wave",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wave.png`
					},
					{
						label: "Weed",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-weed.png`
					},
					{
						label: "Weight",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-weight.png`
					},
					{
						label: "Werewolf",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-werewolf.png`
					},
					{
						label: "Whale",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-whale-alive.png`
					},
					{
						label: "Wine",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wine.png`
					},
					{
						label: "Wine Barrel",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-winebarrel.png`
					},
					{
						label: "Wizard Hat",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-wizardhat.png`
					},
					{
						label: "Yeti",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-bigfoot-yeti.png`
					},
					{
						label: "Zebra",
						value: "0x03",
						imagePath: `${HEAD_PATH}head-zebra.png`
					}
				],
				[GLASSES_SELECTOR]: [
					{
						label: "Black w/ Red Eyes",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-black-eyes-red.png`
					},
					{
						label: "Black RGB",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-black-rgb.png`
					},
					{
						label: "Black",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-black.png`
					},
					{
						label: "Blue",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-blue-med-saturated.png`
					},
					{
						label: "Deep Teal",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-blue.png`
					},
					{
						label: "Frog Green",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-frog-green.png`
					},
					{
						label: "Full Black",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-fullblack.png`
					},
					{
						label: "Green & Blue",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-green-blue-multi.png`
					},
					{
						label: "Grass",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-grey-light.png`
					},
					{
						label: "Grey",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-grey-light.png`
					},
					{
						label: "Guava",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-guava.png`
					},
					{
						label: "Hip Rose",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-hip-rose.png`
					},
					{
						label: "Honey",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-honey.png`
					},
					{
						label: "Magenta",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-magenta.png`
					},
					{
						label: "Orange",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-orange.png`
					},
					{
						label: "Pink & Purple",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-pink-purple-multi.png`
					},
					{
						label: "Red",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-red.png`
					},
					{
						label: "Smoke",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-smoke.png`
					},
					{
						label: "Teal",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-teal.png`
					},
					{
						label: "Watermelon",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-watermelon.png`
					},
					{
						label: "Yellow & Orange",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-yellow-orange-multi.png`
					},
					{
						label: "Yellow",
						value: "0x03",
						imagePath: `${GLASSES_PATH}glasses-square-yellow-saturated.png`
					}
				]
			}
		],
		sentence: "Noun has {0} of {0=>1}",
		info: "Check if a Noun has a specific trait. This can be used to automatically bid on a Noun with your preferred attributes.",
		icon: ShirtIcon,
		primary: true
	},
	isTokenId: {
		address: zeroAddress,
		abi: abis.nouns.isTokenId,
		inputs: parseAbi([abis.nouns.isTokenId])[0]["inputs"],
		options: undefined,
		sentence: "Noun in auction is token {0}",
		info: "Check if a Noun has a specific token ID. This can be used to automatically bid on a specific Noun.",
		icon: HashIcon,
		primary: true
	}
}

export default nouns
