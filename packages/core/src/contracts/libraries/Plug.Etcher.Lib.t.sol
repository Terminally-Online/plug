/// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { Test } from "../abstracts/test/Plug.Test.sol";
import { PlugEtcherLib } from "./Plug.Etcher.Lib.sol";

contract DeploymentTest is Test {
    function test_PlugAssertDeployment() public {
        PlugEtcherLib.plugAssert();
    }

    function test_PlugBooleanDeployment() public {
        PlugEtcherLib.plugBoolean();
    }

    function test_PlugCoercionDeployment() public {
        PlugEtcherLib.plugCoercion();
    }

    function test_PlugDatabaseDeployment() public {
        PlugEtcherLib.plugDatabase();
    }

    function test_PlugEVMDeployment() public {
        PlugEtcherLib.plugEVM();
    }

    function test_PlugFactoryDeployment() public {
        PlugEtcherLib.plugFactory();
    }

    function test_PlugMathDeployment() public {
        PlugEtcherLib.plugMath();
    }

    function test_PlugRewardsDeployment() public {
        PlugEtcherLib.plugRewards();
    }

    function test_PlugSocketDeployment() public {
        PlugEtcherLib.plugSocket();
    }

    function test_PlugTicketDeployment() public {
        PlugEtcherLib.plugTicket();
    }

    function test_PlugTokenDeployment() public {
        PlugEtcherLib.plugToken();
    }

    function test_PlugDeployment() public {
        PlugEtcherLib.plug();
    }
}
