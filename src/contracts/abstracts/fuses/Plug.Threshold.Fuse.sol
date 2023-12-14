//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import { PlugFuseInterface } from "../../interfaces/Plug.Fuse.Interface.sol";
import { PlugTypesLib } from "../../abstracts/Plug.Types.sol";
import { BytesLib } from "../../libraries/BytesLib.sol";

abstract contract ThresholdFuse is PlugFuseInterface {
    using BytesLib for bytes;

    /**
     * See {FuseEnforcer-enforceFuse}.
     */
    function enforceFuse(
        bytes calldata $live,
        PlugTypesLib.Current calldata $current,
        bytes32
    )
        public
        view
        override
        returns (bytes memory $through)
    {
        /// @dev Decode the terms to get the logic operator and threshold.
        (uint256 $operator, uint256 $threshold) = decode($live);

        /// @dev Make sure the block number is before the threshold.
        if ($operator == 0) {
            if ($threshold <= _threshold()) revert("ThresholdFuse:expired-pin");
        }
        /// @dev Make sure the block number is after the threshold.
        else if ($threshold >= _threshold()) {
            revert("ThresholdFuse:early-pin");
        }

        /// @dev Continue the pass through.
        $through = $current.data;
    }

    /**
     * @dev Decode the terms to get the logic operator and threshold.
     */
    function decode(bytes calldata $data) public pure returns (uint128 $operator, uint128 $threshold) {
        /// @dev Retrieve the logic operator set in the terms.
        $operator = $data.toUint128(0);
        /// @dev Move 16 bytes to the right to get the threshold.
        $threshold = $data.toUint128(16);
    }

    /**
     * @dev Encode the logic operator and threshold.
     */
    function encode(uint128 $operator, uint128 $threshold) public pure returns (bytes memory $data) {
        /// @dev Encode the logic operator and threshold.
        $data = abi.encodePacked($operator, $threshold);
    }

    /**
     * @dev Unit denomination of the threshold.
     */
    function _threshold() internal view virtual returns (uint256);
}
