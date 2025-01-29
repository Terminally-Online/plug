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
                  { name: 'to', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'gas', internalType: 'uint256', type: 'uint256' },
                ],
              },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
              { name: 'salt', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [
      {
        name: '$results',
        internalType: 'struct PlugTypesLib.Result',
        type: 'tuple',
        components: [
          { name: 'index', internalType: 'uint8', type: 'uint8' },
          { name: 'error', internalType: 'string', type: 'string' },
        ],
      },
    ],
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
                  { name: 'to', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'gas', internalType: 'uint256', type: 'uint256' },
                ],
              },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
              { name: 'salt', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [
      {
        name: '$results',
        internalType: 'struct PlugTypesLib.Result[]',
        type: 'tuple[]',
        components: [
          { name: 'index', internalType: 'uint8', type: 'uint8' },
          { name: 'error', internalType: 'string', type: 'string' },
        ],
      },
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
// PlugFactory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugFactoryAbi = [
  {
    type: 'function',
    inputs: [{ name: '$salt', internalType: 'bytes', type: 'bytes' }],
    name: 'deploy',
    outputs: [
      { name: '$alreadyDeployed', internalType: 'bool', type: 'bool' },
      { name: '$socketAddress', internalType: 'address', type: 'address' },
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
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'implementation',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'vault',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'salt',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: false,
      },
    ],
    name: 'SocketDeployed',
  },
  {
    type: 'error',
    inputs: [
      { name: '$implementation', internalType: 'address', type: 'address' },
      { name: '$admin', internalType: 'address', type: 'address' },
    ],
    name: 'SaltInvalid',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PlugSocket
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const plugSocketAbi = [
  { type: 'constructor', inputs: [], stateMutability: 'nonpayable' },
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
    inputs: [],
    name: 'domain',
    outputs: [
      {
        name: '$domain',
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
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'domainHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
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
    outputs: [{ name: '$typeHash', internalType: 'bytes32', type: 'bytes32' }],
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
                  { name: 'to', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'gas', internalType: 'uint256', type: 'uint256' },
                ],
              },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
              { name: 'salt', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getLivePlugsHash',
    outputs: [{ name: '$typeHash', internalType: 'bytes32', type: 'bytes32' }],
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
                  { name: 'to', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'gas', internalType: 'uint256', type: 'uint256' },
                ],
              },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
              { name: 'salt', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getLivePlugsSigner',
    outputs: [{ name: '$signer', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      {
        name: '$input',
        internalType: 'struct PlugTypesLib.Plug[]',
        type: 'tuple[]',
        components: [
          { name: 'to', internalType: 'address', type: 'address' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gas', internalType: 'uint256', type: 'uint256' },
        ],
      },
    ],
    name: 'getPlugArrayHash',
    outputs: [{ name: '$typeHash', internalType: 'bytes32', type: 'bytes32' }],
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
          { name: 'to', internalType: 'address', type: 'address' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gas', internalType: 'uint256', type: 'uint256' },
        ],
      },
    ],
    name: 'getPlugHash',
    outputs: [{ name: '$typeHash', internalType: 'bytes32', type: 'bytes32' }],
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
              { name: 'to', internalType: 'address', type: 'address' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
              { name: 'gas', internalType: 'uint256', type: 'uint256' },
            ],
          },
          { name: 'solver', internalType: 'bytes', type: 'bytes' },
          { name: 'salt', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getPlugsDigest',
    outputs: [{ name: '$digest', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'view',
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
              { name: 'to', internalType: 'address', type: 'address' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
              { name: 'gas', internalType: 'uint256', type: 'uint256' },
            ],
          },
          { name: 'solver', internalType: 'bytes', type: 'bytes' },
          { name: 'salt', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'getPlugsHash',
    outputs: [{ name: '$typeHash', internalType: 'bytes32', type: 'bytes32' }],
    stateMutability: 'pure',
  },
  {
    type: 'function',
    inputs: [
      { name: '$owner', internalType: 'address', type: 'address' },
      { name: '$oneClicker', internalType: 'address', type: 'address' },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
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
    inputs: [
      { name: '$oneClickers', internalType: 'address[]', type: 'address[]' },
      { name: '$allowance', internalType: 'bool[]', type: 'bool[]' },
    ],
    name: 'oneClick',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'oneClicker', internalType: 'address', type: 'address' }],
    name: 'oneClickersToAllowed',
    outputs: [{ name: 'allowed', internalType: 'bool', type: 'bool' }],
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
                  { name: 'to', internalType: 'address', type: 'address' },
                  { name: 'data', internalType: 'bytes', type: 'bytes' },
                  { name: 'value', internalType: 'uint256', type: 'uint256' },
                  { name: 'gas', internalType: 'uint256', type: 'uint256' },
                ],
              },
              { name: 'solver', internalType: 'bytes', type: 'bytes' },
              { name: 'salt', internalType: 'bytes', type: 'bytes' },
            ],
          },
          { name: 'signature', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '$solver', internalType: 'address', type: 'address' },
    ],
    name: 'plug',
    outputs: [
      {
        name: '$results',
        internalType: 'struct PlugTypesLib.Result',
        type: 'tuple',
        components: [
          { name: 'index', internalType: 'uint8', type: 'uint8' },
          { name: 'error', internalType: 'string', type: 'string' },
        ],
      },
    ],
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
              { name: 'to', internalType: 'address', type: 'address' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
              { name: 'gas', internalType: 'uint256', type: 'uint256' },
            ],
          },
          { name: 'solver', internalType: 'bytes', type: 'bytes' },
          { name: 'salt', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'plug',
    outputs: [
      {
        name: '$results',
        internalType: 'struct PlugTypesLib.Result',
        type: 'tuple',
        components: [
          { name: 'index', internalType: 'uint8', type: 'uint8' },
          { name: 'error', internalType: 'string', type: 'string' },
        ],
      },
    ],
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
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '$symbol', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'payable',
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
      {
        name: 'implementation',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'Upgraded',
  },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  { type: 'error', inputs: [], name: 'NewOwnerIsZeroAddress' },
  { type: 'error', inputs: [], name: 'NoHandoverRequest' },
  {
    type: 'error',
    inputs: [
      { name: '$index', internalType: 'uint8', type: 'uint8' },
      { name: '$reason', internalType: 'string', type: 'string' },
    ],
    name: 'PlugFailed',
  },
  { type: 'error', inputs: [], name: 'Reentrancy' },
  { type: 'error', inputs: [], name: 'Unauthorized' },
  { type: 'error', inputs: [], name: 'UnauthorizedCallContext' },
  { type: 'error', inputs: [], name: 'UpgradeFailed' },
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
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useReadPlugFactory = /*#__PURE__*/ createUseReadContract({
  abi: plugFactoryAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"getAddress"`
 */
export const useReadPlugFactoryGetAddress = /*#__PURE__*/ createUseReadContract(
  { abi: plugFactoryAbi, functionName: 'getAddress' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"initCodeHash"`
 */
export const useReadPlugFactoryInitCodeHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugFactoryAbi,
    functionName: 'initCodeHash',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWritePlugFactory = /*#__PURE__*/ createUseWriteContract({
  abi: plugFactoryAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugFactoryAbi}__ and `functionName` set to `"deploy"`
 */
export const useWritePlugFactoryDeploy = /*#__PURE__*/ createUseWriteContract({
  abi: plugFactoryAbi,
  functionName: 'deploy',
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useSimulatePlugFactory = /*#__PURE__*/ createUseSimulateContract({
  abi: plugFactoryAbi,
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
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__
 */
export const useWatchPlugFactoryEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugFactoryAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugFactoryAbi}__ and `eventName` set to `"SocketDeployed"`
 */
export const useWatchPlugFactorySocketDeployedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugFactoryAbi,
    eventName: 'SocketDeployed',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__
 */
export const useReadPlugSocket = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"domain"`
 */
export const useReadPlugSocketDomain = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'domain',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"domainHash"`
 */
export const useReadPlugSocketDomainHash = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'domainHash',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getEIP712DomainHash"`
 */
export const useReadPlugSocketGetEip712DomainHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getEIP712DomainHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getLivePlugsHash"`
 */
export const useReadPlugSocketGetLivePlugsHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getLivePlugsHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getLivePlugsSigner"`
 */
export const useReadPlugSocketGetLivePlugsSigner =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getLivePlugsSigner',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getPlugArrayHash"`
 */
export const useReadPlugSocketGetPlugArrayHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getPlugArrayHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getPlugHash"`
 */
export const useReadPlugSocketGetPlugHash = /*#__PURE__*/ createUseReadContract(
  { abi: plugSocketAbi, functionName: 'getPlugHash' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getPlugsDigest"`
 */
export const useReadPlugSocketGetPlugsDigest =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getPlugsDigest',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"getPlugsHash"`
 */
export const useReadPlugSocketGetPlugsHash =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'getPlugsHash',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"name"`
 */
export const useReadPlugSocketName = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'name',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"oneClickersToAllowed"`
 */
export const useReadPlugSocketOneClickersToAllowed =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'oneClickersToAllowed',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPlugSocketOwner = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"ownershipHandoverExpiresAt"`
 */
export const useReadPlugSocketOwnershipHandoverExpiresAt =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'ownershipHandoverExpiresAt',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"proxiableUUID"`
 */
export const useReadPlugSocketProxiableUuid =
  /*#__PURE__*/ createUseReadContract({
    abi: plugSocketAbi,
    functionName: 'proxiableUUID',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"symbol"`
 */
export const useReadPlugSocketSymbol = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'symbol',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"version"`
 */
export const useReadPlugSocketVersion = /*#__PURE__*/ createUseReadContract({
  abi: plugSocketAbi,
  functionName: 'version',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__
 */
export const useWritePlugSocket = /*#__PURE__*/ createUseWriteContract({
  abi: plugSocketAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useWritePlugSocketCancelOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useWritePlugSocketCompleteOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"initialize"`
 */
export const useWritePlugSocketInitialize =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"oneClick"`
 */
export const useWritePlugSocketOneClick = /*#__PURE__*/ createUseWriteContract({
  abi: plugSocketAbi,
  functionName: 'oneClick',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"plug"`
 */
export const useWritePlugSocketPlug = /*#__PURE__*/ createUseWriteContract({
  abi: plugSocketAbi,
  functionName: 'plug',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePlugSocketRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useWritePlugSocketRequestOwnershipHandover =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePlugSocketTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"upgradeToAndCall"`
 */
export const useWritePlugSocketUpgradeToAndCall =
  /*#__PURE__*/ createUseWriteContract({
    abi: plugSocketAbi,
    functionName: 'upgradeToAndCall',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__
 */
export const useSimulatePlugSocket = /*#__PURE__*/ createUseSimulateContract({
  abi: plugSocketAbi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"cancelOwnershipHandover"`
 */
export const useSimulatePlugSocketCancelOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'cancelOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"completeOwnershipHandover"`
 */
export const useSimulatePlugSocketCompleteOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'completeOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"initialize"`
 */
export const useSimulatePlugSocketInitialize =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'initialize',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"oneClick"`
 */
export const useSimulatePlugSocketOneClick =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'oneClick',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"plug"`
 */
export const useSimulatePlugSocketPlug =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'plug',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePlugSocketRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"requestOwnershipHandover"`
 */
export const useSimulatePlugSocketRequestOwnershipHandover =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'requestOwnershipHandover',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePlugSocketTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link plugSocketAbi}__ and `functionName` set to `"upgradeToAndCall"`
 */
export const useSimulatePlugSocketUpgradeToAndCall =
  /*#__PURE__*/ createUseSimulateContract({
    abi: plugSocketAbi,
    functionName: 'upgradeToAndCall',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugSocketAbi}__
 */
export const useWatchPlugSocketEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: plugSocketAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugSocketAbi}__ and `eventName` set to `"OwnershipHandoverCanceled"`
 */
export const useWatchPlugSocketOwnershipHandoverCanceledEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugSocketAbi,
    eventName: 'OwnershipHandoverCanceled',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugSocketAbi}__ and `eventName` set to `"OwnershipHandoverRequested"`
 */
export const useWatchPlugSocketOwnershipHandoverRequestedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugSocketAbi,
    eventName: 'OwnershipHandoverRequested',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugSocketAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPlugSocketOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugSocketAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link plugSocketAbi}__ and `eventName` set to `"Upgraded"`
 */
export const useWatchPlugSocketUpgradedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: plugSocketAbi,
    eventName: 'Upgraded',
  })
