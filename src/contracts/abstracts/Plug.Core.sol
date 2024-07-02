// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {console2} from 'forge-std/console2.sol';

import {PlugTypes} from './Plug.Types.sol';
import {PlugLib, PlugTypesLib, PlugAddressesLib} from '../libraries/Plug.Lib.sol';
import {PlugConnectorInterface} from '../interfaces/Plug.Connector.Interface.sol';

/**
 * @title PlugCore
 * @notice The core logic for executing Plugs and managing the state of the
 *         Plugs that are being executed. All execution logic is handled
 *         through the bundled transactions except for revocation as all
 *         executions should have the capability to be revoked.
 * @author @nftchance (chance@onplug.io)
 */
abstract contract PlugCore is PlugTypes {
	/// @dev Keep track of which Plugs have been revoked.
	mapping(bytes32 => bool) public isRevoked;

	/**
	 * @notice Manage the revocation state of a bundle of Plugs.
	 * @param $plugsHash The hash of the Plugs to revoke.
	 * @param $isRevoked The state to set the Plugs to.
	 */
	function _revoke(bytes32 $plugsHash, bool $isRevoked) internal {
		/// @dev Update the internal state of the Plugs to reflect
		///      the expected revocation state.
		isRevoked[$plugsHash] = $isRevoked;

		/// @dev Announce an update of the revocation state to make in-app
		///      management straightforward.
		emit PlugLib.PlugsRevocationUpdated($plugsHash, $isRevoked);
	}

	/**
	 * @notice Execute a bundle of Plugs.
	 * @param $plugs The Plugs to execute containing the bundle and side effects.
	 * @param $solver Encoded data defining the Solver and compensation.
	 * @param $gas Snapshot of gas at the start of interaction.
	 * @return $results The return data of the plugs.
	 */
	function _plug(
		PlugTypesLib.Plugs calldata $plugs,
		address $solver,
		uint256 $gas
	) internal returns (PlugTypesLib.Result[] memory $results) {
		/// @dev Hash the body of the object to ensure the integrity of
		///      the (bundle of) Plugs that are being executed.
		bytes32 plugsHash = getPlugsHash($plugs);

		/// @dev Revert if the Plugs have been revoked before executing,
		///      otherwise allow processing to continue.
		if (isRevoked[plugsHash]) revert PlugLib.PlugsRevoked();

		/// @dev Load the Plug stack into memory for cheaper access.
		uint256 length = $plugs.plugs.length;
		$results = new PlugTypesLib.Result[](length);

		/// @dev Save the object into memory to avoid multiple creations
		///      of the same object.
		PlugTypesLib.Plug calldata plug;

		/// @dev Iterate over the Plugs that are held within this bundle
		///      an execute each of them. Each respectively may be a
		///      condition being enforced or an outcome focused transaction.
		for (uint256 i; i < length; i++) {
			/// @dev Place the active Plug in the shorter reference stack.
			plug = $plugs.plugs[i];

			/// @dev If the call has an associated value, ensure the contract
			///      has enough balance to cover the cost of the call.
			if (address(this).balance < plug.value) {
				revert PlugLib.ValueInvalid(
					plug.target,
					plug.value,
					address(this).balance
				);
			}

			($results[i].success, $results[i].result) = plug.target.call{
				value: plug.value
			}(plug.data[1:]);

			/// @dev If the call failed, bubble up the revert reason if needed.
			PlugLib.bubbleRevert($results[i].success, $results[i].result);
		}

		/// @dev Pay the Solver for the gas used if it was not open-access.
		if ($plugs.solver.length != 0) {
			/// @dev Unpack the solver data from the encoded Solver data.
			(
				uint96 maxPriorityFeePerGas,
				uint96 maxFeePerGas,
				address solver
			) = abi.decode($plugs.solver, (uint96, uint96, address));

			/// @dev Confirm the Solver is allowed to execute the transaction.
			///      This is done here instead of a modifier so that the gas
			///      snapshot accounts for the additional gas cost of the require.
			if (solver != $solver) {
				revert PlugLib.SolverInvalid(solver, $solver);
			}

			/// @dev Calculate the gas price based on the current block.
			uint256 value = maxPriorityFeePerGas + block.basefee;
			/// @dev Determine which gas price to use based on if it is a legacy
			///      transaction (on a chain that does not support it) or if the
			///      the transaction is submit post EIP-1559.
			value = maxFeePerGas == maxPriorityFeePerGas
				? maxFeePerGas
				: maxFeePerGas < value
					? maxFeePerGas
					: value;

			/// @dev Augment the native gas price with the Solver "gas" fee.
			value = ($gas - gasleft()) * value;

			/// @dev Transfer the money the Solver is owed and confirm it
			///      the transfer is successful.
			(bool success, ) = solver.call{value: value}('');
			if (success == false) {
				revert PlugLib.CompensationFailed(solver, value);
			}
		}

		emit PlugLib.PlugsExecuted(plugsHash, $results);
	}
}
