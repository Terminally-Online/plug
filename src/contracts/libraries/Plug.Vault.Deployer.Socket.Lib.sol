// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

library PlugVaultDeployerSocketLib {
	event SocketDeployed(
		address indexed $deployer,
		address indexed $vault,
		bytes32 $salt
	);
}
