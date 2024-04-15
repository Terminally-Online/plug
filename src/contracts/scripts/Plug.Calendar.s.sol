/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * @title Plug.Calendar Deployment
 * @dev Deploy a Plug.Calendar to a new chain using the immutable
 *      Create2 factory for constant addresses across all major EVM chains.
 * @notice To deploy the most up to date version of Plug.Calendar, you can always just run
 *         this script and everything will be deployed as configured.
 */
contract PlugCalendarDeployment is Script {
    function run() external {
        vm.startBroadcast();

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_CALENDAR_SALT, PlugEtcherLib.PLUG_CALENDAR_INITCODE
        );

        vm.stopBroadcast();
    }
}
