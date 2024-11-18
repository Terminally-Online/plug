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
