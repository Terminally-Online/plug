// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Receiver } from "solady/accounts/Receiver.sol";
import { Ownable } from "solady/auth/Ownable.sol";
import { PlugSwapper, PlugLib } from "../abstracts/Plug.Swapper.sol";

/**
 * @title PlugTreasury
 * @notice The Treasury contract that receives fees from the Plug
 *         framework. While the owner of the treasury can execute
 *         arbitrary transactions, the Treasury also has a built-in
 *         Swapper to enable streamlined fee collection while allowing
 *         the tokens to remain in the Treasury rather than needing
 *         another token transfer.
 */
contract PlugTreasury is Receiver, Ownable, PlugSwapper {
    /**
     * @notice Initialize the contract with the owner.
     * @param $owner The address of the owner.
     */
    function initialize(address $owner) public {
        _initializeOwner($owner);
    }

    /**
     * @notice Set the targets allowed to be executed by the Treasury.
     * @param $targets The targets to set the allowed status for.
     * @param $allowed The allowed status to set.
     */
    function setTargetsAllowed(
        address[] calldata $targets,
        bool $allowed
    )
        public
        virtual
        onlyOwner
    {
        for (uint256 i; i < $targets.length; i++) {
            targetToAllowed[$targets[i]] = $allowed;
        }
    }

    /**
     * @notice Execute multiple calls that are not required to succeed.
     * @param $targets The targets to call.
     * @param $values The values to send with the calls.
     * @param $datas The data to send with the calls.
     * @return $successes The success status of each call.
     * @return $results The results of each call.
     */
    function execute(
        address[] calldata $targets,
        uint256[] calldata $values,
        bytes[] calldata $datas
    )
        public
        virtual
        onlyOwner
        returns (bool[] memory $successes, bytes[] memory $results)
    {
        /// @dev Take a snapshot of the number of calls in the array.
        uint256 length = $targets.length;

        /// @dev Instantiate the results array with the length of the calls array.
        $successes = new bool[](length);
        $results = new bytes[](length);

        /// @dev Loop through all of the calls in the array.
        for (uint256 i; i < length; i++) {
            /// @dev Execute the transaction from within the array and save the response
            ///      of success and failure reason into the result.
            ($successes[i], $results[i]) = $targets[i].call{ value: $values[i] }($datas[i]);
        }
    }

    /**
     * See {Ownable-_initializeOwner}.
     */
    function _guardInitializeOwner() internal pure override returns (bool $guard) {
        $guard = true;
    }
}
