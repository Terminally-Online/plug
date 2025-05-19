// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {
    Test,
    PlugLib,
    PlugEtcherLib,
    PlugTypesLib,
    Plug,
    PlugFactory,
    PlugSocket,
    PlugMockEcho
} from "../abstracts/test/Plug.Test.sol";
import { Ownable } from "solady/auth/Ownable.sol";
import "forge-std/console2.sol";
import { ECDSA } from "solady/utils/ECDSA.sol";

contract PlugSocketTest is Test {
    function setUp() public virtual {
        setUpPlug();
    }

    function test_name() public {
        assertEq(socket.name(), "Plug Socket");
    }

    function test_symbol() public {
        assertEq(socket.symbol(), "PLUGS");
    }

    function test_domain() public {
        assertNotEq(socket.domainHash(), "");
    }

    function testRevert_Initialize_Again() public {
        vm.deal(address(socket), 100 ether);
        vm.expectRevert("PlugTypes:already-initialized");
        socket.initialize(signer, oneClicker);
    }

    function test_owner_Implementation() public {
        assertEq(socketImplementation.owner(), address(1));
    }

    function test_owner() public {
        assertEq(socket.owner(), signer);
    }

    function test_transferOwnership() public {
        vm.prank(signer);
        socket.transferOwnership(_randomNonZeroAddress());
    }

    function testRevert_transferOwnership() public {
        vm.expectRevert(Ownable.Unauthorized.selector);
        socket.transferOwnership(_randomNonZeroAddress());
    }

    function test_domainHashSet() public {
        // Deploy a new vault
        socket = deployVault();
        
        // Get the expected domain hash by constructing it manually
        bytes32 expectedDomainHash = socket.getEIP712DomainHash(
            PlugTypesLib.EIP712Domain({
                name: socket.name(),
                version: socket.version(),
                chainId: block.chainid,
                verifyingContract: address(socket)
            })
        );

        // Compare the actual domain hash with the expected one
        assertEq(
            socket.domainHash(),
            expectedDomainHash,
            "Domain hash should match expected value"
        );
        
        // Verify it's not zero
        assertTrue(socket.domainHash() != 0x0, "Domain hash should not be zero");
    }

    function test_debugSignatureValidation() public {
        address solverAddr = 0xE09C3f7a144D3b1Bef42499486cA5C36d20ec5d1;
        address socketAddr = 0xFc49633D9e97a489E6E4CE63D21Ff98D5918C96d;
        
        PlugTypesLib.Plugs memory testPlugs;
        testPlugs.socket = socketAddr;
        
        PlugTypesLib.Plug[] memory plugArray = new PlugTypesLib.Plug[](1);
        plugArray[0].data = hex"a9059cbb0000000000000000000000000bb5d848487b10f8cfba21493c8f6d47e8a8b17c000000000000000000000000000000000000000000000000000000174876e800";
        plugArray[0].updates = new PlugTypesLib.Update[](0);
        testPlugs.plugs = plugArray;
        
        // Set solver and salt to exact same bytes as Go
        testPlugs.solver = hex"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e09c3f7a144d3b1bef42499486ca5c36d20ec5d1";
        testPlugs.salt = hex"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fc49633d9e97a489e6e4ce63d21ff98d5918c96d000000000000000000000000e09c3f7a144d3b1bef42499486ca5c36d20ec5d100000000000000000000000000000000006dbf3eb0da1c6b0629866440addc0a";

 bytes32 PLUGS_TYPEHASH =
        0x05b2ab8b8c7ceee9902f5288470f7189883657d476121976b1079d47722718a2;

        bytes32 plugArrayHash = socket.getPlugArrayHash(testPlugs.plugs);
        console2.log("Components being packed:");
        console2.log("1. PLUGS_TYPEHASH:", vm.toString(PLUGS_TYPEHASH));
        console2.log("2. socket:", vm.toString(testPlugs.socket));
        console2.log("3. plugArrayHash:", vm.toString(plugArrayHash));
        console2.log("4. solverHash:", vm.toString(keccak256(testPlugs.solver)));
        console2.log("5. saltHash:", vm.toString(keccak256(testPlugs.salt)));

        console2.log("testPlugs.socket:", vm.toString(testPlugs.socket));
        console2.log("testPlugs.plugs[0].data:", vm.toString(testPlugs.plugs[0].data));
        console2.log("testPlugs.solver:", vm.toString(testPlugs.solver));
        console2.log("testPlugs.salt:", vm.toString(testPlugs.salt));


        bytes32 plugsHash = socket.getPlugsHash(testPlugs);
        console2.log("plugsHash:", vm.toString(plugsHash));
        
        bytes32 domainHashValue = socket.getEIP712DomainHash(
            PlugTypesLib.EIP712Domain({
                name: "Plug Socket",
                version: "0.0.1",
                chainId: 8453,
                verifyingContract: socketAddr
            })
        );
        console2.logBytes32(domainHashValue);
        
        bytes memory prefix = hex"1901";
        console2.logBytes(prefix);
        console2.logBytes32(domainHashValue);
        console2.logBytes32(plugsHash);

        bytes memory digestInput = abi.encodePacked("\x19\x01", domainHashValue, plugsHash);
        console2.logBytes(digestInput);

        bytes32 finalDigest = keccak256(digestInput);
        console2.logBytes32(finalDigest);

        bytes memory signature = hex"55380155db822987674a0128637b0aa9e0d52056a33a8058a51d28965a16861762a97bed8fb148751b7a816725e0917a7d3e46d898fd158dfce33036aa89eabc1b";
        require(signature.length == 65, "Invalid signature length");

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }

        address recovered = ecrecover(finalDigest, v, r, s);
        console2.log("Recovered address:", recovered);

        address recovered2 = ECDSA.recover(finalDigest, signature);
        console2.log("Recovered address (ECDSA):", recovered2);
    }
}
