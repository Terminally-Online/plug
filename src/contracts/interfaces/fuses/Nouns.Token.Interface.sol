// SPDX-License-Identifier: MIT

pragma solidity 0.8.18;

interface NounsTokenInterface {
    event Approval(
        address indexed owner,
        address indexed approved,
        uint256 indexed tokenId
    );
    event ApprovalForAll(
        address indexed owner, address indexed operator, bool approved
    );
    event DelegateChanged(
        address indexed delegator,
        address indexed fromDelegate,
        address indexed toDelegate
    );
    event DelegateVotesChanged(
        address indexed delegate,
        uint256 previousBalance,
        uint256 newBalance
    );
    event DescriptorLocked();
    event DescriptorUpdated(address descriptor);
    event MinterLocked();
    event MinterUpdated(address minter);
    event NounBurned(uint256 indexed tokenId);
    event NounCreated(
        uint256 indexed tokenId, INounsSeeder.Seed seed
    );
    event NoundersDAOUpdated(address noundersDAO);
    event OwnershipTransferred(
        address indexed previousOwner, address indexed newOwner
    );
    event SeederLocked();
    event SeederUpdated(address seeder);
    event Transfer(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );

    function DELEGATION_TYPEHASH() external view returns (bytes32);

    function DOMAIN_TYPEHASH() external view returns (bytes32);

    function approve(address to, uint256 tokenId) external;

    function balanceOf(address owner)
        external
        view
        returns (uint256);

    function burn(uint256 nounId) external;

    function checkpoints(
        address,
        uint32
    )
        external
        view
        returns (uint32 fromBlock, uint96 votes);

    function contractURI() external view returns (string memory);

    function dataURI(uint256 tokenId)
        external
        view
        returns (string memory);

    function decimals() external view returns (uint8);

    function delegate(address delegatee) external;

    function delegateBySig(
        address delegatee,
        uint256 nonce,
        uint256 expiry,
        uint8 v,
        bytes32 r,
        bytes32 s
    )
        external;

    function delegates(address delegator)
        external
        view
        returns (address);

    function descriptor() external view returns (address);

    function getApproved(uint256 tokenId)
        external
        view
        returns (address);

    function getCurrentVotes(address account)
        external
        view
        returns (uint96);

    function getPriorVotes(
        address account,
        uint256 blockNumber
    )
        external
        view
        returns (uint96);

    function isApprovedForAll(
        address owner,
        address operator
    )
        external
        view
        returns (bool);

    function isDescriptorLocked() external view returns (bool);

    function isMinterLocked() external view returns (bool);

    function isSeederLocked() external view returns (bool);

    function lockDescriptor() external;

    function lockMinter() external;

    function lockSeeder() external;

    function mint() external returns (uint256);

    function minter() external view returns (address);

    function name() external view returns (string memory);

    function nonces(address) external view returns (uint256);

    function noundersDAO() external view returns (address);

    function numCheckpoints(address) external view returns (uint32);

    function owner() external view returns (address);

    function ownerOf(uint256 tokenId)
        external
        view
        returns (address);

    function proxyRegistry() external view returns (address);

    function renounceOwnership() external;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    )
        external;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory _data
    )
        external;

    function seeder() external view returns (address);

    function seeds(uint256)
        external
        view
        returns (
            uint48 background,
            uint48 body,
            uint48 accessory,
            uint48 head,
            uint48 glasses
        );

    function setApprovalForAll(
        address operator,
        bool approved
    )
        external;

    function setContractURIHash(string memory newContractURIHash)
        external;

    function setDescriptor(address _descriptor) external;

    function setMinter(address _minter) external;

    function setNoundersDAO(address _noundersDAO) external;

    function setSeeder(address _seeder) external;

    function supportsInterface(bytes4 interfaceId)
        external
        view
        returns (bool);

    function symbol() external view returns (string memory);

    function tokenByIndex(uint256 index)
        external
        view
        returns (uint256);

    function tokenOfOwnerByIndex(
        address owner,
        uint256 index
    )
        external
        view
        returns (uint256);

    function tokenURI(uint256 tokenId)
        external
        view
        returns (string memory);

    function totalSupply() external view returns (uint256);

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    )
        external;

    function transferOwnership(address newOwner) external;

    function votesToDelegate(address delegator)
        external
        view
        returns (uint96);
}

interface INounsSeeder {
    struct Seed {
        uint48 background;
        uint48 body;
        uint48 accessory;
        uint48 head;
        uint48 glasses;
    }
}
