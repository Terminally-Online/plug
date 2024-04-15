/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../abstracts/test/Plug.Test.sol";
import { PlugEtcherLib } from "./Plug.Etcher.Lib.sol";

contract DeploymentTest is Test {
    function test_PlugBalanceSemiFungibleDeployment() public {
        PlugEtcherLib.plugBalanceSemiFungible();
    }

    function test_PlugBalanceDeployment() public {
        PlugEtcherLib.plugBalance();
    }

    function test_PlugBaseFeeDeployment() public {
        PlugEtcherLib.plugBaseFee();
    }

    function test_PlugBlockNumberDeployment() public {
        PlugEtcherLib.plugBlockNumber();
    }

    function test_PlugCalendarDeployment() public {
        PlugEtcherLib.plugCalendar();
    }

    function test_PlugFactoryDeployment() public {
        PlugEtcherLib.plugFactory();
    }

    function test_PlugFraxlendAPYDeployment() public {
        PlugEtcherLib.plugFraxlendAPY();
    }

    function test_PlugLimitedCallsDeployment() public {
        PlugEtcherLib.plugLimitedCalls();
    }

    function test_PlugNounsBidDeployment() public {
        PlugEtcherLib.plugNounsBid();
    }

    function test_PlugNounsIdDeployment() public {
        PlugEtcherLib.plugNounsId();
    }

    function test_PlugNounsTraitDeployment() public {
        PlugEtcherLib.plugNounsTrait();
    }

    function test_PlugRevocationDeployment() public {
        PlugEtcherLib.plugRevocation();
    }

    function test_PlugTimestampDeployment() public {
        PlugEtcherLib.plugTimestamp();
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
