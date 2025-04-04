// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {PlugTypesLib} from '../abstracts/Plug.Types.sol';
import {PlugAddressesLib} from './Plug.Addresses.Lib.sol';

library PlugLib {
	/////////////////////////////////////////////////
	//                     PLUG                    //
	/////////////////////////////////////////////////
	event PlugResult(uint8 index, bytes32 plugsHash, PlugTypesLib.Result reason);

	error PlugFailed(uint8 $index, string $reason);

	/////////////////////////////////////////////////
	//                    SOCKET                   //
	/////////////////////////////////////////////////

	event SocketDeployed(
		address indexed implementation,
		address indexed vault,
		bytes32 salt
	);
	event SocketOwnershipTransferred(
		address indexed previousOwner,
		address indexed newOwner,
		bytes32 imageHash
	);

	error SocketAddressInvalid(address $intended, address $socket);
	error SocketAddressEmpty(address $socket);

	error SaltInvalid(address $implementation, address $admin);
	error CallerInvalid(address $expected, address $reality);
	error RouterInvalid(address $reality);
	error TypeInvalid(uint8 $reality);
	error CompensationFailed(address $recipient, uint256 $value);

	/////////////////////////////////////////////////
	//                   REWARDS                   //
	/////////////////////////////////////////////////

	event NewRewardPeriod(
		uint256 indexed period,
		bytes32 merkleRoot,
		uint256 totalAmount
	);
	event RewardClaimed(
		uint256 indexed period,
		address indexed user,
		uint256 amount
	);

	error InvalidMerkleProof();
	error PeriodNotInitialized();
	error RewardsAlreadyClaimed();
	error InsufficientRewardBalance();
	error ZeroAmount();
}
