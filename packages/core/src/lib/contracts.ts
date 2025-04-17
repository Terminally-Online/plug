export const contracts = [
    {
        "name": "PlugAssert",
        "abi": [
            {
                "type": "function",
                "name": "assertFalse",
                "inputs": [
                    {
                        "name": "condition",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "message",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "outputs": [],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "assertFalse",
                "inputs": [
                    {
                        "name": "condition",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "assertTrue",
                "inputs": [
                    {
                        "name": "condition",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "assertTrue",
                "inputs": [
                    {
                        "name": "condition",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "message",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "outputs": [],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "fail",
                "inputs": [
                    {
                        "name": "message",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "outputs": [],
                "stateMutability": "pure"
            }
        ]
    },
    {
        "name": "PlugBoolean",
        "abi": [
            {
                "type": "function",
                "name": "isAfterTime",
                "inputs": [
                    {
                        "name": "time",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isAnd",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isBeforeTime",
                "inputs": [
                    {
                        "name": "time",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isBetween",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "min",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "max",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isBetweenTimes",
                "inputs": [
                    {
                        "name": "time",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "start",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "end",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isEqual",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isFalse",
                "inputs": [
                    {
                        "name": "value",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isGreaterThan",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isGreaterThanOrEqual",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isImplies",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isLessThan",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isLessThanOrEqual",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isNand",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isNor",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isNot",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isNotEqual",
                "inputs": [
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isOr",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isSameDay",
                "inputs": [
                    {
                        "name": "timestamp1",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "timestamp2",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isTrue",
                "inputs": [
                    {
                        "name": "value",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isWeekday",
                "inputs": [
                    {
                        "name": "timestamp",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isWeekend",
                "inputs": [
                    {
                        "name": "timestamp",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "isXor",
                "inputs": [
                    {
                        "name": "a",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "b",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            }
        ]
    },
    {
        "name": "PlugCoercion",
        "abi": [
            {
                "type": "function",
                "name": "addressToUint160",
                "inputs": [
                    {
                        "name": "y",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint160",
                        "internalType": "uint160"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "boolToInt",
                "inputs": [
                    {
                        "name": "y",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "boolToUint",
                "inputs": [
                    {
                        "name": "y",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "bytes32ToBytes",
                "inputs": [
                    {
                        "name": "y",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "bytes32ToUint",
                "inputs": [
                    {
                        "name": "y",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "bytesToBytes32",
                "inputs": [
                    {
                        "name": "y",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "intToBool",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt128",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int128",
                        "internalType": "int128"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt16",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int16",
                        "internalType": "int16"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt256",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt32",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int32",
                        "internalType": "int32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt64",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int64",
                        "internalType": "int64"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toInt8",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "int8",
                        "internalType": "int8"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint128",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint16",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint16",
                        "internalType": "uint16"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint160",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint160",
                        "internalType": "uint160"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint256",
                "inputs": [
                    {
                        "name": "y",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint32",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint64",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint64",
                        "internalType": "uint64"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint8",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toUint96",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "uint96",
                        "internalType": "uint96"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "uint160ToAddress",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint160",
                        "internalType": "uint160"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "uintToBool",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "uintToBytes32",
                "inputs": [
                    {
                        "name": "y",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "z",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            }
        ]
    },
    {
        "name": "PlugDatabase",
        "abi": [
            {
                "type": "function",
                "name": "TYPE_ADDRESS",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_BOOL",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_BYTES",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_BYTES32",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_INT256",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_STRING",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TYPE_UINT256",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "addressStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "batchSet",
                "inputs": [
                    {
                        "name": "keys",
                        "type": "bytes32[]",
                        "internalType": "bytes32[]"
                    },
                    {
                        "name": "values",
                        "type": "bytes32[]",
                        "internalType": "bytes32[]"
                    },
                    {
                        "name": "typeId",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "boolStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "bytes32Storage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "bytesStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "exists",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "typeId",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "get",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getAddress",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBool",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBytes",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getInt",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getString",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getType",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getUint",
                "inputs": [
                    {
                        "name": "sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "intStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "remove",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "removeWithType",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "typeId",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "set",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setAddress",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setBool",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setBytes",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setInt",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setString",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setUint",
                "inputs": [
                    {
                        "name": "key",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "stringStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "uintStorage",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            }
        ]
    },
    {
        "name": "PlugEVM",
        "abi": [
            {
                "type": "receive",
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "balanceOf",
                "inputs": [
                    {
                        "name": "account",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBlockGasLimit",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBlockHash",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBlockHash",
                "inputs": [
                    {
                        "name": "blockNumber",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getBlockNumber",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getCaller",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getChainId",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getCodeHash",
                "inputs": [
                    {
                        "name": "account",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getGasPrice",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getRemainingGas",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getTimestamp",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "isContract",
                "inputs": [
                    {
                        "name": "account",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "transfer",
                "inputs": [
                    {
                        "name": "recipient",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "gasLimit",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "success",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "returnData",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "transfer",
                "inputs": [
                    {
                        "name": "recipient",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            }
        ]
    },
    {
        "name": "PlugFactory",
        "abi": [
            {
                "type": "function",
                "name": "deploy",
                "inputs": [
                    {
                        "name": "$salt",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$alreadyDeployed",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "$socketAddress",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "getAddress",
                "inputs": [
                    {
                        "name": "$implementation",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$salt",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$vault",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "initCodeHash",
                "inputs": [
                    {
                        "name": "$implementation",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "$initCodeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "event",
                "name": "SocketDeployed",
                "inputs": [
                    {
                        "name": "implementation",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "vault",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "salt",
                        "type": "bytes32",
                        "indexed": false,
                        "internalType": "bytes32"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "SaltInvalid",
                "inputs": [
                    {
                        "name": "$implementation",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$admin",
                        "type": "address",
                        "internalType": "address"
                    }
                ]
            }
        ]
    },
    {
        "name": "PlugMath",
        "abi": [
            {
                "type": "function",
                "name": "add",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "clamp",
                "inputs": [
                    {
                        "name": "value",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "minValue",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "maxValue",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "divide",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "max",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "min",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "modulo",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "multiply",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "power",
                "inputs": [
                    {
                        "name": "base",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "exponent",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "subtract",
                "inputs": [
                    {
                        "name": "a",
                        "type": "int256",
                        "internalType": "int256"
                    },
                    {
                        "name": "b",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "int256",
                        "internalType": "int256"
                    }
                ],
                "stateMutability": "pure"
            }
        ]
    },
    {
        "name": "PlugRewards",
        "abi": [
            {
                "type": "constructor",
                "inputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "cancelOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "claimReward",
                "inputs": [
                    {
                        "name": "$period",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$merkleProof",
                        "type": "bytes32[]",
                        "internalType": "bytes32[]"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "completeOwnershipHandover",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "createRewardPeriod",
                "inputs": [
                    {
                        "name": "$merkleRoot",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "$totalAmount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "currentPeriod",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "fundRewards",
                "inputs": [
                    {
                        "name": "$amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "getRewardBalance",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "hasValidClaim",
                "inputs": [
                    {
                        "name": "$period",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$user",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$merkleProof",
                        "type": "bytes32[]",
                        "internalType": "bytes32[]"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "initialize",
                "inputs": [],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "owner",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownershipHandoverExpiresAt",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "periodMerkleRoots",
                "inputs": [
                    {
                        "name": "period",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "merkleRoot",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "periodTotalAmounts",
                "inputs": [
                    {
                        "name": "period",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "renounceOwnership",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "requestOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "rewardClaimed",
                "inputs": [
                    {
                        "name": "period",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "user",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "claimed",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "rewardToken",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "transferOwnership",
                "inputs": [
                    {
                        "name": "newOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "event",
                "name": "Initialized",
                "inputs": [
                    {
                        "name": "version",
                        "type": "uint64",
                        "indexed": false,
                        "internalType": "uint64"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "NewRewardPeriod",
                "inputs": [
                    {
                        "name": "period",
                        "type": "uint256",
                        "indexed": true,
                        "internalType": "uint256"
                    },
                    {
                        "name": "merkleRoot",
                        "type": "bytes32",
                        "indexed": false,
                        "internalType": "bytes32"
                    },
                    {
                        "name": "totalAmount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverCanceled",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverRequested",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipTransferred",
                "inputs": [
                    {
                        "name": "oldOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "newOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "RewardClaimed",
                "inputs": [
                    {
                        "name": "period",
                        "type": "uint256",
                        "indexed": true,
                        "internalType": "uint256"
                    },
                    {
                        "name": "user",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "AlreadyInitialized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InsufficientRewardBalance",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidInitialization",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidMerkleProof",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NewOwnerIsZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NoHandoverRequest",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NotInitializing",
                "inputs": []
            },
            {
                "type": "error",
                "name": "PeriodNotInitialized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "RewardsAlreadyClaimed",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Unauthorized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "ZeroAmount",
                "inputs": []
            }
        ]
    },
    {
        "name": "PlugSocket",
        "abi": [
            {
                "type": "constructor",
                "inputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "fallback",
                "stateMutability": "payable"
            },
            {
                "type": "receive",
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "cancelOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "completeOwnershipHandover",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "domain",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$domain",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.EIP712Domain",
                        "components": [
                            {
                                "name": "name",
                                "type": "string",
                                "internalType": "string"
                            },
                            {
                                "name": "version",
                                "type": "string",
                                "internalType": "string"
                            },
                            {
                                "name": "chainId",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "verifyingContract",
                                "type": "address",
                                "internalType": "address"
                            }
                        ]
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "domainHash",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getEIP712DomainHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.EIP712Domain",
                        "components": [
                            {
                                "name": "name",
                                "type": "string",
                                "internalType": "string"
                            },
                            {
                                "name": "version",
                                "type": "string",
                                "internalType": "string"
                            },
                            {
                                "name": "chainId",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "verifyingContract",
                                "type": "address",
                                "internalType": "address"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getLivePlugsHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.LivePlugs",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getLivePlugsSigner",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.LivePlugs",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$signer",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getPlugArrayHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.Plug[]",
                        "components": [
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            },
                            {
                                "name": "updates",
                                "type": "tuple[]",
                                "internalType": "struct PlugTypesLib.Update[]",
                                "components": [
                                    {
                                        "name": "start",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "slice",
                                        "type": "tuple",
                                        "internalType": "struct PlugTypesLib.Slice",
                                        "components": [
                                            {
                                                "name": "index",
                                                "type": "uint8",
                                                "internalType": "uint8"
                                            },
                                            {
                                                "name": "start",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "length",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "typeId",
                                                "type": "uint8",
                                                "internalType": "uint8"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getPlugHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Plug",
                        "components": [
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            },
                            {
                                "name": "updates",
                                "type": "tuple[]",
                                "internalType": "struct PlugTypesLib.Update[]",
                                "components": [
                                    {
                                        "name": "start",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "slice",
                                        "type": "tuple",
                                        "internalType": "struct PlugTypesLib.Slice",
                                        "components": [
                                            {
                                                "name": "index",
                                                "type": "uint8",
                                                "internalType": "uint8"
                                            },
                                            {
                                                "name": "start",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "length",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "typeId",
                                                "type": "uint8",
                                                "internalType": "uint8"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getPlugsDigest",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Plugs",
                        "components": [
                            {
                                "name": "socket",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "plugs",
                                "type": "tuple[]",
                                "internalType": "struct PlugTypesLib.Plug[]",
                                "components": [
                                    {
                                        "name": "data",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "updates",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Update[]",
                                        "components": [
                                            {
                                                "name": "start",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "slice",
                                                "type": "tuple",
                                                "internalType": "struct PlugTypesLib.Slice",
                                                "components": [
                                                    {
                                                        "name": "index",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    },
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "length",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "typeId",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "solver",
                                "type": "bytes",
                                "internalType": "bytes"
                            },
                            {
                                "name": "salt",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$digest",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "getPlugsHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Plugs",
                        "components": [
                            {
                                "name": "socket",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "plugs",
                                "type": "tuple[]",
                                "internalType": "struct PlugTypesLib.Plug[]",
                                "components": [
                                    {
                                        "name": "data",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "updates",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Update[]",
                                        "components": [
                                            {
                                                "name": "start",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "slice",
                                                "type": "tuple",
                                                "internalType": "struct PlugTypesLib.Slice",
                                                "components": [
                                                    {
                                                        "name": "index",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    },
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "length",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "typeId",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "solver",
                                "type": "bytes",
                                "internalType": "bytes"
                            },
                            {
                                "name": "salt",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getSliceHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Slice",
                        "components": [
                            {
                                "name": "index",
                                "type": "uint8",
                                "internalType": "uint8"
                            },
                            {
                                "name": "start",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "length",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "typeId",
                                "type": "uint8",
                                "internalType": "uint8"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getUpdateArrayHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.Update[]",
                        "components": [
                            {
                                "name": "start",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "slice",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Slice",
                                "components": [
                                    {
                                        "name": "index",
                                        "type": "uint8",
                                        "internalType": "uint8"
                                    },
                                    {
                                        "name": "start",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "length",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "typeId",
                                        "type": "uint8",
                                        "internalType": "uint8"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getUpdateHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Update",
                        "components": [
                            {
                                "name": "start",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "slice",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Slice",
                                "components": [
                                    {
                                        "name": "index",
                                        "type": "uint8",
                                        "internalType": "uint8"
                                    },
                                    {
                                        "name": "start",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "length",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "typeId",
                                        "type": "uint8",
                                        "internalType": "uint8"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$typeHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "hash",
                "inputs": [
                    {
                        "name": "$livePlugs",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.LivePlugs",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$livePlugsHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "initialize",
                "inputs": [
                    {
                        "name": "$owner",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$oneClicker",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "name",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$name",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "oneClick",
                "inputs": [
                    {
                        "name": "$oneClickers",
                        "type": "address[]",
                        "internalType": "address[]"
                    },
                    {
                        "name": "$allowance",
                        "type": "bool[]",
                        "internalType": "bool[]"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "oneClickersToAllowed",
                "inputs": [
                    {
                        "name": "oneClicker",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "allowed",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "owner",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownershipHandoverExpiresAt",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "plug",
                "inputs": [
                    {
                        "name": "$livePlugs",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.LivePlugs",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "$solver",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plug",
                "inputs": [
                    {
                        "name": "$plugs",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Plugs",
                        "components": [
                            {
                                "name": "socket",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "plugs",
                                "type": "tuple[]",
                                "internalType": "struct PlugTypesLib.Plug[]",
                                "components": [
                                    {
                                        "name": "data",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "updates",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Update[]",
                                        "components": [
                                            {
                                                "name": "start",
                                                "type": "uint256",
                                                "internalType": "uint256"
                                            },
                                            {
                                                "name": "slice",
                                                "type": "tuple",
                                                "internalType": "struct PlugTypesLib.Slice",
                                                "components": [
                                                    {
                                                        "name": "index",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    },
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "length",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "typeId",
                                                        "type": "uint8",
                                                        "internalType": "uint8"
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "solver",
                                "type": "bytes",
                                "internalType": "bytes"
                            },
                            {
                                "name": "salt",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "proxiableUUID",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "renounceOwnership",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "requestOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "symbol",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$symbol",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "transferOwnership",
                "inputs": [
                    {
                        "name": "newOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "upgradeToAndCall",
                "inputs": [
                    {
                        "name": "newImplementation",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "version",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$version",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "event",
                "name": "OwnershipHandoverCanceled",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverRequested",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipTransferred",
                "inputs": [
                    {
                        "name": "oldOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "newOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "Upgraded",
                "inputs": [
                    {
                        "name": "implementation",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "AlreadyInitialized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "FnSelectorNotRecognized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NewOwnerIsZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NoHandoverRequest",
                "inputs": []
            },
            {
                "type": "error",
                "name": "PlugFailed",
                "inputs": [
                    {
                        "name": "$index",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$reason",
                        "type": "string",
                        "internalType": "string"
                    }
                ]
            },
            {
                "type": "error",
                "name": "Reentrancy",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Unauthorized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "UnauthorizedCallContext",
                "inputs": []
            },
            {
                "type": "error",
                "name": "UpgradeFailed",
                "inputs": []
            }
        ]
    },
    {
        "name": "PlugTicket",
        "abi": [
            {
                "type": "constructor",
                "inputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "approve",
                "inputs": [
                    {
                        "name": "account",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "balanceOf",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "baseURI",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "cancelOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "completeOwnershipHandover",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "getApproved",
                "inputs": [
                    {
                        "name": "id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "initialize",
                "inputs": [],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "isApprovedForAll",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "operator",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "mint",
                "inputs": [],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "name",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "owner",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownerOf",
                "inputs": [
                    {
                        "name": "id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownershipHandoverExpiresAt",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "renounceOwnership",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "requestOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "safeTransferFrom",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "safeTransferFrom",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "id",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "setApprovalForAll",
                "inputs": [
                    {
                        "name": "operator",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "isApproved",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setBaseURI",
                "inputs": [
                    {
                        "name": "$baseURI",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "supportsInterface",
                "inputs": [
                    {
                        "name": "interfaceId",
                        "type": "bytes4",
                        "internalType": "bytes4"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "symbol",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "tokenURI",
                "inputs": [
                    {
                        "name": "$id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "totalSupply",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "transferFrom",
                "inputs": [
                    {
                        "name": "$from",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$id",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "transferOwnership",
                "inputs": [
                    {
                        "name": "newOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "event",
                "name": "Approval",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "account",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "id",
                        "type": "uint256",
                        "indexed": true,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "ApprovalForAll",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "operator",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "isApproved",
                        "type": "bool",
                        "indexed": false,
                        "internalType": "bool"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "Initialized",
                "inputs": [
                    {
                        "name": "version",
                        "type": "uint64",
                        "indexed": false,
                        "internalType": "uint64"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverCanceled",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverRequested",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipTransferred",
                "inputs": [
                    {
                        "name": "oldOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "newOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "Transfer",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "to",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "id",
                        "type": "uint256",
                        "indexed": true,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "AccountBalanceOverflow",
                "inputs": []
            },
            {
                "type": "error",
                "name": "AlreadyInitialized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "AlreadyMinted",
                "inputs": []
            },
            {
                "type": "error",
                "name": "BalanceQueryForZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "CallerMustBeContract",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidInitialization",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NewOwnerIsZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NoHandoverRequest",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NonTransferableToken",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NotInitializing",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NotOwnerNorApproved",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TokenAlreadyExists",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TokenDoesNotExist",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TransferFromIncorrectOwner",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TransferToNonERC721ReceiverImplementer",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TransferToZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Unauthorized",
                "inputs": []
            }
        ]
    },
    {
        "name": "PlugToken",
        "abi": [
            {
                "type": "constructor",
                "inputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "DOMAIN_SEPARATOR",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "TOTAL_SUPPLY",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "allowance",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "spender",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "approve",
                "inputs": [
                    {
                        "name": "spender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "balanceOf",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "bridgeUnlock",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "cancelOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "completeOwnershipHandover",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "crosschainBurn",
                "inputs": [
                    {
                        "name": "_from",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "_amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "crosschainMint",
                "inputs": [
                    {
                        "name": "_to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "_amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "decimals",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "initialize",
                "inputs": [],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "name",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "nonces",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "owner",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownershipHandoverExpiresAt",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "permit",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "spender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "deadline",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "v",
                        "type": "uint8",
                        "internalType": "uint8"
                    },
                    {
                        "name": "r",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "s",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "renounceOwnership",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "requestOwnershipHandover",
                "inputs": [],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "senderToAllowed",
                "inputs": [
                    {
                        "name": "",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "setBridgeUnlock",
                "inputs": [
                    {
                        "name": "$unlock",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setSenderAllowed",
                "inputs": [
                    {
                        "name": "$sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$allowed",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "setTransferUnlock",
                "inputs": [
                    {
                        "name": "$unlock",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "supportsInterface",
                "inputs": [
                    {
                        "name": "_interfaceId",
                        "type": "bytes4",
                        "internalType": "bytes4"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "symbol",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "totalSupply",
                "inputs": [],
                "outputs": [
                    {
                        "name": "result",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "transfer",
                "inputs": [
                    {
                        "name": "to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "transferFrom",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "transferOwnership",
                "inputs": [
                    {
                        "name": "newOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "transferUnlock",
                "inputs": [],
                "outputs": [
                    {
                        "name": "",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "event",
                "name": "Approval",
                "inputs": [
                    {
                        "name": "owner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "spender",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "CrosschainBurn",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    },
                    {
                        "name": "sender",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "CrosschainMint",
                "inputs": [
                    {
                        "name": "to",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    },
                    {
                        "name": "sender",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "Initialized",
                "inputs": [
                    {
                        "name": "version",
                        "type": "uint64",
                        "indexed": false,
                        "internalType": "uint64"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverCanceled",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipHandoverRequested",
                "inputs": [
                    {
                        "name": "pendingOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "OwnershipTransferred",
                "inputs": [
                    {
                        "name": "oldOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "newOwner",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "event",
                "name": "Transfer",
                "inputs": [
                    {
                        "name": "from",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "to",
                        "type": "address",
                        "indexed": true,
                        "internalType": "address"
                    },
                    {
                        "name": "amount",
                        "type": "uint256",
                        "indexed": false,
                        "internalType": "uint256"
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "AllowanceOverflow",
                "inputs": []
            },
            {
                "type": "error",
                "name": "AllowanceUnderflow",
                "inputs": []
            },
            {
                "type": "error",
                "name": "AlreadyInitialized",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InsufficientAllowance",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InsufficientBalance",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidInitialization",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidPermit",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NewOwnerIsZeroAddress",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NoHandoverRequest",
                "inputs": []
            },
            {
                "type": "error",
                "name": "NotInitializing",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Permit2AllowanceIsFixedAtInfinity",
                "inputs": []
            },
            {
                "type": "error",
                "name": "PermitExpired",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TotalSupplyOverflow",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Unauthorized",
                "inputs": []
            }
        ]
    },
    {
        "name": "Plug",
        "abi": [
            {
                "type": "function",
                "name": "name",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$name",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "plug",
                "inputs": [
                    {
                        "name": "$livePlugs",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.LivePlugs",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plug",
                "inputs": [
                    {
                        "name": "$livePlugs",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.LivePlugs[]",
                        "components": [
                            {
                                "name": "plugs",
                                "type": "tuple",
                                "internalType": "struct PlugTypesLib.Plugs",
                                "components": [
                                    {
                                        "name": "socket",
                                        "type": "address",
                                        "internalType": "address"
                                    },
                                    {
                                        "name": "plugs",
                                        "type": "tuple[]",
                                        "internalType": "struct PlugTypesLib.Plug[]",
                                        "components": [
                                            {
                                                "name": "data",
                                                "type": "bytes",
                                                "internalType": "bytes"
                                            },
                                            {
                                                "name": "updates",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Update[]",
                                                "components": [
                                                    {
                                                        "name": "start",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "slice",
                                                        "type": "tuple",
                                                        "internalType": "struct PlugTypesLib.Slice",
                                                        "components": [
                                                            {
                                                                "name": "index",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            },
                                                            {
                                                                "name": "start",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "length",
                                                                "type": "uint256",
                                                                "internalType": "uint256"
                                                            },
                                                            {
                                                                "name": "typeId",
                                                                "type": "uint8",
                                                                "internalType": "uint8"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "solver",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes",
                                        "internalType": "bytes"
                                    }
                                ]
                            },
                            {
                                "name": "signature",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "symbol",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$version",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "event",
                "name": "PlugResult",
                "inputs": [
                    {
                        "name": "index",
                        "type": "uint8",
                        "indexed": false,
                        "internalType": "uint8"
                    },
                    {
                        "name": "plugsHash",
                        "type": "bytes32",
                        "indexed": false,
                        "internalType": "bytes32"
                    },
                    {
                        "name": "reason",
                        "type": "tuple",
                        "indexed": false,
                        "internalType": "struct PlugTypesLib.Result",
                        "components": [
                            {
                                "name": "index",
                                "type": "uint8",
                                "internalType": "uint8"
                            },
                            {
                                "name": "error",
                                "type": "string",
                                "internalType": "string"
                            }
                        ]
                    }
                ],
                "anonymous": false
            },
            {
                "type": "error",
                "name": "SocketAddressInvalid",
                "inputs": [
                    {
                        "name": "$intended",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$socket",
                        "type": "address",
                        "internalType": "address"
                    }
                ]
            }
        ]
    }
] as const