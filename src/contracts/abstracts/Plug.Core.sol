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

    /// @notice Enables consumption of the address that gave permission
    ///         to execute this transaction.
    address public grantor;

    /// @notice Enables consumption of the address that was given permission
    ///         to execute this transaction.
    address public granted;

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
     * @notice Update the nonce for a given account and queue.
     * @param $intendedSender The address of the intended sender.
     * @param $protection The replay protection struct.
     */
    function _enforceBreaker(
        address $intendedSender,
        PlugTypesLib.Breaker memory $protection
    )
        internal
    {
        /// @dev Ensure the nonce is in order.
        require(
            $protection.nonce == ++nonce[$intendedSender][$protection.queue],
            "PlugCore:nonce2-out-of-order"
        );
    }

    /**
     * @notice Enforce the fuse of the current plug to confirm the specified
     *         conditions have been met.
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
        if (!success && $fuse.forced) $through.bubbleRevert();

        /// @dev Decode the return data to remove the wrapped bytes in memory.
        $through = abi.decode($through, (bytes));
    }

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
     * @param $plugs The plugs of plugs to execute.
     * @param $sender The address of the sender.
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plug[] calldata $plugs,
        address $sender
    )
        internal
        returns (bytes[] memory $results)
    {
        /// @dev Warm up the results array.
        $results = new bytes[]($plugs.length);

        /// @dev Load the stack.
        uint256 i;
        uint256 ii;
        uint256 iii;
        address pinSigner;
        bytes32 pinHash;

        /// @dev Load the structs into a hot reference.
        PlugTypesLib.Plug memory plug;
        PlugTypesLib.LivePin memory signedPin;
        PlugTypesLib.Pin memory pin;

        /// @dev Iterate over the plugs.
        for (i; i < $plugs.length; i++) {
            /// @dev Load the plug from the plugs.
            plug = $plugs[i];

            /// @dev Reset the hot reference to the pinHash.
            pinHash = 0x0;

            /// @dev If there are no pins, this plug comes from the signer
            if (plug.pins.length == 0) {
                grantor = granted = $sender;
            } else {
                /// @dev Iterate over the authority pins.
                for (ii = 0; ii < plug.pins.length; ii++) {
                    /// @dev Load the pin from the plug.
                    signedPin = plug.pins[ii];

                    /// @dev Determine the signer of the pin.
                    pinSigner = getLivePinSigner(signedPin);

                    /// @dev Implied sending account is the signer of the first pin.
                    if (ii == 0) grantor = granted = pinSigner;

                    /// @dev Ensure the pin signer has authority to grant
                    ///      the claimed pin.
                    require(pinSigner == grantor, "PlugCore:invalid-pin-signer");

                    /// @dev Warm up the pin reference.
                    pin = signedPin.pin;

                    /// @dev Ensure the pin is valid.
                    require(
                        pin.live == pinHash,
                        "PlugCore:invalid-authority-pin-link"
                    );

                    /// @dev Retrieve the packet hash for the pin.
                    pinHash = getLivePinHash(signedPin);

                    /// @dev Loop through all the execution fuses declared in the pin
                    ///      and ensure they are in a state of acceptable execution
                    ///      while building the pass through data based on the nodes.
                    for (iii = 0; iii < pin.fuses.length; iii++) {
                        plug.current.data =
                            _enforceFuse(pin.fuses[iii], plug.current, pinHash);
                    }
                }
            }

            /// @dev Confirm the current is within specification.
            /// @dev This is not done sooner because a fuse may manipulate the
            ///      the declaration of the current.
            _enforceCurrent(plug.current);

            /// @dev Verify the delegate at the end of the pin chain is the signer.
            require(grantor == $sender, "PlugCore:invalid-signer");

            /// @dev Execute the transaction.
            $results[i] = _execute(plug.current, granted);

            /// @dev Clear the grantor slot back to address(0).
            delete grantor;
            /// @dev Clear the granted slot back to address(0).
            delete granted;
        }
    }
}
