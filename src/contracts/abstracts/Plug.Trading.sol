// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTradingInterface } from
    "../interfaces/Plug.Trading.Interface.sol";
import { ModuleAuthUpgradable } from
    "sequence/modules/commons/ModuleAuthUpgradable.sol";

import { PlugLib } from "../libraries/Plug.Lib.sol";

import { ERC721 } from "solady/src/tokens/ERC721.sol";

/**
 * @title Plug Trading
 * @notice Enables the ability to represent Socket ownership through the current
 *         state of an ERC721 that is managed inside the factory that deployed
 *         the Vault. This way, Vaults can be traded on any major marketplace
 *         enabling the ability to spread workflows and earnings of
 *         aforementioned workflows such as points and yield.
 * @author nftchance (chance@onplug.io)
 */
abstract contract PlugTrading is
    PlugTradingInterface,
    ModuleAuthUpgradable
{
    /// @dev The address that houses the ownership information.
    address public ownership;

    /// @dev The address that has permission to route generalized bundles.
    address public router;

    /**
     * @notice Modifier enforcing the caller to be the ownership proxy.
     */
    modifier onlyOwnership() {
        /// @dev Ensure the `caller` is the ownership proxy.
        if (msg.sender != ownership) {
            revert PlugLib.CallerInvalid(ownership, msg.sender);
        }
        _;
    }

    /**
     * @notice Modifier enforcing that  only the owner of the token can call
     *         functions that have this modifier applied.
     */
    modifier onlyOwner() {
        /// @dev Ensure the `caller` is the owner of the Socket.
        if (msg.sender != owner()) {
            revert PlugLib.CallerInvalid(owner(), msg.sender);
        }
        _;
    }

    /**
     * @notice Transfer the ownership of a Socket to a new address when the
     *         NFT is transferred. With this process, due to the signature logic
     *         that is used to verify a signature we must also update the image
     *         hash that contains the encoded definition of the criteria.
     * @param $newOwner The address of the new owner.
     */
    function transferOwnership(address $newOwner)
        public
        virtual
        onlyOwnership
    {
        /// @dev Calculate the image hash based on the new owner. For now, the
        ///      assumption is the definition of a single Socket owner.
        bytes32 expectedImageHash = keccak256(
            abi.encodePacked(
                keccak256(
                    abi.encodePacked(
                        abi.decode(
                            abi.encodePacked(uint96(1), $newOwner),
                            (bytes32)
                        ),
                        uint256(1)
                    )
                ),
                uint256(1)
            )
        );

        /// @dev Update the image hash that is used to verify signatures
        ///      within the execution of a Socket interaction.
        _updateImageHash(expectedImageHash);

        /// @dev Emit the event to signify transfer change as well as a change in
        ///      the image hash so that it can be utilized elsewhere.
        emit PlugLib.SocketOwnershipTransferred(
            owner(), $newOwner, expectedImageHash
        );
    }

    /**
     * @notice Get the owner of the Vault.
     */
    function owner() public view virtual returns (address $owner) {
        $owner = ERC721(ownership).ownerOf(tokenId());
    }

    /**
     * @notice Get the token ID of the Vault.
     */
    function tokenId()
        public
        view
        virtual
        returns (uint256 $tokenId)
    {
        $tokenId = uint256(uint160(address(this)));
    }

    /**
     * @notice Set the address of the ownership proxy which is a ERC721
     *         compliant contract that lives inside of the factory.
     */
    function _initializeOwnership(
        address $ownership,
        address $router
    )
        internal
    {
        /// @dev Check if the inheriting contract requires single-use
        ///      ownership initialization.
        if (_guardInitializeOwnership()) {
            if (ownership != address(0)) {
                revert PlugLib.TradingAlreadyInitialized();
            }
        }

        /// @dev Set the state of the ownership proxy.
        ownership = $ownership;

        /// @dev Set the state of the higher level router.
        router = $router;
    }

    /**
     * @notice Guard the initial ownership of the Vault to ensure that the
     *         ownership is set to the correct address and cannot be changed
     *         one it has been defined. This makes the state immutable in
     *         practice even though the variable itself is not semantically
     *         immutable outside the scope of logic enforcement.
     * @dev By default this function has an empty implementation so that not
     *      all inheriting contracts are required to define this function. If
     *      you would like to enforce single-use ownership initialization, you
     *      can override this function and return `true` to enforce the guard.
     */
    function _guardInitializeOwnership()
        internal
        pure
        virtual
        returns (bool $guard)
    { }
}
