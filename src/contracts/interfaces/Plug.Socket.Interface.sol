//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

/// @dev Shape declarations in the Plug framework.
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

interface PlugSocketInterface {
    /**
     * @notice Allows anyone to submit a plugs of signed plugs for processing.
     * @param $livePlugs The plugs of signed plugs to process.
     * @return $results The return data of each plug executed.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs) external payable returns (bytes[] memory $results);

    /**
     * @notice Allows a smart contract to submit a plugs of plugs for processing,
     *         allowing itself to be the delegate.
     * @param $plugs The plugs of plugs to execute.
     * @return $results The return data of each plug executed.
     */
    function plugContract(PlugTypesLib.Plug[] calldata $plugs) external payable returns (bytes[] memory $results);
}
