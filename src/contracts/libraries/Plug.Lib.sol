// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

library PlugLib {
    address internal constant PLUG_ADDRESS =
        0x00f140e87692075C2D200bf313839Db0d669a5Da;

    address internal constant PLUG_TRADABLE_ADDRESS =
        0x00f140e87692075C2D200bf313839Db0d669a5Da;

    event SocketDeployed(
        address indexed implementation, address indexed vault, bytes32 salt
    );

    /**
     * @notice Bubble up the revert reason revert data.
     * @param $revertData The revert data to extract the reason from.
     */
    function bubbleRevert(bytes memory $revertData) internal pure {
        /// @dev If we won't be able to recover the message, go ahead
        ///      and revert with the default.
        if ($revertData.length < 4) revert("PlugCore:revert");

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
}
