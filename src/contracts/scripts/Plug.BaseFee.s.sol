/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * @title Plug.BaseFee Deployment
 * @dev Deploy a Plug.BaseFee to a new chain using the immutable
 *      Create2 factory for constant addresses across all major EVM chains.
 * @notice To deploy the most up to date version of Plug.BaseFee, you can always just run
 *         this script and everything will be deployed as configured.
 */
contract PlugBaseFeeDeployment is Script {
    function run() external {
        vm.startBroadcast();

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BASE_FEE_SALT, PlugEtcherLib.PLUG_BASE_FEE_INITCODE
        );

        vm.stopBroadcast();
    }
}
