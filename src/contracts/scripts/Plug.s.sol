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
            PlugEtcherLib.PLUG_BALANCE_SEMI_FUNGIBLE_SALT,
            PlugEtcherLib.PLUG_BALANCE_SEMI_FUNGIBLE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BALANCE_SALT, PlugEtcherLib.PLUG_BALANCE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BASE_FEE_SALT, PlugEtcherLib.PLUG_BASE_FEE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BLOCK_NUMBER_SALT, PlugEtcherLib.PLUG_BLOCK_NUMBER_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_CALENDAR_SALT, PlugEtcherLib.PLUG_CALENDAR_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_FACTORY_SALT, PlugEtcherLib.PLUG_FACTORY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_FRAXLEND_APY_SALT, PlugEtcherLib.PLUG_FRAXLEND_APY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_LIMITED_CALLS_SALT, PlugEtcherLib.PLUG_LIMITED_CALLS_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_BID_SALT, PlugEtcherLib.PLUG_NOUNS_BID_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_ID_SALT, PlugEtcherLib.PLUG_NOUNS_ID_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_TRAIT_SALT, PlugEtcherLib.PLUG_NOUNS_TRAIT_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_REVOCATION_SALT, PlugEtcherLib.PLUG_REVOCATION_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TIMESTAMP_SALT, PlugEtcherLib.PLUG_TIMESTAMP_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TREASURY_SALT, PlugEtcherLib.PLUG_TREASURY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_VAULT_SOCKET_SALT, PlugEtcherLib.PLUG_VAULT_SOCKET_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(PlugEtcherLib.PLUG_SALT, PlugEtcherLib.PLUG_INITCODE);

        vm.stopBroadcast();
    }
}
