// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

interface PlugTradingInterface {
    function transferOwnership(address $owner) external;
    function owner() external view returns (address);
}
