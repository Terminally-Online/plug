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
                "type": "constructor",
                "inputs": [
                    {
                        "name": "$owner",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$baseURI",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
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
                "name": "deploy",
                "inputs": [
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
                        "name": "$socket",
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
                "name": "implementations",
                "inputs": [
                    {
                        "name": "",
                        "type": "uint16",
                        "internalType": "uint16"
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
                "name": "setImplementation",
                "inputs": [
                    {
                        "name": "$version",
                        "type": "uint16",
                        "internalType": "uint16"
                    },
                    {
                        "name": "$implementation",
                        "type": "address",
                        "internalType": "address"
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
                        "name": "$symbol",
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
                        "name": "$tokenId",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "outputs": [
                    {
                        "name": "$uri",
                        "type": "string",
                        "internalType": "string"
                    }
                ],
                "stateMutability": "view"
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
                "name": "BalanceQueryForZeroAddress",
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
                        "name": "$to",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "$value",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "$data",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [],
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
                "name": "NotInitializing",
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
        "name": "PlugVaultSocket",
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
                "name": "SET_IMAGE_HASH_TYPE_HASH",
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
                "name": "getCurrentHash",
                "inputs": [
                    {
                        "name": "$input",
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
                    }
                ],
                "outputs": [
                    {
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
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
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getFuseArrayHash",
                "inputs": [
                    {
                        "name": "$input",
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
                ],
                "outputs": [
                    {
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "getFuseHash",
                "inputs": [
                    {
                        "name": "$input",
                        "type": "tuple",
                        "internalType": "struct PlugTypesLib.Fuse",
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
                ],
                "outputs": [
                    {
                        "name": "$hash",
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
                                        "name": "solver",
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
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
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
                    }
                ],
                "outputs": [
                    {
                        "name": "$hash",
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
                    }
                ],
                "outputs": [
                    {
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
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
                                "name": "solver",
                                "type": "bytes",
                                "internalType": "bytes"
                            }
                        ]
                    }
                ],
                "outputs": [
                    {
                        "name": "$hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "stateMutability": "pure"
            },
            {
                "type": "function",
                "name": "imageHash",
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
                "name": "initialize",
                "inputs": [
                    {
                        "name": "$ownership",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "isValidSignature",
                "inputs": [
                    {
                        "name": "_hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "_signatures",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes4",
                        "internalType": "bytes4"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "isValidSignature",
                "inputs": [
                    {
                        "name": "_data",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "_signatures",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "",
                        "type": "bytes4",
                        "internalType": "bytes4"
                    }
                ],
                "stateMutability": "view"
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
                "name": "owner",
                "inputs": [],
                "outputs": [
                    {
                        "name": "$owner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "ownership",
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
                                        "name": "solver",
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
                    },
                    {
                        "name": "$gas",
                        "type": "uint256",
                        "internalType": "uint256"
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
                                "name": "solver",
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
                "name": "signatureRecovery",
                "inputs": [
                    {
                        "name": "_digest",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ],
                "outputs": [
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "weight",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "imageHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "subdigest",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "checkpoint",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ],
                "stateMutability": "view"
            },
            {
                "type": "function",
                "name": "supportsInterface",
                "inputs": [
                    {
                        "name": "_interfaceID",
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
                "stateMutability": "pure"
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
                        "name": "$newOwner",
                        "type": "address",
                        "internalType": "address"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
            },
            {
                "type": "function",
                "name": "updateImageHash",
                "inputs": [
                    {
                        "name": "_imageHash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ],
                "outputs": [],
                "stateMutability": "nonpayable"
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
                "name": "ImageHashUpdated",
                "inputs": [
                    {
                        "name": "newImageHash",
                        "type": "bytes32",
                        "indexed": false,
                        "internalType": "bytes32"
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
                "name": "EmptySignature",
                "inputs": []
            },
            {
                "type": "error",
                "name": "ImageHashIsZero",
                "inputs": []
            },
            {
                "type": "error",
                "name": "InvalidNestedSignature",
                "inputs": [
                    {
                        "name": "_hash",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    },
                    {
                        "name": "_addr",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ]
            },
            {
                "type": "error",
                "name": "InvalidSValue",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "_s",
                        "type": "bytes32",
                        "internalType": "bytes32"
                    }
                ]
            },
            {
                "type": "error",
                "name": "InvalidSignatureFlag",
                "inputs": [
                    {
                        "name": "_flag",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ]
            },
            {
                "type": "error",
                "name": "InvalidSignatureLength",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ]
            },
            {
                "type": "error",
                "name": "InvalidSignatureType",
                "inputs": [
                    {
                        "name": "_type",
                        "type": "bytes1",
                        "internalType": "bytes1"
                    }
                ]
            },
            {
                "type": "error",
                "name": "InvalidVValue",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "_v",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ]
            },
            {
                "type": "error",
                "name": "LowWeightChainedSignature",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "threshold",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "_weight",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ]
            },
            {
                "type": "error",
                "name": "OnlySelfAuth",
                "inputs": [
                    {
                        "name": "_sender",
                        "type": "address",
                        "internalType": "address"
                    },
                    {
                        "name": "_self",
                        "type": "address",
                        "internalType": "address"
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
                "name": "SignerIsAddress0",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    }
                ]
            },
            {
                "type": "error",
                "name": "UnauthorizedCallContext",
                "inputs": []
            },
            {
                "type": "error",
                "name": "UnsupportedSignatureType",
                "inputs": [
                    {
                        "name": "_signature",
                        "type": "bytes",
                        "internalType": "bytes"
                    },
                    {
                        "name": "_type",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "_recoverMode",
                        "type": "bool",
                        "internalType": "bool"
                    }
                ]
            },
            {
                "type": "error",
                "name": "UpgradeFailed",
                "inputs": []
            },
            {
                "type": "error",
                "name": "WrongChainedCheckpointOrder",
                "inputs": [
                    {
                        "name": "_current",
                        "type": "uint256",
                        "internalType": "uint256"
                    },
                    {
                        "name": "_prev",
                        "type": "uint256",
                        "internalType": "uint256"
                    }
                ]
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
                                        "name": "solver",
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
                                        "name": "solver",
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
                        "type": "bytes[][]",
                        "internalType": "bytes[][]"
                    }
                ],
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
    }
] as const