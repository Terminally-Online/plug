package euler

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GetSubAccountAddress(account common.Address, index uint8) common.Address {
	// Convert owner address to big.Int for bitwise operations
    ownerInt := new(big.Int).SetBytes(account[:])
    
    // Create big.Int for accountId
    accountIdInt := new(big.Int).SetUint64(uint64(index))
    
    // XOR operation
    result := new(big.Int).Xor(ownerInt, accountIdInt)
    
    // Convert back to address
    var subAccount common.Address
    resultBytes := result.Bytes()
    
    // Ensure proper padding to 20 bytes (address length)
    copy(subAccount[20-len(resultBytes):], resultBytes)
    
    return subAccount
}