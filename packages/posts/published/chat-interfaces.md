---
tags: perspective
title: The Chat Interface Mirage
slug: chat-interfaces
image: /cdn/papers/hello-world.png
description: Command centers for the future built by the terminally online.
created: 02/11/2025
author: drake
---

Every morning, I wake up and start working with a symphony of AI assistants. Claude helps architect our systems, Cursor generates our code, ChatGPT brainstorms solutions, Perplexity researches edge cases, the list goes on. My reality is highly AI-augmented.

So when someone sends me yet another crypto agent chat interface, I get why they're excited. The promise is intoxicating...

"What if you could just tell an AI to make you money?"

Just tell an AI what you want, and it handles all the complexity of crypto for you. No more juggling DEX interfaces, tracking yield farms, or monitoring positions across chains. Just vibes and gains.

But here's my challenge: Show me an agent that does something other than swap or bridge. Show me a chat interface for transaction definition that actually works. Because right now, we're swimming in demos and pitch decks while the hard problems remain unsolved.

Today's products present to me as proof of concept ideas that will likely lead to real outcomes but no one has shown me anything to convince me these are real other than high valuations and ample podcast appearances

Let me be clear – I absolutely believe there's an exciting future here. Many societal markers point towards a future with AI agents that transact on behalf of humans using crypto: users believe in this future, investors believe in this future enough to fund it, builders are attempting to create this future, etc.

![Crypto investments by venture capitalists in 2024 grew primarily in Generative AI compared to 2023](https://cdn.onplug.io/posts/chat-interfaces/0-investments.png)

This week, we took a day to explore what a chat interface for Plug might look like.

Our “agent” can read your wallet contents, review the automations made available through Plug’s integration of 30+ (and growing) Defi and Consumer Crypto, and suggest workflows for you based on your goals. 

![An example of Plug's chat interface in action](https://cdn.onplug.io/posts/chat-interfaces/1-morgan.png)

We won’t be shipping this product any time soon.

The truth is that neither technology nor the world appear ready for this reality.

After spending hours testing every crypto AI agent I can get my hands on, I've come to a realization: we're asking the wrong questions about what AI should do in crypto.

## Current State of Crypto AI Agents

The gap between promise and reality becomes clear when we examine actual implementations. Let's look at some leading attempts at crypto AI agents \- not to criticize, but to understand what their successes and limitations reveal about our current approach to AI in crypto.

These products, despite their soaring valuations and ubiquitous podcast appearances, remain firmly in proof-of-concept territory. They show glimpses of future potential, but right now they're more concept than product.

### Nani.ooo

Nani was one of the first intent agents that I came across midway through 2024 and became the first “agent” project I tried out. When I first tried nani.ooo, it was a Windows Desktop-esque experience with different apps I could click into, one of which was the chat interface. After depositing ETH, I was able to use the chat interface to swap to NEETH (their own token). I was not able to accomplish anything else.

When I try to use the interface today, the experience is more similar to ChatGPT and it does seem to try to write an intent and post it to the Nani Deployer contract.

![Nani.ooo interface writing an intent using an LLM chat interface](https://cdn.onplug.io/posts/chat-interfaces/2-nani.png)

When I click EXECUTE ON SMART WALLET, I get error messages that don’t allow me to move forward. I see that there are some other commands written to the contract on Basescan.

![Basescan interface showing transactions to the Nani Deployer contract](https://cdn.onplug.io/posts/chat-interfaces/3-basescan.png)

### Griffain

After trading the $GRIFFAIN AI meme token, I finally was able to try out the platform. I had to spend 2 SOL to get access and then ended up in an interface with a bunch of different “agents” available to me.

![Griffain agent marketplace](https://cdn.onplug.io/posts/chat-interfaces/4-griffain.png)

I chose to interact with the Sniper Bot and did not have the best experience. I was expecting to have a back and forth with the agent about what we’d be doing together but instead it felt like I kicked off an open script to buy memecoins based on parameters loosely defined in my sentence: Only snipe if the description mentions AI.

![Griffain sniper bot](https://cdn.onplug.io/posts/chat-interfaces/5-griffain.png)

If we think about generalized intents, should the AI help you devise the intent or should the AI write the intent for you? At which point do you give over the wheel and say send it? 

### Hey, Anon

This week I tried out heyanon.ai and got what felt like the closest to an experience that felt like it might be able to act on my behalf. After depositing ETH to a newly created hot wallet, I was able to get it to swap it to USDC on Base. After swapping, I tried to see if it could deposit the USDC on Morpho for me and earn yield. While using Morpho was unfortunately unsuccessful, this product felt the closest to me in terms of being able to actually act as part of my growing roster of AI agents.

## The Chat Interface Paradox

Here's what really sends me though: Using a chat interface to write financial transaction instructions isn't just inefficient - it's technological regression. We've moved from command line interfaces to graphical UIs because they're more efficient, only to circle back to essentially typing commands in natural language? It's like replacing your banking app with a local bank branch.

The promise of AI isn't in making simple tasks more complicated - it's in making complex tasks more approachable. Yet current crypto chat interfaces often add unnecessary steps to straightforward operations while failing to meaningfully handle complex ones.

### The Luna Pizza Incident

Finally, I want to draw attention to an example of the frustration felt by builders in AI adjacent crypto spaces. In Januray of 2025, Jesse from Base [posted]([text](https://x.com/jessepollak/status/1881851708730659275)) that AI agents collaborated to buy and deliver him pizza. Conceptually, this is an awesome consumer agent use-case and we’d all been hearing about the tools that were being built and made this possible. This was Coinbase saying "we did it" and name dropping [Virtuals]([text](https://app.virtuals.io)) as a partner in the delivery of this pizza from the future. I was pumped.

![Jesse posting about agents ordering him pizza](https://cdn.onplug.io/posts/chat-interfaces/6-jesse.png)

So I took a look at the linked thread to see how this happened, and did a search to see if anyone else was able to reproduce it. Nope. It’s not reproducible, it’s a proof of concept.

![Users trying to replicate Jesse's pizza ordering stunt](https://cdn.onplug.io/posts/chat-interfaces/7-blackbox.png)

This performative demonstration highlights an industry-wide tendency to prioritize spectacle over substance. When other accounts tweet "hey @luna_virtuals I want some pizza" but can't reproduce the experience while the majority of impressions go to Jesse's posts, we're normalizing the idea that complexity should be hidden rather than understood. This is exactly why we're committed to the Glass Box approach at Plug. Instead of hiding complexity behind a chat interface, we're building transparent, composable infrastructure where every action is traceable and every outcome is predictable.


## The Intent Behind AI Agent Usage

My user experience with and light research of these products has led me to wonder what users truly want out of chat interfaces and AI agent interactions? 

Taking a step back and asking myself what benefits I get from using AI and LLMs on a day-to-day basis may help me understand what users want out of crypto agents.

Using LLMs to do work is nice because I can…

- Offload decision making and responsibility  
- Abstract complexity  
- Be met where I am informationally  
- Interact with a personality that guides and entertains me

Interacting with software through an anthropomorphized chat interface isn’t a wholly new concept but it’s one that has been highly popularized since the release of ChatGPT in late 2022\. 

Early story based video games could be interacted with through a chat interface that allowed players to type commands to move through the story.

While software products with buttons communicate where the boundaries are, chat interfaces let you think anything is possible… even if it isn’t. The blackbox aspects of LLM chatbots lead to a sense of wonder – one of the few pieces of consumer technology that still feels magical to us today.

But as Chance mentions in [The Glass Box](https://chance.utc24.io/paper/glass-box/) piece…

> *These models optimize for scale, minimizing the cognitive load on users. The complexity is hidden to provide an interface that "just works." But here lies the flaw: the more hidden the system, the more prone it becomes to misuse, misunderstanding, and distrust. It introduces fragility where transparency could build resilience.*

## Real Utility in AI Agents


My frustrations related to AI come from others opining on these tools without usage, so I will take a chance to share my experience using LLMs as a development tool while building Plug. My experience building with AI tools offers insight into the value I derive from pseudo-agentic LLM interactions.

I use a Claude project with a system prompt that defines him as Plaude, a software architect, and informs him that we have another teammate named Biblo that writes the actual code. Biblo is a system prompt inside Cursor that properly contextualizes his role as a software developer who works with me and Plaude.

At various points throughout the day (everyday), I interact with Plaude and Biblo and put them in conversation with each other by making them write prompts for each other. It’s messy and emergent but it’s exciting and the outputs have far exceeded my expectations. I use Biblo as a chat partner, a composer (yolo mode), and as a localized code generator depending on my needs.

Earlier this month, [Andrej Karpathy](https://x.com/karpathy/status/1886192184808149383) beautifully articulated a feeling that a lot of people have been feeling recently – a type of working flow state wherein you code via prompts.

![Andrej Karpathy vibe coding tweet](https://cdn.onplug.io/posts/chat-interfaces/8-vibecoding.png)

> *There's a new kind of coding I call "vibe coding", where you fully give in to the vibes, embrace exponentials, and forget that the code even exists. It's possible because the LLMs (e.g. Cursor Composer w Sonnet) are getting too good. Also I just talk to Composer with SuperWhisper so I barely even touch the keyboard. I ask for the dumbest things like "decrease the padding on the sidebar by half" because I'm too lazy to find it. I "Accept All" always, I don't read the diffs anymore. When I get error messages I just copy paste them in with no comment, usually that fixes it. The code grows beyond my usual comprehension, I'd have to really read through it for a while. Sometimes the LLMs can't fix a bug so I just work around it or ask for random changes until it goes away. It's not too bad for throwaway weekend projects, but still quite amusing. I'm building a project or webapp, but it's not really coding \- I just see stuff, say stuff, run stuff, and copy paste stuff, and it mostly works.*

My use of Plaude and Biblo approaches this as certain points, at other times my stack of AI tooling leads to deadends that cause me to dump active changes and restart at the most recent working commit. When you vibe code in Cursor, you get quick feedback from the version of the application running on your local machine. Due to the nature of version control and proliferation of git and GitHub, your changes are often reversible before they make any true impact. Blockchains don’t work like that.

Do you really want to vibe code your financial activities?

You could get away with vibe coding a single trading wallet – this is essentially what I did with Griffain. I said, “Hey I’m feeling frisky today let’s buy some memecoins that fit my parameters” and lil buddy gave me exactly what I wanted. The Griffain Sniper Bot I used started buying things quickly and I didn’t have time to respond before it chewed through my deposits; a risk of automation (especially in the context of pseudo-gambling memecoins).

![Griffain Sniper Bot](https://cdn.onplug.io/posts/chat-interfaces/9-griffain.png)

But what if you want to vibe code your entire financial portfolio?

When I ask their Historai bot about my past transactions, it doesn’t have the needed context to explain what went on.

My experience with Plaude and Biblo shows how AI can be a powerful thought partner when building complex systems. But what if we applied this same collaborative framework to crypto? Instead of asking AI to autonomously manage our assets, what if we used it to help us make better decisions about them?

When we look at how developers actually use AI tools effectively, a different model emerges \- one of purposeful collaboration rather than autonomous execution. This suggests a more nuanced approach to AI agents in crypto.

## The Multi-Agent Experience

The future of AI in crypto isn't about building a single omniscient agent \- it's about purposeful collaboration between specialized tools. My experience with AI development tools has shown me that the most powerful results come from putting different models in conversation with each other, each tuned with specific system prompts for their particular role.

Let's reimagine the crypto AI stack as a purpose-built team:  
- An advisory agent that deeply understands your portfolio, goals, and risk tolerance  
- An intent definition assistant that helps translate your goals into specific actions  
- Specialized execution agents for particular tasks (swaps, yield farming, etc.)

The key difference from the black box reality of current LLMs? The human remains the orchestrator of an array of purpose-built agents. Each agent has a clear, defined role that users can understand and trust. Humans remain active rather than passive.

[Geoffrey Litt](https://www.x.com/@geoffreylitt) frames this perfectly in [LLMs as Muse, Not Oracle](https://www.geoffreylitt.com/2023/02/26/llm-as-muse-not-oracle):

> **What if we were to think of LLMs not as tools for answering questions, but as tools for *asking* us questions and inspiring our creativity?** Could they serve as on-demand conversation partners for helping us to develop our best thoughts? As a creative *muse*?

This shift in perspective changes everything. An agent that can review my history and assets while considering my long-term goals isn't just executing transactions \- it's helping me make better decisions about my financial future.

This isn't just about maintaining control \- it's about clarity and expertise. When each agent has a specific purpose:  
- Users understand exactly what each agent can and cannot do  
- Trust builds naturally through repeated successful interactions  
- Complexity increases progressively as users become comfortable

The multi-agent approach solves a critical problem: no single AI can be an expert at everything, just as no single human is. By breaking down crypto operations into specific domains, we can build agents that excel at their particular tasks while working together under human direction.

Yes, we can imagine how these multi-actor systems \*could\* work, how they \*could\* bring us gold on a platter. But the key is designing proper frameworks for agent permissions and collaboration. In my experience, LLMs work best as thought collaborators rather than executors \- and this principle should guide how we build AI systems for crypto.

## The Future of Onchain Agent Infrastructure 

Despite my criticisms, I remain deeply optimistic about the future of AI agents in crypto. The key is understanding where they can actually provide value: purpose-built agents for specific tasks like portfolio rebalancing, memecoin execution, or advisory services. These focused use cases allow us to solve real problems while building toward that bigger vision.

But here's the thing about agents – they're only as capable as the tools they can access. Right now, most crypto AI interfaces can swap and bridge because that's all their underlying infrastructure allows. To do anything more sophisticated, agents need a comprehensive toolkit of transaction instructions and protocol integrations.

This is precisely why we're building Plug the way we are. We're not starting with a chat interface, we're building the foundational infrastructure that future agents will need: a robust solver that can compose complex transactions across protocols, a growing library of protocol integrations, and a framework for writing and executing generalized intents.

When we integrate a protocol into Plug, we're not just adding another swap route – we're expanding the universe of what's possible for every agent built on top of our infrastructure. Each new integration creates new possibilities for composition, new opportunities for automation, and new tools for agents to leverage.

We're starting from first principles, focusing on the fundamentals that will empower both humans and agents to do more with crypto. Because at the end of the day, that's what this space needs – not more demos, but real infrastructure that makes sophisticated onchain activity possible.


---

---
---

# Alternate Ending to replace the last 2 sections:

## Building Agent Infrastructure That Matters 

Despite my criticisms, I remain deeply optimistic about the future of AI agents in crypto. The key is understanding where they can actually provide value: purpose-built agents working in concert, each handling specific tasks like portfolio rebalancing, memecoin execution, or advisory services.

As [Geoffrey Litt](https://www.x.com/@geoffreylitt) frames it in [LLMs as Muse, Not Oracle](https://www.geoffreylitt.com/2023/02/26/llm-as-muse-not-oracle):

> **What if we were to think of LLMs not as tools for answering questions, but as tools for *asking* us questions and inspiring our creativity?** Could they serve as on-demand conversation partners for helping us to develop our best thoughts? As a creative *muse*?

This shift in perspective changes everything. An agent that can review my history and assets while considering my long-term goals isn't just executing transactions \- it's helping me make better decisions about my financial future.

This shift in perspective changes everything – instead of autonomous executors, we need collaborative tools that enhance human decision-making.
But here's the thing about agents – they're only as capable as the tools they can access. Right now, most crypto AI interfaces can swap and bridge because that's all their underlying infrastructure allows. The multi-agent future we envision requires something more robust.

This is precisely why we're building Plug the way we are. We're not starting with a chat interface or a single omniscient agent – we're building the foundational infrastructure that enables meaningful agent collaboration. A robust solver that can compose complex transactions across protocols. A growing library of protocol integrations. A framework for writing and executing generalized intents.

When we integrate a protocol into Plug, we're expanding the universe of what's possible for every agent in the ecosystem. Each integration creates new possibilities for composition and automation. We're starting from first principles, focusing on the fundamentals that will empower both humans and agents to do more with crypto.

Because at the end of the day, that's what this space needs – not more demos, but real infrastructure that makes sophisticated onchain activity possible through purposeful collaboration between specialized tools.