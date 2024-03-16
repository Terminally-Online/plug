// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { PlugExecute } from "./Plug.Execute.sol";
import { PlugTypes, PlugTypesLib } from "./Plug.Types.sol";

/**
 * @title Plug.Core
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugCore is PlugExecute, PlugTypes {
    /**
     * @notice Execute a bundle of Plugs.
     * @param $plugs The plugs to execute.
     * @param $executor The address of the executor.
     * @param $gas Snapshot of gas at the start of interaction.
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $executor,
        uint256 $gas
    )
        internal
        returns (bytes[] memory $results)
    {
        /// @dev Hash the object to use in the Fuses.
        bytes32 plugsHash = getPlugsHash($plugs);

        /// @dev Load the Plug stack.
        PlugTypesLib.Plug[] calldata plugs = $plugs.plugs;
        PlugTypesLib.Current memory current;

        /// @dev Load the loop stack.
        uint256 i;
        uint256 ii;
        uint256 length = plugs.length;
        $results = new bytes[](length);

        /// @dev Iterate over the plugs.
        for (i; i < length; i++) {
            /// @dev Load the plug from the plugs.
            current = plugs[i].current;

            /// @dev Iterate through all the execution fuses declared in the pin
            ///      and ensure they are in a state of acceptable execution
            ///      while building the pass through data based on the nodes.
            for (ii = 0; ii < plugs[i].fuses.length; ii++) {
                (, current.data) =
                    _enforceFuse(plugs[i].fuses[ii], current, plugsHash);
            }

            /// @dev Execute the transaction.
            (, $results[i]) = _execute(current);
        }

        /// @dev Pay the Executor for the gas used and the fee earned if
        ///      it was not the original signer of the Plug bundle.
        if ($executor != address(0)) {
            /// @dev Calculate the gas price based on the current block.
            uint256 value = $plugs.maxPriorityFeePerGas + block.basefee;
            /// @dev Determine which gas price to use based on if it is a legacy
            ///      transaction (on a chain that does not support it) or if the
            ///      the transaction is submit post EIP-1559.
            value = $plugs.maxFeePerGas == $plugs.maxPriorityFeePerGas
                ? $plugs.maxFeePerGas
                : $plugs.maxFeePerGas < value ? $plugs.maxFeePerGas : value;

            /// @dev Augment the native gas price with the Executor "gas" fee.
            value = $plugs.fee + ($gas - gasleft()) * value;

            /// @dev Transfer the money the Executor is owed and confirm it
            ///      the transfer is successful.
            (bool success,) = $executor.call{ value: value }("");
            require(success, "Plug:compensation-failed");
        }
    }
}
