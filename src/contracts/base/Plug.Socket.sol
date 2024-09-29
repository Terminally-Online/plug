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
    /// @dev Import the LibBitmap library from Solady
    using LibBitmap for LibBitmap.Bitmap;

    /// @notice Use the ECDSA library for signature verification.
    using ECDSA for bytes32;

    /// @dev Mapping of one-clickers to their allowed status.
    mapping(address oneClicker => bool allowed) public oneClickersToAllowed;

    /// @dev Bitmap to track used nonces for each signer
    LibBitmap.Bitmap private nonces;

    /*
    * @notice The constructor for the Plug Socket will
    *         initialize to address(1) when not deployed through
    *         a Socket factory.
    */
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
            revert PlugLib.SignatureInvalid();
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
            revert PlugLib.SenderInvalid(msg.sender);
        }
        _;
    }

    /**
     * See {PlugSocketInterface-initialize}.
     */
    function initialize(address $owner, address $oneClicker) public {
        _initializeOwner($owner);

        /// @dev Automatically permission the primary platform one-clicker.
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
        enforceSignature($livePlugs)
        nonReentrant
        returns (PlugTypesLib.Result[] memory $results)
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
        enforceSender
        nonReentrant
        returns (PlugTypesLib.Result[] memory $results)
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
     * @return $results The return data of the plugs.
     */
    function _plug(
        PlugTypesLib.Plugs calldata $plugs,
        address $solver
    )
        internal
        returns (PlugTypesLib.Result[] memory $results)
    {
        /// @dev Hash the body of the object to ensure the integrity of
        ///      the (bundle of) Plugs that are being executed.
        bytes32 plugsHash = getPlugsHash($plugs);

        /// @dev Load the Plug stack into memory for cheaper access.
        uint256 length = $plugs.plugs.length;
        $results = new PlugTypesLib.Result[](length);

        /// @dev Save the object into memory to avoid multiple creations
        ///      of the same object.
        PlugTypesLib.Plug calldata action;

        /// @dev Iterate over the Plugs that are held within this bundle
        ///      an execute each of them. Each respectively may be a
        ///      condition being enforced or an outcome focused transaction.
        for (uint256 i; i < length; i++) {
            /// @dev Place the active Plug in the shorter reference stack.
            action = $plugs.plugs[i];

            /// @dev If the call has an associated value, ensure the contract
            ///      has enough balance to cover the cost of the call.
            if (address(this).balance < action.value) {
                revert PlugLib.ValueInvalid(action.target, action.value, address(this).balance);
            }

            ($results[i].success, $results[i].result) =
                action.target.call{ value: action.value }(action.data[1:]);

            /// @dev If the call failed, bubble up the revert reason if needed.
            PlugLib.bubbleRevert($results[i].success, $results[i].result);
        }

        /// @dev Pay the Solver for the gas used if it was not open-access.
        if ($plugs.solver.length != 0) {
            /// @dev Unpack the solver data from the encoded Solver data.
            (uint48 expiration, address solver) = abi.decode($plugs.solver, (uint48, address));

            /// @dev Confirm the Solver is allowed to execute the transaction.
            ///      This is done here instead of a modifier so that the gas
            ///      snapshot accounts for the additional gas cost of the require.
            if (solver != $solver) {
                revert PlugLib.SolverInvalid(solver, $solver);
            }

            /// @dev Confirm the order provided to the Solver has not expired.
            if (expiration < block.timestamp) {
                revert PlugLib.SolverExpired();
            }
        }

        emit PlugLib.PlugsExecuted(plugsHash, $results);
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
        /// @dev Recover the signer from the signature that was provided.
        address signer = getPlugsDigest($input.plugs).recover($input.signature);

        /// @dev Extract nonce from the salt field (assuming it's the first 12 bytes -- 96 bits).
        uint256 nonce = uint256(uint96(bytes12($input.plugs.salt)));

        /// @dev Confirm the nonce has not been used before.
        if (nonces.get(nonce) == true) {
            revert PlugLib.NonceInvalid();
        }

        /// @dev Use the nonce.
        nonces.set(nonce);

        /// @dev Validate that the signer is allowed within context.
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
