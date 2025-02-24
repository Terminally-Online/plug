---
tags: perspective
title: Chat Interfaces and Crypto Agents
slug: chat-interfaces
image: /cdn/papers/hello-world.png
description: Building the infrastructure required to enable agents to make onchain transactions.
created: 02/17/2025
author: danner
---

Every morning, I wake up and start working with a symphony of AI assistants. Claude helps architect our systems, Cursor generates our code, ChatGPT brainstorms solutions, Perplexity researches edge cases, the list goes on. My reality is highly AI-augmented.

So when someone sends me yet another crypto agent chat interface, I get why they're excited. The promise is intoxicating...

"What if you could just tell an AI to make you money?"

Just tell an AI what you want, and it handles all the complexity of crypto for you. No more juggling DEX interfaces, tracking yield farms, or monitoring positions across chains. Just vibes and gains.

But I'm frustrated because no one has shown me an agent that can do something other than swap or bridge. No one has shown me a chat interface for transaction definition that actually works. Right now, we're swimming in demos and pitch decks while hard problems remain unsolved.

Although it's early, there's an exciting future for these proof of concept ideas. Many societal markers point towards a future with systems that transact on behalf of humans using crypto: users believe in this future, investors believe in this future enough to fund it, builders are attempting to create this future, etc.

![Crypto investments by venture capitalists in 2024 grew primarily in Generative AI compared to 2023](https://cdn.onplug.io/posts/chat-interfaces/0-investments.png)

Last week we budgeted a few hours to explore what a chat interface for Plug might look like. It is scrappy, but it's better than every single market-available "agent" that exists today.

Morgan can read your wallet contents, build and discover Plugs to use, suggest actions based on your goals, run and schedule your Plugs, re-organize the app with a live context-based layout, etc. It only took a ~300 lines of code to significantly improve upon the experience that is readily available today.

![An example of Plug's chat interface in action](https://cdn.onplug.io/posts/chat-interfaces/1-morgan.png)

If you want to try Morgan out someday, sign up for the [Plug waitlist](https://onplug.io).

After testing other crypto AI agents and exploring the chat interface in our product, I've come to a realization: we're asking the wrong questions about what AI should do in crypto.

## Current State of Crypto AI Agents

The gap between promise and reality becomes clear when we examine actual implementations. Let's look at some leading attempts at crypto AI agents to understand what their successes and limitations reveal about our current approach to AI in crypto.

### Nani.ooo

Nani was one of the first intent agents that I came across midway through 2024 and became the first “agent” project I tried out. When I first tried nani.ooo, it was a Windows Desktop-esque experience with different apps I could click into, one of which was the chat interface. After depositing ETH, I was able to use the chat interface to swap to NEETH (their own token).

When I try to use the interface today, the experience is more similar to ChatGPT and it does seem to try to write an intent and post it to the Nani Deployer contract.

![Nani.ooo interface writing an intent using an LLM chat interface](https://cdn.onplug.io/posts/chat-interfaces/2-nani.png)

When I click EXECUTE ON SMART WALLET, I get error messages that don’t allow me to move forward. I see that there are some other commands written to the contract on Basescan.

![Basescan interface showing transactions to the Nani Deployer contract](https://cdn.onplug.io/posts/chat-interfaces/3-basescan.png)

My experience ended here. My understanding is that in it's current state, Nani is more of a set of developer tools than a consumer product.

### Griffain

After trading the $GRIFFAIN AI memecoin, I finally was able to try out the platform. I had to spend 2 SOL to get access and then ended up in an interface with a bunch of different “agents” available to me.

![Griffain agent marketplace](https://cdn.onplug.io/posts/chat-interfaces/4-griffain.png)

I chose to interact with the Sniper Bot and did not have the best experience. I was expecting to have a back and forth with the agent about what we’d be doing together but instead it felt like I kicked off an open script to buy memecoins based on parameters loosely defined in my sentence: Only snipe if the description mentions AI.

![Griffain sniper bot](https://cdn.onplug.io/posts/chat-interfaces/5-griffain.png)

When we think about generalized intents, should the AI help you devise the intent or should the AI write the intent for you? At which point do you give over the wheel and say send it?

## The Chat Interface Oversimplification

Using a chat interface to write financial transaction instructions isn't just inefficient - it's technological regression. We've moved from command line interfaces to graphical UIs because they're more efficient, only to circle back to essentially typing commands in natural language? It's like replacing your banking app with a local bank branch.

The promise of AI isn't in making simple tasks more complicated - it's in making complex tasks more approachable. Yet current crypto chat interfaces often add unnecessary steps to straightforward operations while failing to meaningfully handle complex ones.

## The Intent Behind AI Agent Usage

My user experience with these products has led me to wonder what users truly want out of chat interfaces and AI agent interactions?

Using LLMs in our worfklows adds value because we can…

- Offload decision making and responsibility
- Abstract complexity
- Be met where we are informationally
- Interact with a personality that guides and entertains us

Interacting with software through an anthropomorphized chat interface isn’t a wholly new concept but it’s one that has been highly re-popularized since the release of ChatGPT in late 2022\.

Early story based video games could be interacted with through a chat interface that allowed players to type commands to move through the story.

While software products with buttons communicate where the boundaries are, chat interfaces allow you to think anything is possible… even if it isn’t. The blackbox aspects of LLM chatbots lead to a sense of wonder – one of the few pieces of consumer technology that still feels magical to us today.

But as Chance mentions in [The Glass Box](https://chance.utc24.io/paper/glass-box/):

> _These models optimize for scale, minimizing the cognitive load on users. The complexity is hidden to provide an interface that "just works." But here lies the flaw: the more hidden the system, the more prone it becomes to misuse, misunderstanding, and distrust. It introduces fragility where transparency could build resilience._

By trading an aspect of our autonomy for convenience and wonder when offloading decision making to opaque LLMs, we introduce room for influence and error.

## Real Utility in AI Crypto Agents

In many cases, trading off nuance control to an LLM for speed in output is net positive. Earlier this month, [Andrej Karpathy](https://x.com/karpathy/status/1886192184808149383) beautifully articulated a feeling that a lot of people have been feeling recently – a type of working flow state wherein you code via prompts.

![Andrej Karpathy vibe coding tweet](https://cdn.onplug.io/posts/chat-interfaces/8-vibecoding.png)

> _There's a new kind of coding I call "vibe coding", where you fully give in to the vibes, embrace exponentials, and forget that the code even exists. It's possible because the LLMs (e.g. Cursor Composer w Sonnet) are getting too good. Also I just talk to Composer with SuperWhisper so I barely even touch the keyboard. I ask for the dumbest things like "decrease the padding on the sidebar by half" because I'm too lazy to find it. I "Accept All" always, I don't read the diffs anymore. When I get error messages I just copy paste them in with no comment, usually that fixes it. The code grows beyond my usual comprehension, I'd have to really read through it for a while. Sometimes the LLMs can't fix a bug so I just work around it or ask for random changes until it goes away. It's not too bad for throwaway weekend projects, but still quite amusing. I'm building a project or webapp, but it's not really coding \- I just see stuff, say stuff, run stuff, and copy paste stuff, and it mostly works._

When I use [Cursor](https://www.cursor.com/) and [Claude](https://claude.ai), I sometimes reach this flow state and push features faster than expected. At other times, I find myself in a deadend and have to dump active changes and restart at the most recent working commit.

Vibe coding works in software developments thanks to the ability to revert changes. It's a different story for onchain transactions.

Do you really want to vibe code your financial activities?

You could get away with vibe coding a single trading wallet – this is essentially what I did with Griffain. I said, “Hey I’m feeling frisky today let’s buy some memecoins that fit my parameters” and lil buddy gave me exactly what I wanted.

![Griffain Sniper Bot](https://cdn.onplug.io/posts/chat-interfaces/9-griffain.png)

The Griffain Sniper Bot I used started buying things quickly and I didn’t have time to respond before it chewed through my deposits; a risk of automation (especially in the context of pseudo-gambling memecoins).

My experience using AI in development shows how AI can be a powerful thought partner when building complex systems. What if we applied this same collaborative framework to crypto? Instead of asking AI to autonomously manage our assets, what if we used it to help us make better decisions about them?

A different model emerges \- one of purposeful collaboration rather than autonomous execution. This suggests a more nuanced approach to AI agents in crypto.

## Building Infrastructure for Agents

Right now, most crypto AI interfaces can swap and bridge because that's all their underlying infrastructure allows. The multi-function agentic future we envision requires more work to achieve.

This is precisely why we're building Plug the way we are. We're not starting with a chat interface or a single omniscient agent – we're building the foundational infrastructure that enables meaningful agent collaboration. A robust solver that can compose complex transactions across protocols. A growing library of protocol integrations. A framework for writing and executing generalized intents.

When we integrate a protocol into Plug, we're expanding the universe of what's possible for every agent in the ecosystem. Each integration creates new possibilities for composition and automation. We're starting from first principles, focusing on the fundamentals that will empower both humans and agents to do more with crypto.
