// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";
import { PlugSimulation } from "./Plug.Simulation.sol";
import { PlugTypesLib } from "./Plug.Types.sol";

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance (chance@utc24.io)
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
contract PlugSocket is PlugSocketInterface, PlugSimulation {
    /**
     * See {IPlug-plug}.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs)
        external
        payable
        returns (bytes[] memory $results)
    {
        /// @dev Determine who signed the intent.
        address intentSigner = getLivePlugsSigner($livePlugs);

        /// @dev Invoke the plugs.
        $results = _plug($livePlugs.plugs, intentSigner);
    }

    /**
     * See {IPlug-plugContract}.
     *
     * TODO: Finish the implementation of this make sure it is secure as this
     *       allows existing contracts to declare the execution of an intent
     *       beyond just EOAs and that is a growing usecase now that we not
     *       only have Gnosis Safes, but also EIP-4337.
     */
    function plugContract(PlugTypesLib.Plugs calldata $plugs)
        external
        payable
        returns (bytes[] memory $result)
    {
        $result = _plug($plugs, msg.sender);
    }
}
