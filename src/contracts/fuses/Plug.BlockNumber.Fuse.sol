// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { ThresholdFuse } from "../abstracts/fuses/Plug.Threshold.Fuse.sol";

contract PlugBlockNumberFuse is ThresholdFuse {
    /// @dev Returns the current block number.
    function _threshold() internal view override returns (uint256) {
        return block.number;
    }

    /// @dev Returns the name of the fuse.
    function _name() internal pure virtual override returns (string memory) {
        return "PlugBlockNumberFuse";
    }
}
