---
tags: perspective
title: Infra as Alpha
slug: infra-race
image: /cdn/papers/hello-world.png
description: Your ability to transact limits your ability to capture alpha.
created: 03/10/2025
author: drakedanner
---

## The Infra Race

There's a scene in the Big Short where Brownfield Fund visits JPMorgan with the goal of getting an International Swaps and Derivatives Association (ISDA) agreement to be able to transact in long term derivatives. The owners of Brownfield saw an oppurtunity, but they needed an ISDA to be execute the trade. The only problem? Their AUM of $30m was "about 1 billion 470 million short" of the capital requirement needed to qualify for an ISDA. No ISDA, no trade.

Success in markets is often viewed through the lens of information and execution speed but the example from The Big Short highlights an oft-overlooked third component: Transaction Infrastructure. Even if you have the best information in the world, you may not be able to execute the trade without the right infrastructure.

While information flows and execution speed are important in crypto, this third component is exacerbated in a way that differs from historical understandings. Sure, we can talk about how to know which tokens to buy or which airdrops to hunt, we generally understand that there are MEV bots that trade faster than we can, but we rarely talk about the infrastructure that allows us to transact in the first place.

I've identified 6 tiers of transaction sophistication within DeFi. Let's dig in.

## Levels of Transacting

When we transact onchain, we're competing with other users for block space. Our ability to capture alpha is limited by our transaction capabilities. Are you able to prepare the transaction required and get it into an optimal block? 

But before we get onchain, let's start with the basics.

### Tier 1: Centralized Exchange Users
- Trade on CEXs like Coinbase/Binance
- Use basic market/limit orders
- Abstracted from onchain mechanics
- Front-run by exchanges themselves

These users are most likely spot traders or long term holders. They may utilize some of the Earn products that lend or stake their underlying tokens. They may pay a monthly fee for fee-less trading with soemthing like Coinbase .

### Tier 2: Self-Custody Beginners
- Off-ramp to self-custody wallets (Ledger, Rainbow)
- Use wallet's built-in swap features
- Pay convenience fees unknowingly
- No direct protocol interaction

### Tier 3: Simple Protocol Users
- Interact directly with protocols:
  - Uniswap for swaps
  - Opensea for NFTs
- Higher fees than aggregators
- Limited to single-protocol actions

### Tier 4: Advanced Protocol Users
- Compare rates across lending markets
- Use aggregators for better pricing
- Basic understanding of gas optimization
- Beginning awareness of transaction inefficiencies

### Tier 5: Weapons-Grade Protocol Users
- Understand Ethereum's "Dark Forest" mempool
- Use private RPCs to avoid MEV attacks
- Leverage Flashbots for priority transactions
- Set up Telegram bots for trading opportunities
- Treat transacting as a competitive sport

### Tier 6: Custom Infrastructure Creators
- Build proprietary execution systems
- Create monitoring for data feeds
- Develop specialized transaction bundling
- Automate cross-protocol strategies
- Deploy capital programmatically

## The Infrastructure Gap

- Most DeFi users stuck in Tiers 1-4
- Massive leap from Tier 4 to Tier 5
- Even bigger jump to Tier 6
- Custom infrastructure is prohibitively expensive:
  - Requires deep protocol knowledge
  - Sophisticated development resources
  - Private infrastructure
  - Constant maintenance
  - Significant time investment
- Creates uneven playing field

## Democratizing Transaction Sophistication

- Plug provides Tier 6 infrastructure without engineering requirements
- Features include:
  - Constraint-driven execution
  - Cross-protocol composition
  - MEV protection
  - Scheduled strategies
  - Transparent transaction mechanics

## The Future Belongs to the Infrastructure-Native

- Gap will widen between those with sophisticated infrastructure and those without
- Information advantage becoming less significant
- Transaction infrastructure becoming the durable competitive edge
- Call to action: Join waitlist at onplug.io