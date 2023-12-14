// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { PlugVaultSocket } from "./Plug.Vault.Socket.sol";
import { PlugFactorySocketLib } from "../libraries/Plug.Factory.Socket.Lib.sol";
import { LibClone } from "solady/src/utils/LibClone.sol";

/**
 * @title Plug Factory
 * @notice This contract is responsible for deploying new Plug Vaults that can be used
 *         as personal relays for an individual. The owner can execute transactions through
 *         the vaults, and the vaults can be used to store funds and/or NFTs. The vaults
 *         are deployed using the Beacon Proxy pattern, and the owner can upgrade the
 *         implementation at any time.
 * @author @nftchance (chance@utc24.io)
 * @author @vectorized (https://github.com/Vectorized/solady/blob/main/src/accounts/ERC4337Factory.sol)
 */
contract PlugFactorySocket is PlugSocket {
    /**
     * @notice Initializes a new Plug Vault contract.
     * @dev A vault is instantiated inside the factory in case funds are sent to the
     *      factory by mistake.
     * @param $name The name of the contract
     * @param $version The version of the contract
     */
    constructor(string memory $name, string memory $version) {
        _initializeSocket($name, $version);
    }

    /**
     * @notice Deploy a new Plug contract and initialize it.
     * @param $admin The admin of the vault.
     * @param $salt The salt of the vault.
     * @return $alreadyDeployed Whether or not the vault was already deployed.
     * @return $vault The address of the deployed vault.
     */
    function deploy(
        address $implementation,
        address $admin,
        bytes32 $salt
    )
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $vault)
    {
        /// @dev Make sure the user has provided a valid salt.
        LibClone.checkStartsWith($salt, $admin);

        /// @dev Deploy the new vault using a Beacon Proxy pattern.
        ($alreadyDeployed, $vault) = LibClone.createDeterministicERC1967(msg.value, $implementation, $salt);

        /// @dev If the vault was not already deployed, initialize it.
        if (!$alreadyDeployed) {
            /// @solidity memory-safe-assembly
            assembly {
                /// @dev Store the `$admin` argument.
                mstore(0x14, $admin)
                /// @dev Store the call data for the `initialize(address)` function.
                mstore(0x00, 0xc4d66de8000000000000000000000000)
                if iszero(call(gas(), $vault, 0, 0x10, 0x24, codesize(), 0x00)) {
                    returndatacopy(mload(0x40), 0x00, returndatasize())
                    revert(mload(0x40), returndatasize())
                }
            }

            /// @dev Emit an event for the creation of the Vault to make tracking
            ///		 things easier offchain.
            emit PlugFactorySocketLib.SocketDeployed($implementation, $admin, $salt);
        }
    }

    /**
     * @notice Predict the address of a new Plug Vault.
     * @param $salt The salt of the vault.
     * @return $vault The predicted address of the vault.
     */
    function getAddress(address $implementation, bytes32 $salt) public view returns (address $vault) {
        $vault = LibClone.predictDeterministicAddressERC1967($implementation, $salt, address(this));
    }

    /**
     * @notice Get the init code hash of the vaults.
     * @dev This is used to mine vanity addresses.
     * @return $initCodeHash The init code hash of the vaults.
     */
    function initCodeHash(address $implementation) public view virtual returns (bytes32 $initCodeHash) {
        $initCodeHash = LibClone.initCodeHashERC1967($implementation);
    }
}
