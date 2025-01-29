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

/**
 * @title Plug
 * @notice The core contract for the Plug framework extremely execution paths.
 * @author @nftchance (chance@onplug.io)
 */
contract PlugSocket is
    PlugSocketInterface,
    PlugTypes,
    Ownable,
    Receiver,
    UUPSUpgradeable,
    ReentrancyGuard
{
    using LibBitmap for LibBitmap.Bitmap;
    using ECDSA for bytes32;

    mapping(address oneClicker => bool allowed) public oneClickersToAllowed;

    LibBitmap.Bitmap private nonces;

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
     * @notice Execute a bundle of Plugs.
     * @param $plugs The Plugs to execute containing the bundle and side effects.
     * @param $solver Encoded data defining the Solver and compensation.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $solver
    )
        internal
        returns (PlugTypesLib.Result memory $results)
    {
        if ($plugs.solver.length != 0) {
            (uint48 expiration, address solver) = abi.decode($plugs.solver, (uint48, address));
            if (expiration < block.timestamp) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:solver-expired");
            }
            if (solver != $solver) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugCore:solver-invalid");
            }
        }

        PlugTypesLib.Plug calldata action;
        uint256 length = $plugs.plugs.length;
        for (uint8 i; i < length; i++) {
            action = $plugs.plugs[i];

            (bool success,) = action.to.call{ value: action.value }(action.data[1:]);
            if (!success) {
                revert PlugLib.PlugFailed(i, "PlugCore:plug-failed");
            }
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
     */
    function _enforceSender(address $sender) internal view virtual returns (bool $allowed) {
        $allowed = $sender == owner() || $sender == address(this);
    }

    /**
     * See { UUPSUpgradeable._authorizeUpgrade }
     */
    function _authorizeUpgrade(address) internal virtual override onlyOwner { }

    /**
     * See { PlugTrading._guardInitializeOwnership }
     */
    function _guardInitializeOwnership() internal pure virtual returns (bool $guard) {
        $guard = true;
    }
}
