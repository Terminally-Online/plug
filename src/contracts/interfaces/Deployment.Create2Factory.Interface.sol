// SPDX-License-Identifier: MIT

pragma solidity ^0.8.24;

interface ImmutableCreate2Factory {
    function safeCreate2(
        bytes32 salt,
        bytes calldata initCode
    )
        external
        payable
        returns (address deploymentAddress);

    function findCreate2Address(
        bytes32 salt,
        bytes calldata initCode
    )
        external
        view
        returns (address deploymentAddress);

    function findCreate2AddressViaHash(
        bytes32 salt,
        bytes32 initCodeHash
    )
        external
        view
        returns (address deploymentAddress);
}
