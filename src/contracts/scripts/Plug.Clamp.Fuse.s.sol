/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * @title Plug.Clamp.Fuse Deployment
 * @dev Deploy a Plug.Clamp.Fuse to a new chain using the immutable
 *      Create2 factory for constant addresses across all major EVM chains.
 * @notice To deploy the most up to date version of Plug.Clamp.Fuse, you can always just run
 *         this script and everything will be deployed as configured.
 */
contract PlugClampFuseDeployment is Script {
    function run() external {
        vm.startBroadcast();

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_CLAMP_FUSE_SALT,
            PlugEtcherLib.PLUG_CLAMP_FUSE_INITCODE
        );

        vm.stopBroadcast();
    }
}
