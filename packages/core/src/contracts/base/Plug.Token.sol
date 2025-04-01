// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {SuperChainERC20} from '@op/SuperchainERC20.sol';
import {Initializable} from 'solady/utils/Initializable.sol';
import {Ownable} from 'solady/auth/Ownable.sol';

contract PlugToken is Initializable, Ownable, SuperChainERC20 {
	uint32 public unlock;

	mapping(address => bool) senderToAllowed;

	constructor() {
		_initializeOwner(address(1));
	}

	function initialize(address $owner, uint32 $unlock) initializer {
		_initialize($owner);

		unlock = $unlock;
	}

	function name() public returns (string) {
		return 'Plug';
	}

	function symbol() public returns (string) {
		return 'PLUG';
	}

	function setUnlock(uint32 $unlock) onlyOwner {
		unlock = $unlock;
	}

	function _beforeTokenTransfer(
		address $from,
		address $to,
		uint256 $amount
	) internal virtual {
		bool validSender = block.timestamp >= unlock ||
			$from == address(0) ||
			senderToAllowed[$from];

		require(validSender, 'PlugToken:invalid-sender');
	}
}
