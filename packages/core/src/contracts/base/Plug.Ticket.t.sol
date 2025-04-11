// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;
//
// import { Test } from "../abstracts/test/Plug.Test.sol";
// import { PlugTicket } from "./Plug.Ticket.sol";
// import { PlugLib, PlugAddressesLib } from "../libraries/Plug.Lib.sol";
// import { Ownable } from "solady/auth/Ownable.sol";
// import { PlugMockEcho } from "../mocks/Plug.Mock.Echo.sol";
//
// contract PlugTicketTest is Test {
//     PlugTicket internal ticket;
//     address internal owner;
//     address internal contractAddress;
//     address internal userAddress;
//
//     event Transfer(address indexed from, address indexed to, uint256 indexed id);
//
//     function setUp() public {
//         owner = PlugAddressesLib.PLUG_OWNER_ADDRESS;
//         userAddress = _randomNonZeroAddress();
//
//         // Create a contract to test minting (since only contracts can mint tickets)
//         bytes memory bytecode = type(PlugMockEcho).creationCode;
//         bytes32 salt = keccak256(abi.encodePacked("test contract"));
//         contractAddress = _safeCreate2(salt, bytecode);
//
//         ticket = new PlugTicket();
//         ticket.initialize();
//     }
//
//     function test_constructor() public {
//         PlugTicket newTicket = new PlugTicket();
//         assertEq(newTicket.owner(), address(1));
//     }
//
//     function test_initialize() public {
//         assertEq(ticket.owner(), owner);
//     }
//
//     function testRevert_initialize_Twice() public {
//         vm.expectRevert();
//         ticket.initialize();
//     }
//
//     function test_name() public {
//         assertEq(ticket.name(), "Plug: Ticket");
//     }
//
//     function test_symbol() public {
//         assertEq(ticket.symbol(), "TICKET");
//     }
//
//     function test_baseURI() public {
//         assertEq(ticket.baseURI(), "https://onplug.io/nft/");
//     }
//
//     // function test_setBaseURI() public {
//     //     string memory newBaseURI = "https://newuri.io/ticket/";
//     //     vm.prank(owner);
//     //     ticket.setBaseURI(newBaseURI);
//     //     assertEq(ticket.baseURI(), newBaseURI);
//     // }
//
//     // function testRevert_setBaseURI_NotOwner() public {
//     //     string memory newBaseURI = "https://newuri.io/ticket/";
//     //     vm.expectRevert(Ownable.Unauthorized.selector);
//     //     ticket.setBaseURI(newBaseURI);
//     // }
//
//     function test_mint() public {
//         vm.prank(contractAddress);
//         vm.expectEmit(true, true, true, false);
//         emit Transfer(address(0), contractAddress, 0);
//         ticket.mint();
//
//         assertEq(ticket.totalSupply(), 1);
//         assertEq(ticket.balanceOf(contractAddress), 1);
//         assertEq(ticket.ownerOf(0), contractAddress);
//     }
//
//     function testRevert_mint_NotContract() public {
//         vm.prank(userAddress);
//         vm.expectRevert(PlugLib.CallerMustBeContract.selector);
//         ticket.mint();
//     }
//
//     function testRevert_mint_AlreadyMinted() public {
//         vm.prank(contractAddress);
//         ticket.mint();
//
//         vm.prank(contractAddress);
//         vm.expectRevert(PlugLib.AlreadyMinted.selector);
//         ticket.mint();
//     }
//
//     function test_tokenURI() public {
//         vm.prank(contractAddress);
//         ticket.mint();
//
//         string memory expectedURI = string.concat(ticket.baseURI(), "0");
//         assertEq(ticket.tokenURI(0), expectedURI);
//     }
//
//     function test_transferFrom_ZeroAddress() public {
//         // Minting is effectively a transfer from address(0)
//         vm.prank(contractAddress);
//         ticket.mint();
//
//         assertEq(ticket.ownerOf(0), contractAddress);
//     }
//
//     function test_transferFrom_ToZeroAddress() public {
//         // Set up token first
//         vm.prank(contractAddress);
//         ticket.mint();
//
//         // Burning is transfer to address(0)
//         vm.prank(contractAddress);
//         vm.expectEmit(true, true, true, false);
//         emit Transfer(contractAddress, address(0), 0);
//         ticket.transferFrom(contractAddress, address(0), 0);
//
//         // Check that token is burned by verifying that ownerOf(0) reverts
//         vm.expectRevert();
//         ticket.ownerOf(0);
//     }
//
//     function testRevert_transferFrom_NonZeroAddresses() public {
//         // Set up token first
//         vm.prank(contractAddress);
//         ticket.mint();
//
//         // Try to transfer between non-zero addresses
//         address recipient = _randomNonZeroAddress();
//         vm.prank(contractAddress);
//         vm.expectRevert(PlugLib.NonTransferableToken.selector);
//         ticket.transferFrom(contractAddress, recipient, 0);
//     }
//
//     function test_multipleMints() public {
//         // Create multiple contract addresses to test sequential minting
//         bytes memory bytecode = type(PlugMockEcho).creationCode;
//         bytes32 salt1 = keccak256(abi.encodePacked("test contract 1"));
//         bytes32 salt2 = keccak256(abi.encodePacked("test contract 2"));
//         bytes32 salt3 = keccak256(abi.encodePacked("test contract 3"));
//
//         address contract1 = _safeCreate2(salt1, bytecode);
//         address contract2 = _safeCreate2(salt2, bytecode);
//         address contract3 = _safeCreate2(salt3, bytecode);
//
//         // Mint tokens and verify sequential IDs
//         vm.prank(contract1);
//         ticket.mint();
//         assertEq(ticket.totalSupply(), 1);
//         assertEq(ticket.ownerOf(0), contract1);
//
//         vm.prank(contract2);
//         ticket.mint();
//         assertEq(ticket.totalSupply(), 2);
//         assertEq(ticket.ownerOf(1), contract2);
//
//         vm.prank(contract3);
//         ticket.mint();
//         assertEq(ticket.totalSupply(), 3);
//         assertEq(ticket.ownerOf(2), contract3);
//     }
// }
