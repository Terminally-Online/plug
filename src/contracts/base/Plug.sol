// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PlugInterface } from "../interfaces/Plug.Interface.sol";

import {
    PlugLib,
    PlugTypesLib,
    PlugAddressesLib
} from "../libraries/Plug.Lib.sol";

import { PlugFactoryInterface } from
    "../interfaces/Plug.Factory.Interface.sol";
import { PlugSocketInterface } from
    "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug
 * @notice This contract represents a general purpose relay used to route signed
 *         intents to target Sockets.
 * @dev There is no need to approve assets to this contract as all transactions
 *      are executed through the Socket which will manage its own permissions
 *      that can be safely approved to interact with the assets of another account.
 * @author @nftchance (chance@onplug.io)
 */
contract Plug is PlugInterface {
    /// @dev Define the reference to the factory that enables counterfactual
    ///      deployment through having a presigned bundle of Plugs.
    PlugFactoryInterface factory =
        PlugFactoryInterface(PlugAddressesLib.PLUG_FACTORY_ADDRESS);

    /**
     * See {PlugInterface-plug}.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs)
        public
        payable
        virtual
        returns (bytes[] memory $results)
    {
        /// @dev Pass down the signature components and execute
        ///      the bundle from within the Socket that was declared.
        $results = _socket($livePlugs).plug(
            $livePlugs, msg.sender, gasleft()
        );
    }

    /**
     * See {PlugInterface-plug}.
     */
    function plug(PlugTypesLib.LivePlugs[] calldata $livePlugs)
        public
        payable
        virtual
        returns (bytes[][] memory $results)
    {
        /// @dev Load the stack.
        uint256 i;
        uint256 length = $livePlugs.length;
        $results = new bytes[][](length);

        /// @dev Iterate over the plugs and execute them.
        for (i; i < length; i++) {
            $results[i] = plug($livePlugs[i]);
        }
    }

    /**
     * See {PlugInterface-name}.
     */
    function name() public pure returns (string memory $name) {
        $name = "Plug";
    }

    /**
     * See {PlugInterface-symbol}.
     */
    function symbol() public pure returns (string memory $version) {
        $version = "PLUG";
    }

    /**
     * @notice Initialize the Plug with the factory and the implementation if
     *         it has not been initialized yet, otherwise just use the address
     *         included in the Plug bundle.
     * @param $livePlugs The signed bundle of Plugs being executed.
     * @return $socket The Socket to use.
     */
    function _socket(PlugTypesLib.LivePlugs calldata $livePlugs)
        internal
        virtual
        returns (PlugSocketInterface $socket)
    {
        /// @dev Pull the address of the Socket from the bundle.
        address socketAddress = $livePlugs.plugs.socket;

        /// @dev If the Socket has not yet been deployed, deploy it.
        if (socketAddress.code.length == 0) {
            /// @dev Call the factory that will handle the intent based deployment.
            (, address $socketAddress) =
                factory.deploy($livePlugs.plugs.salt, address(this));

            /// @dev Confirm the Socket was deployed to the right address.
            if (socketAddress != $socketAddress) {
                revert PlugLib.SocketAddressInvalid(
                    socketAddress, $socketAddress
                );
            }
        }

        /// @dev Load the Socket and return it.
        $socket = PlugSocketInterface(socketAddress);
    }
}
