// SPDX-License-Identifier: BUSL-1.1

pragma solidity 0.8.23;

import {PlugSimulation} from './Plug.Simulation.sol';
import {Receiver} from 'solady/src/accounts/Receiver.sol';
import {IPlug} from '../interfaces/IPlug.sol';

import {PlugTypesLib} from './Plug.Types.sol';

/**
 * @title Plug
 * @notice The core contract for the Plug framework that enables
 *         counterfactual revokable pin of extremely
 *         granular pin and execution paths.
 * @author @nftchance (chance@utc24.io)
 * @author @danfinlay (https://github.com/delegatable/delegatable-sol)
 * @author @KamesGeraghty (https://github.com/kamescg)
 */
contract PlugSocket is PlugSimulation, Receiver, IPlug {
	/**
	 * See {IPlug-plug}.
	 */
	function plug(
		PlugTypesLib.LivePlugs calldata $livePlugs
	) external payable returns (bytes[] memory $results) {
        /// @dev Determine who signed the intent.
        address intentSigner = getLivePlugsSigner($livePlugs);

        /// @dev Prevent random people from plugging.
        _enforceSigner(intentSigner);

        /// @dev Load the plugs as a hot reference.
        PlugTypesLib.Plugs calldata plugs = $livePlugs.plugs;

        /// @dev Prevent replay attacks by enforcing replay protection.
        _enforceBreaker(intentSigner, plugs.breaker);

        /// @dev Invoke the plugs.
        $results = _plug(plugs.plugs, intentSigner);
	}

	/**
	 * See {IPlug-plugContract}.
	 */
	function plugContract(
		PlugTypesLib.Plug[] calldata $plugs
	) external payable returns (bytes[] memory $result) {
		$result = _plug($plugs, msg.sender);
	}

    /**
     * @notice Confirm that signer of the intent has permission to declare
     *         the execution of an intent.
     * @dev If you would like to limit the available signers override this
     *      function in your contract with the additional logic.
     */
    function _enforceSigner(address $signer) internal view virtual {}
}
