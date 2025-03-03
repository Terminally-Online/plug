---
tags: perspective
title: Infra as Alpha
slug: infra-race
image: /cdn/papers/infra-race.png
description: Your ability to transact limits your ability to capture alpha.
created: 03/03/2025
author: drakedanner
---

## The Infra Race

There's a scene in the Big Short where Brownfield Fund visits JPMorgan with the goal of getting an International Swaps and Derivatives Association (ISDA) agreement to be able to transact in long term derivatives. The owners of Brownfield saw an opportunity, but they needed an ISDA to be execute the trade. The only problem? Their AUM of $30m was "about 1 billion 470 million short" of the capital requirement needed to qualify for an ISDA. No ISDA, no trade.

![Image from the Big Short discussing the value of the ISDA](https://cdn.onplug.io/posts/infra-race/2-isda.png)

Success in markets is often viewed through the lens of information and execution speed but the example from The Big Short highlights an oft-overlooked third component: Transaction Infrastructure. Even if you have the best information in the world, you may not be able to execute the trade without the right infrastructure.

While information flows and execution speed are important in crypto, this third component is exacerbated in a way that differs from historical understandings. Sure, we can talk about how to know which tokens to buy or which airdrops to hunt, we generally understand that there are MEV bots that trade faster than we can, but we rarely talk about the infrastructure that allows us to transact in the first place.

I've identified 6 tiers of transaction sophistication within DeFi. Let's dig in.

## Levels of Transacting

When we transact onchain, we're competing with other users for block space. Our ability to capture alpha is limited by our transaction capabilities. Are you able to prepare the transaction required and get it into an optimal block?

But before we get onchain, let's start with the basics.

### Tier 1: Centralized Exchange Users

These users are most likely spot traders or long term holders. They may utilize some of the Earn products that lend or stake their underlying tokens. They may pay a monthly fee for fee-less trading with something like Coinbase.

The vast majority of crypto participants never leave this tier. They're comfortable with familiar interfaces that mirror traditional finance platforms, and they accept the trade-offs: higher fees, centralization risk, and limited access to DeFi's innovation. For many, this is a reasonable choice—they prioritize convenience over sovereignty and potential yield.

### Tier 2: Self-Custody Beginners

The leap from Tier 1 to Tier 2 represents one of the most significant psychological barriers in crypto adoption. Self-custody means taking responsibility for your assets. These users have taken their first steps into true crypto ownership, but they're still paying a premium for simplified interfaces that abstract away the complexity of the underlying protocols or failing to identify the highest upside opportunities.

They might swap or stake inside of the Ledger app or the Metamask extension but they aren't seeking out the best rates across the ecosystem.

### Tier 3: Simple Protocol Users

At this tier, users begin to directly interface with the protocols that power DeFi; they're using Uniswap to swap and Blur to find NFTs. They've learned to connect their wallets to dApps and execute basic transactions. They typically interact with one protocol at a time, however, missing opportunities for composability—one of DeFi's most powerful features. Their transactions are often suboptimal, paying higher fees than necessary and frequently falling victim to MEV extraction.

### Tier 4: Advanced Protocol Users

Tier 4 users have developed a working mental model of DeFi's ecosystem. They understand that liquidity is fragmented across protocols and use aggregators to optimize their trades. They've learned to adjust gas settings and may time their transactions to avoid network congestion. These users capture more value than those in lower tiers, but they're still executing transactions manually and reactively rather than programmatically and proactively.

### Tier 5: Weapons-Grade Protocol Users

This is where the playing field begins to tilt dramatically. Tier 5 users have crossed into professional territory, viewing each transaction as a competitive event. They understand that the mempool is a battlefield where transactions compete for inclusion, and they've equipped themselves with specialized tools to gain an edge such as Flashbots MEV Protect. These users have invested significant time in understanding blockchain mechanics beyond the application layer, allowing them to execute strategies that would be impossible for users in lower tiers.

### Tier 6: Custom Infrastructure Developers

At the apex of transaction sophistication are those who have built custom infrastructure tailored to their specific strategies. These are typically institutions, protocols with dedicated development teams, or sophisticated individual traders with programming backgrounds. Their systems monitor onchain and off-chain data continuously, identifying opportunities and executing complex transaction sequences without human intervention. These actors deploy capital programmatically and the alpha they capture is often invisible to lower tiers until well after the opportunity has passed.

## The Infrastructure Gap

This infrastructure gap represents one of DeFi's most significant barriers to true democratization. While we celebrate the removal of traditional gatekeepers, we've created new ones based on technical capability rather than capital or credentials. The resources required to build and maintain Tier 6 infrastructure are beyond the reach of most individuals and smaller organizations, creating a new form of financial exclusion masked as meritocracy.

The gap is widening, not narrowing. As protocols become more complex and interconnected, the knowledge and infrastructure required to interact with them optimally increases exponentially. Each new innovation—from new L2s to cross-chain bridges to novel financial primitives—adds another layer of complexity that must be integrated into existing systems.

## Democratizing Transaction Sophistication

This is where Plug enters the picture. We've built the infrastructure so you don't have to. Our platform provides the sophisticated transaction capabilities of Tier 6 with the accessibility of Tier 3 or 4. We believe that transaction infrastructure should be a utility, not a competitive advantage available only to the technically elite.

By abstracting away the complexity while maintaining transparency and control, we're enabling users to focus on strategy rather than execution mechanics. Whether you're looking to optimize a simple swap or orchestrate a complex cross-protocol strategy, Plug provides the infrastructure to execute it effectively.

## The Future Belongs to the Infrastructure-Enabled

As DeFi continues to evolve, the advantage will increasingly shift toward those who can execute transactions most effectively. Information asymmetry is rapidly diminishing as news spreads instantly on Twitter, Discord, and Telegram. Speed advantages are being commoditized through better block builders and MEV-Share. The edge that remains is the ability to compose complex transactions that capture value across the fragmented DeFi landscape.

The future of DeFi belongs to those who can harness the full power of its composability. Not just understanding it conceptually, but having the infrastructure to execute on that understanding programmatically and reliably. This is the new frontier of financial advantage—not what you know, but how effectively you can act on what you know.

Don't get left behind in the infrastructure race. Join [Plug](https://www.onplug.io/) today and take your first step toward adopting transaction infrastructure built for you.
