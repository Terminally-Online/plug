// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

import { PRBTest } from "@prb/test/PRBTest.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { TestPlus } from "./TestPlus.sol";
import { TestPlug } from "./TestPlug.sol";

abstract contract Test is PRBTest, StdCheats, TestPlus, TestPlug { }
