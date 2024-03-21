// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Receiver } from "solady/src/accounts/Receiver.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

/**
 * @title PlugTreasury
 * @notice The treasury contract that receives fees from the Plug
 *         framework and can execute arbitrary transactions.
 */
contract PlugTreasury is Receiver, Ownable, Initializable {
    /**
     * @notice Initialize the contract with the owner.
     * @param $owner The address of the owner.
     */
    function initialize(address $owner) public initializer {
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
        (bool success,) = $to.call{ value: $value }($data);
        require(success, "PlugTreasury: execution failed");
    }
}
