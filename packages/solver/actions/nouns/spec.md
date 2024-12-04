# Nouns Integration

## Overview

With Nouns our primary focus is a lightbulb moment for less-sophisticated users. Because of this we are mostly focused on establishing understanding and do not expect exceptional usage here. Due to this we will not go crazy with the implementation and only focus on critical functionality to turn the light bulb on.

## Supporting Documentation

- [Nouns Developer Docs](https://docs.ens.domains/contracts)
- [Auction House Contract](https://etherscan.io/address/0x830BD73E4184ceF73443C15111a1DF14e495C706#writeContract)
- [Token Contract](https://etherscan.io/address/0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03#writeContract)
- [Trait Images](https://github.com/nounsDAO/nouns-monorepo/tree/master/packages/nouns-assets/images/v0)

## Contract Interfacing

Most of the Nouns actions take place through the Auction House. However, most constraints actually take place through the Nouns token contract. The Auction House does not contain metavalue information or token id information. Nouns itself reads the onchain value instead of using an API for the value as well as we definitely do not want to be reliant on a ~2 year old API even if they stood one up.

## Scope

| System              | Name                     | Type       | Implemented | Notes                                                                                            |
| :------------------ | :----------------------- | :--------- | :---------- | :----------------------------------------------------------------------------------------------- |
| Nouns Auction House | Bid Amount               | Action     | 12/4/2024   |                                                                                                  |
| Nouns Auction House | Increase Bid             | Action     | 12/4/2024   | Sometimes you only want to outbid the current one by a set amount instead of having a max amount |
| Nouns Token         | Has Trait                | Action     |             | Will be checked by the hash of the trait                                                         |
| Nouns Token         | Is Token Id              | Constraint | 12/4/2024   | Solved for by seeing the last minted token? Believe the active auction returns the last id?      |
| Nouns Token         | Current Bid Within Range | Constraint | 12/4/2024   | Bound the amount willing to bid on the noun                                                      |
