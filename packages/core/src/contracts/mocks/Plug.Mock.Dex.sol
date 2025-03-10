// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { ERC20 } from "solady/tokens/ERC20.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Mock DEX
 * @notice A mock DEX contract for testing the Plug framework
 * @dev This contract simulates a DEX with configurable swap returns
 */
contract PlugMockDex {
    /// @dev Event emitted when a swap is executed
    event SwapExecuted(
        address indexed sender,
        address indexed tokenIn,
        address indexed tokenOut,
        uint256 amountIn,
        uint256 amountOut
    );

    /// @dev Event emitted when the swap rate is configured
    event SwapRateConfigured(address indexed tokenIn, address indexed tokenOut, uint256 rate);

    /// @dev Mapping to store swap rates between token pairs
    mapping(address tokenIn => mapping(address tokenOut => uint256 rate)) public swapRates;

    /// @dev Mapping to store fixed amount returns regardless of input
    mapping(address tokenIn => mapping(address tokenOut => uint256 fixedReturn)) public fixedReturns;

    /**
     * @notice Configure the swap rate between two tokens
     * @param tokenIn The input token address
     * @param tokenOut The output token address
     * @param rate The rate multiplied by 1e18 (e.g., 2e18 means 2 tokenOut per 1 tokenIn)
     */
    function setSwapRate(address tokenIn, address tokenOut, uint256 rate) external {
        swapRates[tokenIn][tokenOut] = rate;
        emit SwapRateConfigured(tokenIn, tokenOut, rate);
    }

    /**
     * @notice Configure a fixed return amount regardless of the input amount
     * @param tokenIn The input token address
     * @param tokenOut The output token address
     * @param amount The fixed amount to return
     */
    function setFixedReturn(address tokenIn, address tokenOut, uint256 amount) external {
        fixedReturns[tokenIn][tokenOut] = amount;
    }

    /**
     * @notice Swap tokens with a configurable rate
     * @param tokenIn The token to swap from
     * @param tokenOut The token to swap to
     * @param amountIn The amount of tokenIn to swap
     * @param minAmountOut The minimum amount of tokenOut expected
     * @return amountOut The amount of tokenOut received
     */
    function swap(
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        uint256 minAmountOut
    )
        external
        returns (uint256 amountOut)
    {
        // Check if a fixed return has been set
        if (fixedReturns[tokenIn][tokenOut] > 0) {
            amountOut = fixedReturns[tokenIn][tokenOut];
        } else {
            // Calculate the amount out based on the configured rate
            uint256 rate = swapRates[tokenIn][tokenOut];
            if (rate == 0) {
                revert PlugLib.PlugFailed(type(uint8).max, "PlugMockDex: rate not set");
            }

            // Calculate amountOut: amountIn * rate / 1e18
            amountOut = (amountIn * rate) / 1e18;
        }

        // Check slippage
        if (amountOut < minAmountOut) {
            revert PlugLib.PlugFailed(type(uint8).max, "PlugMockDex: insufficient output amount");
        }

        // Transfer tokens (in a real DEX but simulated here)
        if (tokenIn != address(0)) {
            try ERC20(tokenIn).transferFrom(msg.sender, address(this), amountIn) {
                // Successfully transferred tokenIn
            } catch {
                // If we can't transfer, we'll just simulate the swap for testing
            }
        }

        if (tokenOut != address(0)) {
            try ERC20(tokenOut).transfer(msg.sender, amountOut) {
                // Successfully transferred tokenOut
            } catch {
                // If we can't transfer, we'll just simulate the swap for testing
            }
        }

        emit SwapExecuted(msg.sender, tokenIn, tokenOut, amountIn, amountOut);
        return amountOut;
    }

    /**
     * @notice Get quote for a swap without executing it
     * @param tokenIn The token to swap from
     * @param tokenOut The token to swap to
     * @param amountIn The amount of tokenIn to swap
     * @return amountOut The expected amount of tokenOut
     */
    function getAmountOut(
        address tokenIn,
        address tokenOut,
        uint256 amountIn
    )
        external
        view
        returns (uint256 amountOut)
    {
        // Check if a fixed return has been set
        if (fixedReturns[tokenIn][tokenOut] > 0) {
            return fixedReturns[tokenIn][tokenOut];
        }

        // Calculate the amount out based on the configured rate
        uint256 rate = swapRates[tokenIn][tokenOut];
        if (rate == 0) {
            revert PlugLib.PlugFailed(type(uint8).max, "PlugMockDex: rate not set");
        }

        // Calculate amountOut: amountIn * rate / 1e18
        return (amountIn * rate) / 1e18;
    }

    /**
     * @notice Allow the contract to receive ETH for swaps
     */
    receive() external payable { }
}
