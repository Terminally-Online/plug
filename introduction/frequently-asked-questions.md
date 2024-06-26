---
head:
  - - meta
    - property: og:title
      content: Frequently Asked Questions
  - - meta
    - name: description
      content: Explore the FAQ of the Plug framework and library.
  - - meta
    - property: og:description
      content: Explore the FAQ of the Plug framework and library.
---

# Frequently Asked Questions

<span style="color: rgba(0,0,0,0.6)">While [Plug](/) is rather simple to use, it's a new user pattern and that may lead to questions that you have never had to ask before. It is understandable that you may have some questions, so below you can find a general breakdown of the questions we've most commonly been asked.</span>

## What is an Intent?

A metaphor will establish this understanding in the simplest way for you. An intent, is a declaration of outcomes that you'd like to see. When you type into Google Maps you only enter the address, not the route you will take. Intents are functionally the same. Contrary to a typical transaction where you declare the exact route you want to take, with an intent you only define the final destination. This way, instead of you managing the complexity a more sophisticated system can calculate the best route and get you there on the quickest path, that has no traffic, avoids construction, and keeps you from being late to your date.

## What is a Solver doing?

When using intents you only declare your outcomes (the final destination). Of course, someone or something still needs to build the transaction that delivers what you're looking for. This is the job of a Solver. There are two distinct types of Solver operations:

1. **Person to Person:** For transactions like swapping tokens you have the ability to avoid direct protocol integration and work with other individuals that have open orders to fill one another. Imagine the simple example where Alice wants to swap 400 **$USDC** to **$ETH** and Bob wants to swap 0.09 **$ETH** to **$USDC** . There are two people that have assets the other wants. This is referred to as a Coincidence of Wants. A Solver works to identify these opportunities and route the assets from one party to another so that each get the best price while having the lowest settlement cost possible.
2. **Person to Protocol:** For transactions that are not swaps, a Solver works to deliver you the best outcomes, but utilizes the state of the ecosystem rather than a specific person. Instead of finding you an order that gives you the best price, if you say **I want to enter a liquidity pool on Uniswap above 80% APY** a Solver will work to deliver that exact outcome.

Put simply, a Solver solves for the transaction that will give you at minimum the outcomes you are looking for, often outperforming the expectations set.

## Are my transactions and Intents private?

Privacy of your onchain activity is important, but delivering your desired outcome is more important. Plug routes through MEV blocker mechanisms when possible, but that's not an option on every chain. For the most part, the lack of privacy has one primary risk: _frontrunning_. Risks related to frontrunning are largely limited to token swaps. For typical protocol actions, there is no upside in frontrunning. Often, there is only downside in having done so. In cases where privacy cannot be fully achieved, pre-confirmation agreements with market makers and Solvers enable you to lock in the price of your asset acquisition completely escaping the concern of being frontrun.

## What are the associated fees?

First things first, you only pay a fee on successful execution. You do not pay for failed transactions. You do not pay when you declare your Intent. Instead, when executing an Intent there is an associated fee that covers that cost of Solving and persistent simulation needed to deliver you the best outcomes. Fees are not static. The fees vary across each chain, type of action, and level of compute consumption, in an effort to mirror the dynamics of your typical fee expenditure. Fees as a whole are designed to only result in a marginal increase in cost while delivering a substantially better result. Due to the dynamic nature, you can always find the fee when you declare an intent.

## Do I lose custody of my assets?

No, users do not lose custody of their assets. Users can even take their Intents and execute it themselves if they prefer. Assets are routed through an onchain contract that has been abstracted to approach “backend router” functionality more than a typical wallet or smart account.

## What wallets are supported?

[Metamask](https://metamask.io/), [Rainbow](https://rainbow.me/), [Coinbase](https://www.coinbase.com/wallet/), [Safe](https://safe.global/), [Rabby](https://rabby.io/), etc. We have broad support for [Wallet Connect](https://walletconnect.com/) and our smart contracts support both EOAs and Smart Accounts. If you can sign into other protocols with your wallet, you can use your wallet with Plug.

## What protocols and chains are supported?

Effectively everything. Due to the generalized nature of our architecture there are very few things that cannot be integrated. You can find a full list of our [active integrations here](/introduction/integrations).

## Wen token?

When the flywheel is spinning hard. No sooner. No later. A critical piece of having a functional token ecosystem is to have a mechanism that provides persistent value accrual. Today, that remains the north star for all tokenization conversations.

## Why haven't you built you built X yet?

We build what users benefit from most. Please don't assume if we haven't implemented something yet it is due to a lack of interest. We index very highly on active feedback and requests from our users. If you have a feature or protocol integration request (**and even general feedback**) [please reach out and tell us](https://twitter.com/onplug_io).

## Is there a public roadmap for what's coming next?

No, there isn't. As a team we do not follow "agile" practices as they create a system that is anything, but agile. Planning every single step to excruciating detail is paradoxical state of software development. We are not here to operate based on predefined objectives. We focus on what improves your experience most. At a higher level we operate with quarterly objectives on top of multi-year north stars.
