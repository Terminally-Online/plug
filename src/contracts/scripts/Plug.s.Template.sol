/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * INSERT DOCUMENTATION
 */
contract PlugDeployment is Script {
    function run() external {
        vm.startBroadcast();

        /// @auto INSERT SEGMENTS

        vm.stopBroadcast();
    }
}
