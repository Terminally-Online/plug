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
    PlugVaultSocket,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";

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
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray);

        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_SignerSolver() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray);

        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidRouter() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray);

        plug = new Plug();
        vm.expectRevert(abi.encodeWithSelector(PlugLib.RouterInvalid.selector, address(plug)));
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidSignature() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    /// @dev Invalid signer private key.
                    0xabc1234,
                    false
                )
            )
        });

        vm.expectRevert(PlugLib.SignatureInvalid.selector);
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_ExternalSolver_NotCompensated() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);
        uint256 preBalance = address(solver).balance;

        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, 0, 0, solver);

        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
        assertEq(preBalance, address(solver).balance);
    }

    function test_PlugEmptyEcho_ExternalSolver_Compensated() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);
        uint256 preBalance = address(vault).balance;

        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, 0.2 ether, 1, solver);

        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
        assertTrue(preBalance - 1 ether > address(vault).balance);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_CompensationFailure() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 0);

        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, 0.2 ether, 24, solver);

        vm.prank(solver);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.ValueInvalid.selector,
                PlugEtcherLib.PLUG_TREASURY_ADDRESS,
                PLUG_VALUE,
                address(vault).balance
            )
        );
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_Invalid() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);

        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs memory livePlugs = createLivePlugs(plugsArray, 0, 0, solver);

        vm.expectRevert(
            abi.encodeWithSelector(PlugLib.SolverInvalid.selector, address(solver), address(this))
        );
        plug.plug(livePlugs);
    }
}
