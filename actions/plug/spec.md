# Plug Integration

## Overview

The definition of Plug actions are a bit unique as they contain generalized function definitions as well as specific Plug actions and constraints. Contained, we need to support general ERC token actions.

## Supporting Documentation

- [ERC20 Specification](https://github.com/ethereum/ERCs/blob/master/ERCS/erc-20.md)
- [ERC721 Specification](https://github.com/ethereum/ERCs/blob/master/ERCS/erc-721.md)
- [ERC1155 Specification](https://github.com/ethereum/ERCs/blob/master/ERCS/erc-1155.md)

## Contract Interfacing

Nothing in this specification interfaces with external materials except token contracts specifically. 

## Scope

| System  | Name         | Type   | Source | Implemented | Notes                                                                |
| :------ | :----------- | :----- | :----- | :---------: | :------------------------------------------------------------------- |
| Native  | Transfer     | Action | Token  |  11/17/24   | Uses 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE as the asset address |
| ERC20   | Transfer     | Action | Token  |  11/17/24   |                                                                      |
| ERC20   | TransferFrom | Action | Token  |             |                                                                      |
| ERC721  | TransferFrom | Action | Token  |             |                                                                      |
| ERC1155 | TransferFrom | Action | Token  |             |                                                                      |
| ERC20   | Approve      | Action | Token  |             |                                                                      |
| ERC721  | Approve      | Action | Token  |             |                                                                      |
| ERC1155 | Approve      | Action | Token  |             |                                                                      |


# Plug Constraints
A set of constraints made available by Plug to ease intent creation and increase control.

### Cooldown

Rate limit / cooldown, etc - TODO




### Token Pricing
Users can check if asset prices are above or below US dollar values.

`If {1} {2} is {3} than {4} USD.`

_If_ **{amount in}** **{token in}** _is_ **{greater, less}** _than_ **{amount out}** _USD._

_If_ **1** **ETH** _is_ **greater** _than_ **3000** _USD._

| Input Slot | Input | Type | Example |
| :--- | :--- | :--- | :--- |
| 1 | amount in | number | 1 |
| 2 | token in | string | ETH |
| 3 | comparison | string | greater, less |
| 4 | amount out | number | 3000 |

### Cross Token Pricing
This allows users to check prices of tokens against each other in a generalized way.

`If {1} {2} is {3} than {4} {5}.`

_If_ **{amount in}** **{token in}** _is_ **{greater, less}** _than_ **{amount out}** **{token out}**_._

_If_ **1** **ETH** _is_ **greater** _than_ **3000** **USDC**_._

| Input Slot | Input | Type | Example |
| :--- | :--- | :--- | :--- |
| 1 | amount in | number | 1 |
| 2 | token in | string | ETH |
| 3 | comparison | string | greater, less |
| 4 | amount out | number | 3000 |
| 5 | token out | string | USDC |

### Token Balances
Users can see if an account opens, adds to, or closes various positions.

`If {1} holds {2} {3} {4}.`

_If_ **{account}** _holds_ **{at least, less than}** **{amount}** **{token}**_._

_If_ **danner.eth** _holds_ **at least** **1** **ETH**_._

| Input Slot | Input | Type | Example |
| :--- | :--- | :--- | :--- |
| 1 | account | address | danner.eth |
| 2 | comparison | string | at least, less than |
| 3 | amount | number | 1 |
| 4 | token | string | ETH |

### NFT Balances
Determine if an account is acquiring, distributing, or holding NFTs from a collection.

`If {1} holds {2} {3} {4}.`

_If_ **{account}** _holds_ **{at least, less than}** **{amount}** **{collection}** _NFT._

_If_ **danner.eth** _holds_ **less than** **1** **Mutant Ape Yacht Club** _NFT._

| Input Slot | Input | Type | Example |
| :--- | :--- | :--- | :--- |
| 1 | account | address | danner.eth |
| 2 | comparison | string | at least, less than |
| 3 | amount | number | 1 |
| 4 | collection | string | Mutant Ape Yacht Club |

### NFT Targeting
Users can see if an account holds a specific NFT.

`If {1} holds {2} {3}.`

_If_ **{account}** _holds_ **{collection}** **{id}**_._

_If_ **danner.eth** _holds_ **Bastard Gan Punk** **#1234**_._

| Input Slot | Input | Type | Example |
| :--- | :--- | :--- | :--- |
| 1 | account | address | danner.eth |
| 2 | collection | string | Bastard Gan Punk |
| 3 | id | string | #1234 |
