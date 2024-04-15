// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { console2 } from "forge-std/console2.sol";

import { PlugTypes } from "./Plug.Types.sol";
import { PlugLib, PlugTypesLib, PlugAddressesLib } from "../libraries/Plug.Lib.sol";
import { PlugConnectorInterface } from "../interfaces/Plug.Connector.Interface.sol";

/**
 * @title PlugCore
 * @author @nftchance (chance@onplug.io)
 */
abstract contract PlugCore is PlugTypes {
    /**
     * @notice Execute a bundle of Plugs.
     * @param $plugs The Plugs to execute containing the bundle and side effects.
     * @param $solver Encoded data defining the Solver and compensation.
     * @param $gas Snapshot of gas at the start of interaction.
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $solver,
        uint256 $gas
    )
        internal
        returns (PlugTypesLib.Result[] memory $results)
    {
        /// @dev Hash the body of the object to ensure the integrity of
        ///      the (bundle of) Plugs that are being executed.
        bytes32 plugsHash = getPlugsHash($plugs);

        /// @dev Load the Plug stack into memory for cheaper access.
        uint256 length = $plugs.plugs.length;
        $results = new PlugTypesLib.Result[](length);

        /// @dev Save the object into memory to avoid multiple creations
        ///      of the same object.
        PlugTypesLib.Plug calldata plug;

        /// @dev Iterate over the Plugs that are held within this bundle
        ///      an execute each of them. Each respectively may be a
        ///      condition being enforced or an outcome focused transaction.
        for (uint256 i; i < length; i++) {
            /// @dev Place the active Plug in the shorter reference stack.
            plug = $plugs.plugs[i];

            /// @dev If the call has an associated value, ensure the contract
            ///      has enough balance to cover the cost of the call.
            if (address(this).balance < plug.value) {
                revert PlugLib.ValueInvalid(plug.target, plug.value, address(this).balance);
            }

            /// @dev Recover the byte that is being used to denote the type of
            ///      Plug being executed. It has to be done this way because
            ///      instead of declaring multiple EIP712 types, we are using
            ///      a single type and encoding the data in a way that is
            ///      recoverable and solvable in a single type and call.
            bytes1 plugType = plug.data[0];

            /// @dev This check is a conditional Plug that requires access
            ///      to the hash and is confirming the current state of some
            ///      onchain data in an external contract.
            if (plugType & 0x01 == plugType) {
                /// @dev Call the Plug to determine that is operating as a
                ///      condition and enforce the outcome of the condition
                ///      if it is not met (reverts).
                ($results[i].success, $results[i].result) = plug.target.call{ value: plug.value }(
                    abi.encodeWithSelector(
                        PlugConnectorInterface.enforce.selector, plug.data[1:], plugsHash
                    )
                );
            }
            /// @dev Make the call to the Plug and bubble up the
            ///      result if it happens to fail.
            else if (plugType & 0x02 == plugType) {
                ($results[i].success, $results[i].result) =
                    plug.target.call{ value: plug.value }(plug.data[1:]);
            }
            /// @dev If an invalid Plug type was provided revert to protect
            ///      against fund siphoning when no work is done.
            else {
                revert PlugLib.TypeInvalid(uint8(plugType));
            }

            /// @dev If the call failed, bubble up the revert reason if needed.
            PlugLib.bubbleRevert($results[i].success, $results[i].result);
        }

        /// @dev Pay the Solver for the gas used if it was not open-access.
        if ($plugs.solver.length != 0) {
            /// @dev Unpack the solver data from the encoded Solver data.
            (uint96 maxPriorityFeePerGas, uint96 maxFeePerGas, address solver) =
                abi.decode($plugs.solver, (uint96, uint96, address));

            /// @dev Confirm the Solver is allowed to execute the transaction.
            ///      This is done here instead of a modifier so that the gas
            ///      snapshot accounts for the additional gas cost of the require.
            if (solver != $solver) {
                revert PlugLib.SolverInvalid(solver, $solver);
            }

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
            (bool success,) = solver.call{ value: value }("");
            if (success == false) {
                revert PlugLib.CompensationFailed(solver, value);
            }
        }
    }
}
