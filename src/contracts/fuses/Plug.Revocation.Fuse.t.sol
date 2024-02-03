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

        router = PlugEtcher.routerSocket();
    }

    function test_enforceFuse() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 0,
            data: encodedTransaction
        });

        /// @dev Instantiate the pin and sign it.
        PlugTypesLib.LivePin[] memory livePins = new PlugTypesLib.LivePin[](1);

        /// @dev Instantiate the revocation fuse and include it in the pin.
        PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](1);
        PlugTypesLib.Fuse memory revocationFuse = PlugTypesLib.Fuse({
            neutral: address(fuse),
            live: bytes("0"),
            forced: true
        });
        fuses[0] = revocationFuse;

        PlugTypesLib.Pin memory pin = PlugTypesLib.Pin({
            neutral: signer,
            live: bytes32(0),
            fuses: fuses,
            salt: bytes32(0),
            forced: true
        });
        digest = router.getPinDigest(pin);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory pinSignature = abi.encodePacked(r, s, v);

        /// @dev Append the pin to the live pins array.
        PlugTypesLib.LivePin memory livePin =
            PlugTypesLib.LivePin({ pin: pin, signature: pinSignature });
        bytes32 pinHash = router.getLivePinHash(livePin);
        assertNotEq(pinHash, bytes32(0));
        address pinSigner = router.getLivePinSigner(livePin);
        assertEq(pinSigner, signer);
        livePins[0] = livePin;

        PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
            pins: livePins,
            current: current,
            forced: true
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
            plugs: plugsArray
        });

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

        /// @dev Instantiate the pin and sign it.
        PlugTypesLib.LivePin[] memory livePins = new PlugTypesLib.LivePin[](1);

        /// @dev Instantiate the revocation fuse and include it in the pin.
        PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](1);
        PlugTypesLib.Fuse memory revocationFuse = PlugTypesLib.Fuse({
            neutral: address(fuse),
            live: bytes("0"),
            forced: true
        });
        fuses[0] = revocationFuse;

        PlugTypesLib.Pin memory pin = PlugTypesLib.Pin({
            neutral: signer,
            live: bytes32(0),
            fuses: fuses,
            salt: bytes32(0),
            forced: true
        });
        digest = router.getPinDigest(pin);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory pinSignature = abi.encodePacked(r, s, v);

        /// @dev Append the pin to the live pins array.
        PlugTypesLib.LivePin memory livePin =
            PlugTypesLib.LivePin({ pin: pin, signature: pinSignature });
        bytes32 pinHash = router.getLivePinHash(livePin);
        assertNotEq(pinHash, bytes32(0));
        address pinSigner = router.getLivePinSigner(livePin);
        assertEq(pinSigner, signer);
        livePins[0] = livePin;

        bytes32 domainHash = router.domainHash();
        address fuseSigner = fuse.getSigner(livePin, domainHash);

        /// @dev Make sure random users can't revoke.
        vm.expectRevert(bytes("PlugRevocationFuse:invalid-revoker"));
        fuse.revoke(livePin, domainHash);

        /// @dev Revoke the pin.
        assertEq(fuseSigner, signer);
        vm.prank(signer);
        fuse.revoke(livePin, domainHash);

        /// @dev Make sure the pin is revoked.
        bool revoked = fuse.isRevoked(pinHash);
        assertTrue(revoked);

        /// @dev Make sure you can't double revoke.
        vm.prank(signer);
        vm.expectRevert(bytes("PlugRevocationFuse:already-revoked"));
        fuse.revoke(livePin, domainHash);

        PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
            pins: livePins,
            current: current,
            forced: true
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
            plugs: plugsArray
        });

        /// @dev Sign the execution.
        digest = router.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = router.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.expectRevert(bytes("PlugRevocationFuse:revoked"));
        router.plug(livePlugs);
    }
}
