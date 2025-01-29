// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { PlugLib } from "../libraries/Plug.Lib.sol";

/**
 * @title Plug Mock Echo
 * @notice A mock contract for testing the Plug framework.
 * @dev This contract is for testing purposes only.
 */
contract PlugMockEcho {
    /// @dev Event emitted when an Echo is invoked.
    event EchoInvoked(address $sender, string $message);

    /**
     * @notice A mock function for testing the framework.
     * @param $message The message to echo.
     */
    function echo(string memory $message) external {
        emit EchoInvoked(msg.sender, $message);
    }

    function emptyEcho() external {
        emit EchoInvoked(msg.sender, "Hello World");
    }

    function revertEcho() external pure {
        revert PlugLib.PlugFailed(type(uint8).max, "PlugMockEcho:revertEcho");
    }

    /**
     * @notice A mock function for testing the framework.
     */
    function mutedEcho(uint256 $echo) external pure returns (uint256 $slot) {
        if ($echo % 8 == 0) {
            $slot = 1;
        }
        $slot = 2;
    }
}
