//SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

interface PlugInterface {
    /**
     * @notice Allows anyone to submit a signed bundle of Plugs for processing.
     * @param $livePlugs The Plug bundle to execute.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs) external payable;

    /**
     * @notice A batch implementation of the sister `plug` function that enables
     *         the ability to execute multiple bundles at once.
     * @param $livePlugs The set of Plug bundles to execute.
     */
    function plug(PlugTypesLib.LivePlugs[] calldata $livePlugs) external payable;
}
