In recent months, [[Intents]] have [risen in popularity](https://www.biconomy.io/post/meta-transactions-account-abstraction-to-intents-evolution-of-web3-ui) among small social pockets of the web3 industry. Enabling a new world of functionality and design space, it's now possible and appreciated by users to deliver a better user experience than was possible previously.

With [[Intents]], the execution of blockchain transactions shifts from imperative executive to declarative execution. The simplest way to think about this, is that a typical (imperative) transaction has value in being executed *now* while a declarative transactions has value in the declared outcome.

In recent years, many transactions required immediate execution. However, that is not true for everything, realistically, very few things. This means, that instead of having to wear the cost of every transaction one may wait to run their transaction until economic incentives are properly aligned. This simple idea, powers an entirely new world of mechanisms and platforms.


---

## Resources

- https://www.paradigm.xyz/2023/06/intents

---

## References

- 

---

Unlike a normal transaction, [[Intents]] are built in a Matchmaking service. With this, an Intent can contain a bundle of transactions much like the model popularized by Flashbots.

![[Pasted image 20230722140305.png]]

All in all, this is not an entirely new idea, though. You may notice that many of this idea corresponds to systems that have been in operation for several years:

1. **Limit Orders**: 100 X may be taken from my account if I receive at least 200 Y.
2. **CowSwap-style Auctions**: same as above, but rely on a third-party or mechanism to match many orders to maximize execution quality.
3. **Gas Sponsorship**: Pay gas in USDC instead of ETH. The intent can only be fulfilled with a matching intent which pays ETH in fees.
4. **Delegation**: Only allow interacting with certain accounts in certain pre-authorized ways. The intent can only be fulfilled if the final transaction respects the access control list specified in the intent.
5. **Transaction Batching**: Allow batching of intents for gas efficiencies.
6. **Aggregators**: Only use “best” price/yield for an action. The intent can be fulfilled by showing a proof that an aggregation over multiple venues was executed and the optimal path was taken.