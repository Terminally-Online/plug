// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugMockEcho } from "./Plug.Mock.Echo.sol";
import { PlugTypes, PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugRouterSocket } from "../sockets/Plug.Router.Socket.sol";

contract PlugMockSocketTest is Test {
    PlugMockEcho internal mock;
    PlugRouterSocket internal router;

    address internal signer;
    uint256 internal signerPrivateKey;

    uint8 internal v;
    bytes32 internal r;
    bytes32 internal s;
    bytes32 internal digest;

    function setUp() public {
        mock = new PlugMockEcho();

        signerPrivateKey = 0xabc123;
        signer = vm.addr(signerPrivateKey);
    }

    function test_Echo() public {
        string memory expected = "Hello World";
        vm.expectEmit(address(mock));
        emit PlugMockEcho.EchoInvoked(address(this), address(this), expected);
        mock.echo(expected);
    }

    function test_EmptyEcho() public {
        vm.expectEmit(address(mock));
        emit PlugMockEcho.EchoInvoked(
            address(this), address(this), "Hello World"
        );
        mock.emptyEcho();
    }

    function test_MutedEcho(uint256 $echo) public view {
        mock.mutedEcho($echo);
    }

    // TODO: Update this work with the changes made to simply break out the need for
    ///      embedded sockets just to recover the sender.

    // function test_GetLivePlugsSigner() public {
    //     /// @dev Encode the transaction that is going to be called.
    //     bytes memory encodedTransaction =
    //         abi.encodeWithSelector(mock.mutedEcho.selector);
    //     PlugTypesLib.Current memory current = PlugTypesLib.Current({
    //         ground: address(mock),
    //         voltage: 0,
    //         data: encodedTransaction
    //     });
    //
    //     /// @dev Instantiate the pin and sign it.
    //     PlugTypesLib.LivePin[] memory livePins = new PlugTypesLib.LivePin[](1);
    //     PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](0);
    //     PlugTypesLib.Pin memory pin = PlugTypesLib.Pin({
    //         neutral: signer,
    //         live: bytes32(0),
    //         fuses: fuses,
    //         salt: bytes32(0),
    //         forced: true
    //     });
    //     digest = mock.getPinDigest(pin);
    //     (v, r, s) = vm.sign(signerPrivateKey, digest);
    //     bytes memory pinSignature = abi.encodePacked(r, s, v);
    //
    //     /// @dev Append the pin to the live pins array.
    //     PlugTypesLib.LivePin memory livePin =
    //         PlugTypesLib.LivePin({ pin: pin, signature: pinSignature });
    //     address pinSigner = mock.getLivePinSigner(livePin);
    //     assertEq(pinSigner, signer);
    //     livePins[0] = livePin;
    //
    //     /// @dev Bundle the plug and sign it.
    //     PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
    //         pins: livePins,
    //         current: current,
    //         forced: true
    //     });
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = Plug;
    //     PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
    //         breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
    //         plugs: plugsArray
    //     });
    //     digest = mock.getPlugsDigest(plugs);
    //     (v, r, s) = vm.sign(signerPrivateKey, digest);
    //     bytes memory plugsSignature = abi.encodePacked(r, s, v);
    //     PlugTypesLib.LivePlugs memory livePlugs =
    //         PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
    //     address plugsSigner = mock.getLivePlugsSigner(livePlugs);
    //     assertEq(plugsSigner, signer);
    // }
    //
    // function test_PlugEmptyEcho_SignerExecutor() public {
    //     /// @dev Encode the transaction that is going to be called.
    //     bytes memory encodedTransaction =
    //         abi.encodeWithSelector(mock.emptyEcho.selector);
    //     PlugTypesLib.Current memory current = PlugTypesLib.Current({
    //         ground: address(mock),
    //         voltage: 0,
    //         data: encodedTransaction
    //     });
    //
    //     /// @dev There are no conditions in this plug meaning a user is executing their own intent.
    //     PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
    //         pins: new PlugTypesLib.LivePin[](0),
    //         current: current,
    //         forced: true
    //     });
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = Plug;
    //
    //     /// @dev Make sure this transaction cannot be replayed.
    //     PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
    //         breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
    //         plugs: plugsArray
    //     });
    //
    //     /// @dev Sign the execution.
    //     digest = mock.getPlugsDigest(plugs);
    //     (v, r, s) = vm.sign(signerPrivateKey, digest);
    //     bytes memory plugsSignature = abi.encodePacked(r, s, v);
    //     PlugTypesLib.LivePlugs memory livePlugs =
    //         PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
    //     address plugsSigner = mock.getLivePlugsSigner(livePlugs);
    //     assertEq(plugsSigner, signer);
    //
    //     /// @dev Execute the plug.
    //     mock.plug(livePlugs);
    // }
    //
    // function test_PlugEmptyEcho_ExternalExecutor() public {
    //     /// @dev Encode the transaction that is going to be called.
    //     bytes memory encodedTransaction =
    //         abi.encodeWithSelector(mock.emptyEcho.selector);
    //     PlugTypesLib.Current memory current = PlugTypesLib.Current({
    //         ground: address(mock),
    //         voltage: 0,
    //         data: encodedTransaction
    //     });
    //
    //     /// @dev There are no conditions in this plug meaning an executor can do anything.
    //     PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
    //         pins: new PlugTypesLib.LivePin[](0),
    //         current: current,
    //         forced: true
    //     });
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = Plug;
    //
    //     /// @dev Make sure this transaction cannot be replayed.
    //     PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
    //         breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
    //         plugs: plugsArray
    //     });
    //
    //     /// @dev Sign the execution.
    //     digest = mock.getPlugsDigest(plugs);
    //     (v, r, s) = vm.sign(signerPrivateKey, digest);
    //     bytes memory plugsSignature = abi.encodePacked(r, s, v);
    //     PlugTypesLib.LivePlugs memory livePlugs =
    //         PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
    //     address plugsSigner = mock.getLivePlugsSigner(livePlugs);
    //     assertEq(plugsSigner, signer);
    //
    //     /// @dev Execute the plug.
    //     vm.expectEmit(address(mock));
    //     emit PlugMockEcho.EchoInvoked(address(mock), signer, "Hello World");
    //     hoax(_randomNonZeroAddress());
    //     mock.plug(livePlugs);
    // }
    //
    // function testFail_PlugMutedEcho() public {
    //     /// @dev Encode the transaction that is going to be called.
    //     bytes memory encodedTransaction =
    //         abi.encodeWithSelector(mock.mutedEcho.selector);
    //     PlugTypesLib.Current memory current = PlugTypesLib.Current({
    //         ground: address(mock),
    //         voltage: 0,
    //         data: encodedTransaction
    //     });
    //
    //     /// @dev There are no conditions in this plug meaning a user is executing their own intent.
    //     PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
    //         pins: new PlugTypesLib.LivePin[](0),
    //         current: current,
    //         forced: true
    //     });
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = Plug;
    //
    //     /// @dev Make sure this transaction cannot be replayed.
    //     PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
    //         breaker: PlugTypesLib.Breaker({ nonce: 1, queue: 0 }),
    //         plugs: plugsArray
    //     });
    //
    //     /// @dev Sign the execution.
    //     digest = mock.getPlugsDigest(plugs);
    //     (v, r, s) = vm.sign(signerPrivateKey, digest);
    //     bytes memory plugsSignature = abi.encodePacked(r, s, v);
    //     PlugTypesLib.LivePlugs memory livePlugs =
    //         PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
    //     address plugsSigner = mock.getLivePlugsSigner(livePlugs);
    //     assertEq(plugsSigner, signer);
    //
    //     /// @dev Execute the plug.
    //     mock.plug(livePlugs);
    // }
}
