// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugFactoryInterface } from
    "../interfaces/Plug.Factory.Interface.sol";
import { PlugTradable } from "../abstracts/Plug.Tradable.sol";

import { PlugLib, PlugTypesLib } from "../libraries/Plug.Lib.sol";
import { LibClone } from "solady/src/utils/LibClone.sol";

import { PlugSocketInterface } from
    "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug Factory
 * @notice This contract is responsible for deploying new Plug Sockets that can be used
 *         as personal accounts for an individual. The owner can execute transactions
 *         through the Sockets. The Sockets are deployed using the Beacon Proxy
 *         pattern, and the owner can upgrade the implementation at any time. On top
 *         of being the deployment mechanism for the Sockets, the Factory also manages
 *         the ownership of the Sockets through the ERC721 standard allowing the
 *         Sockets to be traded on any major marketplace with ease.
 * @author @nftchance (chance@onplug.io)
 */
contract PlugFactory is PlugFactoryInterface, PlugTradable {
    /// @dev The mapping of the implementations of the vaults.
    mapping(uint16 => address) public implementations;

    /**
     * @notice Initialize a reference implementation that will not be
     *         with real intent of consumption.
     */
    constructor() {
        _initializeTradable(address(1), "");
    }

    /**
     * See { PlugFactoryInterface.initialize }
     */
    function initialize(
        address $owner,
        string memory $baseURI,
        address $implementation
    )
        public
        virtual
    {
        /// @dev Configure the starting state of the tradable functionatlity
        ///      that enables non-fungible representation of Socket ownership.
        _initializeTradable($owner, $baseURI);

        /// @dev Set the implementation of the first live version.
        implementations[0] = $implementation;
    }

    /**
     * @notice Set the implementation of a new version of the Socket.
     * @param $version The version of the vault.
     * @param $implementation The implementation of the vault.
     */
    function setImplementation(
        uint16 $version,
        address $implementation
    )
        public
        onlyOwner
    {
        /// @dev Ensure the implementation for this version has not already been set.
        if (implementations[$version] != address(0)) {
            revert PlugLib.ImplementationAlreadyInitialized($version);
        }

        /// @dev Set the implementation of the vault.
        implementations[$version] = $implementation;
    }

    /**
     * See { PlugFactoryInterface.deploy }
     */
    function deploy(
        bytes32 $salt,
        address $router
    )
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $socket)
    {
        /// @dev Recover the version of the Socket being deployed from the end
        ///      of the salt provided.
        uint16 version = uint16(uint256($salt));

        /// @dev Determine the address of the implementation provided the salt.
        address implementation = implementations[version];

        /// @dev Ensure the implementation is valid.
        if (implementation == address(0)) {
            revert PlugLib.ImplementationInvalid(version);
        }

        /// @dev Deploy the new vault using a Beacon Proxy pattern.
        ($alreadyDeployed, $socket) = LibClone
            .createDeterministicERC1967(msg.value, implementation, $salt);

        /// @dev If the vault was not already deployed, initialize it.
        if (!$alreadyDeployed) {
            /// @dev Recover the admin of the Socket by unpacking the signer from the salt.
            ///      The admin is the first 20 bytes of the salt and nonce is the last
            ///      12 bytes of the salt.
            address admin = address(uint160(uint256($salt) >> 96));

            /// @dev Emit an event for the creation of the Vault to make
            ///      tracking things easier offchain.
            emit PlugLib.SocketDeployed(implementation, admin, $salt);

            /// @dev Initialize the Socket with the ownership proxy pointing
            ///      this factory that is deploying the Socket.
            PlugSocketInterface($socket).initialize(
                address(this), $router
            );

            /// @dev Mint the transferable ownership token to the signer that
            ///      created the intent which is implicitly the Socket admin
            ///      with the id of the token as integer representation of
            ///      the address the Socket was deployed to.
            _mint(admin, uint256(uint160(address($socket))));
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
        $vault = LibClone.predictDeterministicAddressERC1967(
            $implementation, $salt, address(this)
        );
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
