// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {LibBytes} from 'solady/utils/LibBytes.sol';

import {PlugLib, PlugTypesLib} from './Plug.Lib.sol';

/**
 * @title Plug Coil Library
 * @notice A library for handling the encoding, decoding, and transformation of coil data
 *         in the Plug framework. The coil is a compact binary format that encodes transaction
 *         data (selector, to, value, data) in a gas-efficient way. It supports both static
 *         (fixed-size) and dynamic (variable-size) data types through a typeId system.
 * @dev The coil is encoded as (selector, to, value, data) where data is a dynamic bytes field.
 *      The encoding follows ABI encoding rules with dynamic fields at the end.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
library PlugCoilLib {
	uint256 private constant WORD = 32;

	/**
	 * TypeId 0: Static Types
	 * - Handles all fixed-size types: uint256, int256, address, bool, bytes32
	 * - Data is stored directly at a fixed offset in the coil
	 * - No length prefix or dynamic sizing needed
	 * - Gas efficient for simple value types
	 * - Examples: uint256(1), address(0x...), bool(true)
	 */
	uint8 private constant TYPE_STATIC = 0x00;

	/**
	 * TypeId 1: Simple Arrays
	 * - Handles arrays of fixed-size types: uint256[], address[], bool[], bytes32[]
	 * - Data format: [length, element1, element2, ...]
	 * - Each element is exactly 32 bytes (padded)
	 * - Gas efficient for common array types
	 * - Examples: uint256[](1,2,3), address[](0x...,0x...)
	 * - Validation: arrayLength * 32 must fit within dataLength
	 */
	uint8 private constant TYPE_ARRAY = 0x01;

	/**
	 * TypeId 2: Complex Arrays
	 * - Handles arrays of variable-size types: struct[], string[], bytes[]
	 * - Data format: [length, element1Size, element1, element2Size, element2, ...]
	 * - Each element can be variable size with its own size prefix
	 * - Flexible for complex data structures
	 * - Examples: struct[](s1,s2), string[]("a","b"), bytes[](0x...,0x...)
	 * - Validation:
	 *   1. Array length must be non-zero
	 *   2. Each element size must be valid
	 *   3. Total size must fit within dataLength
	 */
	uint8 private constant TYPE_COMPLEX_ARRAY = 2;

	/**
	 * TypeId 3: Nested Arrays
	 * - Handles multi-dimensional arrays: uint256[][], address[][]
	 * - Data format: [length, element1, element2, ...]
	 * - Each element is itself an array (32 bytes)
	 * - Used for complex data structures
	 * - Examples: uint256[][]([1,2], [3,4]), address[][]([0x...], [0x...])
	 * - Validation: arrayLength * 32 must fit within dataLength
	 */
	uint8 private constant TYPE_NESTED_ARRAY = 3;

	/**
	 * TypeId 4: Key-Value Pairs
	 * - Handles mapping-like structures: (key, value) pairs
	 * - Data format: [key, value]
	 * - Both key and value are 32 bytes
	 * - Used for simple key-value storage
	 * - Examples: (address => uint256), (bytes32 => address)
	 * - Validation: minimum 64 bytes (two words)
	 */
	uint8 private constant TYPE_KEY_VALUE = 4;

	/**
	 * @notice Decodes a coil bytes into its constituent parts.
	 * @dev The coil is encoded as (selector, to, value, data) where data is a dynamic bytes field.
	 *      The encoding follows ABI encoding rules with dynamic fields at the end.
	 * @param $coil The encoded coil bytes to decode
	 * @return $selector The 8-bit selector identifying the operation type
	 * @return $to The target address for the operation
	 * @return $value The amount of native currency to send with the operation
	 * @return $data The calldata to execute with the operation
	 */
	function decode(
		bytes memory $coil
	)
		internal
		pure
		returns (
			uint8 $selector,
			address $to,
			uint256 $value,
			bytes memory $data
		)
	{
		($selector, $to, $value, $data) = abi.decode(
			$coil,
			(uint8, address, uint256, bytes)
		);
	}

	/**
	 * @notice Extracts the selector from a coil bytes.
	 * @dev The selector is stored at offset 32 in the coil bytes.
	 *      This is a gas-efficient way to access just the selector.
	 * @param $coil The encoded coil bytes
	 * @return $selector The 8-bit selector identifying the operation type
	 */
	function selector(
		bytes memory $coil
	) internal pure returns (uint8 $selector) {
		assembly {
			$selector := mload(add($coil, 32))
		}
	}

	/**
	 * @notice Extracts the target address from a coil bytes.
	 * @dev The target address is stored at offset 64 in the coil bytes.
	 *      This is a gas-efficient way to access just the target address.
	 * @param $coil The encoded coil bytes
	 * @return $to The target address for the operation
	 */
	function to(bytes memory $coil) internal pure returns (address $to) {
		assembly {
			$to := mload(add($coil, 64))
		}
	}

	/**
	 * @notice Extracts the value from a coil bytes.
	 * @dev The value is stored at offset 96 in the coil bytes.
	 *      This is a gas-efficient way to access just the value.
	 * @param $coil The encoded coil bytes
	 * @return $value The amount of native currency to send with the operation
	 */
	function value(bytes memory $coil) internal pure returns (uint256 $value) {
		assembly {
			$value := mload(add($coil, 96))
		}
	}

	/**
	 * @notice Extracts the calldata from a coil bytes.
	 * @dev The calldata is stored as a dynamic bytes field at the end of the coil.
	 *      First reads the offset at position 128, then reads the length and data.
	 * @param $coil The encoded coil bytes
	 * @return $data The calldata to execute with the operation
	 */
	function data(
		bytes memory $coil
	) internal pure returns (bytes memory $data) {
		assembly {
			let offset := mload(add($coil, 128))
			let length := mload(add($coil, add(offset, 32)))
			$data := mload(add($coil, add(offset, 64)))
		}
	}

	/**
	 * @notice Transforms data according to the update specification.
	 * @dev This function handles both static and dynamic data transformations.
	 *      For static data, it directly slices the specified portion.
	 *      For dynamic data, it handles arrays, structs, and key-value pairs.
	 * @param $coil The coil containing the data to transform
	 * @param $i The index of the plug being processed
	 * @param $update The update specification containing slice and type information
	 * @param $data The data to be transformed
	 * @return $charge The transformed data
	 */
	function transform(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) internal pure returns (bytes memory $charge) {
		uint256 area;

		if ($update.slice.typeId == TYPE_STATIC)
			($charge, area) = _static($coil, $i, $update, $data);
		else ($charge, area) = _dynamic($coil, $i, $update, $data);

		$charge = LibBytes.concat(
			LibBytes.concat(LibBytes.slice($data, 0, $update.start), $charge),
			LibBytes.slice($data, area, $data.length)
		);
	}

	/**
	 * @notice Handles static data transformation.
	 * @dev For static data, directly slices the specified portion from the coil.
	 *      Validates that the slice is within bounds and won't cause overflow.
	 * @param $coil The coil containing the data to transform
	 * @param $i The index of the plug being processed
	 * @param $update The update specification containing slice information
	 * @param $data The data to be transformed
	 * @return $charge The transformed static data
	 * @return $area The end position of the transformed data
	 */
	function _static(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) private pure returns (bytes memory $charge, uint256 $area) {
		if ($update.slice.start + $update.slice.length > $coil.length) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreOutOfBounds);
		}
		$charge = LibBytes.slice(
			$coil,
			$update.slice.start,
			$update.slice.start + $update.slice.length
		);
		if ($update.start + $update.slice.length > $data.length) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreWouldOverflow);
		}
		$area = $update.start + $update.slice.length;
	}

	/**
	 * @notice Handles dynamic data transformation.
	 * @dev For dynamic data, handles arrays, structs, and key-value pairs.
	 *      Validates data offsets, lengths, and type-specific requirements.
	 * @param $coil The coil containing the data to transform
	 * @param $i The index of the plug being processed
	 * @param $update The update specification containing slice and type information
	 * @param $data The data to be transformed
	 * @return $charge The transformed dynamic data
	 * @return $area The end position of the transformed data
	 */
	function _dynamic(
		bytes memory $coil,
		uint256 $i,
		PlugTypesLib.Update calldata $update,
		bytes memory $data
	) private pure returns (bytes memory $charge, uint256 $area) {
		uint256 start = $update.start;
		uint256 dataOffset;
		assembly {
			dataOffset := mload(add($coil, add(start, WORD)))
		}
		if (dataOffset >= $coil.length) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreInvalidOffset);
		}
		uint256 dataLength;
		assembly {
			dataLength := mload(add($coil, add(dataOffset, WORD)))
		}
		if (dataLength > type(uint128).max) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreWouldOverflow);
		}
		if (dataOffset + WORD + dataLength > $coil.length) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreInvalidLength);
		}

		if ($update.slice.typeId == TYPE_ARRAY) {
			uint256 arrayLength;
			assembly {
				arrayLength := mload(add($coil, add(dataOffset, WORD)))
			}
			if (arrayLength * WORD > dataLength) {
				revert PlugLib.PlugFailed(
					$i,
					PlugLib.PlugCoreArrayLengthInvalid
				);
			}
		} else if ($update.slice.typeId == TYPE_COMPLEX_ARRAY) {
			uint256 arrayLength;
			assembly {
				arrayLength := mload(add($coil, add(dataOffset, WORD)))
			}
			if (arrayLength == 0) {
				revert PlugLib.PlugFailed(
					$i,
					PlugLib.PlugCoreArrayLengthInvalid
				);
			}
			uint256 totalSize;
			uint256 currentOffset = dataOffset + WORD;
			for (uint256 ii; ii < arrayLength; ii++) {
				uint256 elementSize;
				assembly {
					elementSize := mload(add($coil, add(currentOffset, WORD)))
				}
				if (elementSize == 0 || elementSize > type(uint128).max) {
					revert PlugLib.PlugFailed(
						$i,
						PlugLib.PlugCoreInvalidLength
					);
				}
				if (elementSize > dataLength) {
					revert PlugLib.PlugFailed(
						$i,
						PlugLib.PlugCoreInvalidLength
					);
				}
				totalSize += elementSize + WORD;
				currentOffset += elementSize + WORD;
			}
			if (totalSize > dataLength) {
				revert PlugLib.PlugFailed($i, PlugLib.PlugCoreWouldOverflow);
			}
		} else if ($update.slice.typeId == TYPE_NESTED_ARRAY) {
			uint256 arrayLength;
			assembly {
				arrayLength := mload(add($coil, add(dataOffset, WORD)))
			}
			if (arrayLength * WORD > dataLength) {
				revert PlugLib.PlugFailed(
					$i,
					PlugLib.PlugCoreArrayLengthInvalid
				);
			}
		} else if ($update.slice.typeId == TYPE_KEY_VALUE && dataLength < 64) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreKeyValueTooSmall);
		}

		$charge = LibBytes.slice(
			$coil,
			dataOffset + WORD,
			dataOffset + WORD + dataLength
		);

		if (start + WORD + dataLength > $data.length) {
			revert PlugLib.PlugFailed($i, PlugLib.PlugCoreWouldOverflow);
		}
		$area = start + WORD + dataLength;
	}
}
