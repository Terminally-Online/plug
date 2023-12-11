// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugCore} from './Plug.Core.sol';
import {PlugTypesLib} from './Plug.Types.sol';
import {IFuse} from '../interfaces/IFuse.sol';
import {PlugSimulationLib} from '../libraries/Plug.Simulation.Lib.sol';

/**
 * @title PlugSimulation
 * @notice The simulation contract that enables the ability to simulate an entire
 *         bundle at once or through a segmented route of execution using
 *         explicit plug keys.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugSimulation is PlugCore {
	/**
	 * @notice Simulate the entire bundle of LivePlugs.
	 * @param $livePlugs The bundle of LivePlugs to process in the simulation.
	 * @return results An array of results describing the status of the bundle.
	 */
	function simulate(
		PlugTypesLib.LivePlugs[] calldata $livePlugs
	) public view virtual returns (PlugSimulationLib.Result[] memory) {
		uint256 resultsLength = length($livePlugs);

		if (resultsLength == 0) return new PlugSimulationLib.Result[](0);

		PlugSimulationLib.Result[]
			memory results = new PlugSimulationLib.Result[](resultsLength);

		for (uint8 i; i < $livePlugs.length; i++) {
			PlugTypesLib.LivePlugs calldata livePlugs = $livePlugs[i];

			for (uint8 j; j < livePlugs.plugs.plugs.length; j++) {
				PlugTypesLib.Plug memory plug = livePlugs.plugs.plugs[j];

				for (uint8 k; k < plug.pins.length; k++) {
					PlugTypesLib.LivePin memory livePin = plug.pins[k];

					for (uint8 l; l < livePin.pin.fuses.length; l++) {
						results[results.length - resultsLength--] = simulate(
							plug,
							livePin,
							livePin.pin.fuses[l]
						);
					}
				}
			}
		}

		return results;
	}

	/**
	 * @notice Simulate the execution of a bundle of live plugs through explicit indexes.
	 * @param $indexes The indexes of the plugs to simulate.
	 *        uint8 | 00 | livePlugsIndex
	 *        uint8 | 08 | plugsIndex
	 *        uint8 | 16 | pinIndex
	 *        uint8 | 24 | fuseIndex
	 * @param $livePlugs The live plugs to simulate.
	 * @return $results The results of the simulation.
	 */
	function simulate(
		uint32[] memory $indexes,
		PlugTypesLib.LivePlugs[] calldata $livePlugs
	) public view virtual returns (PlugSimulationLib.Result[] memory) {
		if ($indexes.length == 0) return new PlugSimulationLib.Result[](0);

		PlugSimulationLib.Result[]
			memory $results = new PlugSimulationLib.Result[]($indexes.length);

		for (uint256 i; i < $indexes.length; ) {
			uint32 index = $indexes[i];

			PlugTypesLib.Plug memory plug = $livePlugs[uint8(index)]
				.plugs
				.plugs[uint8(index >> 8)];
			PlugTypesLib.LivePin memory livePin = plug.pins[uint8(index >> 16)];

			$results[i] = simulate(
				plug,
				livePin,
				livePin.pin.fuses[uint8(index >> 24)]
			);

			unchecked {
				++i;
			}
		}

		return $results;
	}

	/**
	 * @notice Simulate the execution of a specific Plug Fuse.
	 * @param $plug The plug to simulate.
	 * @param $livePin The live pin to simulate.
	 * @param $fuse The fuse to simulate.
	 * @return $result The result of the simulation.
	 */
	function simulate(
		PlugTypesLib.Plug memory $plug,
		PlugTypesLib.LivePin memory $livePin,
		PlugTypesLib.Fuse memory $fuse
	) public view virtual returns (PlugSimulationLib.Result memory $result) {
		bytes32 pinHash = getLivePinHash($livePin);

		(bool $success, bytes memory $callback) = address($fuse.neutral)
			.staticcall(
				abi.encodeWithSelector(
					IFuse($fuse.neutral).enforceFuse.selector,
					$fuse.live,
					$plug.current,
					pinHash
				)
			);

		$result = PlugSimulationLib.Result({
			success: $success,
			callback: $callback
		});
	}

	/**
	 * @notice Overloaded implementation of indexes to push automatic resolution for
	 *         the length of the bundle rather than building it "offchain" before calling.
	 * @param $plugs The raw plugs to simulate.
	 * @return $indexes The indexes of the plugs to simulate.
	 *         uint8 | 00 | livePlugsIndex
	 *         uint8 | 08 | plugsIndex
	 *         uint8 | 16 | pinIndex
	 *         uint8 | 24 | fuseIndex
	 */
	function indexes(
		PlugTypesLib.Plug[] calldata $plugs
	) public view virtual returns (uint24[] memory) {
		return indexes(uint8(length($plugs)), $plugs);
	}

	/**
	 * @notice Get the indexes of the plugs to simulate.
	 * @param $fusesLength The length of the total number of fuses in this Plug bundle.
	 * @param $plugs The raw plugs to simulate.
	 * @return $indexes The indexes of the plugs to simulate.
	 *         uint8 | 00 | livePlugsIndex
	 *         uint8 | 08 | plugsIndex
	 *         uint8 | 16 | pinIndex
	 *         uint8 | 24 | fuseIndex
	 */
	function indexes(
		uint8 $fusesLength,
		PlugTypesLib.Plug[] calldata $plugs
	) public view virtual returns (uint24[] memory) {
		if ($fusesLength == 0) return new uint24[](0);

		uint24[] memory fuseIndexes = new uint24[]($fusesLength);

		for (uint8 i; i < $plugs.length; i++) {
			PlugTypesLib.Plug memory plug = $plugs[i];

			for (uint8 j; j < plug.pins.length; j++) {
				PlugTypesLib.Pin memory pin = plug.pins[j].pin;

				for (uint8 k; k < pin.fuses.length; k++) {
					fuseIndexes[fuseIndexes.length - $fusesLength--] = uint24(
						(k << 16) | (j << 8) | (i << 0)
					);
				}
			}
		}

		return fuseIndexes;
	}

	/**
	 * @notice Overloaded implementation of indexes to push automatic resolution for
	 *         the length of the bundle rather than building it "offchain" before calling.
	 * @param $livePlugs The live plugs to simulate.
	 * @return $indexesArray The indexes of the plugs to simulate.
	 *         uint8 | 00 | livePlugsIndex
	 *         uint8 | 08 | plugsIndex
	 *         uint8 | 16 | pinIndex
	 *         uint8 | 24 | fuseIndex
	 */
	function indexes(
		PlugTypesLib.LivePlugs[] calldata $livePlugs
	) public view virtual returns (uint32[] memory) {
		return indexes(length($livePlugs), $livePlugs);
	}

	/**
	 * @notice Get the indexes of the plugs to simulate.
	 * @param $livePlugs The live plugs to simulate.
	 * @return $indexesArray The indexes of the plugs to simulate.
	 *         uint8 | 00 | livePlugsIndex
	 *         uint8 | 08 | plugsIndex
	 *         uint8 | 16 | pinIndex
	 *         uint8 | 24 | fuseIndex
	 */
	function indexes(
		uint256 $fusesLength,
		PlugTypesLib.LivePlugs[] calldata $livePlugs
	) public view virtual returns (uint32[] memory) {
		uint256 indexesLength = length($livePlugs);

		if (indexesLength == 0) return new uint32[](0);

		uint32[] memory $indexesArray = new uint32[](indexesLength);

		for (uint256 i; i < indexesLength; i++) {
			uint24[] memory fuseIndexes = indexes(
				uint8($fusesLength),
				$livePlugs[i].plugs.plugs
			);

			/// @dev Loop through every index and place the livePlugs
			///		 index as the first 8 bits.
			for (uint8 j; j < fuseIndexes.length; j++) {
				/// @dev Append the livePlugs index to the fuse index.
				$indexesArray[i] = uint32(fuseIndexes[j] << 8) | j;
			}
		}

		return $indexesArray;
	}

	/**
	 * @notice Determine the length of responses this bundle of Plugs requires.
	 * @param $plugs The bundle of Plugs being simulated.
	 * @return $total The number of responses to expect for this bundle.
	 **/
	function length(
		PlugTypesLib.Plug[] calldata $plugs
	) public view virtual returns (uint256 $total) {
		for (uint8 i; i < $plugs.length; i++) {
			PlugTypesLib.Plug memory plug = $plugs[i];

			for (uint8 j; j < plug.pins.length; j++) {
				$total += plug.pins[j].pin.fuses.length;
			}
		}
	}

	/**
	 * @notice Determine the length of responses this bundle of Plugs requires.
	 * @dev This function is overloaded with the primary types used to interact.
	 * @param $livePlugs The bundle of LivePlugs being simulated.
	 * @return $total The number of responses to expect for this bundle.
	 **/
	function length(
		PlugTypesLib.LivePlugs[] calldata $livePlugs
	) public view virtual returns (uint256 $total) {
		for (uint8 i; i < $livePlugs.length; i++) {
			$total += length($livePlugs[i].plugs.plugs);
		}
	}
}
