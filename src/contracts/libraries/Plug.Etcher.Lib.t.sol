/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../abstracts/test/Plug.Test.sol";
import { PlugEtcherLib } from "./Plug.Etcher.Lib.sol";

contract DeploymentTest is Test {
    function test_PlugBalanceFuseDeployment() public {
        PlugEtcherLib.plugBalanceFuse();
    }

    function test_PlugBalanceSemiFungibleFuseDeployment() public {
        PlugEtcherLib.plugBalanceSemiFungibleFuse();
    }

    function test_PlugBaseFeeFuseDeployment() public {
        PlugEtcherLib.plugBaseFeeFuse();
    }

    function test_PlugBlockNumberFuseDeployment() public {
        PlugEtcherLib.plugBlockNumberFuse();
    }

    function test_PlugCalendarFuseDeployment() public {
        PlugEtcherLib.plugCalendarFuse();
    }

    function test_PlugFactoryDeployment() public {
        PlugEtcherLib.plugFactory();
    }

    function test_PlugFraxlendAPYFuseDeployment() public {
        PlugEtcherLib.plugFraxlendAPYFuse();
    }

    function test_PlugLimitedCallsFuseDeployment() public {
        PlugEtcherLib.plugLimitedCallsFuse();
    }

    function test_PlugNounsBidFuseDeployment() public {
        PlugEtcherLib.plugNounsBidFuse();
    }

    function test_PlugNounsIdFuseDeployment() public {
        PlugEtcherLib.plugNounsIdFuse();
    }

    function test_PlugNounsTraitFuseDeployment() public {
        PlugEtcherLib.plugNounsTraitFuse();
    }

    function test_PlugRevocationFuseDeployment() public {
        PlugEtcherLib.plugRevocationFuse();
    }

    function test_PlugTimestampFuseDeployment() public {
        PlugEtcherLib.plugTimestampFuse();
    }

    function test_PlugTreasuryDeployment() public {
        PlugEtcherLib.plugTreasury();
    }

    function test_PlugVaultSocketDeployment() public {
        PlugEtcherLib.plugVaultSocket();
    }

    function test_PlugDeployment() public {
        PlugEtcherLib.plug();
    }
}
