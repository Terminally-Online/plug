// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugSocketInterface } from
    "../interfaces/Plug.Socket.Interface.sol";
import { PlugCore } from "./Plug.Core.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { PlugTypesLib } from "./Plug.Types.sol";

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance (chance@onplug.io)
 */
abstract contract PlugSocket is
    PlugSocketInterface,
    PlugCore,
    ReentrancyGuard
{
    /**
     * See {PlugSocketInterface-plug}.
     */
    function plug(
        PlugTypesLib.LivePlugs calldata $livePlugs,
        address $solver,
        uint256 $gas
    )
        external
        payable
        virtual
        enforceRouter
        enforceSignature($livePlugs)
        nonReentrant
        returns (bytes[] memory $results)
    {
        $results = _plug($livePlugs.plugs, $solver, $gas);
    }

    /**
     * See {PlugSocketInterface-plug}.
     */
    function plug(PlugTypesLib.Plugs calldata $plugs)
        external
        payable
        virtual
        nonReentrant
        returns (bytes[] memory $results)
    {
        // TODO: Make sure only intended signers have permission to submit here.

        $results = _plug($plugs, address(0), 0);
    }
}
