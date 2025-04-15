// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { PlugTypesLib } from "../abstracts/Plug.Types.sol";
import { PlugAddressesLib } from "./Plug.Addresses.Lib.sol";

library PlugLib {
    /////////////////////////////////////////////////
    //                     PLUG                    //
    /////////////////////////////////////////////////
    event PlugResult(
        uint8 index, bytes32 plugsHash, PlugTypesLib.Result reason
    );

    error PlugFailed(uint256 $index, string $reason);

    /////////////////////////////////////////////////
    //                    SOCKET                   //
    /////////////////////////////////////////////////

    event SocketDeployed(
        address indexed implementation, address indexed vault, bytes32 salt
    );
    event SocketOwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner,
        bytes32 imageHash
    );

    error SocketAddressInvalid(address $intended, address $socket);
    error SocketAddressEmpty(address $socket);

    error SaltInvalid(address $implementation, address $admin);
    error CallerInvalid(address $expected, address $reality);
    error RouterInvalid(address $reality);
    error TypeInvalid(uint8 $reality);
    error CompensationFailed(address $recipient, uint256 $value);

    string internal constant PlugCoreSignatureInvalid =
        "PlugCore:signature-invalid";
    string internal constant PlugCoreSenderInvalid = "PlugCore:sender-invalid";
    string internal constant PlugCoreSolverMalformed =
        "PlugCore:solver-malformed";
    string internal constant PlugCoreSolverExpired = "PlugCore:solver-expired";
    string internal constant PlugCoreSolverInvalid = "PlugCore:solver-invalid";
    string internal constant PlugCoreNonceInvalid = "PlugCore:nonce-invalid";
    string internal constant PlugCorePlugFailed = "PlugCore:plug-failed";
    string internal constant PlugCoreOutOfBounds = "PlugCore:out-of-bounds";
    string internal constant PlugCoreWouldOverflow = "PlugCore:would-overflow";
    string internal constant PlugCoreInvalidOffset = "PlugCore:invalid-offset";
    string internal constant PlugCoreInvalidLength = "PlugCore:invalid-length";
    string internal constant PlugCoreArrayLengthInvalid =
        "PlugCore:array-length-invalid";
    string internal constant PlugCoreStructTooSmall =
        "PlugCore:struct-too-small";
    string internal constant PlugCoreKeyValueTooSmall =
        "PlugCore:key-value-too-small";
    string internal constant PlugCoreTypeInvalid = "PlugCore:type-invalid";

    /////////////////////////////////////////////////
    //                   REWARDS                   //
    /////////////////////////////////////////////////

    event NewRewardPeriod(
        uint256 indexed period, bytes32 merkleRoot, uint256 totalAmount
    );
    event RewardClaimed(
        uint256 indexed period, address indexed user, uint256 amount
    );

    error InvalidMerkleProof();
    error PeriodNotInitialized();
    error RewardsAlreadyClaimed();
    error InsufficientRewardBalance();
    error ZeroAmount();
}
