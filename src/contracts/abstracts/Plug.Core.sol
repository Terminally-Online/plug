// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypes, PlugTypesLib } from "./Plug.Types.sol";
import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugErrors } from "../libraries/Plug.Errors.sol";

/**
 * @title Plug Core
 * @notice The core contract for the Plug framework that enables
 *         counterfactual intent execution with granular conditional
 *         verification and execution.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugCore is PlugTypes {
    using PlugErrors for bytes;

    /**
     * @notice Confirm that signer of the intent has permission to declare
     *         the execution of an intent.
     * @dev If you would like to limit the available signers override this
     *      function in your contract with the additional logic.
     */
    function _enforceSigner(address $signer) internal view virtual { }

    /**
     * @notice Enforce the fuse of the current plug to confirm
     *         the specified conditions have been met.
     * @param $fuse The fuse to enforce.
     * @param $current The state of the transaction to execute.
     * @param $pinHash The hash of the pin.
     * @return $through The return data of the fuse.
     */
    function _enforceFuse(
        PlugTypesLib.Fuse memory $fuse,
        PlugTypesLib.Current memory $current,
        bytes32 $pinHash
    )
        internal
        returns (bytes memory $through)
    {
        /// @dev Warm up the slot for the return data.
        bool success;

        /// @dev Call the Fuse to determine if it is valid.
        (success, $through) = $fuse.neutral.call(
            abi.encodeWithSelector(
                PlugFuseInterface.enforceFuse.selector,
                $fuse.live,
                $current,
                $pinHash
            )
        );

        /// @dev If the Fuse failed and is not optional, bubble up the revert.
        if (!success) $through.bubbleRevert();

        /// @dev Decode the return data to remove the wrapped bytes in memory.
        $through = abi.decode($through, (bytes));
    }

    /**
     * @notice Possibly restrict the capability of the defined
     *         execution path dependent on larger external factors
     *         such as only allowing a transaction to be executed
     *         on the socket itself.
     * @param $current The current state of the transaction.
     */
    function _enforceCurrent(PlugTypesLib.Current memory $current)
        internal
        view
        virtual
    { }

    /**
     * @notice Execution a built transaction.
     * @param $current The current state of the transaction.
     * @return $result The return data of the transaction.
     */
    function _execute(
        PlugTypesLib.Current memory $current,
        address $sender
    )
        internal
        returns (bytes memory $result)
    {
        /// @dev Build the final call data.
        bytes memory full = abi.encodePacked($current.data, $sender);

        /// @dev Warm up the slot for the return data.
        bool success;

        /// @dev Make the external call with a standard call.
        (success, $result) = address($current.ground).call{
            gas: gasleft(),
            value: $current.voltage
        }(full);

        /// @dev If the call failed, bubble up the revert reason if possible.
        if (!success) $result.bubbleRevert();
    }

    /**
     * @notice Execute an array of plugs
     * @param $plugs The plugs to execute.
     * @param $sender The address of the sender.
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $sender
    )
        internal
        returns (bytes[] memory $results)
    {
        /// @dev Prevent random people from plugging.
        _enforceSigner($sender);

        /// @dev Load the plugs from the live plugs.
        PlugTypesLib.Plug[] memory plugs = $plugs.plugs;

        /// @dev Load the stack.
        PlugTypesLib.Plug memory plug;
        uint256 i;
        uint256 ii;
        uint256 length = plugs.length;
        $results = new bytes[](length);

        /// @dev Unique hash of the Plug bundle being executed.
        bytes32 plugsHash = getPlugsHash($plugs);

        /// @dev Iterate over the plugs.
        for (i; i < length; i++) {
            /// @dev Load the plug from the plugs.
            plug = plugs[i];

            /// @dev Iterate through all the execution fuses declared in the pin
            ///      and ensure they are in a state of acceptable execution
            ///      while building the pass through data based on the nodes.
            for (ii = 0; ii < plug.fuses.length; ii++) {
                plug.current.data =
                    _enforceFuse(plug.fuses[ii], plug.current, plugsHash);
            }

            /// @dev Confirm the current is within specification.
            /// @dev This is not done sooner because a fuse may manipulate the
            ///      the declaration of the current.
            _enforceCurrent(plug.current);

            /// @dev Execute the transaction.
            $results[i] = _execute(plug.current, $sender);
        }
    }
}
