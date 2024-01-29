// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugInitializable } from "../abstracts/Plug.Initializable.sol";

/**
 * @title Plug Router Socket
 * @notice This contract represents a general purpose relay socket that can be
 *         used to route transactions to other contracts.
 * @notice Do not approve assets to this contract as anyone can sign and/or
 *         execute transactions which means they can use your approvals.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugRouterSocket is PlugInitializable {
    /**
     * @notice Initializes a new Plug Router Socket contract.
     */
    constructor() PlugInitializable() { }

    /**
     * @notice Name used for the domain separator.
     */
    function name() public pure override returns (string memory) {
        return "PlugRouterSocket";
    }
}
