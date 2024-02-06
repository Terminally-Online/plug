// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugMockEcho } from "../mocks/Plug.Mock.Echo.sol";
import { PlugTypes, PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugCore } from "../abstracts/Plug.Core.sol";
import { PlugRevocationFuse } from "./Plug.Revocation.Fuse.sol";
import { PlugRouterSocket } from "../sockets/Plug.Router.Socket.sol";
import { PlugEtcher } from "../utils/Plug.Etcher.sol";

contract PlugRevocationFuseTest is Test {
    PlugMockEcho internal mock;
    PlugRevocationFuse internal fuse;
    PlugRouterSocket internal router;

    address internal signer;
    uint256 internal signerPrivateKey;

    uint8 internal v;
    bytes32 internal r;
    bytes32 internal s;
    bytes32 internal digest;

    function setUp() public {
        mock = new PlugMockEcho();
        fuse = new PlugRevocationFuse();

        signerPrivateKey = 0xabc123;
        signer = vm.addr(signerPrivateKey);

        // router = PlugEtcher.routerSocket();
        router = etchRouterSocket();
    }

    function etchRouterSocket() internal returns (PlugRouterSocket) {
        vm.etch(
            0x00b09C89Ace100AB7A4Dc47ebfBd1E7997920062,
            address(new PlugRouterSocket()).code
        );
        return PlugRouterSocket(
            payable(0x00b09C89Ace100AB7A4Dc47ebfBd1E7997920062)
        );
    }

    function test_enforceFuse() public {
        /// @dev Encode the transaction that is going to be called.
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 0,
            data: abi.encodeWithSelector(mock.emptyEcho.selector)
        });

        /// @dev Prepare the fuses.
        PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](1);
        fuses[0] = PlugTypesLib.Fuse({
            neutral: address(fuse),
            live: fuse.encode(signer)
        });

        /// @dev Prepare the Plugs.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({ current: current, fuses: fuses });

        PlugTypesLib.Plugs memory plugs =
            PlugTypesLib.Plugs({ plugs: plugsArray, salt: bytes32(0) });

        /// @dev Sign the execution.
        digest = router.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = router.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        router.plug(livePlugs);
    }

    function testRevert_enforceFuse_revoked() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 0,
            data: encodedTransaction
        });

        /// @dev Instantiate the revocation fuse and include it in the pin.
        PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](1);
        PlugTypesLib.Fuse memory revocationFuse = PlugTypesLib.Fuse({
            neutral: address(fuse),
            live: fuse.encode(signer)
        });
        fuses[0] = revocationFuse;

        bytes32 domainHash = router.domainHash();

        PlugTypesLib.Plug memory Plug =
            PlugTypesLib.Plug({ current: current, fuses: fuses });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs =
            PlugTypesLib.Plugs({ plugs: plugsArray, salt: bytes32(0) });

        /// @dev Sign the execution.
        digest = router.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = router.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Make sure random users can't revoke.
        vm.expectRevert(bytes("PlugRevocationFuse:invalid-revoker"));
        fuse.revoke(livePlugs, domainHash, true);

        /// @dev Revoke the pin.
        vm.prank(signer);
        fuse.revoke(livePlugs, domainHash, true);

        /// @dev Make sure the pin is revoked.
        bool revoked =
            fuse.isRevoked(signer, router.getPlugsHash(livePlugs.plugs));
        assertTrue(revoked);

        /// @dev Make sure you can't double revoke.
        vm.prank(signer);
        vm.expectRevert(bytes("PlugRevocationFuse:same-state"));
        fuse.revoke(livePlugs, domainHash, true);

        /// @dev Execute the plug.
        vm.expectRevert(bytes("PlugRevocationFuse:revoked"));
        router.plug(livePlugs);
    }
}
