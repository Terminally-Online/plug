export const contracts = [
    {
        "name": "PlugBaseFeeFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "outputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            }
        ]
    },
    {
        "name": "PlugBlockNumberFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "outputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            }
        ]
    },
    {
        "name": "PlugClampFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$min",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$max",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$min",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$max",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
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
                        "name": "$implementation",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$admin",
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
                        "name": "$alreadyDeployed",
                        "type": "bool",
                        "internalType": "bool"
                    },
                    {
                        "name": "$vault",
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
            }
        ]
    },
    {
        "name": "PlugLimitedCallsFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$terms",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$callCount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$callCount",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "$terms",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "$plugsHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "nonpayable"
            }
        ]
    },
    {
        "name": "PlugNounsIdFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
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
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$value",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            }
        ]
    },
    {
        "name": "PlugNounsTraitFuse",
        "abi": [
            {
                "type": "constructor",
                "inputs": [
                    {
                        "name": "$art",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "ACCESSORY_SELECTOR",
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
                "name": "BACKGROUND_SELECTOR",
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
                "name": "BODY_SELECTOR",
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
                "name": "GLASSES_SELECTOR",
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
                "name": "HEAD_SELECTOR",
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
                "name": "decode",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$selector",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "$trait",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$selector",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "$trait",
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
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "nounTrait",
                "inputs": [
                    {
                        "name": "$selector",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$traitHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
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
                "name": "setArt",
                "inputs": [
                    {
                        "name": "$art",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
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
                "type": "error",
                "name": "AlreadyInitialized",
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
                "name": "Unauthorized",
                "inputs": []
            }
        ]
    },
    {
        "name": "PlugRevocationFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$sender",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$sender",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "$plugsHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "isRevoked",
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
                "name": "revoke",
                "inputs": [
                    {
                        "name": "$plugsHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "$revoked",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            }
        ]
    },
    {
        "name": "PlugTimestampFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$operator",
                        "type": "uint128",
                        "internalType": "uint128"
                    },
                    {
                        "name": "$threshold",
                        "type": "uint128",
                        "internalType": "uint128"
                    }
                ],
                "outputs": [
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            }
        ]
    },
    {
        "name": "PlugWindowFuse",
        "abi": [
            {
                "type": "function",
                "name": "decode",
                "inputs": [
                    {
                        "name": "$schedule",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "$startTime",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$repeatsEvery",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$duration",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$daysOfWeek",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "encode",
                "inputs": [
                    {
                        "name": "$startTime",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$repeatsEvery",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$duration",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$daysOfWeek",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "outputs": [
                    {
                        "name": "$schedule",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "enforceFuse",
                "inputs": [
                    {
                        "name": "$live",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$current",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Current",
                        "components": [
                            {
                                "name": "target",
                                "type": "address",
                                "internalType": "address"
                            },
                            {
                                "name": "value",
                                "type": "uint256",
                                "internalType": "uint256"
                            },
                            {
                                "name": "data",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    },
                    {
                        "name": "",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$through",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "isWithinWindow",
                "inputs": [
                    {
                        "name": "$startTime",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$repeatsEvery",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$duration",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$daysOfWeek",
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
                "name": "isWithinWindow",
                "inputs": [
                    {
                        "name": "$schedule",
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
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "toWindow",
                "inputs": [
                    {
                        "name": "$schedule",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "$window",
                        "type": "tuple",
                        "internalType": "struct WindowFuseLib.Window",
                        "components": [
                            {
                                "name": "periods",
                                "type": "tuple[]",
                                "internalType": "struct WindowFuseLib.Period[]",
                                "components": [
                                    {
                                        "name": "startTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    },
                                    {
                                        "name": "endTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toWindow",
                "inputs": [
                    {
                        "name": "$startTime",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$duration",
                        "type": "uint32",
                        "internalType": "uint32"
                    },
                    {
                        "name": "$daysOfWeek",
                        "type": "uint8",
                        "internalType": "uint8"
                    }
                ],
                "outputs": [
                    {
                        "name": "$window",
                        "type": "tuple",
                        "internalType": "struct WindowFuseLib.Window",
                        "components": [
                            {
                                "name": "periods",
                                "type": "tuple[]",
                                "internalType": "struct WindowFuseLib.Period[]",
                                "components": [
                                    {
                                        "name": "startTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    },
                                    {
                                        "name": "endTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "toWindows",
                "inputs": [
                    {
                        "name": "$schedule",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$n",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "outputs": [
                    {
                        "name": "$windows",
                        "type": "tuple[]",
                        "internalType": "struct WindowFuseLib.Window[]",
                        "components": [
                            {
                                "name": "periods",
                                "type": "tuple[]",
                                "internalType": "struct WindowFuseLib.Period[]",
                                "components": [
                                    {
                                        "name": "startTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    },
                                    {
                                        "name": "endTime",
                                        "type": "uint32",
                                        "internalType": "uint32"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "$cursor",
                        "type": "uint32",
                        "internalType": "uint32"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "error",
                "name": "WindowCaveatViolation",
                "inputs": []
            },
            {
                "type": "error",
                "name": "WindowLackingDays",
                "inputs": []
            },
            {
                "type": "error",
                "name": "WindowLackingDuration",
                "inputs": []
            },
            {
                "type": "error",
                "name": "WindowLackingStartTime",
                "inputs": []
            },
            {
                "type": "error",
                "name": "WindowLackingSufficientRepeatsEvery",
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
                                                "name": "current",
                                                "type": "tuple",
                                                "internalType": "struct PlugTypesLib.Current",
                                                "components": [
                                                    {
                                                        "name": "target",
                                                        "type": "address",
                                                        "internalType": "address"
                                                    },
                                                    {
                                                        "name": "value",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "data",
                                                        "type": "bytes",
                                                        "internalType": "bytes"
                                                    }
                                                ]
                                            },
                                            {
                                                "name": "fuses",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Fuse[]",
                                                "components": [
                                                    {
                                                        "name": "target",
                                                        "type": "address",
                                                        "internalType": "address"
                                                    },
                                                    {
                                                        "name": "data",
                                                        "type": "bytes",
                                                        "internalType": "bytes"
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes32",
                                        "internalType": "bytes32"
                                    },
                                    {
                                        "name": "fee",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "maxFeePerGas",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "maxPriorityFeePerGas",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "executor",
                                        "type": "address",
                                        "internalType": "address"
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
                        "name": "$results",
                        "type": "bytes[]",
                        "internalType": "bytes[]"
                    }
                ],
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
                                                "name": "current",
                                                "type": "tuple",
                                                "internalType": "struct PlugTypesLib.Current",
                                                "components": [
                                                    {
                                                        "name": "target",
                                                        "type": "address",
                                                        "internalType": "address"
                                                    },
                                                    {
                                                        "name": "value",
                                                        "type": "uint256",
                                                        "internalType": "uint256"
                                                    },
                                                    {
                                                        "name": "data",
                                                        "type": "bytes",
                                                        "internalType": "bytes"
                                                    }
                                                ]
                                            },
                                            {
                                                "name": "fuses",
                                                "type": "tuple[]",
                                                "internalType": "struct PlugTypesLib.Fuse[]",
                                                "components": [
                                                    {
                                                        "name": "target",
                                                        "type": "address",
                                                        "internalType": "address"
                                                    },
                                                    {
                                                        "name": "data",
                                                        "type": "bytes",
                                                        "internalType": "bytes"
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "name": "salt",
                                        "type": "bytes32",
                                        "internalType": "bytes32"
                                    },
                                    {
                                        "name": "fee",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "maxFeePerGas",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "maxPriorityFeePerGas",
                                        "type": "uint256",
                                        "internalType": "uint256"
                                    },
                                    {
                                        "name": "executor",
                                        "type": "address",
                                        "internalType": "address"
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
                        "name": "$results",
                        "type": "bytes[][]",
                        "internalType": "bytes[][]"
                    }
                ],
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
            }
        ]
    }
] as const