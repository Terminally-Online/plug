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

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_FACTORY_SALT, PlugEtcherLib.PLUG_FACTORY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_SOCKET_SALT, PlugEtcherLib.PLUG_SOCKET_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TREASURY_SALT, PlugEtcherLib.PLUG_TREASURY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(PlugEtcherLib.PLUG_SALT, PlugEtcherLib.PLUG_INITCODE);

        vm.stopBroadcast();
    }
}
