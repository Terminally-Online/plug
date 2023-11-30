// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.19;

import {IFuse} from '../interfaces/IFuse.sol';
import {PlugCore} from './PlugCore.sol';

import {PlugSimulationHelpers} from '../libraries/SimulationHelpers.sol';

abstract contract Simulation is PlugCore {
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
	 * @notice Simulate the execution of a bundle of live plugs.
	 * @param $livePlugs The live plugs to simulate.
	 * @param $indexes The indexes of the plugs to simulate.
	 *        uint8 | 00 | livePlugsIndex
	 *        uint8 | 08 | plugsIndex
	 *        uint8 | 16 | pinIndex
	 *        uint8 | 24 | fuseIndex
	 * @return $results The results of the simulation.
	 */
	function simulate(
		LivePlugs[] calldata $livePlugs,
		uint32[] calldata $indexes
	) external view returns (PlugSimulationHelpers.Result[] memory) {
		PlugSimulationHelpers.Result[]
			memory $results = new PlugSimulationHelpers.Result[](
				$indexes.length
			);

		for (uint256 i; i < $indexes.length; ) {
			uint32 index = $indexes[i];

			Plug memory plug = $livePlugs[uint8(index)].plugs.plugs[
				uint8(index >> 8)
			];
			LivePin memory livePin = plug.pins[uint8(index >> 16)];
			Fuse memory fuse = livePin.pin.fuses[uint8(index >> 24)];

			bytes32 pinHash = getLivePinHash(livePin);

			(bool $success, bytes memory $callback) = address(fuse.neutral)
				.staticcall(
					abi.encodeWithSelector(
						IFuse(fuse.neutral).enforceFuse.selector,
						fuse.live,
						plug.current,
						pinHash
					)
				);

			$results[index] = PlugSimulationHelpers.Result({
				success: $success,
				callback: $callback
			});

			unchecked {
				++i;
			}
		}

		return $results;
	}

	/**
	 * @notice Get the indexes of the plugs to simulate.
	 * @param $livePlugs The live plugs to simulate.
	 * @return $indexes The indexes of the plugs to simulate.
	 *         uint8 | 00 | plugsIndex
	 *         uint8 | 08 | pinIndex
	 *         uint8 | 16 | fuseIndex
	 */
	function indexes(
		LivePlugs calldata $livePlugs
	) public payable virtual returns (uint24[] memory $indexes) {
		for (uint256 i; i < $livePlugs.plugs.plugs.length; i++) {
			Plug memory plug = $livePlugs.plugs.plugs[i];

			for (uint256 j; j < plug.pins.length; j++) {
				Pin memory pin = plug.pins[j].pin;

				for (uint256 k; k < pin.fuses.length; k++) {
					$indexes[$indexes.length] = uint24(
						(uint8(k) << 16) | (uint8(j) << 8) | uint8(i)
					);
				}
			}
		}
	}

	/**
	 * @notice Get the indexes of the plugs to simulate.
	 * @param $livePlugs The live plugs to simulate.
	 * @param $indexesLength The length of the total number of fuses in this bundle.
	 * @return $indexesArray The indexes of the plugs to simulate.
	 *         uint8 | 00 | livePlugsIndex
	 *         uint8 | 08 | plugsIndex
	 *         uint8 | 16 | pinIndex
	 *         uint8 | 24 | fuseIndex
	 */
	function indexes(
		LivePlugs[] calldata $livePlugs,
		uint256 $indexesLength
	) public payable virtual returns (uint32[] memory) {
		uint32[] memory $indexesArray = new uint32[]($indexesLength);

		for (uint256 i; i < $indexesLength; ) {
			uint24[] memory fuseIndexes = indexes($livePlugs[i]);

			/// @dev Loop through every index and place the livePlugs
			///		 index as the first 8 bits.
			for (uint8 j; j < fuseIndexes.length; ) {
				uint24 fuseIndex = fuseIndexes[j];

				/// @dev Append the livePlugs index to the fuse index.
				$indexesArray[i] = uint32(fuseIndex << 8) | j;

				unchecked {
					++i;
					++j;
				}
			}
		}

		return $indexesArray;
	}
}
