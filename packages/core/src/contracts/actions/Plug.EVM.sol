// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.23;

import { SafeTransferLib } from "solady/utils/SafeTransferLib.sol";

/**
 * @title Plug EVM
 * @notice A collection of EVM-specific utilities for interacting with
 *         blockchain data and functions.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugEVM {
    /**
     * @notice Get the native balance of an address
     * @param account The address to check balance for
     * @return result The native token balance in wei
     */
    function balanceOf(address account) public view returns (uint256 result) {
        return account.balance;
    }

    /**
     * @notice Get the current block timestamp
     * @return result The current block timestamp as a uint256
     */
    function getTimestamp() public view returns (uint256 result) {
        return block.timestamp;
    }

    /**
     * @notice Get the current block number
     * @return result The current block number as a uint256
     */
    function getBlockNumber() public view returns (uint256 result) {
        return block.number;
    }

    /**
     * @notice Get the current block hash
     * @return result The current block hash
     */
    function getBlockHash() public view returns (bytes32 result) {
        return blockhash(block.number - 1);
    }

    /**
     * @notice Get a specific previous block hash (up to 256 blocks in the past)
     * @param blockNumber The block number to get the hash for
     * @return result The requested block hash
     */
    function getBlockHash(uint256 blockNumber) public view returns (bytes32 result) {
        require(
            blockNumber < block.number && blockNumber >= block.number - 256,
            "PlugEVM:block-hash-unavailable"
        );
        return blockhash(blockNumber);
    }

    /**
     * @notice Safely transfer native tokens to a recipient using Solady's audited SafeTransferLib
     * @param recipient The address to receive the tokens
     * @param amount The amount to send in wei
     */
    function transfer(address recipient, uint256 amount) public payable {
        require(address(this).balance >= amount, "PlugEVM:insufficient-balance");

        SafeTransferLib.safeTransferETH(recipient, amount);
    }

    /**
     * @notice Safely transfer native tokens to a recipient with a gas limit
     * @dev Uses Solady's forceSafeTransferETH which handles gas stipend internally
     * @param recipient The address to receive the tokens
     * @param amount The amount to send in wei
     * @param gasLimit The gas stipend for the transfer
     * @return success Always returns true as Solady will revert if the transfer fails
     * @return returnData Empty bytes as Solady handles error internally
     */
    function transfer(
        address recipient,
        uint256 amount,
        uint256 gasLimit
    )
        public
        payable
        returns (bool success, bytes memory returnData)
    {
        require(address(this).balance >= amount, "PlugEVM:insufficient-balance");

        SafeTransferLib.forceSafeTransferETH(recipient, amount, gasLimit);

        return (true, new bytes(0));
    }

    /**
     * @notice Get the current chain ID
     * @return result The current chain ID
     */
    function getChainId() public view returns (uint256 result) {
        return block.chainid;
    }

    /**
     * @notice Get the address that is currently executing this code
     * @return result The current execution address
     */
    function getCaller() public view returns (address result) {
        return msg.sender;
    }

    /**
     * @notice Get the gas price of the current transaction
     * @return result The gas price in wei
     */
    function getGasPrice() public view returns (uint256 result) {
        return tx.gasprice;
    }

    /**
     * @notice Get the gas limit of the current block
     * @return result The block gas limit
     */
    function getBlockGasLimit() public view returns (uint256 result) {
        return block.gaslimit;
    }

    /**
     * @notice Get the remaining gas in the current execution
     * @return result The remaining gas
     */
    function getRemainingGas() public view returns (uint256 result) {
        return gasleft();
    }

    /**
     * @notice Check if an address is a contract
     * @param account The address to check
     * @return result True if the address is a contract, false if it is an EOA
     */
    function isContract(address account) public view returns (bool result) {
        uint256 size;
        assembly {
            size := extcodesize(account)
        }
        return size > 0;
    }

    /**
     * @notice Get the code hash of an address
     * @param account The address to check
     * @return result The code hash of the address
     */
    function getCodeHash(address account) public view returns (bytes32 result) {
        return account.codehash;
    }

    /**
     * @notice Receive function to allow the contract to receive ETH
     */
    receive() external payable { }
}
