# Actions

<span style="color: rgba(0,0,0,0.6)">By combining comprehensive [constraints](/concepts/atoms/constraints) with targeted [actions](/concepts/atoms/actions), Plug empowers users to have full control over both the execution process and its results, significantly optimizing the transaction experience across multiple blockchain environments.</span>

## Types of Actions

To achieve the benefits of a deeply generalized framework, Plug supports three different kinds of actions:

- [**Imperative:**](#imperative) declare precisely the transaction to execute.
- [**Declarative:**](#declarative) declare what the final result will be when a transaction executes.
- [**Snapshot:**](#snapshot) snapshot current onchain state for use by later constraints and actions.

### Imperative

Imperative actions are akin to giving precise, step-by-step cooking instructions.

Imagine you're in the kitchen with a recipe that tells you to preheat the oven to 350 degrees, mix flour and sugar in a bowl, and then bake for exactly 25 minutes. Each step is explicitly defined, and there's no room for interpretation—just follow the directions to achieve the desired outcome, which in this case is baking a cake.

Similarly, in the context of Plug, an imperative action involves defining exact blockchain transactions. For instance, sending 5 ETH from wallet A to wallet B at a specified time. The system executes exactly what's commanded without considering any other possibilities or outcomes.

### Declarative

Conversely, declarative actions are more like telling a professional chef to prepare a delicious dessert. You don’t specify how to make it; you’re just interested in having a delightful end result. The chef uses their expertise to decide whether to make a cake, pie, or something entirely different. They choose the ingredients and techniques based on the desired outcome of a great dessert.

In Plug, declarative actions operate similarly. You state what you want to achieve—for example, owning a specific NFT. The system then figures out the best route to take, whether it’s participating in an auction or buying it directly at the current market price. It automates the decision-making process within the constraints provided to achieve the stated goal.

These action types offer distinct advantages depending on the situation. Imperative actions provide control and predictability, ensuring that transactions are executed precisely as intended. Declarative actions, on the other hand, offer flexibility and efficiency, as the system can optimize the steps to achieve the desired outcome based on the current state and available opportunities in the blockchain environment.

### Snapshot

Snapshot actions are similar to capturing a snapshot with a camera at a significant moment, preserving specific details for future reference. In Plug, snapshot actions record the state of the blockchain at a chosen point in a transaction sequence.

This function is particularly useful for transactions where outcomes depend on changes in the blockchain state over the transaction's duration. For example, if a transaction requires confirming an increase in token balance by 50 tokens, a snapshot action would first capture the initial token balance. The transaction proceeds, resulting in token acquisition, and upon completion, subsequent actions or constraints check the new balance against the snapshot. This confirms that the balance has indeed increased by at least 50 tokens, as required.

By employing snapshot actions, you can orchestrate complex, condition-based transaction paths that adapt to both real-time changes and historical states. This ensures dynamic and accurate execution of smart contracts, tailored to precise needs and scenarios.

Snapshot actions involve capturing and utilizing the current state of the blockchain to inform transaction execution at a later time in the process of transaction execution.


