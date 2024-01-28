// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../tests/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugAllowedMethodsFuse } from "./Plug.AllowedMethods.Fuse.sol";

contract PlugAllowedMethodsFuseTest is Test {
    PlugAllowedMethodsFuse internal fuse;

    bytes4 functionSignature = PlugAllowedMethodsFuse.decode.selector;
    PlugTypesLib.Current current = PlugTypesLib.Current({
        ground: address(fuse),
        voltage: 0,
        data: abi.encode(PlugAllowedMethodsFuse.decode.selector)
    });
    bytes32 pinHash = bytes32("0");

    function setUp() public virtual {
        fuse = new PlugAllowedMethodsFuse();
    }

    function testSingle_enforceFuse() public {
        bytes4[] memory signatures = new bytes4[](1);
        signatures[0] = PlugAllowedMethodsFuse.decode.selector;
        bytes memory terms = fuse.encode(signatures);
        bytes4[] memory decodedSignatures = fuse.decode(terms);
        assertEq(decodedSignatures[0], functionSignature);
        fuse.enforceFuse(terms, current, pinHash);
    }

    function testMultiple_enforceFuse() public {
        bytes4[] memory signatures = new bytes4[](2);
        signatures[0] = PlugAllowedMethodsFuse.encode.selector;
        signatures[1] = PlugAllowedMethodsFuse.decode.selector;
        bytes memory terms = fuse.encode(signatures);
        bytes4[] memory decodedSignatures = fuse.decode(terms);
        assertEq(decodedSignatures[0], PlugAllowedMethodsFuse.encode.selector);
        assertEq(decodedSignatures[1], PlugAllowedMethodsFuse.decode.selector);
        fuse.enforceFuse(terms, current, pinHash);
    }

    function testRevert_enforceFuse() public {
        bytes4[] memory signatures = new bytes4[](1);
        signatures[0] = PlugAllowedMethodsFuse.encode.selector;
        bytes memory terms = fuse.encode(signatures);
        vm.expectRevert(bytes("PlugAllowedMethodsFuse:method-not-allowed"));
        fuse.enforceFuse(terms, current, pinHash);
    }
}
