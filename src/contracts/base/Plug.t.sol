// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { Test } from "../utils/Test.sol";
import { TestPlug } from "../utils/TestPlug.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugFactory } from "../base/Plug.Factory.sol";
import { Plug } from "./Plug.sol";
import { PlugVaultSocket } from "../sockets/Plug.Vault.Socket.sol";
import { PlugMockEcho } from "../mocks/Plug.Mock.Echo.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

import { Initializable } from "solady/src/utils/Initializable.sol";

contract PlugTest is Test {
    event EchoInvoked(address $sender, string $message);

    PlugFactory internal factory;
    Plug internal plug;
    PlugVaultSocket internal vaultImplementation;
    PlugVaultSocket internal vault;
    PlugMockEcho internal mock;

    address factoryOwner;
    string baseURI = "https://onplug.io/metadata/";

    address internal signer;
    uint256 internal signerPrivateKey;

    function setUp() public virtual {
        factoryOwner = _randomNonZeroAddress();

        signerPrivateKey = 0xabc123;
        signer = vm.addr(signerPrivateKey);

        plug = etchPlug();
        vaultImplementation = new PlugVaultSocket();
        factory = new PlugFactory(factoryOwner, baseURI);
        mock = new PlugMockEcho();

        vm.prank(factoryOwner);
        factory.setImplementation(0, address(vaultImplementation));
        (, address vaultAddress) =
            factory.deploy(bytes32(abi.encodePacked(signer, uint96(0))));

        vault = PlugVaultSocket(payable(vaultAddress));
    }

    function etchPlug() internal returns (Plug) {
        vm.etch(PlugEtcherLib.PLUG_ADDRESS, address(new Plug()).code);
        return Plug(payable(PlugEtcherLib.PLUG_ADDRESS));
    }

    function test_name() public {
        assertEq(plug.name(), "Plug");
    }

    function test_symbol() public {
        assertEq(plug.symbol(), "PLUG");
    }

    function test_PlugEmptyEcho_SignerSolver() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });

        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidRouter() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug on an invalid router.
        plug = new Plug();
        vm.expectRevert(bytes("Plug:invalid-router"));
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerSolver_InvalidSignature() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: bytes("")
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature =
            sign(vault.getPlugsHash(plugs), address(vault), 0xabc1234, false);

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug with an invalid signer.
        vm.expectRevert(bytes("Plug:invalid-signature"));
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_ExternalSolver_NotCompensated() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);

        uint256 preBalance = address(solver).balance;

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            solver: abi.encode(uint96(0), uint96(0), solver)
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        vm.prank(solver);
        plug.plug(livePlugs);

        uint256 postBalance = address(solver).balance;
        assertEq(preBalance, postBalance);
    }

    function test_PlugEmptyEcho_ExternalSolver_Compensated() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0.2 ether), uint96(1), solver)
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(vault), "Hello World");
        /// @dev Make sure the compensation successfully changes hands.
        uint256 preBalance = address(vault).balance;
        vm.prank(solver);
        plug.plug(livePlugs);
        uint256 postBalance = address(vault).balance;
        /// @dev Check if it's greater than to account for the solver fee.
        assertTrue(preBalance - 1 ether > postBalance);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_CompensationFailure()
        public
    {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 0);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0.2 ether), uint96(24), solver)
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug.
        vm.prank(solver);
        vm.expectRevert("Plug:compensation-failed");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_ExternalSolver_Invalid() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.emptyEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev There are no conditions in this plug meaning a user is executing
        ///      their own intent.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(vault), 100 ether);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            solver: abi.encode(uint96(0), uint96(0), address(solver))
        });

        /// @dev Sign the execution.
        bytes memory plugsSignature = sign(
            vault.getPlugsHash(plugs), address(vault), signerPrivateKey, false
        );

        PlugTypesLib.LivePlugs memory livePlugs = PlugTypesLib.LivePlugs({
            plugs: plugs,
            signature: pack(plugsSignature, 1, 1, 1)
        });

        /// @dev Execute the plug.
        vm.expectRevert(bytes("Plug:invalid-solver"));
        plug.plug(livePlugs);
    }
}
