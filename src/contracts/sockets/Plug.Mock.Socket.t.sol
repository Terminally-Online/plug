// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PRBTest } from "@prb/test/PRBTest.sol";
import { console2 } from "forge-std/console2.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { TestPlus } from "../tests/TestPlus.sol";

import { PlugMockSocket } from "./Plug.Mock.Socket.sol";
import { PlugTypes, PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugCore } from "../abstracts/Plug.Core.sol";

contract PlugMockSocketTest is PRBTest, StdCheats, TestPlus {
    PlugMockSocket internal mock;

    address internal signer;
    uint256 internal signerPrivateKey;

    uint8 internal v;
    bytes32 internal r;
    bytes32 internal s;
    bytes32 internal digest;

    function setUp() public {
		mock = new PlugMockSocket('PlugMockSocket', '0.0.0');

		signerPrivateKey = 0xabc123;
		signer = vm.addr(signerPrivateKey);
    }

    function test_Echo() public {
		string memory expected = 'Hello World';
		vm.expectEmit(address(mock));
		emit PlugMockSocket.EchoInvoked(address(this), address(this), expected);
		mock.echo(expected);
    }

    function test_EmptyEcho() public {
		vm.expectEmit(address(mock));
		emit PlugMockSocket.EchoInvoked(address(this), address(this), 'Hello World');
		mock.emptyEcho();
    }

    function testFail_MutedEcho() public {
		vm.expectRevert('EchoMuted');
		mock.mutedEcho();
    }

    function test_GetLivePlugsSigner() public {
	    PlugTypesLib.Fuse[] memory fuses = new PlugTypesLib.Fuse[](0);
		PlugTypesLib.Pin memory pin = PlugTypesLib.Pin({
			neutral: signer,
			live: bytes32(0),
			fuses: fuses,
			salt: bytes32(0),
			forced: true
		});

		digest = mock.getPinDigest(pin);
		(v, r, s) = vm.sign(signerPrivateKey, digest);

		bytes memory pinSignature = abi.encodePacked(r, s, v);
		PlugTypesLib.LivePin memory livePin = PlugTypesLib.LivePin({
			pin: pin,
			signature: pinSignature
		});
		address pinSigner = mock.getLivePinSigner(livePin);
		assertEq(pinSigner, signer);

        bytes memory encodedTransaction = abi.encodeWithSelector(
            mock.mutedEcho.selector
        );
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 0,
            data: encodedTransaction
        });
        PlugTypesLib.LivePin[] memory livePins = new PlugTypesLib.LivePin[](1);
        livePins[0] = livePin;
        PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
            pins: livePins,
            current: current,
            forced: true
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            breaker: PlugTypesLib.Breaker({
                nonce: 1,
                queue: 0
            }),
            plugs: plugsArray
        });
        digest = mock.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: plugsSignature
        });
        address plugsSigner = mock.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);
    }

    function test_PlugEmptyEcho() public {
        bytes memory encodedTransaction = abi.encodeWithSelector(
            mock.emptyEcho.selector
        );
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 21000,
            data: encodedTransaction
        });
        PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
            pins: new PlugTypesLib.LivePin[](0),
            current: current,
            forced: true
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            breaker: PlugTypesLib.Breaker({
                nonce: 1,
                queue: 0
            }),
            plugs: plugsArray
        });
        digest = mock.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: plugsSignature
        });
        address plugsSigner = mock.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = livePlugs;
        mock.plug(livePlugsArray);
    }


    function testFail_PlugMutedEcho() public {
        bytes memory encodedTransaction = abi.encodeWithSelector(
            mock.mutedEcho.selector
        );
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            ground: address(mock),
            voltage: 21000,
            data: encodedTransaction
        });
        PlugTypesLib.Plug memory Plug = PlugTypesLib.Plug({
            pins: new PlugTypesLib.LivePin[](0),
            current: current,
            forced: true
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = Plug;
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            breaker: PlugTypesLib.Breaker({
                nonce: 1,
                queue: 0
            }),
            plugs: plugsArray
        });
        digest = mock.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: plugsSignature
        });
        address plugsSigner = mock.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = livePlugs;
        mock.plug(livePlugsArray);
    }
}
