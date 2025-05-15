// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { Test } from "forge-std/Test.sol";
import { PlugMockDex } from "./Plug.Mock.Dex.sol";
import { PlugMockERC20 } from "./Plug.Mock.ERC20.sol";

/**
 * @title PlugMockDexTest
 * @notice Tests for the PlugMockDex contract
 */
contract PlugMockDexTest is Test {
    PlugMockDex internal dex;
    PlugMockERC20 internal tokenA;
    PlugMockERC20 internal tokenB;

    address internal user;
    uint256 internal constant INITIAL_AMOUNT = 1000 ether;

    function setUp() public {
        dex = new PlugMockDex();
        tokenA = new PlugMockERC20();
        tokenB = new PlugMockERC20();

        user = address(0x1);
        vm.startPrank(user);

        // Mint some tokens to the user
        tokenA.mint(user, INITIAL_AMOUNT);
        tokenB.mint(user, INITIAL_AMOUNT);

        // Approve the DEX to spend tokens
        tokenA.approve(address(dex), type(uint256).max);
        tokenB.approve(address(dex), type(uint256).max);

        vm.stopPrank();
    }

    function test_SetSwapRate() public {
        // Set a 2:1 swap rate (2 tokenB for 1 tokenA)
        uint256 rate = 2 ether; // 2e18
        dex.setSwapRate(address(tokenA), address(tokenB), rate);

        // Verify rate was set correctly
        assertEq(dex.swapRates(address(tokenA), address(tokenB)), rate);
    }

    function test_SetFixedReturn() public {
        // Set a fixed return of 100 tokenB for any amount of tokenA
        uint256 fixedAmount = 100 ether;
        dex.setFixedReturn(address(tokenA), address(tokenB), fixedAmount);

        // Verify fixed return was set correctly
        assertEq(
            dex.fixedReturns(address(tokenA), address(tokenB)), fixedAmount
        );
    }

    function test_GetAmountOut_WithRate() public {
        // Set a 2:1 swap rate
        uint256 rate = 2 ether;
        dex.setSwapRate(address(tokenA), address(tokenB), rate);

        // Calculate expected amount for 10 tokenA (should get 20 tokenB)
        uint256 amountIn = 10 ether;
        uint256 expectedAmountOut = (amountIn * rate) / 1 ether;

        // Verify getAmountOut returns the expected amount
        uint256 amountOut =
            dex.getAmountOut(address(tokenA), address(tokenB), amountIn);
        assertEq(amountOut, expectedAmountOut);
    }

    function test_GetAmountOut_WithFixedReturn() public {
        // Set both a rate and a fixed return
        dex.setSwapRate(address(tokenA), address(tokenB), 2 ether);
        dex.setFixedReturn(address(tokenA), address(tokenB), 50 ether);

        // Fixed return should have priority
        uint256 amountOut =
            dex.getAmountOut(address(tokenA), address(tokenB), 10 ether);
        assertEq(amountOut, 50 ether);
    }

    function test_Swap_WithRate() public {
        // Set a 2:1 swap rate
        uint256 rate = 2 ether;
        dex.setSwapRate(address(tokenA), address(tokenB), rate);

        uint256 amountIn = 10 ether;
        uint256 expectedAmountOut = (amountIn * rate) / 1 ether;

        uint256 userABalanceBefore = tokenA.balanceOf(user);
        uint256 userBBalanceBefore = tokenB.balanceOf(user);

        // Add tokens to the DEX so it can actually perform the transfer
        tokenB.mint(address(dex), expectedAmountOut);

        vm.prank(user);
        uint256 amountOut = dex.swap(
            address(tokenA),
            address(tokenB),
            amountIn,
            expectedAmountOut - 1 ether // Allow some slippage
        );

        assertEq(amountOut, expectedAmountOut);

        // After the swap, the user should have less tokenA and more tokenB
        assertEq(tokenA.balanceOf(user), userABalanceBefore - amountIn);
        assertEq(tokenB.balanceOf(user), userBBalanceBefore + amountOut);
    }

    function test_Swap_WithFixedReturn() public {
        // Set a fixed return amount
        uint256 fixedAmount = 25 ether;
        dex.setFixedReturn(address(tokenA), address(tokenB), fixedAmount);

        uint256 amountIn = 10 ether;

        uint256 userABalanceBefore = tokenA.balanceOf(user);
        uint256 userBBalanceBefore = tokenB.balanceOf(user);

        // Add tokens to the DEX so it can actually perform the transfer
        tokenB.mint(address(dex), fixedAmount);

        vm.prank(user);
        uint256 amountOut = dex.swap(
            address(tokenA),
            address(tokenB),
            amountIn,
            fixedAmount // Exact amount expected
        );

        assertEq(amountOut, fixedAmount);

        // After the swap, the user should have less tokenA and the fixed amount more of tokenB
        assertEq(tokenA.balanceOf(user), userABalanceBefore - amountIn);
        assertEq(tokenB.balanceOf(user), userBBalanceBefore + fixedAmount);
    }

    function test_RevertWhen_InsufficientOutput() public {
        // Set a 2:1 swap rate
        dex.setSwapRate(address(tokenA), address(tokenB), 2 ether);

        uint256 amountIn = 10 ether;
        uint256 minAmountOut = 30 ether; // Expect 30, but will only get 20

        vm.expectRevert();
        vm.prank(user);
        dex.swap(address(tokenA), address(tokenB), amountIn, minAmountOut);
    }
}
