// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugAddressesLib } from "./Plug.Addresses.Lib.sol";

library PlugLib {
    event SocketDeployed(address indexed implementation, address indexed vault, bytes32 salt);

    event SocketOwnershipTransferred(
        address indexed previousOwner, address indexed newOwner, bytes32 imageHash
    );
    event PlugResult(uint256 index, bytes32 livePlugsHash, PlugTypesLib.Result reason);

    error NotImplemented();

    error SocketAddressInvalid(address $intended, address $socket);
    error SocketAddressEmpty(address $socket);

    error SaltInvalid(address $implementation, address $admin);
    error CallerInvalid(address $expected, address $reality);
    error RouterInvalid(address $reality);
    error TypeInvalid(uint8 $reality);

    error PlugFailed(uint8 $index, string $reason);

    error CompensationFailed(address $recipient, uint256 $value);

    error ThresholdInvalid();
    error ThresholdExceeded(uint256 $expected, uint256 $reality);
    error ThresholdInsufficient(uint256 $expected, uint256 $reality);

    error TargetInvalid();

    error TokenAllowanceInvalid();
    error TokenBalanceInvalid();
}
