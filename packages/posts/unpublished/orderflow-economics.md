---
tags: knowledge
title: Economics of Blockchain Orderflow
slug: economics-of-blockchain-orderflow
image: /cdn/papers/economics-of-blockchain-orderflow.png
description: How value accrues in crypto markets through the lens of orderflow economics - from wallets to dexes and everything in between.
created: 04/08/2025
author: nftchance
---

[COMMENT: add section about incentivized order flow - Layer 3 quests, Boost, Divvi]

# Economics of Blockchain Orderflow

TLDR: In blockchain markets, value extraction follows a hierarchical pattern where those closest to users (wallets) can squeeze margins from every layer below them. This creates a recurring market structure where aggregators compress margins for protocols, resulting in commoditization at the protocol level. Understanding this market structure is crucial for builders and investors to strategically position products where value actually accrues, rather than just building "better technology" in a vacuum.

## The Hidden Power Structure of Crypto Markets

When someone makes a swap in crypto, they're participating in a complex value extraction system that most users never see. The question of who captures value in this system isn't just academic—it dictates which products become sustainable businesses and which are relegated to commodity status.

In traditional finance, we've seen decades of evolution where market makers pay for retail order flow through brokers like Robinhood. This same pattern is emerging in crypto, though with unique characteristics shaped by the open and permissionless nature of blockchains.

What makes this particularly fascinating is the recursive nature of the squeeze. Each layer in the stack extracts maximum value from all layers beneath it, with those at the top enjoying the most favorable position. The market forces at work here aren't crypto-specific—they're manifestations of universal economic principles around commoditization and value capture that we've seen play out across industries for centuries.

Let's break down where value actually accrues in crypto markets today:

![Value Extraction Hierarchy](https://cdn.onplug.io/posts/economics-of-blockchain-orderflow/value-extraction-hierarchy.png)

## The Wallet Advantage

Wallets sit at the top of the value chain for a simple reason: they own the user relationship. When Phantom integrates a swap feature directly in their wallet, they're not just improving user experience—they're cutting out swap applications entirely, capturing value that would otherwise go to specialized swap applications.

This creates a powerful position where wallets can:

1. Extract fees directly from user transactions
2. Sell orderflow to aggregators or builders competing for their users
3. Squeeze margins from every layer below them in the stack

A wallet's ability to integrate features that were once the domain of standalone applications creates a gravitational pull that fundamentally shapes the market. This isn't unique to crypto—it's the same dynamic that allows Apple to extract 30% from app developers through the App Store. The pattern is consistent: whoever owns the user relationship has leverage over everyone else.

The wallet advantage is particularly acute in crypto because of the intimate relationship between a wallet and a user's assets. Unlike other software where switching costs might be lower, changing wallets involves a fundamental reconsideration of how you manage your digital assets. This creates stickiness that wallets can leverage to extract more value over time.

We're still early in this evolution, but we're already seeing wallets like Phantom, Rainbow, and MetaMask expand their feature sets to include swaps, bridges, and even primitive financial services. They're not doing this solely for user convenience—they're doing it because they understand the economics of their position in the value chain.

## The Aggregator Squeeze

One layer below wallets sit aggregators—platforms that collect and compare quotes from multiple protocols to find the best price for users. While aggregators appear to benefit users with better prices, they simultaneously create commoditization pressure on the protocols beneath them.

Consider what happens when an aggregator like DefiLlama's swap interface compares prices across multiple DEXes:

1. The aggregator queries multiple DEXes for quotes on the same swap
2. It displays and routes to the cheapest option
3. Each DEX must compete on price alone to win the trade

This creates a classic commodity dynamic where DEXes are forced to minimize fees and maximize efficiency just to remain competitive. As one participant in our discussion noted:

> "I can't have my DEX quote the aggregator a 3% fee without the aggregator choosing a cheaper DEX."

The result? Margin compression for the DEXes themselves. The market structure forces them to compete primarily on price, limiting their ability to differentiate and capture value.

What's particularly interesting about aggregators is their dual nature—they create value for users while simultaneously destroying margin for protocols. This tension creates a marketplace where the best aggregators become valuable businesses despite (or perhaps because of) their role in compressing margins below them.

The brilliance of the aggregator model is that it directly aligns with user incentives. Users naturally want the best price, and aggregators deliver this by creating competition among protocols. This alignment with user interests makes aggregators difficult to displace, even as they squeeze the margins of the protocols whose liquidity they depend on.

## Protocol Commoditization

At the bottom of this hierarchy sit the actual swap protocols—Uniswap, Curve, and others. Despite being the fundamental infrastructure that enables trading, they increasingly find themselves in a commoditized position.

When an aggregator forces protocols to compete primarily on price, it becomes increasingly difficult for protocols to maintain margins. The protocols that built the original infrastructure become commodities, and their ability to capture value decreases over time.

This explains why Uniswap has never "flipped the fee switch" despite years of speculation about token value accrual. The current market structure simply doesn't allow for significant fee extraction at the protocol level without losing market share to competitors.

The tragedy for protocols is that their success in creating liquid, efficient markets directly contributes to their commoditization. The better a protocol becomes at its core function, the more it resembles its competitors from a user's perspective. This creates a perverse incentive where differentiation through technical excellence often accelerates the path to commoditization.

Some protocols attempt to escape this trap through tokenomics or liquidity incentives, but these are often temporary solutions. The fundamental market structure continues to push protocols toward commoditization as aggregators and wallets extract the majority of value.

The protocols that will survive this commoditization pressure are those that can create genuine moats—unique liquidity profiles that can't be easily replicated or specialized functionality that serves niche use cases better than general-purpose alternatives.

## The MEV Pipeline

Maximal Extractable Value (MEV) forms another layer in this value hierarchy. MEV represents the profit that can be extracted by reordering, including, or excluding transactions within a block. In the context of swaps, this often manifests as the value captured by block builders and searchers who identify profitable opportunities from pending transactions.

Wallets again sit in a privileged position here. By controlling the flow of transactions from users, they can sell this orderflow to builders much like Robinhood sells orderflow to market makers. The competitive dynamics among block builders ensure that wallets can extract maximum value.

This creates another recursive squeeze:

1. Wallets sell orderflow to builders
2. Builders compete to offer the best prices to wallets
3. Competition drives down builder margins
4. Value flows upward to the wallet

What makes MEV particularly fascinating is how it creates an entire shadow economy around transaction ordering. This economy has its own competitive dynamics, specialized participants, and value flows that most users never see or understand.

The MEV landscape also demonstrates the constant tension between centralization and decentralization in crypto. While blockchains themselves may be decentralized, the economic forces at work naturally push toward centralization of value capture at the points where users interface with the system.

The sophistication of MEV extraction has evolved dramatically in recent years, from simple arbitrage bots to complex bundle optimization and auction mechanisms. Despite this evolution, the fundamental pattern remains consistent: those who control user access extract the most value.

## Non-Toxic Orderflow: Retail's Hidden Value

One crucial distinction that often goes unrecognized is the difference between toxic and non-toxic orderflow. When sophisticated traders execute transactions, they typically possess information or speed advantages that make their orderflow "toxic" to market makers. Conversely, retail users generally trade for convenience rather than informational advantage, making their orderflow "non-toxic" and therefore more valuable.

This explains why Solana's retail-heavy ecosystem has seen such intense competition for orderflow. Retail users generate the most valuable type of orderflow—trades that don't pose informational risks to market makers. A wallet with predominantly retail users can extract significantly more value from builders and aggregators than one serving sophisticated traders.

The implications are profound: platforms that successfully attract retail users have an inherent advantage in this market structure, potentially extracting more value per transaction than those serving professionals.

What makes this particularly interesting is how it inverts the usual dynamic in financial markets. In traditional markets, institutional trading is often seen as more prestigious and profitable than retail. In crypto, the opposite can be true—retail orderflow often commands a premium because it's more predictable and less likely to cause losses for market makers.

[COMMENT: isn't institutional trading sort of like being the market maker that consolidates the user orders? I think you're trying to get across the idea: Non-toxic retail order flow is more preferred by market makers than privliged whales who may cause market maker losses -- ie: Robinhood generates retail order flow and sells it to Citadel who trades at an institutional level]

This dynamic creates interesting incentives for wallets and applications to focus on simplicity and accessibility rather than catering to sophisticated users with complex feature sets. The most valuable customers [COMMENT: most valuable for who?] might not be the whales with millions in assets, but everyday users making small, predictable trades.

## Intent Markets: The Next Evolution

Intent-based trading represents the next evolution in this market structure. Rather than submitting exact transactions with specific parameters, users express what they want to achieve (their "intent"), and the market executes it in the most optimal way.

This shifts the dynamics in several ways:

1. Intent solvers compete to fulfill user intents most efficiently
2. Wallets can extract value by auctioning these intents
3. The competitive market for intent solving drives efficiency

As one participant noted, this dynamic should theoretically resemble the Robinhood model, where retail intents are particularly valuable. However, the market for crypto intents remains nascent, making current value flows less predictable than established models.

What makes intent-based markets particularly fascinating is how they abstract away even more complexity from users. In a traditional transaction, users specify exactly how they want to execute a trade. With intents, they simply specify what they want to achieve, leaving the how to competitive solvers.

This abstraction creates new layers of potential value extraction:

1. Interfaces where users express intents 
2. Marketplaces where these intents are auctioned
3. Solvers who compete to fulfill intents efficiently

The critical question becomes: who will control the interface where users express these intents? Current trends suggest wallets will maintain their privileged position, but new entrants could disrupt this hierarchy.

Intent-based markets also create interesting questions around user data and privacy. When users express intents rather than submitting specific transactions, they potentially reveal more about their broader financial goals. This data itself becomes valuable, creating another potential axis of value extraction.

The promise of intent-based markets is greater efficiency and better outcomes for users. The reality may be a more complex redistribution of value extraction, with new players emerging to capture portions of the intent solving and routing market.

## Strategic Positioning for Builders and VCs

Understanding this market structure has profound implications for where builders should focus and how VCs should allocate capital. The market structure explanation clarifies why VCs might prefer funding an aggregator over a DEX, or a wallet over an aggregator—each step up the hierarchy offers greater potential for sustainable value capture.

This doesn't mean there isn't opportunity at every layer. A protocol that can create a genuine liquidity moat or an aggregator with significantly better execution can still build a sustainable business. However, they'll likely capture less value than similarly positioned projects higher in the stack.

As one participant eloquently noted:

> "If I am a solo founder with no VC, I'd rather own the aggregator than the DEX for this specific market structure."

The strategic implications here are profound:

1. **For Protocols:** Focus on creating genuine moats through unique liquidity profiles or specialized functionality that can't be easily commoditized.

2. **For Aggregators:** Maximize efficiency to win in a competitive marketplace while exploring ways to move up the stack through direct user relationships.

3. **For Wallets:** Leverage your position at the top of the hierarchy to extract maximum value while continuing to expand feature sets to capture more user activity.

4. **For VCs:** Recognize that pure technical excellence isn't enough—position in the value hierarchy matters more than marginal technical improvements.

The most successful builders will recognize these dynamics and position accordingly, rather than simply building better technology in a position destined for commoditization.

## Disrupting the Structure

The most transformative opportunities in crypto won't come from incremental improvements within the existing market structure but from fundamentally reshaping it. Companies like Amazon, Google, and Airbnb didn't just outcompete others in established markets—they created entirely new market structures that they could dominate.

For crypto builders, this means thinking beyond building marginally better DEXes or aggregators. The truly revolutionary opportunities lie in:

1. Creating new market structures where value accrues differently
2. Finding areas where existing structures break down
3. Developing moats that can't be easily commoditized

This is where real innovation happens—not in building marginally better implementations of existing models, but in rethinking the fundamental structure of how value flows through the system.

What would a genuinely disruptive model look like? Perhaps a system where users collectively own the interfaces they use, short-circuiting the traditional extraction hierarchy. Or a model where intents are matched in private pools before reaching public markets, creating new privacy-preserving mechanisms for trade execution.

The possibilities are endless, but the key insight remains: the biggest opportunities won't come from playing the existing game better, but from changing the game entirely.

The Solana ecosystem offers a fascinating case study here. While currently charging what many consider excessive fees (around 3%) [COMMENT: Who is charging this fee?], the ecosystem remains in early stages. As markets mature, we should expect competition to drive these fees down, just as it did in traditional finance. The market that eventually emerges will likely be more efficient, with razor-thin margins for most participants and concentrated value capture for those who own the user relationship.

## Lessons from Adjacent Industries

Parallels to this market structure exist throughout technology and finance:

1. Apple's App Store extracting 30% from developers
2. Google's search dominance allowing it to capture most ad value
3. Robinhood selling retail orderflow to market makers
4. Cable companies bundling channels and extracting value from content creators
5. Credit card networks extracting fees from merchants desperate for access to consumers

The consistent pattern is that user access becomes the most valuable commodity in any market, with those controlling this access extracting maximum value from all participants desperate to reach those users.

Crypto's permissionless nature doesn't fundamentally change this dynamic—it simply reshapes how the competition plays out. The protocols may be open-source and the networks decentralized, but the economic realities of user acquisition and value extraction remain stubbornly consistent.

What's particularly interesting is how quickly crypto markets converge on these patterns despite their relative youth. The same forces that shaped traditional markets over decades are compressing into years in crypto, creating rapid evolution toward similar end states.

## Beyond Pure Technology

Perhaps the most important insight from this analysis is that technical excellence alone doesn't guarantee success in crypto markets. Building the "best" DEX or the most efficient aggregator won't necessarily create the most valuable business if your position in the market structure limits your ability to capture value.

As one participant reflected:

> "My inclination as a dev is the same—my first instinct is to build the best tech, but if I want people to depend on me for livelihood it would be irresponsible to not learn about business."

This mirrors experiences from poker, where the most successful players aren't necessarily those with the best technical skills, but those who strategically choose which games to play. Similarly, in crypto, choosing the right position in the market structure can be more important than technical excellence within a commoditized position.

The most successful builders in crypto will combine technical understanding with strategic insight. They'll recognize not just how to build better systems, but where to position those systems within the broader economic hierarchy to capture maximum value.

This dual competency—understanding both the technical and economic dimensions of crypto—remains rare. Most builders focus exclusively on technical excellence, missing the critical economic context that determines whether their technically excellent solutions will create sustainable value.

## Liquidity Monopolies and Market Evolution

An interesting question raised in the conversation was whether liquidity monopolies could disrupt this hierarchy. In theory, a protocol with a dominant liquidity position could reverse the usual dynamic, forcing aggregators and wallets to route through them regardless of fees.

The counterargument presented was that crypto assets constantly change, with new assets starting from zero liquidity, making sustained liquidity monopolies difficult to maintain. While this argument has merit, it oversimplifies the reality, particularly for stablecoin pairs and major assets where liquidity tends to concentrate.

Liquidity monopolies can and do exist in crypto, but they're typically transient and confined to specific trading pairs. The permissionless nature of crypto makes perfect monopolies difficult to sustain, as competitors can always fork code and incentivize new liquidity pools.

However, network effects around liquidity are real. The more liquidity a protocol attracts, the better prices it can offer, attracting more volume, which attracts more liquidity. This virtuous cycle creates natural tendencies toward concentration, even in an open ecosystem.

The evolution of these dynamics over time will be fascinating to watch. As markets mature, we might see increasing specialization, with different protocols dominating different niches rather than competing for the same general-purpose liquidity pools.

## ERC-7702 and the Disruption of Wallet Lock-in

[COMMENT: its starting feel long at this point and like we are going into a new line of thinking which is more technical than anything above. Should 7702 stuff be it's own piece?]

A significant disruption to the current orderflow market structure is on the horizon with the implementation of ERC-7702 and its companion standard ERC-7779. These Ethereum standards, expected to launch with the Pectra upgrade in Q1 2025, have the potential to fundamentally alter the dynamics of value capture in the crypto ecosystem.

### Breaking the Wallet Monopoly

Currently, wallets sit at the top of the value extraction hierarchy largely because of user lock-in. Once a user creates an EOA (Externally Owned Account) wallet, migrating to another wallet means either transferring all assets (incurring gas fees) or importing the private key to another application (creating security concerns). This friction creates sticky user relationships that wallets leverage to extract value.

ERC-7702 changes this dynamic by allowing EOAs to temporarily execute as smart contract wallets for single transactions without requiring users to migrate to an entirely new wallet. This capability is further enhanced by ERC-7779, which creates a standardized framework for wallet interoperability, allowing users to seamlessly switch between wallet providers while maintaining all their advanced functionality.

The implications for the current value extraction hierarchy are profound:

1. **Reduced Wallet Leverage**: When users can migrate between wallets in one click without sacrificing functionality or paying gas fees, wallets lose their position of power. The ability to extract fees will face competitive pressure as switching costs approach zero.

2. **Value Shift to User Experience**: Without lock-in as a moat, wallets will need to compete on genuine user experience and feature innovation, potentially shifting value to those who can create the most intuitive and powerful interfaces rather than those who simply acquired users first.

3. **Protocol Resurgence**: As wallet monopolies weaken, protocols may regain some negotiating power, potentially allowing them to retain more of the value they create rather than seeing it all captured by the layers above.

### Cross-Chain Implications

The disruption extends beyond just wallet switching. With ERC-7702 and ERC-7779, users can leverage "chain abstraction" and execute seamless transactions across multiple chains without manually switching networks in their wallet. This creates a new paradigm where multi-chain strategies become accessible to everyday users, further breaking down the walled gardens that have fragmented liquidity.

In this new landscape, the ability to optimize across chains becomes a competitive advantage. Aggregators that can identify the best opportunities across multiple chains will gain prominence, potentially displacing single-chain wallets as the primary interface for users.

### The New Competitive Landscape

The traditional hierarchy may be replaced by a more complex ecosystem where:

1. **Composable Middleware**: Components that work across multiple wallets and chains could capture significant value by being universally accessible.

2. **Interoperable Smart Accounts**: The ease with which users can switch between implementations creates pressure for wallet developers to focus on genuine innovation rather than lock-in.

3. **Cross-Chain Aggregators**: Services that can optimize across chains may emerge as the new power centers, replacing single-chain wallets at the top of the hierarchy.

4. **Open Intent Marketplaces**: Intent-based systems that operate independently of specific wallets could flourish, creating new dynamics for value capture.

What's particularly fascinating is how these standards could effectively democratize the wallet layer, forcing competition on features rather than user acquisition, and potentially redistributing value across the stack in ways that are more aligned with actual utility creation.

[COMMENT: I see how this is related to breaking down the power that we've established Wallets have by owning the relationship with the user]

## The Mature State of Crypto Markets

As crypto markets mature, we should expect:

1. Increased commoditization at the protocol layer
2. Intense competition for user relationships with reduced lock-in effects
3. Margin compression throughout the stack
4. Value redistribution toward genuine utility and UX rather than just user capture
5. Specialized niches where unique value propositions can sustain higher margins

This maturation process will likely mirror what we've seen in traditional finance, where margins compressed over time as markets became more efficient. The 3% fees common in parts of crypto today will eventually seem as antiquated as the high commissions stockbrokers charged before the advent of discount brokers.

The end state might resemble today's financial markets, where:

1. Transaction costs approach zero for common operations
2. Value accrues to those with unique data, user relationships, or specialized capabilities
3. Scale becomes increasingly important for commoditized services
4. Vertical integration allows capturing multiple layers of the value chain

Standards like ERC-7702, ERC-7779, and continued evolution of intent-based trading will reshape these dynamics, potentially creating a more user-centric ecosystem where value flows more equitably to those creating genuine utility rather than simply occupying privileged positions in the value extraction hierarchy.

## Implications for the Future

[COMMENT: Feel like we're doing this list for a second time, further making me think this could be a follow on piece that is more technically inclined]

The market structure insights we've explored have significant implications for how crypto evolves:

1. **For Users:** Increasing efficiency and lower costs as competition intensifies at all layers of the stack.

2. **For Builders:** A strategic imperative to position within the value hierarchy rather than focusing solely on technical excellence.

3. **For Investors:** A framework for evaluating investments based not just on technical quality but on positioning within the value extraction hierarchy.

4. **For the Ecosystem:** Continued tension between the ideals of decentralization and the economic realities of value capture.

The most profound changes will likely come from projects that fundamentally rethink the structure itself. Just as Bitcoin created an entirely new value paradigm, future innovations might create entirely new market structures that redistribute value in unexpected ways.

For now, though, the pattern remains clear: value flows upward to those who control user access, creating recursive margin compression at each layer below.

## Conclusion

The economics of blockchain orderflow reveals a consistent pattern: value flows upward to those who control user access, creating recursive margin compression at each layer below. While this creates challenges for protocols and other infrastructure providers, it also creates clear opportunities for those who can position themselves strategically within this hierarchy.

Understanding this market structure isn't just academic—it's essential for anyone building in the space. The most successful projects won't necessarily be the most technically impressive, but those that position themselves where value naturally accrues in this evolving economic hierarchy.

As the market matures and competition intensifies, this dynamic will only become more pronounced. The winners won't just build better technology—they'll build businesses positioned to capture the value that technology creates.

The true disruptors won't be those who play the existing game better, but those who change the rules entirely. Just as the early internet protocols created value but captured little of it, today's blockchain protocols may find themselves in a similar position unless they can create defensible moats or move up the stack to capture more user relationships directly.

For builders, investors, and users alike, recognizing these patterns is the first step toward navigating them successfully. The market structure isn't fixed—it's evolving constantly as new models emerge and existing players adapt. The opportunities lie not just in understanding today's structure, but in imagining and creating tomorrow's.
