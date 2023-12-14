// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import { PlugTypes, PlugTypesLib } from "./Plug.Types.sol";
import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugErrors } from "../libraries/Plug.Errors.sol";

/**
 * @title Plug Core
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract PlugCore is PlugTypes {
    using PlugErrors for bytes;

    /// @notice Multi-dimensional account pin nonce management.
    mapping(address => mapping(uint256 => uint256)) public nonce;

    /**
     * @notice Determine the address representing the message sender in the
     *         current context. This is important for pins, as the
     *         message sender may be the framework itself, in which case
     *         the sender must be extracted from the data.
     * @return $sender The address of the message sender.
     */
    function _msgSender() internal view virtual returns (address $sender) {
        /// @dev If the message sender is the framework, we need to extract the
        ///      sender from the data.
        if (msg.sender == address(this)) {
            /// @dev Load the data as a hot reference.
            bytes memory array = msg.data;

            /// @dev Load the length of the data as a hot reference.
            uint256 index = array.length;

            assembly {
                /// @dev Load the sender from the data by applying a
                ///      bitwise AND operation to the data and the
                ///      maximum uint256 value, keeping only the last
                ///      20 bytes (160 bits) (address size).
                $sender :=
                    and(
                        /// @dev Load the bytes at the computer pointer.
                        mload(
                            /// @dev Computes the sum of the starting address of array and index.
                            ///      This effectively points to the end of array.
                            add(array, index)
                        ),
                        0xffffffffffffffffffffffffffffffffffffffff
                    )
            }
        }
        /// @dev Otherwise, the sender is the message sender.
        else {
            $sender = msg.sender;
        }
    }

    /**
     * @notice Update the nonce for a given account and queue.
     * @param $intendedSender The address of the intended sender.
     * @param $protection The replay protection struct.
     */
    function _enforceBreaker(address $intendedSender, PlugTypesLib.Breaker memory $protection) internal {
        /// @dev Ensure the nonce is in order.
        require($protection.nonce == ++nonce[$intendedSender][$protection.queue], "PlugCore:nonce2-out-of-order");
    }

    function _enforceFuse(
        PlugTypesLib.Fuse memory $fuse,
        PlugTypesLib.Current memory $current,
        bytes32 $pinHash
    )
        internal
        returns (bytes memory $through)
    {
        /// @dev Warm up the success variable.
        bool success;

        /// @dev Call the Fuse to determine if it is valid.
        (success, $through) = address($fuse.neutral).call(
            abi.encodeWithSelector(PlugFuseInterface($fuse.neutral).enforceFuse.selector, $fuse.live, $current, $pinHash)
        );

        /// @dev If the Fuse failed and is not optional, bubble up the revert.
        if (!success && $fuse.forced) $through.bubbleRevert();
    }

    /**
     * @notice Execution a built transaction.
     * @param $to The address of the contract to execute.
     * @param $data The data to execute on the contract.
     * @param $voltage The value to send with the transaction.
     * @param $sender The address of the sender.
     * @return $result The return data of the transaction.
     */
    function _execute(
        address $to,
        bytes memory $data,
        uint256 $voltage,
        address $sender
    )
        internal
        returns (bytes memory $result)
    {
        /// @dev Build the final call data.
        bytes memory full = abi.encodePacked($data, $sender);

        /// @dev Warm up the slot for the return data.
        bool success;

        /// @dev Make the external call with a standard call.
        (success, $result) = address($to).call{ gas: gasleft(), value: $voltage }(full);

        /// @dev If the call failed, bubble up the revert reason if possible.
        if (!success) $result.bubbleRevert();
    }

    /**
     * @notice Execute an array of plugs
     * @param $plugs The plugs of plugs to execute.
     * @param $sender The address of the sender.
     * @return $results The return data of the plugs.
     */
    function _plug(PlugTypesLib.Plug[] calldata $plugs, address $sender) internal returns (bytes[] memory $results) {
        /// @dev Warm up the results array.
        $results = new bytes[]($plugs.length);

        /// @dev Load the stack.
        uint256 i;
        uint256 j;
        uint256 k;
        address canGrant;
        address intendedSender;
        address pinSigner;
        bytes32 pinHash;

        /// @dev Load the structs into a hot reference.
        PlugTypesLib.Plug memory plug;
        PlugTypesLib.LivePin memory signedPin;
        PlugTypesLib.Pin memory pin;
        PlugTypesLib.Current memory current;

        /// @dev Iterate over the plugs.
        for (i; i < $plugs.length; i++) {
            /// @dev Load the plug from the plugs.
            plug = $plugs[i];

            /// @dev Reset the hot reference to the pinHash.
            pinHash = 0x0;

            /// @dev Load the transaction from the plug.
            current = plug.current;

            /// @dev If there are no pins, this plug comes from the signer
            if (plug.pins.length == 0) {
                canGrant = intendedSender = $sender;
            } else {
                /// @dev Iterate over the authority pins.
                for (j = 0; j < plug.pins.length; j++) {
                    /// @dev Load the pin from the plug.
                    signedPin = plug.pins[j];

                    /// @dev Determine the signer of the pin.
                    pinSigner = getLivePinSigner(signedPin);

                    /// @dev Implied sending account is the signer of the first pin.
                    if (j == 0) canGrant = intendedSender = pinSigner;

                    /// @dev Ensure the pin signer has authority to grant
                    ///      the claimed pin.
                    require(pinSigner == canGrant, "PlugCore:invalid-pin-signer");

                    /// @dev Warm up the pin reference.
                    pin = signedPin.pin;

                    /// @dev Ensure the pin is valid.
                    require(pin.live == pinHash, "PlugCore:invalid-authority-pin-link");

                    /// @dev Retrieve the packet hash for the pin.
                    pinHash = getLivePinHash(signedPin);

                    /// @dev Loop through all the execution fuses declared in the pin
                    ///      and ensure they are in a state of acceptable execution
                    ///      while building the pass through data based on the nodes.
                    for (k = 0; k < pin.fuses.length; k++) {
                        current.data = _enforceFuse(pin.fuses[k], current, pinHash);
                    }
                }
            }

            /// @dev Verify the delegate at the end of the pin chain is the signer.
            require(canGrant == $sender, "PlugCore:invalid-signer");

            /// @dev Execute the transaction.
            $results[i] = _execute(current.ground, current.data, current.voltage, intendedSender);
        }
    }
}
