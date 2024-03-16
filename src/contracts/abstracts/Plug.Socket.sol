// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.24;

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";
import { PlugCore } from "./Plug.Core.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { PlugTypesLib } from "./Plug.Types.sol";

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugSocket is
    PlugSocketInterface,
    PlugCore,
    ReentrancyGuard
{
    /**
     * See {PlugSocketInterface-signer}.
     */
    function signer(PlugTypesLib.LivePlugs calldata $livePlugs)
        external
        view
        returns (address $signer)
    {
        /// @dev Determine the address that signed the Plug bundle.
        $signer = getLivePlugsSigner($livePlugs);
    }

    /**
     * See {PlugSocketInterface-plug}.
     *
     * @dev Process the Plug bundle with an external Executor.
     */
    function plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $signer,
        uint256 $gas
    )
        external
        payable
        virtual
        enforceRouter
        enforceSigner($signer)
        nonReentrant
        returns (bytes[] memory $results)
    {
        $results = _plug($plugs, $plugs.executor, $gas);
    }

    /**
     * See {PlugSocketInterface-plug}.
     *
     * @dev Process the Plug bundle without an external Executor.
     */
    function plug(PlugTypesLib.Plugs calldata $plugs)
        external
        payable
        virtual
        enforceSigner(msg.sender)
        nonReentrant
        returns (bytes[] memory $results)
    {
        $results = _plug($plugs, address(0), 0);
    }
}
