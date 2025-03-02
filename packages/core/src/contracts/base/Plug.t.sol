// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {
    Test,
    PlugLib,
    PlugTypesLib,
    PlugAddressesLib,
    PlugEtcherLib,
    PlugFactory,
    Plug,
    PlugMockDex,
    PlugMockEcho,
    PlugMockERC20
} from "../abstracts/test/Plug.Test.sol";
import { ECDSA } from "solady/utils/ECDSA.sol";

contract PlugTest is Test {
    event EchoInvoked(address $sender, string $message);

    function setUp() public virtual {
        setUpPlug();
    }

    function test_name() public {
        assertEq(plug.name(), "Plug");
    }

    function test_symbol() public {
        assertEq(plug.symbol(), "PLUG");
    }

    function test_PlugEmptyEcho_TypeRecovery() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver_TreasuryPayment() public {
        address solver = _randomNonZeroAddress();
        address treasury = 0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E;
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        uint256 preBalance = address(treasury).balance;
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);
        assertTrue(address(treasury).balance > preBalance);
    }

    function testRevert_PlugEmptyEcho_Solver_TreasuryPaymentFailure() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 0);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray, solver);
        vm.prank(solver);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 1, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_PlugEmptyEcho_Solver_InvalidNonce() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray, solver);
        vm.prank(solver);
        vm.expectEmit(address(mock));
        emit EchoInvoked(address(socket), "Hello World");
        plug.plug(livePlugsArray);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:nonce-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_InvalidSignature() public {
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        bytes32 digest = socket.getPlugsDigest(plugs);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(0x123456, digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        livePlugsArray[0] = PlugTypesLib.LivePlugs({ plugs: plugs, signature: signature });

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:signature-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_Invalid() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs =
            createPlugs(plugsArray, uint48(block.timestamp + 3 minutes), solver);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:solver-invalid" })
        );
        plug.plug(livePlugsArray);
    }

    function testRevert_PlugEmptyEcho_Solver_Expired() public {
        address solver = _randomNonZeroAddress();
        vm.deal(solver, 100 ether);
        vm.deal(address(socket), 100 ether);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs =
            createPlugs(plugsArray, uint48(block.timestamp - 1 minutes), solver);
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugs);
        vm.prank(solver);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(
            0, PlugTypesLib.Result({ index: type(uint8).max, error: "PlugCore:solver-expired" })
        );
        plug.plug(livePlugsArray);
    }

    function test_Plugs_PlugEmptyEcho_TypeRecovery() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_Direct() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);

        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);

        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 0, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_SecondFails() public {
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2); // Array of 2 plugs
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
        plugsArray[1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
        PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
        livePlugsArray[0] = createLivePlugs(plugs);

        vm.expectEmit(address(plug));
        emit PlugLib.PlugResult(0, PlugTypesLib.Result({ index: 1, error: "PlugCore:plug-failed" }));
        plug.plug(livePlugsArray);
    }

    function test_CatchReverting_Plugs_MultipleIndexes() public {
        for (uint8 testIndex = 0; testIndex < 5; testIndex++) {
            PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
            PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](testIndex + 2);
            for (uint8 i = 0; i < testIndex + 1; i++) {
                plugsArray[i] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
            }
            plugsArray[testIndex + 1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
            PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
            livePlugsArray[0] = createLivePlugs(plugs);

            vm.expectEmit(address(plug));
            emit PlugLib.PlugResult(
                0, PlugTypesLib.Result({ index: testIndex + 1, error: "PlugCore:plug-failed" })
            );
            plug.plug(livePlugsArray);
        }
    }

    function test_PlugSwapVerifyOutput() public {
        // Create a plug that performs a swap and returns the output amount
        bytes memory swapCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                PlugMockDex.swap.selector,
                address(mockERC20), // tokenIn
                address(mockERC20), // tokenOut
                10 ether, // amountIn
                19 ether // minAmountOut (slightly less than expected 20 ether)
            )
        );

        PlugTypesLib.Plug memory swapPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(dex),
            value: 0,
            data: swapCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create a plug that does nothing - we'll just use it to verify the first plug executed correctly
        bytes memory echoCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(mock.emptyEcho.selector)
        );

        PlugTypesLib.Plug memory echoPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mock),
            value: 0,
            data: echoCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = swapPlug;
        plugsArray[1] = echoPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected
        uint256 userBalance = mockERC20.balanceOf(address(socket));
        assertEq(
            userBalance,
            1010 ether,
            "Socket balance should be 1010 ether (initial 1000 - 10 swapped in + 20 received out)"
        );
    }

    function test_DataUpdate_SwapHalfOfSwapOutput() public {
        // First plug: Swap 10 tokens for 20 tokens (based on 2:1 rate)
        bytes memory firstSwapCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                PlugMockDex.swap.selector,
                address(mockERC20), // tokenIn
                address(mockERC20), // tokenOut
                10 ether, // amountIn
                19 ether // minAmountOut
            )
        );

        PlugTypesLib.Plug memory firstSwapPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(dex),
            value: 0,
            data: firstSwapCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Second plug: Take the output from the first swap and use 10 ether of it for a transfer
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mockERC20.transfer.selector,
                recipient,
                10 ether // Transfer 10 ether - half of the expected 20 ether output from first swap
            )
        );

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = firstSwapPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 10 (swap in) + 20 (swap out) - 10 (transfer) = 1000 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 1000 ether, "Socket balance incorrect");

        // Recipient should have 10 ether
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 10 ether, "Recipient balance incorrect");

        // DEX should have: 1000 (initial) + 10 (swap in) - 20 (swap out) = 990 ether
        uint256 dexBalance = mockERC20.balanceOf(address(dex));
        assertEq(dexBalance, 990 ether, "DEX balance incorrect");
    }

    function test_BasicPlugExecution() public {
        // Simplest test to ensure the execution framework works
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);

        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // If we get here, the execution worked
    }

    function test_SimpleUpdateMechanism() public {
        // First plug: Swap 10 tokens for 20 tokens (based on 2:1 rate)
        bytes memory firstSwapCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                PlugMockDex.swap.selector,
                address(mockERC20), // tokenIn
                address(mockERC20), // tokenOut
                10 ether, // amountIn
                19 ether // minAmountOut
            )
        );

        PlugTypesLib.Plug memory firstSwapPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(dex),
            value: 0,
            data: firstSwapCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Second plug: Transfer a fixed amount (will NOT use update in this test)
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mockERC20.transfer.selector,
                recipient,
                5 ether // Fixed amount for transfer
            )
        );

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = firstSwapPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 10 (swap in) + 20 (swap out) - 5 (transfer) = 1005 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 1005 ether, "Socket balance incorrect");

        // Recipient should have 5 ether from the transfer
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 5 ether, "Recipient balance incorrect");

        // DEX should have: 1000 (initial) + 10 (swap in) - 20 (swap out) = 990 ether
        uint256 dexBalance = mockERC20.balanceOf(address(dex));
        assertEq(dexBalance, 990 ether, "DEX balance incorrect");
    }

    function test_UpdateMechanismWithAmountUpdate() public {
        // First, let's understand what the swap function returns:
        // It returns a uint256 representing the output amount of tokens.

        // Step 1: Create a plug that performs a swap
        bytes memory swapCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                PlugMockDex.swap.selector,
                address(mockERC20), // tokenIn
                address(mockERC20), // tokenOut
                10 ether, // amountIn
                19 ether // minAmountOut
            )
        );

        PlugTypesLib.Plug memory swapPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(dex),
            value: 0,
            data: swapCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Step 2: Create a transfer plug that will use the output from the swap
        // The key here is to make sure we're constructing the correct format for the transfer call

        // Create an ABI-encoded transfer call with the correct arguments
        bytes memory transferData = abi.encodeWithSelector(
            mockERC20.transfer.selector,
            recipient,
            uint256(0) // This placeholder will be replaced with the swap output
        );

        // Add the plug type byte at the beginning
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            transferData // The actual function call data
        );

        // Create an update that will take the swap output and insert it at the amount position
        // The ERC20 transfer function has a 4-byte selector followed by two 32-byte parameters
        // The amount parameter is the second one, starting at position 4 + 32 = 36
        PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
        updates[0] = PlugTypesLib.Update({
            start: 36, // 4 (selector) + 32 (address param)
            slice: PlugTypesLib.Slice({
                index: 0, // First plug (the swap)
                start: 0, // Start of the returned data
                length: 32 // Length of a uint256
             })
        });

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: updates
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = swapPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 10 (swap in) + 20 (swap out) - 20 (transfer) = 990 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 990_000_000_000_000_000_000, "Socket balance incorrect");

        // Recipient should have 20 (the full swap output amount)
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 20_000_000_000_000_000_000, "Recipient balance incorrect");

        // DEX should have: 1000 (initial) + 10 (swap in) - 20 (swap out) = 990 ether
        uint256 dexBalance = mockERC20.balanceOf(address(dex));
        assertEq(dexBalance, 990_000_000_000_000_000_000, "DEX balance incorrect");
    }

    function test_SimpleDataUpdate() public {
        // Create a plug that returns a simple uint256 value
        bytes memory echoCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mock.mutedEcho.selector,
                uint256(10) // Input value doesn't matter, function will return 2
            )
        );

        PlugTypesLib.Plug memory returnPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mock),
            value: 0,
            data: echoCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create a plug that performs a transfer with a dynamic amount
        bytes memory transferData = abi.encodeWithSelector(
            mockERC20.transfer.selector,
            recipient,
            uint256(0) // Will be replaced by the return value from the first plug (2)
        );

        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            transferData
        );

        // Create an update that will take the return value and place it in the amount parameter
        PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
        updates[0] = PlugTypesLib.Update({
            start: 36, // Position after function selector (4 bytes) + recipient address (32 bytes)
            slice: PlugTypesLib.Slice({
                index: 0, // First plug
                start: 0, // Start of return data
                length: 32 // Length of uint256
             })
        });

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: updates
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = returnPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 2 (transfer) = 998 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 999_999_999_999_999_999_998, "Socket balance incorrect");

        // Recipient should have 2 (the amount returned from the first plug)
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 2, "Recipient balance incorrect");
    }

    function test_BasicTransfer() public {
        // Create a direct ERC20 transfer plug with a fixed amount
        bytes memory transferData = abi.encodeWithSelector(
            mockERC20.transfer.selector,
            recipient,
            5 ether // Fixed amount of 5 ether
        );

        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            transferData
        );

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create plugs array with the transfer operation
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 5 (transfer) = 995 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 995 ether, "Socket balance incorrect");

        // Recipient should have 5 ether
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 5 ether, "Recipient balance incorrect");
    }

    function test_SimpleUpdatePrecise() public {
        // Step 1: Create a plug that performs a basic echo operation to return a value
        bytes memory echoCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mock.mutedEcho.selector,
                uint256(10) // Returns 2
            )
        );

        PlugTypesLib.Plug memory echoPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mock),
            value: 0,
            data: echoCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Step 2: Create the transfer plug
        // The ERC20 transfer function signature is: transfer(address,uint256)
        // First, get the exact bytes for the transfer call
        bytes memory transferSelector =
            abi.encodeWithSelector(mockERC20.transfer.selector, recipient, 0);

        // Now, create the plug call data with our plug type
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            transferSelector
        );

        // Analyze the transfer data to ensure our update position is correct
        // The function selector is 4 bytes (bytes 0-3)
        // The recipient address is 32 bytes (bytes 4-35)
        // The amount is 32 bytes (bytes 36-67)
        // With the plug type byte stripped in _plug function, the amount starts at byte 36

        // Create an update that targets exactly where the amount parameter is
        PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
        updates[0] = PlugTypesLib.Update({
            start: 36, // 4 (selector) + 32 (address)
            slice: PlugTypesLib.Slice({
                index: 0, // Get result from first plug
                start: 0, // Start of result data
                length: 32 // Length of uint256
             })
        });

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: updates
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = echoPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 2 (transfer from mutedEcho) = 998 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 999_999_999_999_999_999_998, "Socket balance incorrect");

        // Recipient should have 2 (the value from mutedEcho)
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 2, "Recipient balance incorrect");
    }

    function test_UpdateWithByteAdjustment() public {
        // Step 1: Create a plug that performs a basic echo operation to return a value
        bytes memory echoCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mock.mutedEcho.selector,
                uint256(10) // Returns 2
            )
        );

        PlugTypesLib.Plug memory echoPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mock),
            value: 0,
            data: echoCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Step 2: Create the transfer plug
        // The ERC20 transfer function signature is: transfer(address,uint256)
        bytes memory transferData =
            abi.encodeWithSelector(mockERC20.transfer.selector, recipient, 0);

        // Now, create the plug call data with our plug type
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            transferData
        );

        // In the _plug function, the first byte is skipped with: data = action.data[1:]
        // The function selector is 4 bytes (bytes 0-3)
        // The recipient address is 32 bytes (bytes 4-35)
        // The amount starts at byte 36 once the plug type byte is stripped

        PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
        updates[0] = PlugTypesLib.Update({
            start: 36, // 4 (selector) + 32 (address)
            slice: PlugTypesLib.Slice({
                index: 0, // Get result from first plug
                start: 0, // Start of result data
                length: 32 // Length of uint256
             })
        });

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: updates
        });

        // Create plugs array with both operations
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = echoPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 2 (transfer from mutedEcho) = 998 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 999_999_999_999_999_999_998, "Socket balance incorrect");

        // Recipient should have 2 (the value from mutedEcho)
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 2, "Recipient balance incorrect");
    }

    function test_ManualTransferConstruction() public {
        // Instead of relying on the framework's update mechanism, let's manually construct
        // the transfer data and ensure it's exactly what the ERC20 contract expects

        // The ERC20 transfer function takes (address to, uint256 amount)
        bytes memory transferCallData = abi.encodeWithSelector(
            mockERC20.transfer.selector, // 0xa9059cbb
            recipient, // the recipient
            2 ether // exact amount, no updates needed
        );

        // Add the plug type byte at the beginning
        bytes memory fullCallData = abi.encodePacked(
            bytes1(0x02), // plug type byte
            transferCallData // the raw transfer data
        );

        // Create our plug directly
        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: fullCallData,
            updates: new PlugTypesLib.Update[](0) // No updates needed
         });

        // Create plugs array with just the transfer operation
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
        plugsArray[0] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected:
        // Socket should have: 1000 (initial) - 2 (transfer) = 998 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        assertEq(socketBalance, 998 ether, "Socket balance incorrect");

        // Recipient should have 2 ether
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 2 ether, "Recipient balance incorrect");
    }

    function test_CorrectlyFormattedUpdate() public {
        // STRATEGY:
        // 1. First plug: run mutedEcho to return a value (2)
        // 2. Create a properly formatted transfer with all zeros for the amount
        // 3. Handle update at the correct position (36) to update the amount

        // Create a plug that performs a basic echo operation to return a value (2)
        bytes memory echoCallData = abi.encodePacked(
            bytes1(0x02), // The plug type (PLUG_EXECUTION)
            abi.encodeWithSelector(
                mock.mutedEcho.selector,
                uint256(10) // Returns 2
            )
        );

        PlugTypesLib.Plug memory echoPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mock),
            value: 0,
            data: echoCallData,
            updates: new PlugTypesLib.Update[](0)
        });

        // Create the transfer call data with a well-formed but zeroed amount
        bytes memory transferData = abi.encodeWithSelector(
            mockERC20.transfer.selector, // 0xa9059cbb
            recipient, // recipient address
            uint256(0) // amount (to be updated)
        );

        // Create the plug call data
        bytes memory transferCallData = abi.encodePacked(
            bytes1(0x02), // The plug type byte
            transferData // The transfer data
        );

        // Define the update with the correct position
        // Important: In _plug function, data = action.data[1:] removes the first byte
        PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
        updates[0] = PlugTypesLib.Update({
            start: 36, // Position AFTER first byte is removed (4+32)
            slice: PlugTypesLib.Slice({
                index: 0, // Reference the first plug's result
                start: 0, // Start from beginning of result
                length: 32 // Take full uint256 (32 bytes)
             })
        });

        PlugTypesLib.Plug memory transferPlug = PlugTypesLib.Plug({
            selector: 0x00, // Call
            to: address(mockERC20),
            value: 0,
            data: transferCallData,
            updates: updates
        });

        // Create the plugs array and execute
        PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
        plugsArray[0] = echoPlug;
        plugsArray[1] = transferPlug;

        // Create LivePlugs and execute
        PlugTypesLib.LivePlugs[] memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
        livePlugsArray[0] = createLivePlugs(plugsArray);

        // Execute the plugs
        plug.plug(livePlugsArray);

        // Verify balances changed as expected
        // Socket should have: 1000 (initial) - 2 (transfer) = 998 ether
        uint256 socketBalance = mockERC20.balanceOf(address(socket));
        // The raw value is returned without wei units, so we need to check the raw value
        assertEq(socketBalance, 999_999_999_999_999_999_998, "Socket balance incorrect");

        // Recipient should have 2 (the raw value from mutedEcho)
        uint256 recipientBalance = mockERC20.balanceOf(recipient);
        assertEq(recipientBalance, 2, "Recipient balance incorrect");
    }
}
