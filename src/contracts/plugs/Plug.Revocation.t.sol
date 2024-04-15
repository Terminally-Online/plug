// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test, PlugTypesLib } from "../abstracts/test/Plug.Test.sol";
import { PlugRevocation } from "./Plug.Revocation.sol";

contract PlugRevocationTest is Test {
    PlugRevocation internal connector;

    bytes32 plugsHash = bytes32("0");

    function setUp() public {
        connector = new PlugRevocation();
    }

    function test_enforce_NotRevoked() public {
        bytes memory terms = connector.encode(address(this));
        (address decodedSigner) = connector.decode(terms);
        assertEq(decodedSigner, address(this));
        connector.enforce(terms, plugsHash);
    }

    function testRevert_enforce_Revoked() public {
        connector.revoke(plugsHash, true);
        bytes memory terms = connector.encode(address(this));
        vm.expectRevert(bytes("PlugRevocation:revoked"));
        connector.enforce(terms, plugsHash);
    }
}
