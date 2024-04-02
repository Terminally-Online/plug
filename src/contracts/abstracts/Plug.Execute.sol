// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugEnforce } from "./Plug.Enforce.sol";
import { PlugTypes, PlugTypesLib } from "./Plug.Types.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";

abstract contract PlugExecute is PlugEnforce {
    /**
     * @notice Execution a pre-defined transaction stemming from a signed
     *         bundle of Plugs for this domain.
     * @param $current The current state of the transaction.
     * @return $success If the transaction was successful.
     * @return $result The return data of the transaction.
     */
    function _execute(PlugTypesLib.Current memory $current)
        internal
        enforceCurrent($current)
        returns (bool $success, bytes memory $result)
    {
        /// @dev Make the external call with a standard call.
        ($success, $result) = address($current.target).call{
            value: $current.value
        }($current.data);

        /// @dev If the call failed, bubble up the revert reason if possible.
        PlugLib.bubbleRevert($success, $result);
    }
}
