// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";
import { PlugCore } from "./Plug.Core.sol";
import { PlugEnforce } from "./Plug.Enforce.sol";
import { ReentrancyGuard } from "solady/utils/ReentrancyGuard.sol";
import { PlugTypesLib } from "./Plug.Types.sol";

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance (chance@onplug.io)
 */
abstract contract PlugSocket is PlugSocketInterface, PlugCore, PlugEnforce, ReentrancyGuard {
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
        returns (PlugTypesLib.Result[] memory $results)
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
        enforceSender
        nonReentrant
        returns (PlugTypesLib.Result[] memory $results)
    {
        $results = _plug($plugs, address(0), 0);
    }

    /**
     * See {PlugSocketInterface-revoke}.
     */
    function revoke(bytes32 $plugsHash, bool $isRevoked) public virtual;
}
