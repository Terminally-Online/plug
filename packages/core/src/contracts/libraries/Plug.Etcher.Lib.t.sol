/// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import { Test } from "../abstracts/test/Plug.Test.sol";
import { PlugEtcherLib } from "./Plug.Etcher.Lib.sol";

contract DeploymentTest is Test {
    function test_PlugFactoryDeployment() public {
        PlugEtcherLib.plugFactory();
    }

    function test_PlugSocketDeployment() public {
        PlugEtcherLib.plugSocket();
    }

    function test_PlugTreasuryDeployment() public {
        PlugEtcherLib.plugTreasury();
    }

    function test_PlugDeployment() public {
        PlugEtcherLib.plug();
    }
}
