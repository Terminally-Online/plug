/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * @title Plug.Nouns.Id Deployment
 * @dev Deploy a Plug.Nouns.Id to a new chain using the immutable
 *      Create2 factory for constant addresses across all major EVM chains.
 * @notice To deploy the most up to date version of Plug.Nouns.Id, you can always just run
 *         this script and everything will be deployed as configured.
 */
contract PlugNounsIdDeployment is Script {
    function run() external {
        vm.startBroadcast();

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_ID_SALT, PlugEtcherLib.PLUG_NOUNS_ID_INITCODE
        );

        vm.stopBroadcast();
    }
}
