// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import {Test, PlugCoilLib, PlugLib, PlugTypesLib, PlugAddressesLib, PlugEtcherLib, PlugFactory, Plug, PlugMockDex, PlugMockEcho, PlugMockERC20} from '../abstracts/test/Plug.Test.sol';
import {ECDSA} from 'solady/utils/ECDSA.sol';
import {ERC20} from 'solady/tokens/ERC20.sol';

contract PlugTest is Test {
	event EchoInvoked(address $sender, string $message);

	function setUp() public virtual {
		setUpPlug();
	}

	function test_name() public {
		assertEq(plug.name(), 'Plug');
	}

	function test_symbol() public {
		assertEq(plug.symbol(), 'PLUG');
	}

	function test_PlugEmptyEcho() public {
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugs);
		plug.plug(livePlugsArray);
	}

	function test_PlugEmptyEchoWithSolver() public {
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugsArray);
		vm.expectEmit(address(mock));
		emit EchoInvoked(address(socket), 'Hello World');
		plug.plug(livePlugsArray);
	}

	function test_PlugEmptyEchoWithSolverAndTreasuryPayment() public {
		address solver = _randomNonZeroAddress();
		address treasury = 0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E;
		vm.deal(solver, 100 ether);
		vm.deal(address(socket), 100 ether);
		uint256 preBalance = address(treasury).balance;
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugsArray);
		vm.prank(solver);
		vm.expectEmit(address(mock));
		emit EchoInvoked(address(socket), 'Hello World');
		plug.plug(livePlugsArray);
		assertTrue(address(treasury).balance > preBalance);
	}

	function testRevert_PlugEmptyEchoWithSolverAndTreasuryPaymentFailure()
		public
	{
		address solver = _randomNonZeroAddress();
		vm.deal(solver, 100 ether);
		vm.deal(address(socket), 0);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugsArray, solver);
		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.prank(solver);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({index: 1, error: 'PlugCore:plug-failed'})
		);
		plug.plug(livePlugsArray);
	}

	function test_PlugEmptyEchoWithSolverAndInvalidNonce() public {
		address solver = _randomNonZeroAddress();
		vm.deal(solver, 100 ether);
		vm.deal(address(socket), 100 ether);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugsArray, solver);
		vm.prank(solver);
		vm.expectEmit(address(mock));
		emit EchoInvoked(address(socket), 'Hello World');
		plug.plug(livePlugsArray);

		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({
				index: type(uint8).max,
				error: 'PlugCore:nonce-invalid'
			})
		);
		plug.plug(livePlugsArray);
	}

	function testRevert_PlugEmptyEchoWithSolverAndInvalidSignature() public {
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugs);
		bytes32 digest = socket.getPlugsDigest(plugs);
		(uint8 v, bytes32 r, bytes32 s) = vm.sign(0x123456, digest);
		bytes memory signature = abi.encodePacked(r, s, v);
		livePlugsArray[0] = PlugTypesLib.LivePlugs({
			plugs: plugs,
			signature: signature
		});

		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({
				index: type(uint8).max,
				error: 'PlugCore:signature-invalid'
			})
		);
		plug.plug(livePlugsArray);
	}

	function testRevert_PlugEmptyEchoWithInvalidSolver() public {
		address solver = _randomNonZeroAddress();
		vm.deal(solver, 100 ether);
		vm.deal(address(socket), 100 ether);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		plugsArray[1] = createPlug(PLUG_VALUE, PLUG_EXECUTION);
		PlugTypesLib.Plugs memory plugs = createPlugs(
			plugsArray,
			uint48(block.timestamp + 3 minutes),
			solver
		);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugs);
		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({
				index: type(uint8).max,
				error: 'PlugCore:solver-invalid'
			})
		);
		plug.plug(livePlugsArray);
	}

	function testRevert_PlugEmptyEchoWithExpiredSolver() public {
		address solver = _randomNonZeroAddress();
		vm.deal(solver, 100 ether);
		vm.deal(address(socket), 100 ether);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		PlugTypesLib.Plugs memory plugs = createPlugs(
			plugsArray,
			uint48(block.timestamp - 1 minutes),
			solver
		);
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugs);
		vm.prank(solver);

		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({
				index: type(uint8).max,
				error: 'PlugCore:solver-expired'
			})
		);
		plug.plug(livePlugsArray);
	}

	function test_RevertingPlug() public {
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](1);

		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);

		PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
		livePlugsArray[0] = createLivePlugs(plugs);

		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({index: 0, error: 'PlugCore:plug-failed'})
		);
		plug.plug(livePlugsArray);
	}

	function test_RevertingPlugWithFollowingIndex() public {
		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2); // Array of 2 plugs
		plugsArray[0] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
		plugsArray[1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
		PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
		livePlugsArray[0] = createLivePlugs(plugs);

		bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
		vm.expectEmit(address(plug));
		emit PlugLib.PlugResult(
			0,
			livePlugsHash,
			PlugTypesLib.Result({index: 1, error: 'PlugCore:plug-failed'})
		);
		plug.plug(livePlugsArray);
	}

	function test_RevertingPlugAtMultipleIndexes() public {
		for (uint8 testIndex = 0; testIndex < 5; testIndex++) {
			PlugTypesLib.LivePlugs[]
				memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
			PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](
				testIndex + 2
			);
			for (uint8 i = 0; i < testIndex + 1; i++) {
				plugsArray[i] = createPlug(PLUG_NO_VALUE, PLUG_EXECUTION);
			}
			plugsArray[testIndex + 1] = createPlug(PLUG_NO_VALUE, PLUG_REVERT);
			PlugTypesLib.Plugs memory plugs = createPlugs(plugsArray);
			livePlugsArray[0] = createLivePlugs(plugs);

			bytes32 livePlugsHash = socket.hash(livePlugsArray[0]);
			vm.expectEmit(address(plug));
			emit PlugLib.PlugResult(
				0,
				livePlugsHash,
				PlugTypesLib.Result({
					index: testIndex + 1,
					error: 'PlugCore:plug-failed'
				})
			);
			plug.plug(livePlugsArray);
		}
	}

	// function useInputs(uint8 index) public pure returns (uint8) {
	// 	return index;
	// }
	//
	// function useOutputs(uint8 index) public pure returns (uint8) {
	// 	return index + 1;
	// }

	function test_PlugTransferWithBalanceCoil() public {
		address recipient = _randomNonZeroAddress();
		uint256 initialSocketBalance = mockERC20.balanceOf(address(socket));
		uint256 initialRecipientBalance = mockERC20.balanceOf(recipient);

		PlugTypesLib.Plug[] memory plugsArray = new PlugTypesLib.Plug[](2);

		plugsArray[0] = PlugTypesLib.Plug({
			data: abi.encode(
				uint8(0x03),
				address(mockERC20),
				uint256(0),
				abi.encodeWithSelector(
					mockERC20.balanceOf.selector,
					address(socket)
				)
			),
			updates: new PlugTypesLib.Update[](0)
		});

		PlugTypesLib.Update[] memory updates = new PlugTypesLib.Update[](1);
		updates[0] = PlugTypesLib.Update({
			start: 196,
			slice: PlugTypesLib.Slice({
				index: 1,
				start: 0,
				length: 32,
				typeId: 0x00
			})
		});
		plugsArray[1] = PlugTypesLib.Plug({
			data: abi.encode(
				uint8(0x00),
				address(mockERC20),
				uint256(0),
				abi.encodeWithSelector(
					mockERC20.transfer.selector,
					recipient,
					uint256(0)
				)
			),
			updates: updates
		});

		PlugTypesLib.LivePlugs[]
			memory livePlugsArray = new PlugTypesLib.LivePlugs[](1);
		livePlugsArray[0] = createLivePlugs(plugsArray);
		plug.plug(livePlugsArray);

		assertEq(mockERC20.balanceOf(address(socket)), 0);
		assertEq(
			mockERC20.balanceOf(recipient),
			initialRecipientBalance + initialSocketBalance
		);
	}
}
