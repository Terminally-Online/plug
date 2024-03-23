// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Receiver } from "solady/src/accounts/Receiver.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title PlugTreasury
 * @notice The treasury contract that receives fees from the Plug
 *         framework and can execute arbitrary transactions.
 */
contract PlugTreasury is Receiver, Ownable {
    /**
     * @notice Initialize the contract with the owner.
     * @param $owner The address of the owner.
     */
    function initialize(address $owner) public {
        _initializeOwner($owner);
    }

    /**
     * @notice Execute an arbitrary transaction from the Treasury.
     * @param $to The address to send the transaction to.
     * @param $value The amount of value to send.
     * @param $data The data to send.
     */
    function execute(
        address $to,
        uint256 $value,
        bytes memory $data
    )
        external
        onlyOwner
    {
        (bool success, bytes memory reason) =
            $to.call{ value: $value }($data);
        PlugLib.bubbleRevert(success, reason);
    }

    /**
     * See {Ownable-_initializeOwner}.
     */
    function _guardInitializeOwner()
        internal
        pure
        override
        returns (bool $guard)
    {
        $guard = true;
    }
}
