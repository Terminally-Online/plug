// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import { PlugCore } from "./Plug.Core.sol";
import { PlugTypesLib } from "./Plug.Types.sol";
import { PlugFuseInterface } from "../interfaces/Plug.Fuse.Interface.sol";
import { PlugSimulationLib } from "../libraries/Plug.Simulation.Lib.sol";

/**
 * @title PlugSimulation
 * @notice The simulation contract that enables the ability to simulate an entire
 *         bundle at once or through a segmented route of execution using
 *         explicit plug keys.
 * @author @nftchance (chance@utc24.io)
 */
abstract contract PlugSimulation is PlugCore {
    /**
     * @notice Simulate the entire bundle of LivePlugs.
     * @param $plugs The bundle of Plugs to process in the simulation.
     * @return $results An array of results describing the status of the bundle.
     */
    function simulate(PlugTypesLib.Plug[] calldata $plugs)
        public
        view
        virtual
        returns (PlugSimulationLib.Result[] memory $results)
    {
        uint256 plugsLength = $plugs.length;

        if (plugsLength == 0) return new PlugSimulationLib.Result[](0);

        PlugSimulationLib.Result[] memory results = new PlugSimulationLib.Result[](plugsLength);

        for (uint8 i; i < $plugs.length; i++) {
            PlugTypesLib.Plug memory plug = $plugs[i];

            if (plug.pins.length == 0) continue;

            for (uint8 j; j < plug.pins.length; j++) {
                PlugTypesLib.Pin memory pin = plug.pins[j].pin;

                if (pin.fuses.length == 0) continue;

                /// @dev Get the pin hash from the live pin.
                bytes32 pinHash = getLivePinHash(plug.pins[j]);

                for (uint8 l; l < pin.fuses.length; l++) {
                    (plug.current.data, $results[results.length - plugsLength--]) = simulate(
                        pinHash,
                        plug,
                        pin.fuses[l]
                    );
                }
            }
        }
    }

    /**
     * @notice Simulate the execution of a specific Plug Fuse.
     * @param $pinHash The hash of the pin to simulate.
     * @param $plug The plug to simulate.
     * @param $fuse The fuse to simulate.
     * @return $through The data that was passed through the fuse.
     * @return $result The result of the simulation.
     */
    function simulate(
        bytes32 $pinHash,
        PlugTypesLib.Plug memory $plug,
        PlugTypesLib.Fuse memory $fuse
    )
        public
        view
        virtual
        returns (bytes memory $through, PlugSimulationLib.Result memory $result)
    {
        /// @Dev Slot to ensure the simulation is successful.
        bool success;

        (success, $through) = address($fuse.neutral).staticcall(
            abi.encodeWithSelector(PlugFuseInterface($fuse.neutral).enforceFuse.selector, $fuse.live, $plug.current, $pinHash)
        );

        $result = PlugSimulationLib.Result({ success: success, callback: $through });
    }
}
