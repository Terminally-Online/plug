// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";
import { PlugTypes } from "../abstracts/Plug.Types.sol";
import { Ownable } from "solady/auth/Ownable.sol";
import { Receiver } from "solady/accounts/Receiver.sol";
import { UUPSUpgradeable } from "solady/utils/UUPSUpgradeable.sol";
import { ReentrancyGuard } from "solady/utils/ReentrancyGuard.sol";
import { LibBitmap } from "solady/utils/LibBitmap.sol";
import { ECDSA } from "solady/utils/ECDSA.sol";
import { PlugLib, PlugTypesLib } from "../libraries/Plug.Lib.sol";
import { LibBytes } from "solady/utils/LibBytes.sol";

/**
 * @title PlugSocket
 * @notice The "account" contract for a user of Plug.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugSocket is
    PlugSocketInterface,
    PlugTypes,
    Ownable,
    Receiver,
    UUPSUpgradeable,
    ReentrancyGuard
{
    using ECDSA for bytes32;
    using LibBitmap for LibBitmap.Bitmap;

    uint256 private constant WORD = 32;
    uint256 private constant TYPE_CALL = 0x00;
    uint256 private constant TYPE_DELEGATECALL = 0x01;
    uint256 private constant TYPE_CALL_WITH_VALUE = 0x02;
    uint256 private constant TYPE_STATICCALL = 0x03;

    LibBitmap.Bitmap private nonces;

    mapping(address oneClicker => bool allowed) public oneClickersToAllowed;

    constructor() {
        initialize(address(1), address(1));
    }

    /**
     * @notice Modifier to enforce the signer of the transaction.
     * @dev Apply to this to functions that are designed to execute a bundle
     *      of Plugs regardless of whether through a Router or or direct access.
     * @param $input The LivePlugs the definition of execution as well as the
     *               signature used to verify the execution permission.
     */
    modifier enforceSignature(PlugTypesLib.LivePlugs calldata $input) {
        if (_enforceSignature($input) == false) {
            revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:signature-invalid");
        }
        _;
    }

    /**
     * @notice Modifier to enforce the sender of the transaction.
     * @dev Apply to this to functions that are designed to execute a bundle
     *      of Plugs directly from the sender or a recursive call from the contract.
     */
    modifier enforceSender() {
        if (_enforceSender(msg.sender) == false) {
            revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:sender-invalid");
        }
        _;
    }

    /**
     * See {PlugSocketInterface-initialize}.
     */
    function initialize(address $owner, address $oneClicker) public {
        _initializeOwner($owner);

        if ($oneClicker != address(0)) {
            oneClickersToAllowed[$oneClicker] = true;
        }
    }

    /**
     * See {PlugSocketInterface-plug}.
     */
    function plug(
        PlugTypesLib.LivePlugs calldata $livePlugs,
        address $solver
    )
        external
        payable
        virtual
        nonReentrant
        enforceSignature($livePlugs)
        returns (PlugTypesLib.Result memory $results)
    {
        $results = _plug($livePlugs.plugs, $solver);
    }

    /**
     * See {PlugSocketInterface-plug}.
     */
    function plug(PlugTypesLib.Plugs calldata $plugs)
        external
        payable
        virtual
        nonReentrant
        enforceSender
        returns (PlugTypesLib.Result memory $results)
    {
        $results = _plug($plugs, address(0));
    }

    /**
     * @notice Enable specific addresses to build the final route of the Plug.
     * @param $oneClickers The address of the one clicker.
     * @param $allowance The allowance of the one clicker.
     */
    function oneClick(
        address[] calldata $oneClickers,
        bool[] calldata $allowance
    )
        public
        virtual
        onlyOwner
    {
        for (uint256 i; i < $oneClickers.length; i++) {
            oneClickersToAllowed[$oneClickers[i]] = $allowance[i];
        }
    }

    /**
     * See { PlugSocket-name }
     */
    function name() public pure override returns (string memory $name) {
        $name = "Plug Socket";
    }

    /**
     * See { PlugSocket-version }
     */
    function version() public pure override returns (string memory $version) {
        $version = "0.0.1";
    }

    /**
     * @notice Execute a set of Plugs.
     * @param $plugs The Plugs to execute containing the bundle and side effects.
     * @param $solver Encoded data defining the Solver and compensation.
     * @return $results The results of the execution.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $solver
    )
        internal
        returns (PlugTypesLib.Result memory $results)
    {
        if ($plugs.solver.length != 0) {
            if ($plugs.solver.length < 0x40) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:solver-malformed");
            }
            if (uint256(LibBytes.loadCalldata($plugs.solver, 0x00)) < block.timestamp) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:solver-expired");
            }
            if (address(uint160(uint256(LibBytes.loadCalldata($plugs.solver, 0x20)))) != $solver) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:solver-invalid");
            }
        }

        uint256 length = $plugs.plugs.length;
        PlugTypesLib.Plug calldata currentPlug;
        bytes memory data;

        bool success;
        bytes[] memory results = new bytes[](length);

        uint256 updatesLength;
        uint256 ii;

        PlugTypesLib.Update calldata update;
        PlugTypesLib.Slice calldata slice;
        bytes memory coil;
        bytes memory charge;
        uint256 start;
        uint256 dataOffset;
        uint256 dataLength;
        uint256 area;
        uint256 arrayLength;

        for (uint8 i; i < length; i++) {
            currentPlug = $plugs.plugs[i];
            data = currentPlug.data;
            updatesLength = currentPlug.updates.length;

            for (ii = 0; ii < updatesLength; ii++) {
                update = currentPlug.updates[ii];
                slice = update.slice;
                coil = results[slice.index];

                if (slice.typeId == 0) {
                    if (slice.start + slice.length > coil.length) {
                        revert PlugLib.PlugFailed(i, "PlugCore:out-of-bounds");
                    }
                    charge = LibBytes.slice(coil, slice.start, slice.start + slice.length);
                    if (update.start + slice.length > data.length) {
                        revert PlugLib.PlugFailed(i, "PlugCore:would-overflow");
                    }
                    area = update.start + slice.length;
                } else {
                    start = update.start;

                    assembly {
                        dataOffset := mload(add(coil, add(start, WORD)))
                    }
                    if (dataOffset >= coil.length) {
                        revert PlugLib.PlugFailed(i, "PlugCore:invalid-offset");
                    }

                    assembly {
                        dataLength := mload(add(coil, add(dataOffset, WORD)))
                    }
                    if (dataOffset + WORD + dataLength > coil.length) {
                        revert PlugLib.PlugFailed(i, "PlugCore:invalid-length");
                    }

                    if (slice.typeId == 1 || slice.typeId == 4) {
                        assembly {
                            arrayLength := mload(add(coil, add(dataOffset, WORD)))
                        }
                        if (arrayLength * 32 > dataLength) {
                            revert PlugLib.PlugFailed(i, "PlugCore:array-length-invalid");
                        }
                    } else if (slice.typeId == 3) {
                        if (dataLength < 32) {
                            revert PlugLib.PlugFailed(i, "PlugCore:struct-too-small");
                        }
                    } else if (slice.typeId == 5 && dataLength < 64) {
                        revert PlugLib.PlugFailed(i, "PlugCore:key-value-too-small");
                    }

                    charge = LibBytes.slice(coil, dataOffset + WORD, dataOffset + WORD + dataLength);

                    if (start + WORD + dataLength > data.length) {
                        revert PlugLib.PlugFailed(i, "PlugCore:would-overflow");
                    }
                    area = start + WORD + dataLength;
                }

                // More efficient data concatenation using a single memory allocation strategy
                // would be preferable, but we're constrained by LibBytes implementation
                data = LibBytes.concat(
                    LibBytes.concat(LibBytes.slice(data, 0, update.start), charge),
                    LibBytes.slice(data, area, data.length)
                );
            }

            // Execute the appropriate call type based on selector
            if (currentPlug.selector == TYPE_DELEGATECALL) {
                (success, results[i]) = currentPlug.to.delegatecall(data);
            } else if (currentPlug.selector == TYPE_CALL) {
                (success, results[i]) = currentPlug.to.call(data);
            } else if (currentPlug.selector == TYPE_CALL_WITH_VALUE) {
                (success, results[i]) = currentPlug.to.call{ value: currentPlug.value }(data);
            } else if (currentPlug.selector == TYPE_STATICCALL) {
                (success, results[i]) = currentPlug.to.staticcall(data);
            }

            if (!success) revert PlugLib.PlugFailed(i, "PlugCore:plug-failed");
        }

        $results = PlugTypesLib.Result({ index: type(uint8).max, error: "" });
    }

    /**
     * @notice Confirm that signer has permission to declare execution of a
     *         Plug bundle on the parent-socket that inherits this contract.
     * @dev Inheriting contracts must implement the logic of this function to make
     *      sure that only signatures intended for this scope are allowed.
     * @param $input The LivePlugs object that contains the Plugs object as well as
     *               the signature defining the permission to execute the bundle.
     * @return $allowed True if the signature is valid, false otherwise.
     */
    function _enforceSignature(PlugTypesLib.LivePlugs calldata $input)
        internal
        virtual
        returns (bool $allowed)
    {
        address signer = getPlugsDigest($input.plugs).recover($input.signature);
        uint256 nonce = uint256(uint96(bytes12($input.plugs.salt)));
        if (nonces.get(nonce) == true) {
            revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:nonce-invalid");
        }
        nonces.set(nonce);
        $allowed = oneClickersToAllowed[signer] || owner() == signer;
    }

    /**
     * @notice Confirm that the sender of the transaction is allowed as the Socket
     *         was directly interacted with.
     * @dev Inheriting contracts must implement the logic of this function to make
     *      sure that only senders intended for this scope are allowed.
     * @param $sender The sender of the transaction.
     * @return $allowed True if the sender is allowed, false otherwise.
     */
    function _enforceSender(address $sender) internal view virtual returns (bool $allowed) {
        $allowed = $sender == owner() || $sender == address(this);
    }

    /**
     * See { PlugTrading._guardInitializeOwnership }
     */
    function _guardInitializeOwnership() internal pure virtual returns (bool $guard) {
        $guard = true;
    }

    /**
     * See { UUPSUpgradeable._authorizeUpgrade }
     */
    function _authorizeUpgrade(address) internal virtual override onlyOwner { }
}
