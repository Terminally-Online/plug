// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {LibBytes} from 'solady/utils/LibBytes.sol';

import {PlugLib, PlugTypesLib} from './Plug.Lib.sol';

library PlugCoilLib {
	uint256 private constant WORD = 32;

	function to(bytes memory $coil) internal returns (address $to) {}

	function data(bytes memory $coil) internal returns (bytes memory $data) {}

	function value(bytes memory $coil) internal returns (bytes memory $value) {}

	function _static(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) internal pure returns (bytes memory $charge, uint256 $area) {
		if ($update.slice.start + $update.slice.length > $coil.length) {
			revert PlugLib.PlugFailed($i, 'PlugCore:out-of-bounds');
		}
		$charge = LibBytes.slice(
			$coil,
			$update.slice.start,
			$update.slice.start + $update.slice.length
		);
		if ($update.start + $update.slice.length > $data.length) {
			revert PlugLib.PlugFailed($i, 'PlugCore:would-overflow');
		}
		$area = $update.start + $update.slice.length;
	}

	function _dynamic(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) internal pure returns (bytes memory $charge, uint256 $area) {
		uint256 start = $update.start;
		uint256 dataOffset;
		assembly {
			dataOffset := mload(add($coil, add(start, WORD)))
		}
		if (dataOffset >= $coil.length) {
			revert PlugLib.PlugFailed($i, 'PlugCore:invalid-offset');
		}
		uint256 dataLength;
		assembly {
			dataLength := mload(add($coil, add(dataOffset, WORD)))
		}
		if (dataOffset + WORD + dataLength > $coil.length) {
			revert PlugLib.PlugFailed($i, 'PlugCore:invalid-length');
		}

		if ($update.slice.typeId == 1 || $update.slice.typeId == 4) {
			uint256 arrayLength;
			assembly {
				arrayLength := mload(add($coil, add(dataOffset, WORD)))
			}
			if (arrayLength * 32 > dataLength) {
				revert PlugLib.PlugFailed($i, 'PlugCore:array-length-invalid');
			}
		} else if ($update.slice.typeId == 3) {
			if (dataLength < 32) {
				revert PlugLib.PlugFailed($i, 'PlugCore:struct-too-small');
			}
		} else if ($update.slice.typeId == 5 && dataLength < 64) {
			revert PlugLib.PlugFailed($i, 'PlugCore:key-value-too-small');
		}
		$charge = LibBytes.slice(
			$coil,
			dataOffset + WORD,
			dataOffset + WORD + dataLength
		);

		if (start + WORD + dataLength > $data.length) {
			revert PlugLib.PlugFailed($i, 'PlugCore:would-overflow');
		}
		$area = start + WORD + dataLength;
	}

	/**
	 * @notice Coil the data for the given update.
	 * @param $coil The coil to coil the data for.
	 * @param $i The index of the Plug.
	 * @param $update The update to coil the data for.
	 * @param $data The data to coil the data for.
	 * @return $charge The charged data.
	 */
	function insert(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) internal pure returns (bytes memory $charge) {
		uint256 area;

		if ($update.slice.typeId == 0)
			($charge, area) = _static($coil, $i, $update, $data);
		else ($charge, area) = _dynamic($coil, $i, $update, $data);

		$charge = LibBytes.concat(
			LibBytes.concat(LibBytes.slice($data, 0, $update.start), $charge),
			LibBytes.slice($data, area, $data.length)
		);
	}
}
