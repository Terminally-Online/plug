// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";

/**
 * @title Plug Vault Socket
 * @author @nftchance (chance@utc24.io)
 */
contract PlugVaultSocket is PlugSocket, Receiver, Ownable {
    /// @dev Bit related shifts and masks for access management.
    uint8 internal constant DEFAULT_ACCESS = 0x10;
    uint8 internal constant ACCESS = 0x1;
    uint8 internal constant ACCESS_DENIED = 0x0;
    uint8 internal constant SIGNER_SHIFT = 0x4;

    /// @notice Mapping that holds the access definition of addresses.
    /// @dev This stores two "boolean" values in a single uint9 slot
    ///      to save on storage costs. The first 4 bits are for the
    ///      router, and the last 4 bits are for the signer.
    mapping(uint160 => uint8) public access;

    /*
     * @notice The constructor for the Plug Vault Socket will
     *         initialize to address(1) when not deployed through
     *         a Socket factory.
     */
    constructor() {
        initialize(address(1));
    }

    /**
     * @notice Initializes a new Plug Vault Socket contract.
     */
    function initialize(address $owner) public {
        /// @dev Initialize the owner.
        _initializeOwner($owner);
        /// @dev Initialize the Plug Socket.
        _initializePlug();
    }

    /**
     * See { Ownable-transferOwnership }
     *
     * @dev The typical transfer ownership is overridden to automatically manage
     *      the access state of the previous and new owner without requiring
     *      the new owner to manually add themselves as an allowed direct signer.
     *      Ownership of a Vault implicitly grants owner() operational access to
     *      add themselves anyways, so this is just a convenience feature as well
     *      as prevents any malicious previous owners from signing on a vault that
     *      they have lost access to.
     * @dev Note that the function appears to be public howver the super function
     *      already has the onlyOwner modifier applied to it, so it is effectively
     *      only callable by the owner.
     */
    function transferOwnership(address $owner)
        public
        payable
        virtual
        override
    {
        /// @dev Convert the existing owner to a uint160 to avoid double storage reads.
        uint160 owner = uint160(owner());

        /// @dev Remove the previous owner from the default-set list of signers
        ///      while maintaing their potential state as a router.
        access[owner] = access[owner] & ACCESS;

        /// @dev Add the new owner as a signer to enable seamless
        ///      direct fulfillment of a Plug bundle without the need
        ///      for a signature or adding themselves manually
        ///      while maintaining their potential state as a router.
        access[uint160($owner)] = DEFAULT_ACCESS | (access[owner] & ACCESS);

        /// @dev Proceed with the normal ownership transfer.
        super.transferOwnership($owner);
    }

    /**
     * @notice Set the access of a router or signer.
     * @param $address The address to set the access for.
     * @param $allowance The bitpacked allowance to set for the address.
     */
    function setAccess(
        address $address,
        uint8 $allowance
    )
        public
        virtual
        onlyOwner
    {
        _setAccess($address, $allowance);
    }

    /**
     * @notice Helper view function that can be used to build the packed
     *         access value for a router or signer.
     * @param $isRouter If the address is a router.
     * @param $isSigner If the address is a signer.
     * @return $access The packed access value.
     */
    function getAccess(
        bool $isRouter,
        bool $isSigner
    )
        public
        pure
        returns (uint8 $access)
    {
        /// @dev Set the router value.
        $access = $isRouter ? ACCESS : ACCESS_DENIED;
        /// @dev Set the signer value preserving the value of the
        ///      router flag previously set.
        $access |= $isSigner ? ACCESS << SIGNER_SHIFT : ACCESS_DENIED;
    }

    /**
     * @notice Helper view function that can be used to unpack the access
     *         value for a router or signer.
     * @param $address The address to get the access for.
     * @return $isRouter If the address is a router.
     * @return $isSigner If the address is a signer.
     */
    function getAccess(address $address)
        public
        view
        returns (bool $isRouter, bool $isSigner)
    {
        /// @dev Retrieve the state from storage.
        uint8 $access = access[uint160($address)];
        /// @dev Unpack the router and signer flags.
        $isRouter = _enforceAccess($access);
        $isSigner = _enforceAccess($access >> SIGNER_SHIFT);
    }

    /**
     * See { PlugSocket-name }
     */
    function name() public pure override returns (string memory $name) {
        $name = "PlugVaultSocket";
    }

    /**
     * See { PlugSocket-version }
     */
    function version() public pure override returns (string memory $version) {
        $version = "0.0.1";
    }

    /**
     * See { Ownable-_initializeOwner }
     */
    function _initializeOwner(address $owner) internal override {
        /// @dev Add the owner as a signer to enable seamless
        ///      direct fulfillment of a Plug bundle without
        ///      the need for a signature.
        _setAccess($owner, DEFAULT_ACCESS);

        /// @dev Proceed with normal owner initialization.
        super._initializeOwner($owner);
    }

    /**
     * @notice Internal management of the access a router or signer has.
     * @dev Note that you cannot toggle off the canonical router. All participants
     *      of the Plug ecosystem are expected to follow pie-growing principles.
     *      Even by explicitly setting the router to false, the router will still
     *      be considered a router. This is to prevent any malicious actors from
     *      attempting to disable the router and disrupt the ecosystem.
     * @dev If you would like to append additional logic to this function, you can
     *      override it in a derived contract and call super._setAccess() to ensure
     *      that the access state is properly managed.
     * @param $address The address to manage access for.
     * @param $allowance The bitpacked allowance to set for the address.
     */
    function _setAccess(address $address, uint8 $allowance) internal virtual {
        /// @dev Set the packed access state for the address.
        access[uint160($address)] = $allowance;
    }

    /**
     * See { PlugEnforce._enforceRouter }
     */
    function _enforceRouter(address $router)
        internal
        view
        override
        returns (bool $allowed)
    {
        /// @dev Confirm the router is allowed by recovering the packed access
        ///      state as well as checking if the router is the canonical router.
        $allowed = _enforceAccess(access[uint160($router)])
            || super._enforceRouter($router);
    }

    /**
     * See { PlugEnforce._enforceSigner }
     */
    function _enforceSigner(address $signer)
        internal
        view
        override
        returns (bool $allowed)
    {
        /// @dev Confirm the signer is allowed by recovering the packed access state.
        $allowed = _enforceAccess(access[uint160($signer)] >> SIGNER_SHIFT);
    }

    /**
     * @notice Enforce the access of a router or signer.
     * @param $state The state of the access.
     * @return $allowed If the access is allowed.
     */
    function _enforceAccess(uint8 $state)
        internal
        pure
        returns (bool $allowed)
    {
        /// @dev Confirm the masked state is equal to the flag state.
        $allowed = $state & ACCESS == ACCESS;
    }
}
