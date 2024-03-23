// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

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

    function test_PlugEmptyEcho_SignerSolver() public {
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidRouter() public {
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        plug = new Plug();
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.RouterInvalid.selector, address(plug)
            )
        );
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidSignature() public {
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(vault.getPlugsHash(plugs), address(vault), 0xabc1234, false),
                1,
                1,
                1
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

        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: abi.encode(uint96(0), uint96(0), solver)
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        vm.prank(solver);
        plug.plug(livePlugs);
        assertEq(preBalance, address(solver).balance);
    }

    function test_PlugEmptyEcho_ExternalSolver_Compensated() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);
        uint256 preBalance = address(vault).balance;

        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0.2 ether), uint96(1), solver)
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");

        vm.prank(solver);
        plug.plug(livePlugs);
        assertTrue(preBalance - 1 ether > address(vault).balance);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_CompensationFailure()
        public
    {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 0);

        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector)
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0.2 ether), uint96(24), solver)
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        vm.prank(solver);
        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.CompensationFailed.selector,
                PlugAddressesLib.PLUG_TREASURY_ADDRESS,
                1 ether
            )
        );
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_Invalid() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);

        bytes memory encodedTransaction =
            abi.encodeWithSelector(PlugMockEcho.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0), uint96(0), address(solver))
        });
        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(
                sign(
                    vault.getPlugsHash(plugs),
                    address(vault),
                    signerPrivateKey,
                    false
                ),
                1,
                1,
                1
                )
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                PlugLib.SolverInvalid.selector, address(solver), address(this)
            )
        );
        plug.plug(livePlugs);
    }
}
