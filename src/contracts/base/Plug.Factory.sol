// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { PlugTradable } from "../abstracts/Plug.Tradable.sol";

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";
import { LibClone } from "solady/src/utils/LibClone.sol";

/**
 * @title Plug Factory
 * @notice This contract is responsible for deploying new Plug Vaults that can be used
 *         as personal relays for an individual. The owner can execute transactions
 *         through the Sockets. The Sockets are deployed using the Beacon Proxy
 *         pattern, and the owner can upgrade the implementation at any time. On top
 *         of being the deployment mechanism for the Sockets, the Factory also manages
 *         the ownership of the Sockets through the ERC721 standard allowing the
 *         Sockets to be traded on any major marketplace with ease.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugFactory is PlugTradable {
    mapping(uint160 => uint256) public nonce;

    constructor(
        address $owner,
        string memory $baseURI
    )
        PlugTradable($owner, $baseURI)
    { }

    /**
     * @notice Deploy a new Plug contract and initialize it.
     * @param $implementation The implementation of the vault.
     * @param $livePlugs The bundle of Plugs that will be used to initialize the Socket.
     * @return $alreadyDeployed Whether or not the Socket was already deployed.
     * @return $socket The address of the deployed Socket.
     */
    function deploy(
        address $implementation,
        PlugTypesLib.LivePlugs calldata $livePlugs
    )
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $socket)
    {
        /// @dev Get the salt from the livePlugs.
        bytes32 salt = $livePlugs.plugs.salt;

        /// @dev Deploy the new vault using a Beacon Proxy pattern.
        ($alreadyDeployed, $socket) = LibClone.createDeterministicERC1967(
            msg.value, $implementation, salt
        );

        /// @dev If the vault was not already deployed, initialize it.
        if (!$alreadyDeployed) {
            /// @dev Cast the address to a Socket interface.
            PlugSocketInterface socket = PlugSocketInterface($socket);

            /// @dev Recover the admin of the Socket by retrieving the signer
            ///      of the first intent that flowed through the Socket.
            address admin = socket.signer($livePlugs);

            _afterDeploy($implementation, admin, salt, socket);
        }
    }

    /**
     * @notice Deploy a new Socket and initialize it.
     * @dev This version is used to interface with directly enabling the ability
     *      to deploy multiple Sockets at once from a single Plug bundle.
     * @param $implementation The implementation of the vault.
     * @param $admin The admin of the Socket.
     * @param $salt The salt of the Socket.
     * @return $alreadyDeployed Whether or not the Socket was already deployed.
     * @return $socket The address of the deployed Socket.
     */
    function deploy(
        address $implementation,
        address $admin,
        bytes32 $salt
    )
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $socket)
    {
        /// @dev Deploy the new vault using a Beacon Proxy pattern.
        ($alreadyDeployed, $socket) = LibClone.createDeterministicERC1967(
            msg.value, $implementation, $salt
        );

        /// @dev If the vault was not already deployed, initialize it.
        if (!$alreadyDeployed) {
            _afterDeploy(
                $implementation, $admin, $salt, PlugSocketInterface($socket)
            );
        }
    }

    /**
     * @notice Handle the state updates after the deployment of a new Socket.
     * @param $implementation The implementation of the vault.
     * @param $admin The admin of the Socket.
     * @param $salt The salt of the Socket.
     * @param $socket The deployed Socket.
     */
    function _afterDeploy(
        address $implementation,
        address $admin,
        bytes32 $salt,
        PlugSocketInterface $socket
    )
        internal
        virtual
    {
        /// @dev Make sure the user has provided a valid salt.
        LibClone.checkStartsWith($salt, $admin);

        /// @dev Initialize the Socket with the ownership proxy pointing
        ///      this factory that is deploying the Socket.
        $socket.initialize(address(this));

        /// @dev Emit an event for the creation of the Vault to make
        ///      tracking things easier offchain.
        emit PlugLib.SocketDeployed($implementation, $admin, $salt);

        /// @dev Mint the transferable ownership token to the signer that
        ///      created the intent which is implicitly the Socket admin.
        mint($admin, address($socket));
    }

    /**
     * @notice Predict the address of a new Plug Vault.
     * @param $salt The salt of the vault.
     * @return $vault The predicted address of the vault.
     */
    function getAddress(
        address $implementation,
        bytes32 $salt
    )
        public
        view
        returns (address $vault)
    {
        $vault = LibClone.predictDeterministicAddressERC1967(
            $implementation, $salt, address(this)
        );
    }

    /**
     * @notice Get the init code hash of the vaults.
     * @dev This is used to mine vanity addresses.
     * @return $initCodeHash The init code hash of the vaults.
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
