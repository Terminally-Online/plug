package types

import (
	"encoding/json"
	"fmt"
	"solver/bindings/morpho_router"
	"strings"

	"encoding/base64"

	"github.com/ethereum/go-ethereum/common"
)

/**
* MorphoTargetParams acts as an intermediary data structure that we can store as a serialized
* string so that it can be embedded in an option and passed into the action without needing to
* make a read call to get the params for the market from within the action.
 */
type MorphoTargetParams struct {
	morpho_router.MarketParams
	TargetId string `json:"targetId" description:"The target ID is either the market uuid, or the address of the vault depending on the type."`
}

// SerializeToCompactString serializes the struct into a compact string format
// If MarketParams is empty/undefined, it will be represented as an empty JSON object {}
func (u MorphoTargetParams) SerializeToCompactString() string {
	// Serialize the MarketParams to JSON, which will handle empty values appropriately
	paramsJSON, err := json.Marshal(u.MarketParams)
	if err != nil {
		// If marshaling fails for some reason, use an empty object
		paramsJSON = []byte("{}")
	}

	// Use base64 encoding to avoid issues with special characters in the JSON
	encodedParams := base64.StdEncoding.EncodeToString(paramsJSON)

	// Combine the targetId and encoded params
	return fmt.Sprintf("%s:%s", u.TargetId, encodedParams)
}

// DeserializeFromCompactString parses the compact string format back to MorphoTargetParams
func DeserializeFromCompactString(data string) (MorphoTargetParams, error) {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		return MorphoTargetParams{}, fmt.Errorf("invalid compact format: %s", data)
	}

	var result MorphoTargetParams
	result.TargetId = parts[0]

	// Decode the base64 encoded params
	paramsJSON, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return MorphoTargetParams{}, fmt.Errorf("failed to decode params: %w", err)
	}

	// Handle empty JSON object case
	if string(paramsJSON) == "{}" || string(paramsJSON) == "" {
		// Empty market params, just return with the targetId
		return result, nil
	}

	// Parse the market params JSON
	if err := json.Unmarshal(paramsJSON, &result.MarketParams); err != nil {
		return MorphoTargetParams{}, fmt.Errorf("failed to deserialize market params: %w", err)
	}

	return result, nil
}

// IsMarketParamsDefined checks if the MarketParams has meaningful values (i.e. this is a market, not a vault)
func (u MorphoTargetParams) IsMarketParamsDefined() bool {
	return u.LoanToken != common.Address{} ||
		u.CollateralToken != common.Address{} ||
		u.Oracle != common.Address{} ||
		u.Irm != common.Address{} ||
		(u.Lltv != nil && u.Lltv.Sign() != 0)
}
