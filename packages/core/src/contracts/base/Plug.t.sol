// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugAddressesLib,
    PlugEtcherLib,
    PlugFactory,
    Plug,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";
import { ECDSA } from "solady/utils/ECDSA.sol";

contract PlugTest is Test {
    event EchoInvoked(address $sender, string $message);

    function setUp() public virtual {
        setUpPlug();
    }

    function test_name() public {
        assertEq(plug.name(), "Plug");
    }

    function test_symbol() public {
        assertEq(plug.symbol(), "PLUG");
    }

    function test_PlugEmptyEcho_TypeRecovery() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver_TreasuryPayment() public {
        address solver = _randomNonZeroAddress();
        address treasury = 0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E;
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        uint256 preBalance = address(treasury).balance;
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);
        assertTrue(address(treasury).balance > preBalance);
    }

    function testRevert_PlugEmptyEcho_Solver_TreasuryPaymentFailure() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 0);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray, solver);
        vm.prank(solver);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 1, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver_InvalidNonce() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray, solver);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:nonce-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_InvalidSignature() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        bytes32 digest = socket.getPlugsDigest(plugs);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(0x123456, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        livePlugsArray[0] = PlugTypesLib.LivePlugs({ plugs: plugs, signature: signature });

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:signature-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_Invalid() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs =
            createPlugs(plugsArray, uint48(block.timestamp + 3 minutes), solver);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:solver-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_Expired() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs =
            createPlugs(plugsArray, uint48(block.timestamp - 1 minutes), solver);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        vm.prank(solver);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:solver-expired" })
        );
        plug.plug(livePlugsArray);
    }

    function test_Plugs_PlugEmptyEcho_TypeRecovery() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_Direct() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);

        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);

        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 0, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_SecondFails() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2); // Array of 2 plugs
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 1, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_MultipleIndexes() public {
        for (uint8 testIndex = 0; testIndex < 5; testIndex++) {
            PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
            PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](testIndex + 2);
            for (uint8 i = 0; i < testIndex + 1; i++) {
                plugsArray[i] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
            }
            plugsArray[testIndex + 1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
            PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
            livePlugsArray[0] = createLivePlugs(plugs);

            vm.expectEmit(address(plug));
            emit PlugLib.PlugResult(
                0, PlugTypesLib.Result({ index: testIndex + 1, error: "PlugCore:plug-failed" })
            );
            plug.plug(livePlugsArray);
        }
    }

    // function test_PlugGasLimit_Success() public {
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
    //
    //     PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
    //     livePlugsArray[0] = createLivePlugs(plugsArray);
    //
    //     vm.expectEmit(address(mock));
    //     emit EchoInvoked(address(socket), "Hello World");
    //     plug.plug(livePlugsArray);
    // }

    // function test_PlugGasLimit_TooLow() public {
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
    //     plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION, 1000);
    //
    //     PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
    //     livePlugsArray[0] = createLivePlugs(plugsArray);
    //
    //     vm.expectEmit(address(plug));
    //     emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 0, error: "PlugCore:plug-failed" }));
    //     plug.plug(livePlugsArray);
    // }
    //
    // function test_PlugGasLimit_MultiplePlugs() public {
    //     PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
    //     plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION, 100_000);
    //     plugsArray[1] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION, 50_000);
    //
    //     PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
    //     livePlugsArray[0] = createLivePlugs(plugsArray);
    //
    //     vm.expectEmit(address(mock));
    //     emit EchoInvoked(address(socket), "Hello World");
    //     vm.expectEmit(address(mock));
    //     emit EchoInvoked(address(socket), "Hello World");
    //     plug.plug(livePlugsArray);
    // }
}
