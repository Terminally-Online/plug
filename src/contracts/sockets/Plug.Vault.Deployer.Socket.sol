// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import {PlugSocket} from '../abstracts/Plug.Socket.sol';
import {Ownable} from 'solady/src/auth/Ownable.sol';
import {ERC1967Factory} from 'solady/src/utils/ERC1967Factory.sol';

/**
 * @title Plug Vault Socket Deployer
 * @notice This contract is responsible for deploying new Plug Vaults that can be used
 *         as personal relays for the owner. The owner can execute transactions through
 *         the vaults, and the vaults can be used to store funds and/or NFTs. The vaults
 *         are deployed using the Beacon Proxy pattern, and the owner can upgrade the
 *         implementation at any time.
 * @author @nftchance (chance@utc24.io)
 */
contract PlugVaultSocketDeployer is PlugSocket, ERC1967Factory {
	/// @dev The address of the active Plug Vault implementation.
	address implementation;

	/// @dev The nonce of the sender to determine the next vault address.
	mapping(address => uint256) senderToNonce;

	constructor(
		address $implementation
	) PlugSocket('PlugVaultSocket', '0.0.1') {
		/// @dev Set the version of vaults to deploy.
		implementation = $implementation;
	}

	/**
	 * @notice Deploy a new Plug Vault that can be used as a personal relay.
	 * @param $data The data to initialize the vault with.
	 * @return $vault The address of the deployed vault.
	 */
	function deploy(bytes calldata $data) external returns (address $vault) {
		$vault = deployDeterministicAndCall(
			implementation,
			_msgSender(),
			keccak256(
				abi.encodePacked(_msgSender(), senderToNonce[_msgSender()]++)
			),
			$data
		);
	}

	/**
	 * @notice Predict the address of a new Plug Vault.
	 * @param $admin The admin of the vault.
	 * @param $nonce The nonce of the admin.
	 * @return $address The predicted address of the vault.
	 */
	function getAddress(
		address $admin,
		uint256 $nonce
	) public view returns (address) {
		bytes32 salt = keccak256(abi.encodePacked($admin, $nonce));

		return predictDeterministicAddress(salt);
	}

	/**
	 * @notice Get all of the addresses of the vaults for a given admin.
	 * @param $admin The admin of the vaults.
	 * @return $addresses The addresses of the vaults.
	 */
	function getAddresses(
		address $admin
	) public view returns (address[] memory) {
		uint256 vaults = senderToNonce[$admin];

		address[] memory addresses = new address[](vaults);

		for (uint256 i = 0; i < vaults; i++) {
			addresses[i] = getAddress($admin, i);
		}
	}
}
