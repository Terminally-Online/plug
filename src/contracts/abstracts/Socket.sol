// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {IPlug} from '../interfaces/IPlug.sol';
import {PlugCore} from './PlugCore.sol';

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
abstract contract Socket is PlugCore, IPlug {
	/**
	 * @notice Instantiates a new Plug contract.
	 * @param $name The name of the contract
	 * @param $version The version of the contract
	 */
	constructor(
		string memory $name,
		string memory $version
	) PlugCore($name, $version) {}

	/**
	 * See {IPlug-plugContract}.
	 */
	function plugContract(
		Plug[] calldata $plugs
	) external returns (bool $success) {
		$success = _plug($plugs, msg.sender);
	}

	/**
	 * See {IPlug-plug}.
	 */
	function plug(
		LivePlugs[] calldata $livePlugs
	) external returns (bool $success) {
		/// @dev Load the stack.
		uint256 i;

		/// @dev Loop through the signed plugs.
		for (i; i < $livePlugs.length; ) {
			/// @dev Load the signed intent as a hot reference.
			LivePlugs calldata livePlugs = $livePlugs[i];

			/// @dev Determine who signed the intent.
			address intentSigner = getLivePlugsSigner(livePlugs);

			/// @dev Load the plugs as a hot reference.
			Plugs calldata plugs = livePlugs.plugs;

			/// @dev Prevent replay attacks by enforcing replay protection.
			_enforceBreaker(intentSigner, plugs.breaker);

			/// @dev Invoke the plugs.
			$success = _plug(plugs.plugs, intentSigner);

			unchecked {
				++i;
			}
		}
	}
}
