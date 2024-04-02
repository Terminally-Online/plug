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
            PlugEtcherLib.PLUG_BALANCE_FUSE_SALT,
            PlugEtcherLib.PLUG_BALANCE_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BALANCE_SEMI_FUNGIBLE_FUSE_SALT,
            PlugEtcherLib.PLUG_BALANCE_SEMI_FUNGIBLE_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BASE_FEE_FUSE_SALT,
            PlugEtcherLib.PLUG_BASE_FEE_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_BLOCK_NUMBER_FUSE_SALT,
            PlugEtcherLib.PLUG_BLOCK_NUMBER_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_CALENDAR_FUSE_SALT,
            PlugEtcherLib.PLUG_CALENDAR_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_FACTORY_SALT,
            PlugEtcherLib.PLUG_FACTORY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_FRAXLEND_APY_FUSE_SALT,
            PlugEtcherLib.PLUG_FRAXLEND_APY_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_LIMITED_CALLS_FUSE_SALT,
            PlugEtcherLib.PLUG_LIMITED_CALLS_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_BID_FUSE_SALT,
            PlugEtcherLib.PLUG_NOUNS_BID_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_ID_FUSE_SALT,
            PlugEtcherLib.PLUG_NOUNS_ID_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_NOUNS_TRAIT_FUSE_SALT,
            PlugEtcherLib.PLUG_NOUNS_TRAIT_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_REVOCATION_FUSE_SALT,
            PlugEtcherLib.PLUG_REVOCATION_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TIMESTAMP_FUSE_SALT,
            PlugEtcherLib.PLUG_TIMESTAMP_FUSE_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_TREASURY_SALT,
            PlugEtcherLib.PLUG_TREASURY_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_VAULT_SOCKET_SALT,
            PlugEtcherLib.PLUG_VAULT_SOCKET_INITCODE
        );

        PlugEtcherLib.FACTORY.safeCreate2(
            PlugEtcherLib.PLUG_SALT, PlugEtcherLib.PLUG_INITCODE
        );

        vm.stopBroadcast();
    }
}
