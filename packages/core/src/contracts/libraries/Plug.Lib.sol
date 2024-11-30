// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugAddressesLib } from "./Plug.Addresses.Lib.sol";

library PlugLib {
    event SocketDeployed(address indexed implementation, address indexed vault, bytes32 salt);

    event SocketOwnershipTransferred(
        address indexed previousOwner, address indexed newOwner, bytes32 imageHash
    );

    event PlugsExecuted(bytes32 indexed $plugsHash, PlugTypesLib.Result[] $results);

    error NotImplemented();

    error SocketAddressInvalid(address $intended, address $socket);
    error SocketAddressEmpty(address $socket);

    error SaltInvalid(address $implementation, address $admin);
    error NonceInvalid();
    error CallerInvalid(address $expected, address $reality);
    error SenderInvalid(address $reality);
    error RouterInvalid(address $reality);
    error SignatureInvalid();
    error SolverInvalid(address $expected, address $reality);
    error SolverExpired();
    error TypeInvalid(uint8 $reality);
    error ValueInvalid(address $recipient, uint256 $expected, uint256 $reality);

    error PlugFailed();

    error CompensationFailed(address $recipient, uint256 $value);

    error ThresholdInvalid();
    error ThresholdExceeded(uint256 $expected, uint256 $reality);
    error ThresholdInsufficient(uint256 $expected, uint256 $reality);

    error TargetInvalid();

    error TokenAllowanceInvalid();
    error TokenBalanceInvalid();

    /**
     * @notice Bubble up the revert reason revert data from an internal call
     *         that would typically revert without surfacing the reason.
     * @param $revertData The revert data to extract the reason from.
     */
    function bubbleRevert(bytes memory $revertData) internal pure {
        /// @dev If we won't be able to recover the message, go ahead
        ///      and revert with the default.
        if ($revertData.length < 4) {
            revert PlugFailed();
        }

        bytes4 errorSelector;
        assembly {
            errorSelector := mload(add($revertData, 0x20))
        }

        /// @dev Panic(uint256) (>=0.8.0)
        if (errorSelector == bytes4(0x4e487b71)) {
            string memory reason = "PlugCore:target-panicked-0x";
            uint256 errorCode;

            assembly {
                errorCode := mload(add($revertData, 0x24))
                let reasonWord := mload(add(reason, 0x20))
                // [0..9] is converted to ['0'..'9']
                // [0xa..0xf] is not correctly converted to ['a'..'f']
                // but since panic code doesn't have those cases, we will ignore them for now!
                let e1 := add(and(errorCode, 0xf), 0x30)
                let e2 := shl(8, add(shr(4, and(errorCode, 0xf0)), 0x30))
                reasonWord :=
                    or(
                        and(
                            reasonWord,
                            0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000
                        ),
                        or(e2, e1)
                    )
                mstore(add(reason, 0x20), reasonWord)
            }

            revert(reason);
        }

        /// @dev Error(string) (>= 0.7.0)
        /// @dev Custom errors (>= 0.8.0)
        uint256 len = $revertData.length;
        assembly {
            revert(add($revertData, 32), len)
        }
    }

    /**
     * @notice Helper function to bubble up a revert reason if a condition is not met.
     * @param $reason The revert reason to surface.
     */
    function bubbleRevert(bool $success, bytes memory $reason) internal pure {
        /// @dev Confirm the call was successful.
        if (!$success) {
            /// @dev Go ahead and surface the revert reason as a failure would have
            ///      returned an error rather than the expected typed message.
            bubbleRevert($reason);
        }
    }
}
