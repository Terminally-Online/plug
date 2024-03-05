// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugFactory } from "../base/Plug.Factory.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";
import { PlugVaultSocket } from "./Plug.Vault.Socket.sol";
import { PlugMockEcho } from "../mocks/Plug.Mock.Echo.sol";

import { Initializable } from "solady/src/utils/Initializable.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract PlugVaultSocketTest is Test {
    PlugVaultSocket internal implementation;
    PlugVaultSocket internal vault;
    PlugFactory internal factory;
    PlugMockEcho internal mock;

    address internal signer;
    uint256 internal signerPrivateKey;

    uint8 internal v;
    bytes32 internal r;
    bytes32 internal s;
    bytes32 internal digest;

    function setUp() public virtual {
        signerPrivateKey = 0xabc123;
        signer = vm.addr(signerPrivateKey);

        implementation = new PlugVaultSocket();
        factory = new PlugFactory();

        (, address vaultAddress) =
            factory.deploy(address(implementation), address(this), bytes32(0));
        vault = PlugVaultSocket(payable(vaultAddress));
    }

    function etchRouterSocket() internal returns (Plug) {
        vm.etch(PlugEtcherLib.PLUG_ADDRESS, address(new Plug()).code);
        return Plug(payable(PlugEtcherLib.PLUG_ADDRESS));
    }

    function test_SingletonUse(uint256) public {
        vm.deal(address(vault), 100 ether);
        vm.expectRevert(bytes("PlugTypes:already-initialized"));
        vault.initialize(address(this));
    }

    function test_name() public {
        assertEq(vault.name(), "PlugVaultSocket");
    }

    function test_symbol() public {
        assertEq(vault.symbol(), "PVS");
    }

    function test_owner_Implementation() public {
        assertEq(implementation.owner(), address(1));
    }

    function test_owner() public {
        assertEq(vault.owner(), address(this));
    }

    function test_owner_getAccess() public {
        (bool isRouter, bool isSigner) = vault.getAccess(address(this));
        assertTrue(isSigner);
        assertFalse(isRouter);
    }

    function test_setAccess_Both() public {
        bool isRouter = true;
        bool isSigner = true;
        signer = _randomNonZeroAddress();
        uint8 access = vault.getAccess(isRouter, isSigner);
        vault.setAccess(signer, access);
        (isSigner, isRouter) = vault.getAccess(signer);
        assertTrue(isSigner);
        assertTrue(isRouter);
    }

    function test_setAccess_IsRouter() public {
        bool isRouter = true;
        bool isSigner = false;
        signer = _randomNonZeroAddress();
        uint8 access = vault.getAccess(isRouter, isSigner);
        vault.setAccess(signer, access);
        (isRouter, isSigner) = vault.getAccess(signer);
        assertFalse(isSigner);
        assertTrue(isRouter);
    }

    function test_setAccess_IsSigner() public {
        bool isRouter = false;
        bool isSigner = true;
        signer = _randomNonZeroAddress();
        uint8 access = vault.getAccess(isRouter, isSigner);
        vault.setAccess(signer, access);
        (isRouter, isSigner) = vault.getAccess(signer);
        assertFalse(isRouter);
        assertTrue(isSigner);
    }

    function testRevert_setAccess_Unauthorized() public {
        bool isRouter = false;
        bool isSigner = true;
        signer = _randomNonZeroAddress();
        uint8 access = vault.getAccess(isRouter, isSigner);
        vm.prank(_randomNonZeroAddress());
        vm.expectRevert(Ownable.Unauthorized.selector);
        vault.setAccess(signer, access);
    }

    function test_transferOwnership() public {
        bool isRouter;
        bool isSigner;
        address newOwner = _randomNonZeroAddress();
        vault.transferOwnership(newOwner);
        (isRouter, isSigner) = vault.getAccess(address(this));
        assertFalse(isSigner);
        (isRouter, isSigner) = vault.getAccess(newOwner);
        assertTrue(isSigner);
    }

    function testRevert_testOwnership_Unauthorized() public {
        vm.prank(_randomNonZeroAddress());
        vm.expectRevert(Ownable.Unauthorized.selector);
        vault.transferOwnership(_randomNonZeroAddress());
    }

    function test_GetLivePlugsSigner() public {
        /// @dev Encode the transaction that is going to be called.
        bytes memory encodedTransaction =
            abi.encodeWithSelector(mock.mutedEcho.selector);
        PlugTypesLib.Current memory current = PlugTypesLib.Current({
            target: address(mock),
            value: 0,
            data: encodedTransaction
        });

        /// @dev Bundle the Plug and sign it.
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = PlugTypesLib.Plug({
            current: current,
            fuses: new PlugTypesLib.Fuse[](0)
        });

        PlugTypesLib.Plugs memory plugs = PlugTypesLib.Plugs({
            socket: address(this),
            plugs: plugsArray,
            salt: bytes32(0),
            fee: 0,
            maxFeePerGas: 0,
            maxPriorityFeePerGas: 0,
            executor: address(0)
        });

        digest = vault.getPlugsDigest(plugs);
        (v, r, s) = vm.sign(signerPrivateKey, digest);
        bytes memory plugsSignature = abi.encodePacked(r, s, v);

        PlugTypesLib.LivePlugs memory livePlugs =
            PlugTypesLib.LivePlugs({ plugs: plugs, signature: plugsSignature });

        address plugsSigner = vault.getLivePlugsSigner(livePlugs);
        assertEq(plugsSigner, signer);
    }
}
