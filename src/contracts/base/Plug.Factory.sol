// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugFactoryInterface } from "../interfaces/Plug.Factory.Interface.sol";

import { PlugLib } from "../libraries/Plug.Lib.sol";
import { LibClone } from "solady/utils/LibClone.sol";

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug Factory
 * @notice This contract is responsible for deploying new Plug Sockets that can be used
 *         as personal accounts for an individual. The Sockets are deployed using the
 *         Beacon Proxy pattern, and the owner can upgrade the implementation at any time.
 * @author @nftchance (chance@onplug.io)
 */
contract PlugFactory is PlugFactoryInterface {
    /**
     * See { PlugFactoryInterface.deploy }
     */
    function deploy(bytes calldata $salt)
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $socketAddress)
    {
        /// @dev Recover the packed implementation and admin from the salt.
        (uint96 nonce, address admin, address oneClicker, address implementation) =
            abi.decode($salt, (uint96, address, address, address));

        /// @dev Ensure the implementation is valid.
        if (implementation == address(0) || admin == address(0)) {
            revert PlugLib.SaltInvalid(implementation, admin);
        }

        /// @dev Create the deployment salt using the nonce and the admin.
        bytes32 salt = bytes32(abi.encodePacked(uint96(nonce), bytes20(admin)));

        /// @dev Deploy the new vault using a Beacon Proxy pattern.
        ($alreadyDeployed, $socketAddress) =
            LibClone.createDeterministicERC1967(msg.value, implementation, salt);

        /// @dev If the vault was not already deployed, initialize it.
        if (!$alreadyDeployed) {
            /// @dev Emit an event for the creation of the Socket to make
            ///      tracking things easier offchain.
            emit PlugLib.SocketDeployed(implementation, admin, salt);

            /// @dev Initialize the Socket with the ownership proxy pointing
            ///      this factory that is deploying the Socket.
            PlugSocketInterface($socketAddress).initialize(admin, oneClicker);
        }
    }

    /**
     * See { PlugFactoryInterface.getAddress }
     */
    function getAddress(
        address $implementation,
        bytes32 $salt
    )
        public
        view
        returns (address $vault)
    {
        $vault = LibClone.predictDeterministicAddressERC1967($implementation, $salt, address(this));
    }

    /**
     * See { PlugFactoryInterface.initCodeHash }
     */
    function initCodeHash(address $implementation)
        public
        view
        virtual
        returns (bytes32 $initCodeHash)
    {
        $initCodeHash = LibClone.initCodeHashERC1967($implementation);
    }
}
