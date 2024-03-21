//SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

/// @dev Shape declarations in the Plug framework.
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

interface PlugSocketInterface {
    /**
     * @notice Initialize the Socket with the ownership proxy of the Socket.
     * @param $ownership The address of the owner of the Socket.
     */
    function initialize(address $ownership) external;

    /**
     * @notice Allows anyone to submit a plugs of signed plugs for processing.
     * @notice This version of the function will always be called by the Router.
     * @param $livePlugs The Plug bundle to execute.
     * @param $solver The address of the Solver.
     * @param $gas The gas to execute the plugs.
     * @return $results The return data of each plug executed.
     */
    function plug(
        PlugTypesLib.LivePlugs calldata $livePlugs,
        address $solver,
        uint256 $gas
    )
        external
        payable
        returns (bytes[] memory $results);

    /**
     * @notice Allows anyone to submit a plugs of signed plugs for processing.
     * @notice This version of the function will always be called by the Router.
     * @param $plugs The Plug bundle to execute.
     * @return $results The return data of each plug executed.
     */
    function plug(PlugTypesLib.Plugs calldata $plugs)
        external
        payable
        returns (bytes[] memory $results);
}
