// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugExecute } from "./Plug.Execute.sol";
import { PlugTypes, PlugTypesLib } from "./Plug.Types.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title PlugCore
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugCore is PlugExecute {
    /**
     * @notice Distribute the fee earned to the platform and/or Solver.
     * @param $recipient The address of the recipient.
     * @param $value The amount of value to send.
     */
    function _compensate(address $recipient, uint256 $value) internal {
        /// @dev Transfer the money the Solver is owed and confirm it
        ///      the transfer is successful.
        (bool success,) = $recipient.call{ value: $value }("");
        require(success, "Plug:compensation-failed");
    }

    /**
     * @notice Execute a bundle of Plugs.
     * @param $plugs The plugs to execute.
     * @param $solver The address of the Solver.
     * @param $gas Snapshot of gas at the start of interaction.
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $solver,
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

        /// @dev Pay the platform fee if it there is an associated fee.
        if ($plugs.fee != 0) {
            _compensate(PlugLib.PLUG_TREASURY_ADDRESS, $plugs.fee);
        }

        /// @dev Pay the Solver for the gas used if it was not open-access.
        if ($plugs.solver.length != 0) {
            /// @dev Unpack the solver data from the encoded slot.
            (uint96 maxPriorityFeePerGas, uint96 maxFeePerGas, address solver) =
                abi.decode($plugs.solver, (uint96, uint96, address));

            /// @dev Confirm the Solver is allowed to execute the transaction.
            ///      This is done here instead of a modifier so that the gas
            ///      snapshot accounts for the additional gas cost of the require.
            require(solver == $solver, "Plug:invalid-solver");

            /// @dev Calculate the gas price based on the current block.
            uint256 value = maxPriorityFeePerGas + block.basefee;
            /// @dev Determine which gas price to use based on if it is a legacy
            ///      transaction (on a chain that does not support it) or if the
            ///      the transaction is submit post EIP-1559.
            value = maxFeePerGas == maxPriorityFeePerGas
                ? maxFeePerGas
                : maxFeePerGas < value ? maxFeePerGas : value;

            /// @dev Augment the native gas price with the Solver "gas" fee.
            value = ($gas - gasleft()) * value;

            /// @dev Transfer the money the Solver is owed and confirm it
            ///      the transfer is successful.
            _compensate(solver, value);
        }
    }
}
