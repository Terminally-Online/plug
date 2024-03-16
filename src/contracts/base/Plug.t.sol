// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Test } from "../utils/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugFactory } from "../base/Plug.Factory.sol";
import { Plug } from "./Plug.sol";
import { PlugVaultSocket } from "../sockets/Plug.Vault.Socket.sol";
import { PlugMockEcho } from "../mocks/Plug.Mock.Echo.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

import { Initializable } from "solady/src/utils/Initializable.sol";

contract PlugTest is Test {
    PlugFactory internal factory;
    Plug internal plug;
    PlugVaultSocket internal vaultImplementation;
    PlugVaultSocket internal vault;
    PlugMockEcho internal mock;

    address factoryOwner;
    string baseURI = "https://onplug.io/metadata/";

    address internal signer;
    uint256 internal signerPrivateKey;

    uint8 internal v;
    bytes32 internal r;
    bytes32 internal s;
    bytes32 internal digest;

    function setUp() public virtual {
        factoryOwner = _randomNonZeroAddress();

        signerPrivateKey = 0xabc123;
        signer = vm.addr(signerPrivateKey);

        plug = etchPlug();
        vaultImplementation = new PlugVaultSocket();
        factory = new PlugFactory(factoryOwner, baseURI);
        mock = new PlugMockEcho();

        (, address vaultAddress) =
            factory.deploy(address(vaultImplementation), signer, bytes32(0));
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

    function test_PlugEmptyEcho_SignerExecutor() public {
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
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: address(0)
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit PlugMockEcho.EchoInvoked(address(vault), "Hello World");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerExecutor_InvalidRouter() public {
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
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: address(0)
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug on an invalid router.
        plug = new Plug();
        vm.expectRevert(bytes("Plug:invalid-router"));
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_SignerExecutor_InvalidSigner() public {
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
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: address(0)
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        signerPrivateKey = 0xdef456;
        signer = vm.addr(signerPrivateKey);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug with an invalid signer.
        vm.expectRevert(bytes("Plug:invalid-signer"));
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_ExternalExecutor_NotCompensated() public {
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

        address executor = _randomNonZeroAddress();
        vm.deal(executor, 100 ether);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: executor
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit PlugMockEcho.EchoInvoked(address(vault), "Hello World");
        vm.prank(executor);
        plug.plug(livePlugs);
    }

    function test_PlugEmptyEcho_ExternalExecutor_Compensated() public {
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

        address executor = _randomNonZeroAddress();
        vm.deal(executor, 100 ether);
        vm.deal(address(vault), 100 ether);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: executor
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.expectEmit(address(mock));
        emit PlugMockEcho.EchoInvoked(address(vault), "Hello World");
        vm.prank(executor);
        /// @dev Make sure the compensation successfully changes hands.
        uint256 preBalance = address(vault).balance;
        plug.plug(livePlugs);
        uint256 postBalance = address(vault).balance;
        assertEq(preBalance - 1 ether, postBalance);
    }

    function testRevert_PlugEmptyEcho_ExternalExecutor_CompensationFailure()
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

        address executor = _randomNonZeroAddress();
        vm.deal(executor, 100 ether);
        vm.deal(address(vault), 0);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: executor
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.prank(executor);
        vm.expectRevert("Plug:compensation-failed");
        plug.plug(livePlugs);
    }

    function testRevert_PlugEmptyEcho_ExternalExecutor_Invalid() public {
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

        address executor = _randomNonZeroAddress();
        vm.deal(executor, 100 ether);
        vm.deal(address(vault), 100 ether);

        /// @dev Make sure this transaction cannot be replayed.
        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(vault),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 1 ether,
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: executor
        });

        /// @dev Sign the execution.
        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);
        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });
        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);

        /// @dev Execute the plug.
        vm.expectRevert(bytes("Plug:invalid-executor"));
        plug.plug(livePlugs);
    }
}
