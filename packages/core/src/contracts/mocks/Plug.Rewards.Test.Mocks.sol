// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { ERC20 } from "solady/tokens/ERC20.sol";
import { SafeTransferLib } from "solady/utils/SafeTransferLib.sol";

// Mocks needed for testing PlugToken with PlugRewards

// Mock SuperchainERC20 implementation
abstract contract SuperchainERC20 is ERC20 {
    function crosschainMint(address $to, uint256 $amount) external returns (bool) {
        _mint($to, $amount);
        return true;
    }

    function crosschainBurn(address $from, uint256 $amount) external returns (bool) {
        _burn($from, $amount);
        return true;
    }
}

// Mock PredeployAddresses
library PredeployAddresses {
    address public constant SUPERCHAIN_TOKEN_BRIDGE =
        address(0x9876543210987654321098765432109876543210);
}
