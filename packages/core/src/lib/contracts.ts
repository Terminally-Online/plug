export const contracts = [
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
                "outputs": [
                    {
                        "name": "$results",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.Result[]",
                        "components": [
                            {
                                "name": "success",
                                "type": "bool",
                                "internalType": "bool"
                            },
                            {
                                "name": "result",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
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
                        "name": "$results",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.Result[]",
                        "components": [
                            {
                                "name": "success",
                                "type": "bool",
                                "internalType": "bool"
                            },
                            {
                                "name": "result",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
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
                "name": "PlugsExecuted",
                "inputs": [
                    {
                        "name": "$plugsHash",
                        "type": "bytes32",
                        "indexed": true,
                        "internalType": "bytes32"
                    },
                    {
                        "name": "$results",
                        "type": "tuple[]",
                        "indexed": false,
                        "internalType": "struct PlugTypesLib.Result[]",
                        "components": [
                            {
                                "name": "success",
                                "type": "bool",
                                "internalType": "bool"
                            },
                            {
                                "name": "result",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
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
                "name": "NonceInvalid",
                "inputs": []
            },
            {
                "type": "error",
                "name": "PlugFailed",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Reentrancy",
                "inputs": []
            },
            {
                "type": "error",
                "name": "SenderInvalid",
                "inputs": [
                    {
                        "name": "$reality",
                        "type": "address",
                        "internalType": "address"
                    }
                ]
            },
            {
                "type": "error",
                "name": "SignatureInvalid",
                "inputs": []
            },
            {
                "type": "error",
                "name": "SolverExpired",
                "inputs": []
            },
            {
                "type": "error",
                "name": "SolverInvalid",
                "inputs": [
                    {
                        "name": "$expected",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$reality",
                        "type": "address",
                        "internalType": "address"
                    }
                ]
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
            },
            {
                "type": "error",
                "name": "ValueInvalid",
                "inputs": [
                    {
                        "name": "$recipient",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$expected",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$reality",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ]
            }
        ]
    },
    {
        "name": "PlugTreasury",
        "abi": [
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
                "name": "execute",
                "inputs": [
                    {
                        "name": "$targets",
                        "type": "address[]",
                        "internalType": "address[]"
                    },
                    {
                        "name": "$values",
                        "type": "uint256[]",
                        "internalType": "uint256[]"
                    },
                    {
                        "name": "$datas",
                        "type": "bytes[]",
                        "internalType": "bytes[]"
                    }
                ],
                "outputs": [
                    {
                        "name": "$successes",
                        "type": "bool[]",
                        "internalType": "bool[]"
                    },
                    {
                        "name": "$results",
                        "type": "bytes[]",
                        "internalType": "bytes[]"
                    }
                ],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "initialize",
                "inputs": [
                    {
                        "name": "$owner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
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
                "name": "plugNative",
                "inputs": [
                    {
                        "name": "$target",
                        "type": "address",
                        "internalType": "address payable"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$fee",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plugNativeToToken",
                "inputs": [
                    {
                        "name": "$tokenIn",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$target",
                        "type": "address",
                        "internalType": "address payable"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$fee",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plugToken",
                "inputs": [
                    {
                        "name": "$tokenOut",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$target",
                        "type": "address",
                        "internalType": "address payable"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$sell",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$fee",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plugTokenToNative",
                "inputs": [
                    {
                        "name": "$tokenOut",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$target",
                        "type": "address",
                        "internalType": "address payable"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$sell",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$fee",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
            },
            {
                "type": "function",
                "name": "plugTokenToToken",
                "inputs": [
                    {
                        "name": "$tokenOut",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$tokenIn",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$target",
                        "type": "address",
                        "internalType": "address payable"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "$sell",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$fee",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [],
                "stateMutability": "payable"
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
                "name": "setTargetsAllowed",
                "inputs": [
                    {
                        "name": "$targets",
                        "type": "address[]",
                        "internalType": "address[]"
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
                "name": "targetToAllowed",
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
                "name": "PlugFailed",
                "inputs": []
            },
            {
                "type": "error",
                "name": "Reentrancy",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TargetInvalid",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TokenAllowanceInvalid",
                "inputs": []
            },
            {
                "type": "error",
                "name": "TokenBalanceInvalid",
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
                        "name": "$results",
                        "type": "tuple[][]",
                        "internalType": "struct PlugTypesLib.Result[][]",
                        "components": [
                            {
                                "name": "success",
                                "type": "bool",
                                "internalType": "bool"
                            },
                            {
                                "name": "result",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
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
                        "name": "$results",
                        "type": "tuple[]",
                        "internalType": "struct PlugTypesLib.Result[]",
                        "components": [
                            {
                                "name": "success",
                                "type": "bool",
                                "internalType": "bool"
                            },
                            {
                                "name": "result",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
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