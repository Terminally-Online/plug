---
tags: knowledge
title: A Truly Gasless Experience
slug: a-truly-gasless-experience
image: /cdn/papers/a-truly-gasless-experience.png
description: When it comes to the blockchain, fees are inescapable. Gas fees. App fees. Swap fees. Everything has a fee. But... what if it didn't have to be that way?
created: 04/10/2025
author: nftchance
---

When it comes to the blockchain, fees are inescapable. Gas fees. App fees. Swap fees. Everything has a fee. But... what if it didn't have to be that way?

I myself have paid over $250,000 real United States Dollars purely for gas over the last 5 years. Fees suck. So, we are removing them entirely so that you never have to think about them again.

It's not a small endeavour, but it has an extreme payoff. Thankfully though, in the last six months we have had the insight, knowledge and expertise of the industry leading experts on incentives.

By working with teams that live and breathe "transaction incentives" we have been able to build a system where value flows instantly and respectively. 

A new market structure for transaction execution has been realized: Zero user-side fees at all. With the new incentive vehicles in place, apps are finally equipped with everything needed to adjust the way value flows through their ecosystem. A world where users do not have to pay exorbitant fees is achievable.

Let's dive in.

## The Current System

Until very recently the market structure for blockchain orderflow has followed a hierarchical pattern that creates specific incentive alignments.

Value tends to flow upward through a defined structure where those closest to users capture the most value. This means the traditional hierarchy looks like:

1. **Wallets** (top position, own the user relationship)
2. **Aggregators** (middleware routing to protocols)
3. **Protocols** (foundational layer of onchain functionality)

In this structure, wallets leverage their user relationships to capture as much value and attention as possible. Meanwhile, living a level lower, aggregators create price pressure for protocols beneath them, leading to margin compression and commoditization at the protocol level as we've seen in the DEX (Decentralized Exchange) market. 

As things compress, the wallet has historically been the major winner. This structure made sense in early markets, but creates limitations for user experience and protocol growth that we're now ready to overcome as an industry.

## The Game Changers: New Incentive Mechanisms

So what's changed? Nothing, overnight. Instead, momentum has been quietly building over the last few years and is now starting to peak.

This movement has resulted in an increase in builders' ability to negate blockchain fees, thanks to three distinct models:

1. Automatic Reward Distributor: Pays based on contribution.
2. Action Marketplace: Pays based on fulfillment.
3. Transaction-Level Rewards: Pays based on completion.

The nature of the three are incredibly unique from one another. But, in a fully functional system combine to create an incredible system of incentive orchestration.

### Divvi: The Automatic Rewards Distributor

The newest kid on the block, [Divvi brings something entirely new to the market](https://www.divvi.xyz/blog/apps): Direct impact valuation and compensation.

Divvi makes this possible by letting protocols define what they care about (TVL, transaction volume, unique users) and then automatically rewarding the interfaces that move those metrics. All tracked and verified on-chain.

![Flow chart demonstrating how Divvi helps Web3 builders earn by directing activity to protocols](https://cdn.onplug.io/posts/a-truly-gasless-experience/value-flow-divvi.png)

[When a transaction created in our application is directed through a Divvi supported protocol](https://docs.divvi.xyz/protocol/rewards), we're compensated. This seemingly simple vehicle allows us to internalize the costs of using the blockchain without charging you, the user a gas fee, much less a platform fee.

"Protocols now automatically pay the apps that bring them users" is not something I expected to be able to say this year, but here we are.

No negotiations. No middlemen. Just instant revenue sharing based on actual impact. The protocol captures value they wouldn't have otherwise and you have a costless experience doing so. A rare positive-sum contribution to the space.

### Royco: The Action Marketplace

Having been around a bit longer, Royco brings a different angle to the incentive game with their "Incentivized Action Markets" or IAMs.

Think of it as the first real marketplace for user actions. Protocols list exactly what they want users to do and how much they're willing to pay for it. For example: [Deposit USDC into Rings to Mint and Stake scUSD](https://app.royco.org/market/146/0/0x7d1f2a66eabf9142dd30d1355efcbfd4cfbefd2872d24ca9855641434816a525)

Notably here, both the deposit and stake are happening in the same transaction. At the same time, you have markets like: [Deposit xUSD for 30 days](https://app.royco.org/market/146/0/0xfcd798abefe4f9784e8f7ce3019c5e567e85687235ce0ce61c27271ba97d26cd) where the liquidity deposited cannot be moved for the agreed upon time or you will forfeit rewards.

These actions can be fulfilled by users (or through interfaces like ours) who can then claim the rewards.

![Flow chart demonstrating how Royco's Intent Action Markets can be used by Plug to decrease user costs](https://cdn.onplug.io/posts/a-truly-gasless-experience/value-flow-royco.png)

Incentives are explicit and have specific stipulations that must be followed. It's right there on-chain - "We'll pay X for Y action" - and that transparency changes everything about how we can build user experiences.

Want users to provide liquidity? Bridge assets? Try your new feature? Put it on the marketplace with a clear price tag. No more guessing what actions are valuable - the market establishes that in real-time. It's a beautiful system where there is opportunity outside of typical yield mechanisms such as borrow and lend markets.

Where Royco and IAMs really start to provide value to Plug users, is the system's ability to route transactions through Royco. When you're performing an action and it's possible to get increased rewards, the Solver automatically routes your transaction through the appropriate market to maximize incentive exposure.

### Boost & Layer3: The Transaction-Level Reward Engine

Unlike the newer incentive systems, quest platforms like [Boost](https://boost.xyz/) and [Layer3](https://app.layer3.xyz/quests) have been around longer, quietly perfecting completionist rewards for specific actions.

What many don't realize is how perfectly suited these platforms are for subsidizing transaction costs. While they were originally designed for user acquisition, we've transformed them into a core component of our zero-fee infrastructure.

![Flow chart demonstrating how Boost and Layer3 quest products can be used by Plug to decrease user costs](https://cdn.onplug.io/posts/a-truly-gasless-experience/value-flow-boost.png)

Where questing platforms truly shine for Plug's zero-fee model is their immediate and stackable rewards. While a single quest reward might seem small, our system automatically identifies and claims all eligible quests for every transaction you make. This creates an instant pool of rewards that we use to cover your gas costs on the spot. For quests that we cannot fill in a standard route we can surface the quest to our end-user in a few lines of code and have them automatically capture every fit opportunity.

The economics only work at scale—individually and require immediate fill-and-pay opportunities. A single quest might not cover gas, but when pooled together across thousands of users, these rewards create a sustainable funding source for everyone's transactions.

What's particularly exciting is how we've worked with protocols to create quest programs specifically targeting gas coverage. These programs reward interfaces like Plug for handling user gas fees, creating a sustainable cycle where everyone benefits.

## A Gasless Blockchain Ecosystem

At Plug, we're not just surfacing these apps to our users. Here's how it works:

1. **We Capture Incentives Others Miss**: While most apps focus on one incentive source, we're tapped into all of them simultaneously. When you use Plug, we're collecting rewards from multiple protocols for the same action across as many incentive layers as possible.

2. **We Find Routes That Maximize Incentives**: Our routing algorithm doesn't just find the best price—it finds the path that generates the most incentives. This creates a multiplier effect where a single user action can earn rewards from several sources.

3. **We Use Those Incentives To Cover Your Costs**: Instead of pocketing all these rewards, we allocate a portion to cover your gas fees, bridge costs, and other transaction expenses. The result? Transactions that appear completely free from your perspective.

This isn't theoretical. We've built a system where the protocols themselves are effectively paying for your transactions, all because we're able to capture and redistribute the value correctly. The great thing is in practice it is as simple as it sounds! Let me break down a simple example where you want to swap Token A for Token B.

**Traditional path:**

1.  You pay gas + app fees + slippage
2.  You get Token B

**Plug path:**

1.  We route your transaction through Protocol Y, which pays us an incentive
2.  We use part of that incentive to cover your gas
3.  You get Token B without paying any fees

The future we're building isn't just about eliminating fees. The future we're creating brings an entirely new way for blockchain interactions to work as smoothly as Web2, but with all the benefits of Web3.

The technical implementation has a few nuances and additional mechanisms are needed for this to work at scale, but who cares about that? The point is that everyone wins. You get free transactions. Protocols get the user activity they want. And we have the ability to generate constant onchain activity that's more economically viable than anywhere else.

No more worrying about gas. No more fee shock. Just seamless transactions that happen exactly when you want them to. The future isn't just close, it is here.
