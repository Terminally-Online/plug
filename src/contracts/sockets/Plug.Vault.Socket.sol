// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugSocket } from "../abstracts/Plug.Socket.sol";
import { PlugTrading } from "../abstracts/Plug.Trading.sol";
import { Receiver } from "solady/accounts/Receiver.sol";
import { UUPSUpgradeable } from "solady/utils/UUPSUpgradeable.sol";
import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { MerkleProofLib } from "solady/utils/MerkleProofLib.sol";

/**
 * @title Plug Vault Socket
 * @author @nftchance (chance@onplug.io)
 */
contract PlugVaultSocket is PlugSocket, PlugTrading, Receiver, UUPSUpgradeable {
    /*
    * @notice The constructor for the Plug Vault Socket will
    *         initialize to address(1) when not deployed through
    *         a Socket factory.
    */
    constructor() {
        initialize(address(1), address(1));
    }

    /**
     * @notice Initializes a new Plug Vault Socket.
     * @param $ownership The address of the ownership proxy.
     * @param $router The address of the router.
     */
    function initialize(address $ownership, address $router) public {
        /// @dev Associate the ownership proxy as the factory that deployed
        ///      the contract at the same time as deployment.
        _initializeOwnership($ownership, $router);
    }

    /**
     * See { PlugSocket-name }
     */
    function name() public pure override returns (string memory $name) {
        $name = "Plug Vault Socket";
    }

    /**
     * See { PlugSocket-version }
     */
    function version() public pure override returns (string memory $version) {
        $version = "0.0.1";
    }

    /**
     * @notice Confirm that the only specified routers can execute the transaction.
     * @dev If you would like to limit the available routers override this
     *      function in your contract with the additional logic.
     * @param $router The router of the transaction.
     */
    function _enforceRouter(address $router)
        internal
        view
        virtual
        override
        returns (bool $allowed)
    {
        $allowed = $router == router;
    }

    /**
     * See { PlugEnforce._enforceSignature }
     */
    function _enforceSignature(PlugTypesLib.LivePlugs calldata $input)
        internal
        view
        virtual
        override
        returns (bool $allowed)
    {
        /// @dev The hash of the Plugs bundle that is used to represent the unique
        ///      state of declaration and execution intent.
        bytes32 plugsHash = getPlugsHash($input.plugs);

        /// @dev The last bit denotes whether it is a standard signature
        ///      or a merkle proof signature.
        bytes1 signatureType = $input.signature[0];

        /// @dev Utilize a standard signature recovery method that is only designed
        ///      to support one domain and intent at a time.
        if (signatureType & 0x03 == signatureType) {
            ($allowed,) = _signatureValidation(plugsHash, $input.signature);
        }
        /// @dev Utilize a merkle proof signature recovery method that holds several
        ///      domains and intents at a time inside a single signature.
        else if (signatureType & 0x04 == signatureType) {
            /// @dev Recover the merkle proof data from the packed signature.
            (bytes32 root, bytes32[] memory proof, bytes memory signature) =
                abi.decode($input.signature[1:], (bytes32, bytes32[], bytes));

            /// @dev Ensure the merkle tree contains the data of the signed bundle.
            require(
                MerkleProofLib.verify(proof, root, getPlugsHash($input.plugs)),
                "PlugTypes:invalid-proof"
            );

            /// @dev Calculate the offset needed to extract solely the signature from
            ///      the packed state of the `signature` data provided.
            uint256 offset = proof.length * 32 + 161;

            ($allowed,) =
                _signatureValidation(plugsHash, $input.signature[offset:offset + signature.length]);
        }
    }

    /**
     * See { UUPSUpgradeable._authorizeUpgrade }
     */
    function _authorizeUpgrade(address) internal virtual override onlyOwner { }

    /**
     * See { PlugTrading._guardInitializeOwnership }
     */
    function _guardInitializeOwnership() internal pure virtual override returns (bool $guard) {
        $guard = true;
    }
}
