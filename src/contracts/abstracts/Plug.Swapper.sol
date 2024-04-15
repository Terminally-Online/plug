// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { ReentrancyGuard } from "solady/utils/ReentrancyGuard.sol";
import { PlugLib } from "../libraries/Plug.Lib.sol";
import { PlugBalanceInterface } from "../interfaces/Plug.Balance.Interface.sol";

import { SafeTransferLib } from "solady/utils/SafeTransferLib.sol";

/**
 * @title Plug Swapper
 * @notice The Swapper enables secure execution of transactions that consume some
 *         form of value, often in return for another form of value. As the primary
 *         top-level target, the Swapper is responsible for ensuring that Plug
 *         has a secure mechanism to collect a fee upon successful execution.
 */
abstract contract PlugSwapper is ReentrancyGuard {
    /// @dev Storage reference to track which contracts can be interacted with through
    ///      the Swapper. Without this, a malicious call could be made to directly call the
    ///      assets held within and transfer them out having bypassed any swap mechanism.
    mapping(address => bool) public targetToAllowed;

    /// @dev Modifier to confirm that all transactions are only being executed on
    ///      active targets that are allowed to be interacted with.
    modifier onlyActiveTarget(address $target) {
        /// @dev If the target is not on the allowance list then revert as it
        ///      is potentially an unsafe call.
        if (targetToAllowed[$target] == false) {
            revert PlugLib.TargetInvalid();
        }
        _;
    }

    /**
     * @notice Run a transaction that consumes Native tokens and collect a fee in the
     *         Native token that is being transacted with as a form of value.
     * @param $target The address of the contract to call.
     * @param $data The data to send to the target.
     * @param $fee The fee to charge in native token amount.
     */
    function plugNative(
        address payable $target,
        bytes calldata $data,
        uint256 $fee
    )
        public
        payable
        virtual
        nonReentrant
        onlyActiveTarget($target)
    {
        /// @dev Calculate the amount of the native token that is available to use
        ///      when swapping the token.
        uint256 outNative = msg.value - $fee;

        /// @dev Submit the call that is going to consume the Native currency
        ///      and return a non-enforced amount of Token B using the amount
        ///      of the Native currency supplied (without fees) included which
        ///      results in the fee in Native currency remaining in the Swapper
        ///      for future withdrawal.
        (bool success, bytes memory reason) = $target.call{ value: outNative }($data);
        /// @dev If the call was not successful, bubble up the revert.
        PlugLib.bubbleRevert(success, reason);
    }

    /**
     * @notice Run a transaction that consumes ERC20 tokens and collect a fee in the
     *         ERC20 token that is being transacted with as a form of value.
     * @param $tokenOut The address of the token being swapped.
     * @param $target The address of the contract to call.
     * @param $data The data to send to the target.
     * @param $sell The amount of tokens to sell.
     * @param $fee The fee to charge in $tokenOut amount.
     */
    function plugToken(
        address $tokenOut,
        address payable $target,
        bytes calldata $data,
        uint256 $sell,
        uint256 $fee
    )
        public
        payable
        virtual
        nonReentrant
        onlyActiveTarget($target)
    {
        /// @dev Retrieve the tokens being swapped from the transaction caller (a Socket)
        ///      and transfer them to this contract.
        SafeTransferLib.safeTransferFrom($tokenOut, msg.sender, address(this), $sell);

        /// @dev Give the target contract allowance to move the token being swapped up to the
        ///      amount that the `caller` can move after the fee is accounted for.
        SafeTransferLib.safeApprove($tokenOut, $target, $sell - $fee);

        /// @dev Submit the call that is going to consume Token B using the amount of the Native
        ///      currency supplied (without fees) included which results in the fee in Native
        ///      currency remaining in the Swapper for future withdrawal.
        (bool success, bytes memory reason) = $target.call{ value: msg.value }($data);
        /// @dev If the call was not successful, bubble up the revert. Otherwise,
        ///      continue on with the execution.
        PlugLib.bubbleRevert(success, reason);

        /// @dev If the target did not use all of the tokens within allowance, revert
        ///      as there was an issue and/or imbalance somewhere in the execution.
        if (PlugBalanceInterface($tokenOut).allowance(address(this), $target) != 0) {
            revert PlugLib.TokenAllowanceInvalid();
        }
    }

    /**
     * @notice Run a transaction that consumes Native tokens (Token A) and returns
     *         ERC20 tokens (Token B) and collects fees in the native token (Token A)
     *         that are left sitting in the Swapper.
     * @param $tokenIn The address of the token being received.
     * @param $target The address of the contract to call.
     * @param $data The data to send to the target.
     * @param $fee The fee to charge in native token amount.
     */
    function plugNativeToToken(
        address $tokenIn,
        address payable $target,
        bytes calldata $data,
        uint256 $fee
    )
        public
        payable
        virtual
        nonReentrant
        onlyActiveTarget($target)
    {
        /// @dev Take a snapshot of how much of the Native currency (Token A) was
        ///      within the contract before this transaction began.
        uint256 preNative = address(this).balance - msg.value;
        /// @dev Calculate the amount of the native token that is available to use
        ///      when swapping the token.
        uint256 outNative = msg.value - $fee;

        /// @dev Create the connected Token B interface for balance retrieval.
        PlugBalanceInterface tokenIn = PlugBalanceInterface($tokenIn);

        /// @dev Take a snapshot of how much of Token B being swapped for was within
        ///      the contract before the transaction is fulfilled.
        uint256 preTokenIn = tokenIn.balanceOf(address(this));

        /// @dev Submit the call that is going to return a non-enforced amount of Token
        ///      B using the amount of the Native currency supplied (without fees)
        ///      included which results in the fee in Native currency remaining
        ///      in the Swapper for future withdrawal.
        (bool success, bytes memory reason) = $target.call{ value: outNative }($data);
        /// @dev If the call was not successful, bubble up the revert. Otherwise,
        ///      continue on with the execution.
        PlugLib.bubbleRevert(success, reason);

        /// @dev Retrieve the amount of Token B this contract now holds.
        uint256 postTokenIn = tokenIn.balanceOf(address(this));

        /// @dev Confirm Token B increased. If the balance is smaller
        ///      or the same as before, revert due to the lack of receipt.
        if (preTokenIn >= postTokenIn) {
            revert PlugLib.TokenBalanceInvalid();
        }

        /// @dev Deliver the acquired Token B to the `caller` (the Socket).
        SafeTransferLib.safeTransfer($tokenIn, msg.sender, postTokenIn - preTokenIn);

        /// @dev Retrieve how many Native tokens this contract now holds
        ///      while having accounted for the fee.
        uint256 postNative = address(this).balance - $fee;

        /// @dev Deliver any surplus of the Native token back to the `caller`.
        if (postNative > preNative) {
            SafeTransferLib.safeTransferETH(msg.sender, postNative - preNative);
        }
    }

    /**
     * @notice Run a transaction that consumes ERC20 tokens (Token A)
     *         and returns ERC20 tokens (Token B).
     * @param $tokenOut The address of the token leaving the callers account.
     * @param $tokenIn The address of the token being received.
     * @param $target The address of the contract to call.
     * @param $data The data to send to the target.
     * @param $sell The amount of tokens to sell.
     * @param $fee The fee to charge in $tokenOut amount.
     */
    function plugTokenToToken(
        address $tokenOut,
        address $tokenIn,
        address payable $target,
        bytes calldata $data,
        uint256 $sell,
        uint256 $fee
    )
        public
        payable
        virtual
        nonReentrant
        onlyActiveTarget($target)
    {
        /// @dev Create the connected Token B interface for balance retrieval.
        PlugBalanceInterface tokenIn = PlugBalanceInterface($tokenIn);

        /// @dev Take a snapshot of how much of Token B being swapped for was within
        ///      the contract before the transaction is fulfilled.
        uint256 preTokenIn = tokenIn.balanceOf(address(this));

        /// @dev Retrieve the tokens being spent from the transaction caller
        ///      (a Socket) and transfer them to this contract so that they can
        ///      be routed to the target destination.
        SafeTransferLib.safeTransferFrom($tokenOut, msg.sender, address(this), $sell);

        /// @dev Give the target contract allowance to move the token being swapped up to the
        ///      amount that the `caller` can move after the fee is accounted for.
        SafeTransferLib.safeApprove($tokenOut, $target, $sell - $fee);

        /// @dev Submit the call that is going to return a non-enforced amount of Token
        ///      B using the amount of the Native currency supplied (without fees)
        ///      included which results in the fee in Native currency remaining
        ///      in the Swapper for future withdrawal.
        (bool success, bytes memory reason) = $target.call{ value: msg.value }($data);
        /// @dev If the call was not successful, bubble up the revert. Otherwise,
        ///      continue on with the execution.
        PlugLib.bubbleRevert(success, reason);

        /// @dev If the target did not use all of the tokens within allowance, revert
        ///      as there was an issue and/or imbalance somewhere in the execution.
        if (PlugBalanceInterface($tokenOut).allowance(address(this), $target) != 0) {
            revert PlugLib.TokenAllowanceInvalid();
        }

        /// @dev Retrieve the amount of Token B this contract now holds.
        uint256 postTokenIn = tokenIn.balanceOf(address(this));

        /// @dev Confirm Token B increased. If the balance is smaller
        ///      or the same as before, revert due to the lack of receipt.
        if (preTokenIn >= postTokenIn) {
            revert PlugLib.TokenBalanceInvalid();
        }

        /// @dev Deliver the swapped Token B to the `caller`.
        SafeTransferLib.safeTransfer($tokenIn, msg.sender, postTokenIn - preTokenIn);
    }

    /**
     * @notice Run a transaction that consumes ERC20s and returns native tokens.
     * @param $tokenOut The address of the token being swapped.
     * @param $target The address of the contract to call.
     * @param $data The data to send to the target.
     * @param $sell The amount of tokens to sell.
     * @param $fee The fee to charge in basis points on native token.
     */
    function plugTokenToNative(
        address $tokenOut,
        address payable $target,
        bytes calldata $data,
        uint256 $sell,
        uint256 $fee
    )
        public
        payable
        virtual
        nonReentrant
        onlyActiveTarget($target)
    {
        /// @dev Take a snapshot of how much of the Native currency (Token B) was
        ///      within the contract before this transaction began.
        uint256 preNative = address(this).balance - msg.value;

        /// @dev Retrieve the tokens being spent from the transaction caller
        ///      (a Socket) and transfer them to this contract so that they can
        ///      be routed to the target destination.
        SafeTransferLib.safeTransferFrom($tokenOut, msg.sender, address(this), $sell);

        /// @dev Give the target contract allowance to move the token being swapped up to the
        ///      amount that the `caller` can move after the fee is accounted for.
        SafeTransferLib.safeApprove($tokenOut, $target, $sell);

        /// @dev Submit the call that is going to return a non-enforced amount of Token
        ///      B using the amount of the Native currency included which results in the fee
        ///      in Native currency remaining in the Swapper for future withdrawal.
        (bool success, bytes memory reason) = $target.call{ value: msg.value }($data);
        /// @dev If the call was not successful, bubble up the revert. Otherwise,
        ///      continue on with the execution.
        PlugLib.bubbleRevert(success, reason);

        /// @dev If the target did not use all of the tokens within allowance, revert
        ///      as there was an issue and/or imbalance somewhere in the execution.
        if (PlugBalanceInterface($tokenOut).allowance(address(this), $target) != 0) {
            revert PlugLib.TokenAllowanceInvalid();
        }

        /// @dev Retrieve the amount of Token B this contract now holds.
        uint256 postNative = address(this).balance;

        /// @dev Confirm Token B increased. If the balance is smaller
        ///      or the same as before, revert due to the lack of receipt.
        if (postNative <= preNative) {
            revert PlugLib.TokenBalanceInvalid();
        }

        /// @dev Determine the increase in native tokens.
        uint256 diffNative = postNative - preNative;

        /// @dev Handle the fee if one is present.
        if ($fee > 0) {
            /// @dev Transfer the native tokens earned excluding the fee to the `caller`.
            SafeTransferLib.safeTransferETH(msg.sender, diffNative - (diffNative * $fee) / 10 ** 18);
        }
        /// @dev If there is no fee, return the entire amount to the sender.
        else if (diffNative > 0) {
            /// @dev Transfer the native tokens earned to the sender.
            SafeTransferLib.safeTransferETH(msg.sender, diffNative);
        }
    }
}
