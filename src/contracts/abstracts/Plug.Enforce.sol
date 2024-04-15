// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypesLib, PlugTypes } from "./Plug.Types.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";
import { PlugConnectorInterface } from "../interfaces/Plug.Connector.Interface.sol";

/**
 * @title Plug Enforce
 * @notice The enforcement mechanisms of Plug to ensure that transactions
 *         are only executed as defined.
 * @dev Inheriting contracts must implement the logic for:
 *      - `_enforceSignature`
 *      - `_enforceRouter`.
 * @author @nftchance (chance@onplug.io)
 */
abstract contract PlugEnforce is PlugTypes {
    /**
     * @notice Modifier to enforce the router of the transaction.
     * @dev Apply to this to functions that are designed to be access by Routers.
     * @dev Implicitly the address is assumed to be the current sender.
     */
    modifier enforceRouter() {
        if (_enforceRouter(msg.sender) == false) {
            revert PlugLib.RouterInvalid(msg.sender);
        }
        _;
    }

    /**
     * @notice Modifier to enforce the signer of the transaction.
     * @dev Apply to this to functions that are designed to execute a bundle
     *      of Plugs regardless of whether through a Router or or direct access.
     * @param $input The LivePlugs the definition of execution as well as the
     *               signature used to verify the execution permission.
     */
    modifier enforceSignature(PlugTypesLib.LivePlugs calldata $input) {
        // require(_enforceSignature($input), "Plug:invalid-signature");
        if (_enforceSignature($input) == false) {
            revert PlugLib.SignatureInvalid();
        }
        _;
    }

    /**
     * @notice Confirm that the only specified routers can execute the transaction.
     * @dev If you would like to limit the available routers override this
     *      function in your contract with the additional logic.
     * @param $router The router of the transaction.
     */
    function _enforceRouter(address $router) internal view virtual returns (bool $allowed);

    /**
     * @notice Confirm that signer has permission to declare execution of a
     *         Plug bundle on the parent-socket that inherits this contract.
     * @dev Inheriting contracts must implement the logic of this function to make
     *      sure that only signatures intended for this scope are allowed.
     * @param $input The LivePlugs object that contains the Plugs object as well as
     *               the signature defining the permission to execute the bundle.
     */
    function _enforceSignature(PlugTypesLib.LivePlugs calldata $input)
        internal
        view
        virtual
        returns (bool $allowed);
}
