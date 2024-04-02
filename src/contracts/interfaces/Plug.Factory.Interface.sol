//SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

interface PlugFactoryInterface {
    /**
     * @notice Initialize this implementation of the factory.
     * @param $owner The address of the owner.
     * @param $baseURI The base URI of the factory.
     */
    function initialize(
        address $owner,
        string memory $baseURI,
        address $implementation
    )
        external;

    /**
     * @notice Deploy a new Socket and initialize it.
     * @dev This version is used to interface with directly enabling the ability
     *      to deploy multiple Sockets at once from a single Plug bundle.
     * @param $salt The salt of the Socket.
     * @return $alreadyDeployed Whether or not the Socket was already deployed.
     * @return $socket The address of the deployed Socket.
     */
    function deploy(
        bytes32 $salt,
        address $router
    )
        external
        payable
        returns (bool $alreadyDeployed, address $socket);

    /**
     * @notice Predict the address of a new Plug Vault.
     * @param $salt The salt of the vault.
     * @return $vault The predicted address of the vault.
     */
    function getAddress(
        address $implementation,
        bytes32 $salt
    )
        external
        view
        returns (address $vault);

    /**
     * @notice Get the init code hash of the vaults.
     * @dev This is used to mine vanity addresses.
     * @return $initCodeHash The init code hash of the vaults.
     */
    function initCodeHash(address $implementation)
        external
        view
        returns (bytes32 $initCodeHash);
}
