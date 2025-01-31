package euler


func GetVerifiedVaults(chainId uint64) ([]Vault, error) {
	var vaultAddresses []string

}

func GetVault(address string, chainId uint64) (Vault, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return Vault{}, err
	}
	for _, vault := range vaults {
		if vault.Address == address {
			return vault, nil
		}
	}
	return Vault{}, fmt.Errorf("vault not found for address: %s", address)
}

