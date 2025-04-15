// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { PlugFactoryInterface } from "../interfaces/Plug.Factory.Interface.sol";

import { PlugLib } from "../libraries/Plug.Lib.sol";
import { LibBytes } from "solady/utils/LibBytes.sol";
import { LibClone } from "solady/utils/LibClone.sol";

import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug Factory
 * @notice This contract is responsible for deploying new Plug Sockets that can be used
 *         as personal accounts for an individual. The Sockets are deployed using the
 *         Beacon Proxy pattern, and the owner can upgrade the implementation at any time.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugFactory is PlugFactoryInterface {
    /**
     * See { PlugFactoryInterface.deploy }
     */
    function deploy(bytes calldata $salt)
        public
        payable
        virtual
        returns (bool $alreadyDeployed, address $socketAddress)
    {
        if ($salt.length < 0x80) revert("PlugCore:salt-malformed");
        uint96 nonce = uint96(uint256(LibBytes.loadCalldata($salt, 0x00)));
        address admin =
            address(uint160(uint256(LibBytes.loadCalldata($salt, 0x20))));
        address oneClicker =
            address(uint160(uint256(LibBytes.loadCalldata($salt, 0x40))));
        address implementation =
            address(uint160(uint256(LibBytes.loadCalldata($salt, 0x60))));
        if (implementation == address(0) || admin == address(0)) {
            revert PlugLib.SaltInvalid(implementation, admin);
        }

        bytes32 salt = bytes32(abi.encodePacked(uint96(nonce), bytes20(admin)));
        ($alreadyDeployed, $socketAddress) =
            LibClone.createDeterministicERC1967(msg.value, implementation, salt);
        if (!$alreadyDeployed) {
            emit PlugLib.SocketDeployed(implementation, admin, salt);
            PlugSocketInterface($socketAddress).initialize(admin, oneClicker);
        }
    }

    /**
     * See { PlugFactoryInterface.getAddress }
     */
    function getAddress(
        address $implementation,
        bytes32 $salt
    )
        public
        view
        returns (address $vault)
    {
        $vault = LibClone.predictDeterministicAddressERC1967(
            $implementation, $salt, address(this)
        );
    }

    /**
     * See { PlugFactoryInterface.initCodeHash }
     */
    function initCodeHash(address $implementation)
        public
        view
        virtual
        returns (bytes32 $initCodeHash)
    {
        $initCodeHash = LibClone.initCodeHashERC1967($implementation);
    }
}
