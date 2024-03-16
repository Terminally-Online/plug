/// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * @title Plug.Timestamp.Fuse Deployment
 * @dev Deploy a Plug.Timestamp.Fuse to a new chain using the immutable
 *      Create2 factory for constant addresses across all major EVM chains.
 * @notice To deploy the most up to date version of Plug.Timestamp.Fuse, you can always just run
 *         this script and everything will be deployed as configured.
 */
contract PlugTimestampFuseDeployment is Script {
    function run() external {
        vm.startBroadcast();

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TIMESTAMP_FUSE_SALT,
            PlugEtcherLib.PLUG_TIMESTAMP_FUSE_INITCODE
        );

        vm.stopBroadcast();
    }
}
