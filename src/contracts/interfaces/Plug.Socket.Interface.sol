//SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.24;

/// @dev Shape declarations in the Plug framework.
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

interface PlugSocketInterface {
    /**
     * @notice Initialize the Socket with the ownership proxy of the Socket.
     * @param $ownership The address of the owner of the Socket.
     */
    function initialize(address $ownership) external;

    /**
     * @notice Get the address of the signer of the bundle of plugs.
     * @param $livePlugs The bundle of plugs to execute.
     * @return $signer The address of the signer of the bundle.
     */
    function signer(PlugTypesLib.LivePlugs calldata $livePlugs)
        external
        view
        returns (address $signer);

    /**
     * @notice Allows anyone to submit a plugs of signed plugs for processing.
     * @notice This version of the function will always be called by the Router.
     * @param $livePlugs The Plug bundle to execute.
     * @param $gas The gas to execute the plugs.
     * @return $results The return data of each plug executed.
     */
    function plug(
        PlugTypesLib.LivePlugs calldata $livePlugs,
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
