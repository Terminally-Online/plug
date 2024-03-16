// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Test } from "../utils/Test.sol";

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugRevocationFuse } from "./Plug.Revocation.Fuse.sol";

contract PlugRevocationFuseTest is Test {
    PlugRevocationFuse internal fuse;
    PlugTypesLib.Current internal current =
        PlugTypesLib.Current({ target: address(fuse), value: 0, data: "0x" });
    bytes32 plugsHash = bytes32("0");

    function setUp() public {
        fuse = new PlugRevocationFuse();
    }

    function test_enforceFuse_NotRevoked() public {
        bytes memory terms = fuse.encode(address(this));
        (address decodedSigner) = fuse.decode(terms);
        assertEq(decodedSigner, address(this));
        fuse.enforceFuse(terms, current, plugsHash);
    }

    function testRevert_enforceFuse_Revoked() public {
        fuse.revoke(plugsHash, true);
        bytes memory terms = fuse.encode(address(this));
        vm.expectRevert(bytes("PlugRevocationFuse:revoked"));
        fuse.enforceFuse(terms, current, plugsHash);
    }
}
