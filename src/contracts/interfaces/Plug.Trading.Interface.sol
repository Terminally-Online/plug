// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

interface PlugTradingInterface {
    function transferOwnership(address $owner) external;
    function owner() external view returns (address);
}
