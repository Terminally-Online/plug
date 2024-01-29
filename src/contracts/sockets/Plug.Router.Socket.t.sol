// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../utils/Test.sol";

import { PlugFactory } from "../utils/Plug.Factory.sol";
import { PlugRouterSocket } from "./Plug.Router.Socket.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

contract PlugVaultSocketTest is Test {
    PlugRouterSocket internal implementation;
    PlugRouterSocket internal router;
    PlugFactory internal factory;

    function setUp() public virtual {
        implementation = new PlugRouterSocket();
        factory = new PlugFactory();

        (, address routerAddress) =
            factory.deploy(address(implementation), address(this), bytes32(0));
        router = PlugRouterSocket(payable(routerAddress));
    }

    function test_SingletonUse(uint256) public {
        vm.deal(address(router), 100 ether);
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        router.initialize(address(this));
    }
}
