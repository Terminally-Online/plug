package zerion

import (
	"fmt"
	"os"
	"solver/internal/utils"
	"strings"
)

type ZerionNonFungibleMedia struct {
	Preview *struct {
		URL string `json:"url"`
	} `json:"preview"`
	Detail *struct {
		URL string `json:"url"`
	} `json:"detail"`
	Video *struct {
		URL string `json:"url"`
	} `json:"video"`
	Icon struct {
		URL string `json:"url"`
	} `json:"icon"`
}

type ZerionNonFungibleMetadata struct {
	Name            string                 `json:"name"`
	TokenID         string                 `json:"token_id"`
	ContractAddress string                 `json:"contract_address"`
	Interface       string                 `json:"interface"`
	Content         ZerionNonFungibleMedia `json:"content"`
	Flags           struct {
		IsSpam bool `json:"is_spam"`
	} `json:"flags"`
}

type ZerionNonFungibleAttributes struct {
	Amount         float64                   `json:"amount"`
	ChangedAt      string                    `json:"changed_at"`
	NFTInfo        ZerionNonFungibleMetadata `json:"nft_info"`
	CollectionInfo struct {
		Name        string                 `json:"name"`
		Description string                 `json:"description"`
		Content     ZerionNonFungibleMedia `json:"content"`
	} `json:"collection_info"`
}

type ZerionNonFungible struct {
	Attributes    ZerionNonFungibleAttributes `json:"attributes"`
	Relationships struct {
		Chain struct {
			Data struct {
				ID string `json:"id"`
			} `json:"data"`
		} `json:"chain"`
	} `json:"relationships"`
}

func GetCollectibles(chains []string, socketId, socketAddress string, limit int) ([]ZerionNonFungible, error) {
	return getCollectibles(chains, socketId, socketAddress, limit, "")
}

func getCollectibles(chains []string, socketId, socketAddress string, limit int, next string) ([]ZerionNonFungible, error) {
	if limit == 0 {
		limit = 100
	}

	address := socketAddress
	if address == "" {
		address = socketId
	}

	url := next
	if url == "" {
		url = fmt.Sprintf(
			"https://api.zerion.io/v1/wallets/%s/nft-positions/?filter[chain_ids]=%s&currency=usd&page[size]=%d",
			address,
			strings.Join(chains, ","),
			limit,
		)
	}

	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"accept":        "application/json",
			"authorization": fmt.Sprintf("Basic %v", os.Getenv("ZERION_KEY")),
		},
		nil,
		nil,
		struct {
			Data  []ZerionNonFungible `json:"data"`
			Links struct {
				Next string `json:"next"`
			}
		}{},
	)
	if err != nil {
		return nil, err
	}

	if response.Links.Next != "" {
		nextNonFungibles, err := getCollectibles(chains, socketId, socketAddress, limit, response.Links.Next)
		if err != nil {
			return nil, err
		}
		response.Data = append(response.Data, nextNonFungibles...)
	}

	return response.Data, nil
}
