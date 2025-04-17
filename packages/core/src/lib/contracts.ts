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