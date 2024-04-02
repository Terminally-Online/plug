// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

interface PlugBalanceInterface {
    function balanceOf(address $holder)
        external
        view
        returns (uint256 $balance);

    function balanceOf(
        address $holder,
        uint256 $tokenId
    )
        external
        view
        returns (uint256);

    function allowance(
        address $holder,
        address $spender
    )
        external
        view
        returns (uint256);
}
