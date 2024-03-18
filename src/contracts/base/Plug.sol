// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { PlugInterface } from "../interfaces/Plug.Interface.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { PlugFactory } from "../base/Plug.Factory.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug
 * @notice This contract represents a general purpose relay socket that can be
 *         used to route transactions to other contracts.
 * @dev There is no need to approve assets to this contract as all transactions
 *      are executed through the socket which will manage its own permissions
 *      can be safely approved to interact with the assets of another account.
 * @author @nftchance (chance@utc24.io)
 */
contract Plug is PlugInterface, Ownable {
    /// @dev The factory that enables automatic Socket deployment.
    PlugFactory factory;

    /// @dev The current version of Sockets that are automatically deployed
    ///      when a new Socket is created when deploying at the time of intent
    ///      fulfillment taking place.
    address implementation;

    /**
     * See {PlugInterface-plug}.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs)
        public
        payable
        virtual
        returns (bytes[] memory $results)
    {
        /// @dev Snapshot how much gas the transaction has.
        uint256 gas = (gasleft() * 63) / 64;

        /// @dev Confirm the executor is allowed to execute the transaction.
        ///      This is done here instead of a modifier so that the gas
        ///      snapshot accounts for the additional gas cost of the require.
        require(
            msg.sender == $livePlugs.plugs.executor
                || $livePlugs.plugs.executor == address(0),
            "Plug:invalid-executor"
        );

        /// @dev Pass down the now-verified signature components and execute
        ///      the bundle from within the Socket that was declared.
        $results =
            _socket($livePlugs.plugs.socket, $livePlugs).plug($livePlugs, gas);
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
     * @param $socketAddress The address of the Socket to use.
     * @param $livePlugs The bundle of plugs to execute.
     * @return $socket The Socket to use.
     */
    function _socket(
        address $socketAddress,
        PlugTypesLib.LivePlugs calldata $livePlugs
    )
        internal
        virtual
        returns (PlugSocketInterface $socket)
    {
        /// @dev Pull the address of the Socket from the bundle.
        address socketAddress = $livePlugs.plugs.socket;

        /// @dev If the Socket has not yet been deployed, deploy it.
        if ($socketAddress.code.length == 0) {
            /// @dev Call the factory that will handle the intent based deployment.
            (, $socketAddress) = factory.deploy(implementation, $livePlugs);

            /// @dev Confirm the Socket was deployed to the right address.
            require(
                $socketAddress == socketAddress, "Plug:invalid-socket-address"
            );
        }

        /// @dev Load the Socket and return it.
        $socket = PlugSocketInterface(socketAddress);
    }
}
