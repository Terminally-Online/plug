import {
  createUseReadContract,
  createUseWriteContract,
  createUseSimulateContract,
  createUseWatchContractEvent,
} from 'wagmi/codegen'

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Plug
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugAbi = [
  {
    type: 'function',
    inputs: [],
    name: 'name',
    outputs: [{ name: '$name', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$livePlugs',
        internalType: 'struct PlugTypesLib.LivePlugs',
        type: 'tuple',
        components: [
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plugs',
            type: 'tuple',
            components: [
              { name: 'socket', internalType: 'address', type: 'address' },
              {
                name: 'plugs',
                internalType: 'struct PlugTypesLib.Plug[]',
                type: 'tuple[]',
                components: [
                  {
                    name: 'current',
                    internalType: 'struct PlugTypesLib.Current',
                    type: 'tuple',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      {
                        name: 'value',
                        internalType: 'uint256',
                        type: 'uint256',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                  {
                    name: 'fuses',
                    internalType: 'struct PlugTypesLib.Fuse[]',
                    type: 'tuple[]',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                ],
              },
              { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
              { name: 'fee', internalType: 'uint256', type: 'uint256' },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [{ name: '$results', internalType: 'bytes[]', type: 'bytes[]' }],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$livePlugs',
        internalType: 'struct PlugTypesLib.LivePlugs[]',
        type: 'tuple[]',
        components: [
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plugs',
            type: 'tuple',
            components: [
              { name: 'socket', internalType: 'address', type: 'address' },
              {
                name: 'plugs',
                internalType: 'struct PlugTypesLib.Plug[]',
                type: 'tuple[]',
                components: [
                  {
                    name: 'current',
                    internalType: 'struct PlugTypesLib.Current',
                    type: 'tuple',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      {
                        name: 'value',
                        internalType: 'uint256',
                        type: 'uint256',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                  {
                    name: 'fuses',
                    internalType: 'struct PlugTypesLib.Fuse[]',
                    type: 'tuple[]',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                ],
              },
              { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
              { name: 'fee', internalType: 'uint256', type: 'uint256' },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [
      { name: '$results', internalType: 'bytes[][]', type: 'bytes[][]' },
    ],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '$version', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'error',
    inputs: [
      { name: '$intended', internalType: 'address', type: 'address' },
      { name: '$socket', internalType: 'address', type: 'address' },
    ],
    name: 'SocketAddressInvalid',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugBaseFeeFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugBaseFeeFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'encode',
    outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdExceeded',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdInsufficient',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugBlockNumberFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugBlockNumberFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'encode',
    outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdExceeded',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdInsufficient',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugClampFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugClampFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [
      { name: '$min', internalType: 'uint256', type: 'uint256' },
      { name: '$max', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$min', internalType: 'uint256', type: 'uint256' },
      { name: '$max', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'encode',
    outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugFactory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugFactoryAbi = [
  { type: 'constructor', inputs: [], stateMutability: 'nonpayable' },
  {
    type: 'function',
    inputs: [
      { name: 'account', internalType: 'address', type: 'address' },
      { name: 'id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'approve',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: 'result', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'cancelOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'completeOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: '$salt', internalType: 'bytes32', type: 'bytes32' },
      { name: '$router', internalType: 'address', type: 'address' },
    ],
    name: 'deploy',
    outputs: [
      { name: '$alreadyDeployed', internalType: 'bool', type: 'bool' },
      { name: '$socket', internalType: 'address', type: 'address' },
    ],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: '$implementation', internalType: 'address', type: 'address' },
      { name: '$salt', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'getAddress',
    outputs: [{ name: '$vault', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'id', internalType: 'uint256', type: 'uint256' }],
    name: 'getApproved',
    outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
    name: 'implementations',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$implementation', internalType: 'address', type: 'address' },
    ],
    name: 'initCodeHash',
    outputs: [
      { name: '$initCodeHash', internalType: 'bytes32', type: 'bytes32' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$owner', internalType: 'address', type: 'address' },
      { name: '$baseURI', internalType: 'string', type: 'string' },
      { name: '$implementation', internalType: 'address', type: 'address' },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'operator', internalType: 'address', type: 'address' },
    ],
    name: 'isApprovedForAll',
    outputs: [{ name: 'result', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'name',
    outputs: [{ name: '$name', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'id', internalType: 'uint256', type: 'uint256' }],
    name: 'ownerOf',
    outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'ownershipHandoverExpiresAt',
    outputs: [{ name: 'result', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'requestOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'safeTransferFrom',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'id', internalType: 'uint256', type: 'uint256' },
      { name: 'data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'safeTransferFrom',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'operator', internalType: 'address', type: 'address' },
      { name: 'isApproved', internalType: 'bool', type: 'bool' },
    ],
    name: 'setApprovalForAll',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: '$baseURI', internalType: 'string', type: 'string' }],
    name: 'setBaseURI',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: '$version', internalType: 'uint16', type: 'uint16' },
      { name: '$implementation', internalType: 'address', type: 'address' },
    ],
    name: 'setImplementation',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'interfaceId', internalType: 'bytes4', type: 'bytes4' }],
    name: 'supportsInterface',
    outputs: [{ name: 'result', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '$symbol', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [{ name: '$tokenId', internalType: 'uint256', type: 'uint256' }],
    name: 'tokenURI',
    outputs: [{ name: '$uri', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transferFrom',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'account',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: true },
    ],
    name: 'Approval',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'operator',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'isApproved',
        internalType: 'bool',
        type: 'bool',
        indexed: false,
      },
    ],
    name: 'ApprovalForAll',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverCanceled',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverRequested',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'oldOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: true },
    ],
    name: 'Transfer',
  },
  { type: 'error', inputs: [], name: 'AccountBalanceOverflow' },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  { type: 'error', inputs: [], name: 'BalanceQueryForZeroAddress' },
  {
    type: 'error',
    inputs: [{ name: '$version', internalType: 'uint16', type: 'uint16' }],
    name: 'ImplementationAlreadyInitialized',
  },
  {
    type: 'error',
    inputs: [{ name: '$version', internalType: 'uint16', type: 'uint16' }],
    name: 'ImplementationInvalid',
  },
  { type: 'error', inputs: [], name: 'NewOwnerIsZeroAddress' },
  { type: 'error', inputs: [], name: 'NoHandoverRequest' },
  { type: 'error', inputs: [], name: 'NotOwnerNorApproved' },
  { type: 'error', inputs: [], name: 'TokenAlreadyExists' },
  { type: 'error', inputs: [], name: 'TokenDoesNotExist' },
  { type: 'error', inputs: [], name: 'TransferFromIncorrectOwner' },
  { type: 'error', inputs: [], name: 'TransferToNonERC721ReceiverImplementer' },
  { type: 'error', inputs: [], name: 'TransferToZeroAddress' },
  { type: 'error', inputs: [], name: 'Unauthorized' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugLimitedCallsFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugLimitedCallsFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$terms', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [{ name: '$callCount', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [{ name: '$callCount', internalType: 'uint256', type: 'uint256' }],
    name: 'encode',
    outputs: [{ name: '$terms', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'nonpayable',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugNounsIdFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugNounsIdFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$live', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '$value', internalType: 'uint256', type: 'uint256' }],
    name: 'encode',
    outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugNounsTraitFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugNounsTraitFuseAbi = [
  {
    type: 'constructor',
    inputs: [{ name: '$art', internalType: 'address', type: 'address' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'ACCESSORY_SELECTOR',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'BACKGROUND_SELECTOR',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'BODY_SELECTOR',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'GLASSES_SELECTOR',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'HEAD_SELECTOR',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'cancelOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'completeOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: '$live', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [
      { name: '$selector', internalType: 'bytes32', type: 'bytes32' },
      { name: '$trait', internalType: 'bytes32', type: 'bytes32' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$selector', internalType: 'bytes32', type: 'bytes32' },
      { name: '$trait', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'encode',
    outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '$selector', internalType: 'bytes32', type: 'bytes32' }],
    name: 'nounTrait',
    outputs: [{ name: '$traitHash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'ownershipHandoverExpiresAt',
    outputs: [{ name: 'result', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'requestOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: '$art', internalType: 'address', type: 'address' }],
    name: 'setArt',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverCanceled',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverRequested',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'oldOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  { type: 'error', inputs: [], name: 'NewOwnerIsZeroAddress' },
  { type: 'error', inputs: [], name: 'NoHandoverRequest' },
  { type: 'error', inputs: [], name: 'Unauthorized' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugRevocationFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugRevocationFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [{ name: '$sender', internalType: 'address', type: 'address' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [{ name: '$sender', internalType: 'address', type: 'address' }],
    name: 'encode',
    outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'isRevoked',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$plugsHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '$revoked', internalType: 'bool', type: 'bool' },
    ],
    name: 'revoke',
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugTimestampFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugTimestampFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    name: 'decode',
    outputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$operator', internalType: 'uint8', type: 'uint8' },
      { name: '$threshold', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'encode',
    outputs: [{ name: '$data', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdExceeded',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'uint256', type: 'uint256' },
      { name: '$reality', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ThresholdInsufficient',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugTreasury
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugTreasuryAbi = [
  { type: 'fallback', stateMutability: 'payable' },
  { type: 'receive', stateMutability: 'payable' },
  {
    type: 'function',
    inputs: [],
    name: 'cancelOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'completeOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      { name: '$to', internalType: 'address', type: 'address' },
      { name: '$value', internalType: 'uint256', type: 'uint256' },
      { name: '$data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'execute',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: '$owner', internalType: 'address', type: 'address' }],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: 'result', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'pendingOwner', internalType: 'address', type: 'address' },
    ],
    name: 'ownershipHandoverExpiresAt',
    outputs: [{ name: 'result', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'requestOwnershipHandover',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverCanceled',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'pendingOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipHandoverRequested',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'oldOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  { type: 'error', inputs: [], name: 'NewOwnerIsZeroAddress' },
  { type: 'error', inputs: [], name: 'NoHandoverRequest' },
  { type: 'error', inputs: [], name: 'Unauthorized' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugVaultSocket
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugVaultSocketAbi = [
  { type: 'constructor', inputs: [], stateMutability: 'nonpayable' },
  { type: 'fallback', stateMutability: 'payable' },
  { type: 'receive', stateMutability: 'payable' },
  {
    type: 'function',
    inputs: [],
    name: 'SET_IMAGE_HASH_TYPE_HASH',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getCurrentHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.EIP712Domain',
        type: 'tuple',
        components: [
          { name: 'name', internalType: 'string', type: 'string' },
          { name: 'version', internalType: 'string', type: 'string' },
          { name: 'chainId', internalType: 'uint256', type: 'uint256' },
          {
            name: 'verifyingContract',
            internalType: 'address',
            type: 'address',
          },
        ],
      },
    ],
    name: 'getEIP712DomainHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Fuse[]',
        type: 'tuple[]',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getFuseArrayHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Fuse',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getFuseHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.LivePlugs',
        type: 'tuple',
        components: [
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plugs',
            type: 'tuple',
            components: [
              { name: 'socket', internalType: 'address', type: 'address' },
              {
                name: 'plugs',
                internalType: 'struct PlugTypesLib.Plug[]',
                type: 'tuple[]',
                components: [
                  {
                    name: 'current',
                    internalType: 'struct PlugTypesLib.Current',
                    type: 'tuple',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      {
                        name: 'value',
                        internalType: 'uint256',
                        type: 'uint256',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                  {
                    name: 'fuses',
                    internalType: 'struct PlugTypesLib.Fuse[]',
                    type: 'tuple[]',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                ],
              },
              { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
              { name: 'fee', internalType: 'uint256', type: 'uint256' },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getLivePlugsHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Plug[]',
        type: 'tuple[]',
        components: [
          {
            name: 'current',
            internalType: 'struct PlugTypesLib.Current',
            type: 'tuple',
            components: [
              { name: 'target', internalType: 'address', type: 'address' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
            ],
          },
          {
            name: 'fuses',
            internalType: 'struct PlugTypesLib.Fuse[]',
            type: 'tuple[]',
            components: [
              { name: 'target', internalType: 'address', type: 'address' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
            ],
          },
        ],
      },
    ],
    name: 'getPlugArrayHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Plug',
        type: 'tuple',
        components: [
          {
            name: 'current',
            internalType: 'struct PlugTypesLib.Current',
            type: 'tuple',
            components: [
              { name: 'target', internalType: 'address', type: 'address' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
            ],
          },
          {
            name: 'fuses',
            internalType: 'struct PlugTypesLib.Fuse[]',
            type: 'tuple[]',
            components: [
              { name: 'target', internalType: 'address', type: 'address' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
            ],
          },
        ],
      },
    ],
    name: 'getPlugHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Plugs',
        type: 'tuple',
        components: [
          { name: 'socket', internalType: 'address', type: 'address' },
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plug[]',
            type: 'tuple[]',
            components: [
              {
                name: 'current',
                internalType: 'struct PlugTypesLib.Current',
                type: 'tuple',
                components: [
                  { name: 'target', internalType: 'address', type: 'address' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                ],
              },
              {
                name: 'fuses',
                internalType: 'struct PlugTypesLib.Fuse[]',
                type: 'tuple[]',
                components: [
                  { name: 'target', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                ],
              },
            ],
          },
          { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
          { name: 'fee', internalType: 'uint256', type: 'uint256' },
          { name: 'solver', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getPlugsHash',
    outputs: [{ name: '$hash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [],
    name: 'imageHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$ownership', internalType: 'address', type: 'address' },
      { name: '$router', internalType: 'address', type: 'address' },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: '_hash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_signatures', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'isValidSignature',
    outputs: [{ name: '', internalType: 'bytes4', type: 'bytes4' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_signatures', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'isValidSignature',
    outputs: [{ name: '', internalType: 'bytes4', type: 'bytes4' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'name',
    outputs: [{ name: '$name', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '$owner', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'ownership',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$livePlugs',
        internalType: 'struct PlugTypesLib.LivePlugs',
        type: 'tuple',
        components: [
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plugs',
            type: 'tuple',
            components: [
              { name: 'socket', internalType: 'address', type: 'address' },
              {
                name: 'plugs',
                internalType: 'struct PlugTypesLib.Plug[]',
                type: 'tuple[]',
                components: [
                  {
                    name: 'current',
                    internalType: 'struct PlugTypesLib.Current',
                    type: 'tuple',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      {
                        name: 'value',
                        internalType: 'uint256',
                        type: 'uint256',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                  {
                    name: 'fuses',
                    internalType: 'struct PlugTypesLib.Fuse[]',
                    type: 'tuple[]',
                    components: [
                      {
                        name: 'target',
                        internalType: 'address',
                        type: 'address',
                      },
                      { name: 'data', internalType: 'bytes', type: 'bytes' },
                    ],
                  },
                ],
              },
              { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
              { name: 'fee', internalType: 'uint256', type: 'uint256' },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '$solver', internalType: 'address', type: 'address' },
      { name: '$gas', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'plug',
    outputs: [{ name: '$results', internalType: 'bytes[]', type: 'bytes[]' }],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$plugs',
        internalType: 'struct PlugTypesLib.Plugs',
        type: 'tuple',
        components: [
          { name: 'socket', internalType: 'address', type: 'address' },
          {
            name: 'plugs',
            internalType: 'struct PlugTypesLib.Plug[]',
            type: 'tuple[]',
            components: [
              {
                name: 'current',
                internalType: 'struct PlugTypesLib.Current',
                type: 'tuple',
                components: [
                  { name: 'target', internalType: 'address', type: 'address' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                ],
              },
              {
                name: 'fuses',
                internalType: 'struct PlugTypesLib.Fuse[]',
                type: 'tuple[]',
                components: [
                  { name: 'target', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                ],
              },
            ],
          },
          { name: 'salt', internalType: 'bytes32', type: 'bytes32' },
          { name: 'fee', internalType: 'uint256', type: 'uint256' },
          { name: 'solver', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [{ name: '$results', internalType: 'bytes[]', type: 'bytes[]' }],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'proxiableUUID',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'router',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '_digest', internalType: 'bytes32', type: 'bytes32' },
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'signatureRecovery',
    outputs: [
      { name: 'threshold', internalType: 'uint256', type: 'uint256' },
      { name: 'weight', internalType: 'uint256', type: 'uint256' },
      { name: 'imageHash', internalType: 'bytes32', type: 'bytes32' },
      { name: 'subdigest', internalType: 'bytes32', type: 'bytes32' },
      { name: 'checkpoint', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '_interfaceID', internalType: 'bytes4', type: 'bytes4' }],
    name: 'supportsInterface',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '$symbol', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'tokenId',
    outputs: [{ name: '$tokenId', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '$newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: '_imageHash', internalType: 'bytes32', type: 'bytes32' }],
    name: 'updateImageHash',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'newImplementation', internalType: 'address', type: 'address' },
      { name: 'data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'upgradeToAndCall',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '$version', internalType: 'string', type: 'string' }],
    stateMutability: 'pure',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'newImageHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: false,
      },
    ],
    name: 'ImageHashUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'implementation',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'Upgraded',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'address', type: 'address' },
      { name: '$reality', internalType: 'address', type: 'address' },
    ],
    name: 'CallerInvalid',
  },
  {
    type: 'error',
    inputs: [
      { name: '$recipient', internalType: 'address', type: 'address' },
      { name: '$value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CompensationFailed',
  },
  { type: 'error', inputs: [], name: 'CurrentInvalid' },
  { type: 'error', inputs: [], name: 'EmptySignature' },
  { type: 'error', inputs: [], name: 'ImageHashIsZero' },
  {
    type: 'error',
    inputs: [
      { name: '_hash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_addr', internalType: 'address', type: 'address' },
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'InvalidNestedSignature',
  },
  {
    type: 'error',
    inputs: [
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
      { name: '_s', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'InvalidSValue',
  },
  {
    type: 'error',
    inputs: [{ name: '_flag', internalType: 'uint256', type: 'uint256' }],
    name: 'InvalidSignatureFlag',
  },
  {
    type: 'error',
    inputs: [{ name: '_signature', internalType: 'bytes', type: 'bytes' }],
    name: 'InvalidSignatureLength',
  },
  {
    type: 'error',
    inputs: [{ name: '_type', internalType: 'bytes1', type: 'bytes1' }],
    name: 'InvalidSignatureType',
  },
  {
    type: 'error',
    inputs: [
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
      { name: '_v', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'InvalidVValue',
  },
  {
    type: 'error',
    inputs: [
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
      { name: 'threshold', internalType: 'uint256', type: 'uint256' },
      { name: '_weight', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'LowWeightChainedSignature',
  },
  {
    type: 'error',
    inputs: [
      { name: '_sender', internalType: 'address', type: 'address' },
      { name: '_self', internalType: 'address', type: 'address' },
    ],
    name: 'OnlySelfAuth',
  },
  { type: 'error', inputs: [], name: 'Reentrancy' },
  {
    type: 'error',
    inputs: [{ name: '$reality', internalType: 'address', type: 'address' }],
    name: 'RouterInvalid',
  },
  { type: 'error', inputs: [], name: 'SignatureInvalid' },
  {
    type: 'error',
    inputs: [{ name: '_signature', internalType: 'bytes', type: 'bytes' }],
    name: 'SignerIsAddress0',
  },
  {
    type: 'error',
    inputs: [
      { name: '$expected', internalType: 'address', type: 'address' },
      { name: '$reality', internalType: 'address', type: 'address' },
    ],
    name: 'SolverInvalid',
  },
  { type: 'error', inputs: [], name: 'TradingAlreadyInitialized' },
  { type: 'error', inputs: [], name: 'UnauthorizedCallContext' },
  {
    type: 'error',
    inputs: [
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
      { name: '_type', internalType: 'uint256', type: 'uint256' },
      { name: '_recoverMode', internalType: 'bool', type: 'bool' },
    ],
    name: 'UnsupportedSignatureType',
  },
  { type: 'error', inputs: [], name: 'UpgradeFailed' },
  {
    type: 'error',
    inputs: [
      { name: '_current', internalType: 'uint256', type: 'uint256' },
      { name: '_prev', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'WrongChainedCheckpointOrder',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugWindowFuse
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugWindowFuseAbi = [
  {
    type: 'function',
    inputs: [{ name: '$schedule', internalType: 'uint256', type: 'uint256' }],
    name: 'decode',
    outputs: [
      { name: '$startTime', internalType: 'uint32', type: 'uint32' },
      { name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
      { name: '$duration', internalType: 'uint32', type: 'uint32' },
      { name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$startTime', internalType: 'uint32', type: 'uint32' },
      { name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
      { name: '$duration', internalType: 'uint32', type: 'uint32' },
      { name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' },
    ],
    name: 'encode',
    outputs: [{ name: '$schedule', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$live', internalType: 'bytes', type: 'bytes' },
      {
        name: '$current',
        internalType: 'struct PlugTypesLib.Current',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'enforceFuse',
    outputs: [{ name: '$through', internalType: 'bytes', type: 'bytes' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: '$startTime', internalType: 'uint32', type: 'uint32' },
      { name: '$repeatsEvery', internalType: 'uint32', type: 'uint32' },
      { name: '$duration', internalType: 'uint32', type: 'uint32' },
      { name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' },
    ],
    name: 'isWithinWindow',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '$schedule', internalType: 'uint256', type: 'uint256' }],
    name: 'isWithinWindow',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '$schedule', internalType: 'uint256', type: 'uint256' }],
    name: 'toWindow',
    outputs: [
      {
        name: '$window',
        internalType: 'struct WindowFuseLib.Window',
        type: 'tuple',
        components: [
          {
            name: 'periods',
            internalType: 'struct WindowFuseLib.Period[]',
            type: 'tuple[]',
            components: [
              { name: 'startTime', internalType: 'uint32', type: 'uint32' },
              { name: 'endTime', internalType: 'uint32', type: 'uint32' },
            ],
          },
        ],
      },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$startTime', internalType: 'uint32', type: 'uint32' },
      { name: '$duration', internalType: 'uint32', type: 'uint32' },
      { name: '$daysOfWeek', internalType: 'uint8', type: 'uint8' },
    ],
    name: 'toWindow',
    outputs: [
      {
        name: '$window',
        internalType: 'struct WindowFuseLib.Window',
        type: 'tuple',
        components: [
          {
            name: 'periods',
            internalType: 'struct WindowFuseLib.Period[]',
            type: 'tuple[]',
            components: [
              { name: 'startTime', internalType: 'uint32', type: 'uint32' },
              { name: 'endTime', internalType: 'uint32', type: 'uint32' },
            ],
          },
        ],
      },
    ],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$schedule', internalType: 'uint256', type: 'uint256' },
      { name: '$n', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'toWindows',
    outputs: [
      {
        name: '$windows',
        internalType: 'struct WindowFuseLib.Window[]',
        type: 'tuple[]',
        components: [
          {
            name: 'periods',
            internalType: 'struct WindowFuseLib.Period[]',
            type: 'tuple[]',
            components: [
              { name: 'startTime', internalType: 'uint32', type: 'uint32' },
              { name: 'endTime', internalType: 'uint32', type: 'uint32' },
            ],
          },
        ],
      },
      { name: '$cursor', internalType: 'uint32', type: 'uint32' },
    ],
    stateMutability: 'view',
  },
  { type: 'error', inputs: [], name: 'WindowCaveatViolation' },
  { type: 'error', inputs: [], name: 'WindowLackingDays' },
  { type: 'error', inputs: [], name: 'WindowLackingDuration' },
  { type: 'error', inputs: [], name: 'WindowLackingStartTime' },
  { type: 'error', inputs: [], name: 'WindowLackingSufficientRepeatsEvery' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// React
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useReadPlug = /*#__PURE__*/ createUseReadContract({ abi: plugAbi })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"name"`
 */
export const useReadPlugName = /*#__PURE__*/ createUseReadContract({
  abi: plugAbi,
  functionName: 'name',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"symbol"`
 */
export const useReadPlugSymbol = /*#__PURE__*/ createUseReadContract({
  abi: plugAbi,
  functionName: 'symbol',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useWritePlug = /*#__PURE__*/ createUseWriteContract({
  abi: plugAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"plug"`
 */
export const useWritePlugPlug = /*#__PURE__*/ createUseWriteContract({
  abi: plugAbi,
  functionName: 'plug',
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugAbi}__
 */
export const useSimulatePlug = /*#__PURE__*/ createUseSimulateContract({
  abi: plugAbi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugAbi}__ and `functionName` set to `"plug"`
 */
export const useSimulatePlugPlug = /*#__PURE__*/ createUseSimulateContract({
  abi: plugAbi,
  functionName: 'plug',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__
 */
export const useReadPlugBaseFeeFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugBaseFeeFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugBaseFeeFuseDecode = /*#__PURE__*/ createUseReadContract(
  { abi: plugBaseFeeFuseAbi, functionName: 'decode' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugBaseFeeFuseEncode = /*#__PURE__*/ createUseReadContract(
  { abi: plugBaseFeeFuseAbi, functionName: 'encode' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBaseFeeFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugBaseFeeFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugBaseFeeFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__
 */
export const useReadPlugBlockNumberFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugBlockNumberFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugBlockNumberFuseDecode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugBlockNumberFuseAbi,
    functionName: 'decode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugBlockNumberFuseEncode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugBlockNumberFuseAbi,
    functionName: 'encode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugBlockNumberFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugBlockNumberFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugBlockNumberFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__
 */
export const useReadPlugClampFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugClampFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugClampFuseDecode = /*#__PURE__*/ createUseReadContract({
  abi: plugClampFuseAbi,
  functionName: 'decode',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugClampFuseEncode = /*#__PURE__*/ createUseReadContract({
  abi: plugClampFuseAbi,
  functionName: 'encode',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugClampFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugClampFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugClampFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useReadPlugFactory = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"balanceOf"`
 */
export const useReadPlugFactoryBalanceOf = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'balanceOf',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"getAddress"`
 */
export const useReadPlugFactoryGetAddress = /*#__PURE__*/ createUseReadContract(
  { abi: plugFactoryAbi, functionName: 'getAddress' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"getApproved"`
 */
export const useReadPlugFactoryGetApproved =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'getApproved',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"implementations"`
 */
export const useReadPlugFactoryImplementations =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'implementations',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"initCodeHash"`
 */
export const useReadPlugFactoryInitCodeHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'initCodeHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"isApprovedForAll"`
 */
export const useReadPlugFactoryIsApprovedForAll =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'isApprovedForAll',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"name"`
 */
export const useReadPlugFactoryName = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'name',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugFactoryOwner = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"ownerOf"`
 */
export const useReadPlugFactoryOwnerOf = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'ownerOf',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"ownershipHandoverExpiresAt"`
 */
export const useReadPlugFactoryOwnershipHandoverExpiresAt =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'ownershipHandoverExpiresAt',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"supportsInterface"`
 */
export const useReadPlugFactorySupportsInterface =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'supportsInterface',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"symbol"`
 */
export const useReadPlugFactorySymbol = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'symbol',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"tokenURI"`
 */
export const useReadPlugFactoryTokenUri = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
  functionName: 'tokenURI',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWritePlugFactory = /*#__PURE__*/ createUseWriteContract({
  abi: plugFactoryAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"approve"`
 */
export const useWritePlugFactoryApprove = /*#__PURE__*/ createUseWriteContract({
  abi: plugFactoryAbi,
  functionName: 'approve',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useWritePlugFactoryCancelOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useWritePlugFactoryCompleteOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"deploy"`
 */
export const useWritePlugFactoryDeploy = /*#__PURE__*/ createUseWriteContract({
  abi: plugFactoryAbi,
  functionName: 'deploy',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"initialize"`
 */
export const useWritePlugFactoryInitialize =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePlugFactoryRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useWritePlugFactoryRequestOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"safeTransferFrom"`
 */
export const useWritePlugFactorySafeTransferFrom =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'safeTransferFrom',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setApprovalForAll"`
 */
export const useWritePlugFactorySetApprovalForAll =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'setApprovalForAll',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setBaseURI"`
 */
export const useWritePlugFactorySetBaseUri =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'setBaseURI',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setImplementation"`
 */
export const useWritePlugFactorySetImplementation =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'setImplementation',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"transferFrom"`
 */
export const useWritePlugFactoryTransferFrom =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'transferFrom',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugFactoryTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugFactoryAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useSimulatePlugFactory = /*#__PURE__*/ createUseSimulateContract({
  abi: plugFactoryAbi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"approve"`
 */
export const useSimulatePlugFactoryApprove =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'approve',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useSimulatePlugFactoryCancelOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useSimulatePlugFactoryCompleteOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"deploy"`
 */
export const useSimulatePlugFactoryDeploy =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'deploy',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"initialize"`
 */
export const useSimulatePlugFactoryInitialize =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePlugFactoryRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useSimulatePlugFactoryRequestOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"safeTransferFrom"`
 */
export const useSimulatePlugFactorySafeTransferFrom =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'safeTransferFrom',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setApprovalForAll"`
 */
export const useSimulatePlugFactorySetApprovalForAll =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'setApprovalForAll',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setBaseURI"`
 */
export const useSimulatePlugFactorySetBaseUri =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'setBaseURI',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"setImplementation"`
 */
export const useSimulatePlugFactorySetImplementation =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'setImplementation',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"transferFrom"`
 */
export const useSimulatePlugFactoryTransferFrom =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'transferFrom',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugFactoryTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugFactoryAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWatchPlugFactoryEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugFactoryAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"Approval"`
 */
export const useWatchPlugFactoryApprovalEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'Approval',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"ApprovalForAll"`
 */
export const useWatchPlugFactoryApprovalForAllEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'ApprovalForAll',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"OwnershipHandoverCanceled"`
 */
export const useWatchPlugFactoryOwnershipHandoverCanceledEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'OwnershipHandoverCanceled',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"OwnershipHandoverRequested"`
 */
export const useWatchPlugFactoryOwnershipHandoverRequestedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'OwnershipHandoverRequested',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPlugFactoryOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"Transfer"`
 */
export const useWatchPlugFactoryTransferEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'Transfer',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useReadPlugLimitedCallsFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugLimitedCallsFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugLimitedCallsFuseDecode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugLimitedCallsFuseAbi,
    functionName: 'decode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugLimitedCallsFuseEncode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugLimitedCallsFuseAbi,
    functionName: 'encode',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useWritePlugLimitedCallsFuse =
  /*#__PURE__*/ createUseWriteContract({ abi: plugLimitedCallsFuseAbi })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useWritePlugLimitedCallsFuseEnforceFuse =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugLimitedCallsFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__
 */
export const useSimulatePlugLimitedCallsFuse =
  /*#__PURE__*/ createUseSimulateContract({ abi: plugLimitedCallsFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugLimitedCallsFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useSimulatePlugLimitedCallsFuseEnforceFuse =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugLimitedCallsFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__
 */
export const useReadPlugNounsIdFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugNounsIdFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugNounsIdFuseDecode = /*#__PURE__*/ createUseReadContract(
  { abi: plugNounsIdFuseAbi, functionName: 'decode' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugNounsIdFuseEncode = /*#__PURE__*/ createUseReadContract(
  { abi: plugNounsIdFuseAbi, functionName: 'encode' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsIdFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugNounsIdFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsIdFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useReadPlugNounsTraitFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugNounsTraitFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"ACCESSORY_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseAccessorySelector =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'ACCESSORY_SELECTOR',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"BACKGROUND_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseBackgroundSelector =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'BACKGROUND_SELECTOR',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"BODY_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseBodySelector =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'BODY_SELECTOR',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"GLASSES_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseGlassesSelector =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'GLASSES_SELECTOR',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"HEAD_SELECTOR"`
 */
export const useReadPlugNounsTraitFuseHeadSelector =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'HEAD_SELECTOR',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugNounsTraitFuseDecode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'decode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugNounsTraitFuseEncode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'encode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugNounsTraitFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"nounTrait"`
 */
export const useReadPlugNounsTraitFuseNounTrait =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'nounTrait',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugNounsTraitFuseOwner =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'owner',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"ownershipHandoverExpiresAt"`
 */
export const useReadPlugNounsTraitFuseOwnershipHandoverExpiresAt =
  /*#__PURE__*/ createUseReadContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'ownershipHandoverExpiresAt',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useWritePlugNounsTraitFuse = /*#__PURE__*/ createUseWriteContract({
  abi: plugNounsTraitFuseAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseCancelOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseCompleteOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePlugNounsTraitFuseRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useWritePlugNounsTraitFuseRequestOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"setArt"`
 */
export const useWritePlugNounsTraitFuseSetArt =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'setArt',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugNounsTraitFuseTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useSimulatePlugNounsTraitFuse =
  /*#__PURE__*/ createUseSimulateContract({ abi: plugNounsTraitFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseCancelOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseCompleteOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePlugNounsTraitFuseRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useSimulatePlugNounsTraitFuseRequestOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"setArt"`
 */
export const useSimulatePlugNounsTraitFuseSetArt =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'setArt',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugNounsTraitFuseTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugNounsTraitFuseAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__
 */
export const useWatchPlugNounsTraitFuseEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugNounsTraitFuseAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipHandoverCanceled"`
 */
export const useWatchPlugNounsTraitFuseOwnershipHandoverCanceledEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugNounsTraitFuseAbi,
    eventName: 'OwnershipHandoverCanceled',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipHandoverRequested"`
 */
export const useWatchPlugNounsTraitFuseOwnershipHandoverRequestedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugNounsTraitFuseAbi,
    eventName: 'OwnershipHandoverRequested',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugNounsTraitFuseAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPlugNounsTraitFuseOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugNounsTraitFuseAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useReadPlugRevocationFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugRevocationFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugRevocationFuseDecode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugRevocationFuseAbi,
    functionName: 'decode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugRevocationFuseEncode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugRevocationFuseAbi,
    functionName: 'encode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugRevocationFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugRevocationFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"isRevoked"`
 */
export const useReadPlugRevocationFuseIsRevoked =
  /*#__PURE__*/ createUseReadContract({
    abi: plugRevocationFuseAbi,
    functionName: 'isRevoked',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useWritePlugRevocationFuse = /*#__PURE__*/ createUseWriteContract({
  abi: plugRevocationFuseAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"revoke"`
 */
export const useWritePlugRevocationFuseRevoke =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugRevocationFuseAbi,
    functionName: 'revoke',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__
 */
export const useSimulatePlugRevocationFuse =
  /*#__PURE__*/ createUseSimulateContract({ abi: plugRevocationFuseAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugRevocationFuseAbi}__ and `functionName` set to `"revoke"`
 */
export const useSimulatePlugRevocationFuseRevoke =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugRevocationFuseAbi,
    functionName: 'revoke',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__
 */
export const useReadPlugTimestampFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugTimestampFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugTimestampFuseDecode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugTimestampFuseAbi,
    functionName: 'decode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugTimestampFuseEncode =
  /*#__PURE__*/ createUseReadContract({
    abi: plugTimestampFuseAbi,
    functionName: 'encode',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTimestampFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugTimestampFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugTimestampFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTreasuryAbi}__
 */
export const useReadPlugTreasury = /*#__PURE__*/ createUseReadContract({
  abi: plugTreasuryAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugTreasuryOwner = /*#__PURE__*/ createUseReadContract({
  abi: plugTreasuryAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"ownershipHandoverExpiresAt"`
 */
export const useReadPlugTreasuryOwnershipHandoverExpiresAt =
  /*#__PURE__*/ createUseReadContract({
    abi: plugTreasuryAbi,
    functionName: 'ownershipHandoverExpiresAt',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__
 */
export const useWritePlugTreasury = /*#__PURE__*/ createUseWriteContract({
  abi: plugTreasuryAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useWritePlugTreasuryCancelOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useWritePlugTreasuryCompleteOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"execute"`
 */
export const useWritePlugTreasuryExecute = /*#__PURE__*/ createUseWriteContract(
  { abi: plugTreasuryAbi, functionName: 'execute' },
)

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"initialize"`
 */
export const useWritePlugTreasuryInitialize =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePlugTreasuryRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useWritePlugTreasuryRequestOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugTreasuryTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugTreasuryAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__
 */
export const useSimulatePlugTreasury = /*#__PURE__*/ createUseSimulateContract({
  abi: plugTreasuryAbi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useSimulatePlugTreasuryCancelOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useSimulatePlugTreasuryCompleteOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"execute"`
 */
export const useSimulatePlugTreasuryExecute =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'execute',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"initialize"`
 */
export const useSimulatePlugTreasuryInitialize =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePlugTreasuryRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useSimulatePlugTreasuryRequestOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugTreasuryAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugTreasuryTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugTreasuryAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugTreasuryAbi}__
 */
export const useWatchPlugTreasuryEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugTreasuryAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugTreasuryAbi}__ and `eventName` set to `"OwnershipHandoverCanceled"`
 */
export const useWatchPlugTreasuryOwnershipHandoverCanceledEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugTreasuryAbi,
    eventName: 'OwnershipHandoverCanceled',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugTreasuryAbi}__ and `eventName` set to `"OwnershipHandoverRequested"`
 */
export const useWatchPlugTreasuryOwnershipHandoverRequestedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugTreasuryAbi,
    eventName: 'OwnershipHandoverRequested',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugTreasuryAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPlugTreasuryOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugTreasuryAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__
 */
export const useReadPlugVaultSocket = /*#__PURE__*/ createUseReadContract({
  abi: plugVaultSocketAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"SET_IMAGE_HASH_TYPE_HASH"`
 */
export const useReadPlugVaultSocketSetImageHashTypeHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'SET_IMAGE_HASH_TYPE_HASH',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getCurrentHash"`
 */
export const useReadPlugVaultSocketGetCurrentHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getCurrentHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getEIP712DomainHash"`
 */
export const useReadPlugVaultSocketGetEip712DomainHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getEIP712DomainHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getFuseArrayHash"`
 */
export const useReadPlugVaultSocketGetFuseArrayHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getFuseArrayHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getFuseHash"`
 */
export const useReadPlugVaultSocketGetFuseHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getFuseHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getLivePlugsHash"`
 */
export const useReadPlugVaultSocketGetLivePlugsHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getLivePlugsHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getPlugArrayHash"`
 */
export const useReadPlugVaultSocketGetPlugArrayHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getPlugArrayHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getPlugHash"`
 */
export const useReadPlugVaultSocketGetPlugHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getPlugHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"getPlugsHash"`
 */
export const useReadPlugVaultSocketGetPlugsHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'getPlugsHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"imageHash"`
 */
export const useReadPlugVaultSocketImageHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'imageHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"isValidSignature"`
 */
export const useReadPlugVaultSocketIsValidSignature =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'isValidSignature',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"name"`
 */
export const useReadPlugVaultSocketName = /*#__PURE__*/ createUseReadContract({
  abi: plugVaultSocketAbi,
  functionName: 'name',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugVaultSocketOwner = /*#__PURE__*/ createUseReadContract({
  abi: plugVaultSocketAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"ownership"`
 */
export const useReadPlugVaultSocketOwnership =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'ownership',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"proxiableUUID"`
 */
export const useReadPlugVaultSocketProxiableUuid =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'proxiableUUID',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"router"`
 */
export const useReadPlugVaultSocketRouter = /*#__PURE__*/ createUseReadContract(
  { abi: plugVaultSocketAbi, functionName: 'router' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"signatureRecovery"`
 */
export const useReadPlugVaultSocketSignatureRecovery =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'signatureRecovery',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"supportsInterface"`
 */
export const useReadPlugVaultSocketSupportsInterface =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'supportsInterface',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"symbol"`
 */
export const useReadPlugVaultSocketSymbol = /*#__PURE__*/ createUseReadContract(
  { abi: plugVaultSocketAbi, functionName: 'symbol' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"tokenId"`
 */
export const useReadPlugVaultSocketTokenId =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'tokenId',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"version"`
 */
export const useReadPlugVaultSocketVersion =
  /*#__PURE__*/ createUseReadContract({
    abi: plugVaultSocketAbi,
    functionName: 'version',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__
 */
export const useWritePlugVaultSocket = /*#__PURE__*/ createUseWriteContract({
  abi: plugVaultSocketAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"initialize"`
 */
export const useWritePlugVaultSocketInitialize =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugVaultSocketAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"plug"`
 */
export const useWritePlugVaultSocketPlug = /*#__PURE__*/ createUseWriteContract(
  { abi: plugVaultSocketAbi, functionName: 'plug' },
)

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugVaultSocketTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugVaultSocketAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"updateImageHash"`
 */
export const useWritePlugVaultSocketUpdateImageHash =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugVaultSocketAbi,
    functionName: 'updateImageHash',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"upgradeToAndCall"`
 */
export const useWritePlugVaultSocketUpgradeToAndCall =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugVaultSocketAbi,
    functionName: 'upgradeToAndCall',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__
 */
export const useSimulatePlugVaultSocket =
  /*#__PURE__*/ createUseSimulateContract({ abi: plugVaultSocketAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"initialize"`
 */
export const useSimulatePlugVaultSocketInitialize =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugVaultSocketAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"plug"`
 */
export const useSimulatePlugVaultSocketPlug =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugVaultSocketAbi,
    functionName: 'plug',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugVaultSocketTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugVaultSocketAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"updateImageHash"`
 */
export const useSimulatePlugVaultSocketUpdateImageHash =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugVaultSocketAbi,
    functionName: 'updateImageHash',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `functionName` set to `"upgradeToAndCall"`
 */
export const useSimulatePlugVaultSocketUpgradeToAndCall =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugVaultSocketAbi,
    functionName: 'upgradeToAndCall',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugVaultSocketAbi}__
 */
export const useWatchPlugVaultSocketEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugVaultSocketAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `eventName` set to `"ImageHashUpdated"`
 */
export const useWatchPlugVaultSocketImageHashUpdatedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugVaultSocketAbi,
    eventName: 'ImageHashUpdated',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugVaultSocketAbi}__ and `eventName` set to `"Upgraded"`
 */
export const useWatchPlugVaultSocketUpgradedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugVaultSocketAbi,
    eventName: 'Upgraded',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__
 */
export const useReadPlugWindowFuse = /*#__PURE__*/ createUseReadContract({
  abi: plugWindowFuseAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"decode"`
 */
export const useReadPlugWindowFuseDecode = /*#__PURE__*/ createUseReadContract({
  abi: plugWindowFuseAbi,
  functionName: 'decode',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"encode"`
 */
export const useReadPlugWindowFuseEncode = /*#__PURE__*/ createUseReadContract({
  abi: plugWindowFuseAbi,
  functionName: 'encode',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"enforceFuse"`
 */
export const useReadPlugWindowFuseEnforceFuse =
  /*#__PURE__*/ createUseReadContract({
    abi: plugWindowFuseAbi,
    functionName: 'enforceFuse',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"isWithinWindow"`
 */
export const useReadPlugWindowFuseIsWithinWindow =
  /*#__PURE__*/ createUseReadContract({
    abi: plugWindowFuseAbi,
    functionName: 'isWithinWindow',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"toWindow"`
 */
export const useReadPlugWindowFuseToWindow =
  /*#__PURE__*/ createUseReadContract({
    abi: plugWindowFuseAbi,
    functionName: 'toWindow',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugWindowFuseAbi}__ and `functionName` set to `"toWindows"`
 */
export const useReadPlugWindowFuseToWindows =
  /*#__PURE__*/ createUseReadContract({
    abi: plugWindowFuseAbi,
    functionName: 'toWindows',
  })
