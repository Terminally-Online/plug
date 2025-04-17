/// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { Script } from "forge-std/Script.sol";
import { Plug } from "../base/Plug.sol";
import { PlugEtcherLib } from "../libraries/Plug.Etcher.Lib.sol";

/**
 * This is the script used to deploy the entire Plug protocol stack at once. Each time upon
 * build the active contracts are automatically populated so that you never have to deal
 * with manual coordination of your contracts or updating the deployment scripts.
 */
contract PlugDeployment is Script {
    function run() external {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(privateKey);

        if (PlugEtcherLib.PLUG_ASSERT_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_ASSERT_SALT, PlugEtcherLib.PLUG_ASSERT_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_BOOLEAN_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_BOOLEAN_SALT, PlugEtcherLib.PLUG_BOOLEAN_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_COERCION_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_COERCION_SALT, PlugEtcherLib.PLUG_COERCION_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_DATABASE_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_DATABASE_SALT, PlugEtcherLib.PLUG_DATABASE_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_EVM_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_EVM_SALT, PlugEtcherLib.PLUG_EVM_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_FACTORY_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_FACTORY_SALT,
                PlugEtcherLib.PLUG_FACTORY_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_MATH_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_MATH_SALT, PlugEtcherLib.PLUG_MATH_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_REWARDS_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_REWARDS_SALT, PlugEtcherLib.PLUG_REWARDS_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_SOCKET_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_SOCKET_SALT,
                PlugEtcherLib.PLUG_SOCKET_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_TICKET_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_TICKET_SALT, PlugEtcherLib.PLUG_TICKET_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_TOKEN_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_TOKEN_SALT, PlugEtcherLib.PLUG_TOKEN_INITCODE
            );
        }

        if (PlugEtcherLib.PLUG_ADDRESS.code.length == 0) {
            PlugEtcherLib.FACTORY.safeCreate2(
                PlugEtcherLib.PLUG_SALT, PlugEtcherLib.PLUG_INITCODE
            );
        }

        vm.stopBroadcast();
    }
}
