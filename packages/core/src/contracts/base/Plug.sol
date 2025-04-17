// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

import { PlugInterface } from "../interfaces/Plug.Interface.sol";

import { PlugLib, PlugTypesLib, PlugAddressesLib } from "../libraries/Plug.Lib.sol";

import { PlugFactoryInterface } from "../interfaces/Plug.Factory.Interface.sol";
import { PlugSocketInterface } from "../interfaces/Plug.Socket.Interface.sol";

/**
 * @title Plug
 * @notice This contract represents a general purpose relay used to route signed
 *         intents to target Sockets.
 * @dev There is no need to approve assets to this contract as all transactions
 *      are executed through the Socket which will manage its own permissions
 *      that can be safely approved to interact with the assets of another account.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract Plug is PlugInterface {
    /// @dev Define the reference to the factory that enables counterfactual
    ///      deployment through having a presigned bundle of Plugs.
    PlugFactoryInterface factory = PlugFactoryInterface(PlugAddressesLib.PLUG_FACTORY_ADDRESS);

    /**
     * See {PlugInterface-plug}.
     */
    function plug(PlugTypesLib.LivePlugs calldata $livePlugs) external payable virtual {
        (PlugTypesLib.Result memory results, bytes32 livePlugsHash) = _plug($livePlugs, msg.sender);
        emit PlugLib.PlugResult(0, livePlugsHash, results);
    }

    /**
     * See {PlugInterface-plug}.
     */
    function plug(PlugTypesLib.LivePlugs[] calldata $livePlugs) external payable virtual {
        PlugTypesLib.Result memory results;
        bytes32 livePlugsHash;
        for (uint8 i; i < $livePlugs.length; i++) {
            (results, livePlugsHash) = _plug($livePlugs[i], msg.sender);
            emit PlugLib.PlugResult(i, livePlugsHash, results);
        }
    }

    /**
     * See {PlugInterface-name}.
     */
    function name() public pure returns (string memory $name) {
        $name = "Plug";
    }

    /**
     * See {PlugInterface-symbol}.
     */
    function symbol() public pure returns (string memory $version) {
        $version = "PLUG";
    }

    /**
     * @notice Initialize the Plug with the factory and the implementation if
     *         it has not been initialized yet, otherwise just use the address
     *         included in the Plug bundle.
     * @param $livePlugs The signed bundle of Plugs being executed.
     * @return $socket The Socket to use.
     */
    function _socket(PlugTypesLib.LivePlugs calldata $livePlugs)
        internal
        virtual
        returns (PlugSocketInterface $socket)
    {
        address socketAddress = $livePlugs.plugs.socket;
        if (socketAddress.code.length == 0) {
            (, address $socketAddress) = factory.deploy($livePlugs.plugs.salt);
            if (socketAddress != $socketAddress) {
                revert PlugLib.SocketAddressInvalid(socketAddress, $socketAddress);
            }
        }
        $socket = PlugSocketInterface(socketAddress);
    }

    /**
     * @notice Internal function to compute the hash of a live plug bundle
     * @param $socket The socket interface to use for hashing
     * @param $livePlugs The signed bundle of Plugs being hashed
     * @return $livePlugsHash The computed hash of the live plugs bundle
     */
    function _hash(
        PlugSocketInterface $socket,
        PlugTypesLib.LivePlugs calldata $livePlugs
    )
        internal
        view
        returns (bytes32 $livePlugsHash)
    {
        $livePlugsHash = $socket.hash($livePlugs);
    }

    /**
     * @notice Internal function to execute the Plug.
     * @param $livePlugs The signed bundle of Plugs being executed.
     * @param $sender The sender of the transaction.
     * @return $results The results of the Plug execution.
     */
    function _plug(
        PlugTypesLib.LivePlugs calldata $livePlugs,
        address $sender
    )
        internal
        returns (PlugTypesLib.Result memory $results, bytes32 $livePlugsHash)
    {
        PlugSocketInterface socket = _socket($livePlugs);

        $livePlugsHash = _hash(socket, $livePlugs);

        try socket.plug($livePlugs, $sender) returns (PlugTypesLib.Result memory _results) {
            $results = _results;
        } catch Error(string memory reason) {
            $results = PlugTypesLib.Result({ index: type(uint8).max - 1, error: reason });
        } catch Panic(uint256 errorCode) {
            $results = PlugTypesLib.Result({
                index: type(uint8).max - 2,
                error: string(abi.encode(errorCode))
            });
        } catch (bytes memory data) {
            if (data.length < 4) {
                return (
                    PlugTypesLib.Result({ index: type(uint8).max - 4, error: "Plug:empty-data" }),
                    $livePlugsHash
                );
            }

            if (bytes4(data) != PlugLib.PlugFailed.selector) {
                return (
                    PlugTypesLib.Result({
                        index: type(uint8).max - 3,
                        error: "Plug:unknown-selector"
                    }),
                    $livePlugsHash
                );
            }

            bytes memory slicedData;
            assembly {
                slicedData := mload(0x40)
                let errLength := sub(mload(data), 4)
                mstore(slicedData, errLength)
                let srcPtr := add(data, 0x24)
                let destPtr := add(slicedData, 0x20)
                for { let offset := 0 } lt(offset, errLength) { offset := add(offset, 0x20) } {
                    mstore(add(destPtr, offset), mload(add(srcPtr, offset)))
                }
                mstore(0x40, add(slicedData, add(0x20, errLength)))
            }
            (uint8 index, string memory reason) = abi.decode(slicedData, (uint8, string));

            $results = PlugTypesLib.Result({ index: index, error: reason });
        }
    }
}
