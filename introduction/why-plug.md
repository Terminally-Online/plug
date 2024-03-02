---
head:
  - - meta
    - property: og:title
      content: Why Plug
  - - meta
    - name: description
      content: A brief preamble on why Plug was built.
  - - meta
    - property: og:description
      content: A brief preamble on why Plug was built.
---

# Why Plug

In the evolving landscape of blockchain technology, the limitations of native EVM transactions have become increasingly apparent. Despite the advent of smart accounts and alternative mempools, the quest for a solution that seamlessly combines composability, extensibility, and modernity has remained largely unfulfilled. This gap in the ecosystem not only results in millions of dollars lost to failed transactions and inefficiencies annually but also stifles the potential for innovation and growth.

## The Persistent Challenges

A decade into the blockchain revolution, users and developers alike continue to grapple with several fundamental issues:

- `Authorization Complexity`: The current model treats authorization on a contract-by-contract basis, leading to a fragmented security landscape. This inconsistency complicates managing permissions across different contracts, hindering seamless interaction within the ecosystem.

- `Sequential Transaction Processing`: The traditional EVM model processes transactions one after another, limiting throughput and exacerbating bottlenecks, especially during peak times. This linear approach fails to leverage the potential for parallel processing, crucial for scaling blockchain applications.

- `Preemptive Costs`: The requirement to pay transaction fees upfront, without assurance of success, places a significant burden on users. This gamble on transaction outcomes introduces unnecessary risk and inefficiency into the system.

- `Scalability Hurdles`: Current methods for managing conditions are cumbersome and not conducive to scalability. As protocols evolve and expand, the lack of a streamlined approach to condition management becomes a critical bottleneck.

- `Vulnerability to State Changes`: Once submitted, transactions are at the mercy of fluctuating network conditions and potential vulnerabilities within smart contracts. This lack of protection exposes users to risks like front-running and transaction manipulation.

## Introducing Plug: A Paradigm Shift

Recognizing these challenges, Plug was conceived as a revolutionary framework designed to redefine EVM transactions. By prioritizing composability, extensibility, and modernity, Plug addresses the inefficiencies plaguing the blockchain space, offering a more logical, efficient, and user-centric approach to transaction management.

## The Declarative Difference

At its core, Plug introduces a shift from passive to active participation in the transaction creation process. Users gain the ability to set precise conditions for transaction execution, transforming the model from rigid and immediate to flexible and strategic.

This not only mitigates the risks associated with upfront costs but also optimizes gas usage and enhances strategic planning capabilities resulting in the best execution outcomes possible.

## Harnessing "If This, Then That" Logic with Plug

At the heart of Plug's innovation is the application of "If This, Then That" (IFTTT) logic, a powerful and intuitive concept that revolutionizes how transactions are executed on the blockchain. This logic allows users to create conditional statements that dictate the execution of transactions, ensuring actions are only taken when specific criteria are met. Here's a closer look at how IFTTT principles empower Plug users:

### Understanding IFTTT in Plug

IFTTT logic in Plug enables users to define triggers ("If This") and actions ("Then That") within the blockchain environment. This approach transforms passive transaction execution into an active, conditional process, where transactions are executed based on real-time data and predefined conditions.

- `Triggers` ("If This"): These are the conditions or events that must occur for a transaction to be initiated. Triggers can be based on a wide range of criteria, such as time-based conditions, market fluctuations, contract states, or external data inputs. For example, a trigger could be set for when a specific cryptocurrency reaches a certain price point.

- `Actions` ("Then That"): These are the transactions or operations that are executed when the trigger conditions are met. Actions can range from simple token transfers to more complex contract interactions, such as executing trades, minting tokens, or triggering smart contract functions.

### The Power of Onchain IFTTT

By integrating IFTTT logic, Plug offers several transformative benefits to blockchain operations:

- `Automated Efficiency`: Users can automate complex strategies and operations, reducing the need for constant monitoring and manual execution. This not only saves time but also ensures that opportunities are never missed due to delays or human error.

- `Strategic Execution`: IFTTT logic allows for strategic planning and execution of transactions. Users can set conditions to optimize for cost (e.g., gas fees), timing, and market position, ensuring that actions are taken under the most favorable circumstances.

- `Enhanced Security`: Conditional execution means transactions are only processed when all criteria are securely met, reducing the risk of unfavorable or unintended outcomes. This adds an extra layer of security and control over blockchain interactions.

- `User Empowerment`: Plug democratizes access to advanced blockchain functionalities, enabling users without extensive technical expertise to create and manage complex transaction conditions. This opens up new possibilities for innovation and participation in the blockchain space.

The applications of IFTTT logic in Plug are vast and varied, ranging from financial transactions, such as automated trading and dynamic pricing models, to operational tasks, like conditional access control and automated governance decisions.

Essentially, Plug's use of IFTTT logic makes the blockchain more adaptable, responsive, and aligned with users' specific needs and strategies.

## Developer Experience

Plug was developed with one thing in mind: **cost to launch.** Too much time is wasted in the crypto development industry by reinventing the wheel and solving complex problems that have not only been solved, but had their answers shared far and wide.

To accomplish this, Plug is around the core concepts of:

- `Types First`: Contrary to what you may expect, type generation and declaration for `Plug` starts with Solidity in the shape of [EIP-712 Type Declarations](https://eips.ethereum.org/EIPS/eip-712#definition-of-hashstruct). The simulation and execution layer of Plug can be updated and iterated upon without impacting the underlying primitives being consumed by end-users.

- `Global Interoprability`: Plug is designed to support all smart contracts regardless of when they were deployed. Without a need for native integration, Plug lives on top of all the existing protocols that exist within the ecosystem allowing core primitives to completely remove the need for conditional validation at every layer enabling genuine [seperation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns).

- `Composability`: With the typical approach, general primitives and protocols required the ability to directly integrate with other pieces of the ecosystem. Importantly, they had to know what they wanted to interface with before even launching. Instead of having to be omniscent, developers have the ability to build transactions that seamlessly interact with multiple protocols at once.

- `Not Opinionated`: Unlike other solutions that offer deem a one-size-fits-all model with an immense amount of opinion baked in, Plug is modular and unopinionated. This design enables you to tailor the framework to meet the unique demands of your specific project without having to deal with the compromises of the standard proposed.

At all steps, Plug is designed to introduce as little new burden and required labor for developers as possible while allowing the focus to remain on the core primitive being built.
