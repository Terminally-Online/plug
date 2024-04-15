//SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugConnectorInterface, PlugTypesLib } from "../../interfaces/Plug.Connector.Interface.sol";
import { PlugThresholdEnforce } from "./Plug.Threshold.Enforce.sol";
import { PlugLib } from "../../libraries/Plug.Lib.sol";

abstract contract PlugThreshold is PlugConnectorInterface, PlugThresholdEnforce {
    /**
     * See {PlugConnectorInterface-enforce}.
     */
    function enforce(bytes calldata $terms, bytes32) public view virtual {
        /// @dev Decode the terms to get the logic operator and threshold.
        (uint8 $operator, uint256 $threshold) = decode($terms);

        /// @dev Enforce the threshold against the current state.
        _enforce($operator, $threshold, _threshold());
    }

    /**
     * @dev Decode the terms to get the logic operator and threshold.
     */
    function decode(bytes calldata $data)
        public
        pure
        virtual
        returns (uint8 $operator, uint256 $threshold)
    {
        ($operator, $threshold) = abi.decode($data, (uint8, uint256));
    }

    /**
     * @dev Encode the logic operator and threshold.
     */
    function encode(
        uint8 $operator,
        uint256 $threshold
    )
        public
        pure
        virtual
        returns (bytes memory $data)
    {
        /// @dev Encode the logic operator and threshold.
        $data = abi.encode($operator, $threshold);
    }

    /**
     * @dev Unit denomination of the threshold.
     */
    function _threshold() internal view virtual returns (uint256) { }
}
