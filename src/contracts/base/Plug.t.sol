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
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugs);
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_Solver() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_Solver_TreasuryPayment() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        uint256 preBalance = address(treasury).balance;
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, solver);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugs);
        assertTrue(address(treasury).balance > preBalance);
    }

    function testRevert_PlugEmptyEcho_Solver_TreasuryPaymentFailure() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 0);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, solver);
        vm.prank(solver);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ValueInvalid.selector,
                PlugEtcherLib.PLUG_TREASURY_ADDRESS,
                PLUG_VALUE,
                address(socket).balance
            )
        );
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_Solver_InvalidNonce() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, solver);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugs);
        vm.expectRevert(PlugLib.NonceInvalid.selector);
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_Solver_InvalidSignature() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        bytes32 digest = socket.getPlugsDigest(plugs);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(0x123456, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: signature });
        vm.expectRevert(PlugLib.SignatureInvalid.selector);
        plug.plug(livePlugs);
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
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugs);
        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.SolverInvalid.selector, address(solver), address(this))
        );
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_Solver_Expired() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs =
            createPlugs(plugsArray, uint48(block.timestamp - 1 minutes), solver);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugs);
        vm.prank(solver);
        vm.expectRevert(PlugLib.SolverExpired.selector);
        plug.plug(livePlugs);
    }
}
