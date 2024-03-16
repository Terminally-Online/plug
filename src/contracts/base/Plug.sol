// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { PlugInterface } from "../interfaces/Plug.Interface.sol";
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
contract Plug is PlugInterface {
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
        uint256 gas = gasleft();

        /// @dev Load the Plug Socket.
        PlugSocketInterface socket =
            PlugSocketInterface($livePlugs.plugs.socket);

        /// @dev Confirm the executor is allowed to execute the transaction.
        ///      This is done here instead of a modifier so that the gas
        ///      snapshot accounts for the additional gas cost of the require.
        require(
            msg.sender == $livePlugs.plugs.executor
                || $livePlugs.plugs.executor == address(0),
            "Plug:invalid-executor"
        );

        /// @dev Recover the address that signed the bundle of Plugs.
        address signer = socket.signer($livePlugs);

        /// @dev Pass down the now-verified signature components and execute
        ///      the bundle from within the Socket that was declared.
        $results = socket.plug($livePlugs.plugs, signer, gas);
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
}
