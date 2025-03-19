---
tags: perspective
title: Abstraction Fixes Fragmentation
slug: abstraction-fixes-fragmentation
image: /cdn/papers/abstraction-fixes-fragmentation.png
description: Over the past 4 years, the Ethereum community has been on a journey that began with an attempt to scale Ethereum and has ended in a fragmented mess.
created: 03/17/2025
author: drakedanner
---

### TLDR:
The Ethereum community sentiment is pretty bad right now after watching Solana eat its lunch despite having a massive cultural headstart. I attribute this to cultural fragmentation caused by L2 scaling attempts that tried to build miniature versions of Ethereum with miniature cultural moats. Teams are working on solving this through chain abstraction. I suggest that consumer application accessibility can go beyond chain abstraction by abstracting away protocol interfaces. 

## The EVM in 2025

Over the past 4 years, the Ethereum community has been on a journey that began with an attempt to scale Ethereum and has ended in a fragmented mess made worse by infighting between cohabitating ecosystems. When layer 1 fees became too expensive for day-to-day transactions, we were met with the promise of layer 2s. This scaling ideal has resulted in a proliferation of rollups and EVM compatible alt layer 1s. We have more blockspace than ever before and we've paid for it dearly through diminished user experience. Without ongoing efforts to abstract the chain away from users, the idea that we may onboard new users to Ethereum presents as laughable.


As of March 2025, DefiLlama lists nearly 250 EVM chains.

Fragmented users and fragmented liquidity lead to fragmented experiences. While early Defi saw flows into major protocols, the eventual rise of transaction costs led protocols to incentivize liquidity on newer, cheaper chains. The only issue? Once these incentives dry up, mercenary yield seekers simply jump to the next incentivized chain.

Over and over again we've seen chains launch to fanfare and then fade into the distance as hotter, younger chains began their own protocol incentive pump cycles. As a user, the EVM has become a shell of the promise that it offered as the world computer. How can we support global financial rails and a single settlement layer when engaging with the system requires moving assets across chains, switching chains in your wallet, and reconsolidating assets into the chain that you prefer.

## Culture to User Ratio Mismatch

For a long time, I viewed Ethereum's moat against Solana and other chains as cultural. Ethereum's community of Defi builders, long lasting NFT artifacts, and DAO proliferation were admirable. In fact, I'd argue that this mix of organizations and use cases was so valuable that the L2s sought to recreate them. Every EVM chain wanted to go viral with an NFT drop, Defi TVL metrics, or onchain based organizations.

All of a sudden, instead of a single strong culture, the EVM was supporting many cultures. On the surface, this appears to be a good thing! A playbook has been discovered to build a culture on a chain, whether or not that culture is a moat, however, is vastly different. Many users that I've talked to over the past few years working in crypto product development have been upfront with me that 1) they have a favorite ecosystem 2) they will try out new ecosystems 3) they try new ecosystems for speculative purposes 4) any gains made in the new ecosystem are swiftly returned to their preferred ecosystem. These users are self-admitted mercenaries who have already chosen their onchain home and leave only to bring back the riches they plunder from elsewhere. Unfortunately, we do not have enough users to support this many cultures.

## User Experience Decline

The explosion of new chains and reality of mercenary users initially presents as downside for chains and upside for users. Chains have to deal with the ebbs and flows of user preferences impacting their primary KPIs (token price, TVL, users, tx volume, etc) and users have the opportunity to explore new spaces and access additional yield or airdrop potential. Most users, however, aren't hunting these airdrops or exploring these spaces. Individuals running large bot networks are often the first-in and first-out of these nascent ecosystems while everyday users might show up to fight for crumbs and then end up leaving dust.

As new apps and protocols choose Ethereum scaling solutions to call home, users are forced to add these chains to the stack of ecosystems they interact with. User assets are spread out all over the place resulting in bridging requirements whenever it comes time to make a large transaction or claim an airdrop.

## Failed Differentiation

As a user, I've personally become extremely fatigued by ongoing launches of new ecosystems and shallow attempts to cultivate cultural moats. I miss the old Ethereum -- an expansive world of sophisticated financial transactions, artistic artifacts, and decentralized autonomous organizations. While holes could be poked in all of these, there was an understanding of a shared virtual machine with a single settlement layer before Ethereum fragmented. None of these new Ethereum scaling solutions have brought anything unique to the table that differentiates them greatly to me -- their largest value proposition is cheaper transactions.

## Chain Abstracted Applications

Due to the immutable nature of these blockchains, we must accept the fragmentation that has occurred. We will never coalesce all assets back to the layer 1 and users will continue to maintain asset balances across these chains. After spending the past 4 years building infrastructure, the time of the application has arrived. The opportunity to bring these experiences and assets together exists at the application layer, where the user interacts with these ecosystems.

Many teams believe that abstracting away chain complexity without sacrificing capabilities represent the next evolution of the EVM ecosystem. When interfaces can seamlessly bridge the gaps between chains, users can focus on strategy rather than logistics. When protocols can be composed across chains through a unified interface, we may recover much of what made early Ethereum magical.

## Protocol Abstracted Interfaces

While chain abstraction makes consumer applications more accessible, it is only a step towards a more user friendly onchain experience rather than the destination. Alternative attempts include interfaces that allow access to and composition of many underlying protocols. Thanks to public smart contracts, development teams are free to create new experiences at the application layer without needing to reinvent or even implement anything at the protocol layer. 

I may even go as far as saying that there is no need for further smart contract development – there are no new primitives to be deployed.

## EVM Ecosystem Unification

The EVM remains the greatest technology-based human coordination system the world has ever seen. Let's not diminish its power by forcing users to navigate its fragmented landscape manually. Instead, let's abstract that complexity and put the full potential of the entire EVM ecosystem at their fingertips.